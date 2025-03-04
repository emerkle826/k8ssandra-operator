/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package reaper

import (
	"context"
	"github.com/go-logr/logr"
	cassdcapi "github.com/k8ssandra/cass-operator/apis/cassandra/v1beta1"
	reaperapi "github.com/k8ssandra/k8ssandra-operator/apis/reaper/v1alpha1"
	"github.com/k8ssandra/k8ssandra-operator/pkg/annotations"
	"github.com/k8ssandra/k8ssandra-operator/pkg/cassandra"
	"github.com/k8ssandra/k8ssandra-operator/pkg/config"
	"github.com/k8ssandra/k8ssandra-operator/pkg/reaper"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// ReaperReconciler reconciles a Reaper object
type ReaperReconciler struct {
	*config.ReconcilerConfig
	client.Client
	Scheme     *runtime.Scheme
	NewManager func() reaper.Manager
}

// +kubebuilder:rbac:groups=reaper.k8ssandra.io,namespace="k8ssandra",resources=reapers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=reaper.k8ssandra.io,namespace="k8ssandra",resources=reapers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="apps",namespace="k8ssandra",resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="core",namespace="k8ssandra",resources=pods;secrets,verbs=get;list;watch
// +kubebuilder:rbac:groups="core",namespace="k8ssandra",resources=services,verbs=get;list;watch;create

func (r *ReaperReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx, "Reaper", req.NamespacedName)

	logger.Info("Starting Reaper reconciliation")

	// Fetch the Reaper instance
	actualReaper := &reaperapi.Reaper{}
	if err := r.Get(ctx, req.NamespacedName, actualReaper); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("Reaper resource not found")
			return ctrl.Result{}, nil
		}
		logger.Info("Failed to fetch Reaper resource")
		return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
	}

	actualReaper = actualReaper.DeepCopy()
	patch := client.MergeFromWithOptions(actualReaper.DeepCopy())

	result, err := r.reconcile(ctx, actualReaper, logger)

	if patchErr := r.Status().Patch(ctx, actualReaper, patch); patchErr != nil {
		logger.Error(patchErr, "Failed to update Reaper status")
	} else {
		logger.Info("Updated Reaper status")
	}

	return result, err
}

func (r *ReaperReconciler) reconcile(ctx context.Context, actualReaper *reaperapi.Reaper, logger logr.Logger) (ctrl.Result, error) {

	actualReaper.Status.Progress = reaperapi.ReaperProgressPending
	actualReaper.Status.SetNotReady()

	actualDc, result, err := r.reconcileDatacenter(ctx, actualReaper, logger)
	if !result.IsZero() || err != nil {
		return result, err
	}

	actualReaper.Status.Progress = reaperapi.ReaperProgressDeploying

	if result, err = r.reconcileDeployment(ctx, actualReaper, actualDc, logger); !result.IsZero() || err != nil {
		return result, err
	}

	if result, err = r.reconcileService(ctx, actualReaper, logger); !result.IsZero() || err != nil {
		return result, err
	}

	actualReaper.Status.Progress = reaperapi.ReaperProgressConfiguring

	if result, err = r.configureReaper(ctx, actualReaper, actualDc, logger); !result.IsZero() || err != nil {
		return result, err
	}

	actualReaper.Status.Progress = reaperapi.ReaperProgressRunning
	actualReaper.Status.SetReady()

	logger.Info("Reaper successfully reconciled")
	return ctrl.Result{}, nil
}

func (r *ReaperReconciler) reconcileDatacenter(
	ctx context.Context,
	actualReaper *reaperapi.Reaper,
	logger logr.Logger,
) (*cassdcapi.CassandraDatacenter, ctrl.Result, error) {
	dcNamespace := actualReaper.Spec.DatacenterRef.Namespace
	if dcNamespace == "" {
		dcNamespace = actualReaper.Namespace
	}
	dcKey := client.ObjectKey{Namespace: dcNamespace, Name: actualReaper.Spec.DatacenterRef.Name}
	logger = logger.WithValues("CassandraDatacenter", dcKey)
	logger.Info("Fetching CassandraDatacenter resource")
	actualDc := &cassdcapi.CassandraDatacenter{}
	if err := r.Get(ctx, dcKey, actualDc); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("Waiting for datacenter to be created")
			return nil, ctrl.Result{RequeueAfter: r.DefaultDelay}, nil
		} else {
			logger.Error(err, "Failed to fetch CassandraDatacenter")
			return nil, ctrl.Result{}, err
		}
	}
	actualDc = actualDc.DeepCopy()
	if !cassandra.DatacenterReady(actualDc) {
		logger.Info("Waiting for datacenter to become ready")
		return nil, ctrl.Result{RequeueAfter: r.DefaultDelay}, nil
	}
	return actualDc, ctrl.Result{}, nil
}

func (r *ReaperReconciler) reconcileDeployment(
	ctx context.Context,
	actualReaper *reaperapi.Reaper,
	actualDc *cassdcapi.CassandraDatacenter,
	logger logr.Logger,
) (ctrl.Result, error) {

	deploymentKey := types.NamespacedName{Namespace: actualReaper.Namespace, Name: actualReaper.Name}
	logger = logger.WithValues("Deployment", deploymentKey)
	logger.Info("Reconciling Reaper Deployment")

	authVars, err := r.collectAuthVars(ctx, actualReaper, logger)
	if err != nil {
		logger.Error(err, "Failed to collect Reaper auth variables")
		return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
	}

	desiredDeployment := reaper.NewDeployment(actualReaper, actualDc, authVars...)

	actualDeployment := &appsv1.Deployment{}
	if err := r.Get(ctx, deploymentKey, actualDeployment); err != nil {
		if errors.IsNotFound(err) {
			if err = controllerutil.SetControllerReference(actualReaper, desiredDeployment, r.Scheme); err != nil {
				logger.Error(err, "Failed to set owner on Reaper Deployment")
				return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
			} else if err = r.Create(ctx, desiredDeployment); err != nil {
				if errors.IsAlreadyExists(err) {
					// the read from the local cache didn't catch that the resource was created
					// already; simply requeue until the cache is up-to-date
					return ctrl.Result{Requeue: true}, nil
				} else {
					logger.Error(err, "Failed to create Reaper Deployment")
					return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
				}
			}
			logger.Info("Reaper Deployment created successfully")
			return ctrl.Result{RequeueAfter: r.DefaultDelay}, nil
		} else {
			logger.Error(err, "Failed to get Reaper Deployment")
			return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
		}
	}

	actualDeployment = actualDeployment.DeepCopy()

	// Check if the deployment needs to be updated
	if !annotations.CompareHashAnnotations(actualDeployment, desiredDeployment) {
		logger.Info("Updating Reaper Deployment")
		resourceVersion := actualDeployment.GetResourceVersion()
		desiredDeployment.DeepCopyInto(actualDeployment)
		actualDeployment.SetResourceVersion(resourceVersion)
		if err := controllerutil.SetControllerReference(actualReaper, actualDeployment, r.Scheme); err != nil {
			logger.Error(err, "Failed to set controller reference on updated Reaper Deployment")
			return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
		} else if err := r.Update(ctx, actualDeployment); err != nil {
			logger.Error(err, "Failed to update Reaper Deployment")
			return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
		} else {
			logger.Info("Reaper Deployment updated successfully")
			return ctrl.Result{RequeueAfter: r.DefaultDelay}, nil
		}
	}

	logger.Info("Reaper Deployment ready")
	return ctrl.Result{}, nil
}

func (r *ReaperReconciler) reconcileService(
	ctx context.Context,
	actualReaper *reaperapi.Reaper,
	logger logr.Logger,
) (ctrl.Result, error) {
	serviceKey := types.NamespacedName{Namespace: actualReaper.Namespace, Name: reaper.GetServiceName(actualReaper.Name)}
	logger = logger.WithValues("Service", serviceKey)
	logger.Info("Reconciling Reaper Service")
	desiredService := reaper.NewService(serviceKey, actualReaper)
	actualService := &corev1.Service{}
	if err := r.Client.Get(ctx, serviceKey, actualService); err != nil {
		if errors.IsNotFound(err) {
			// create the service
			if err = controllerutil.SetControllerReference(actualReaper, desiredService, r.Scheme); err != nil {
				logger.Error(err, "Failed to set controller reference on Reaper Service")
				return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
			}
			logger.Info("Creating Reaper service")
			if err = r.Client.Create(ctx, desiredService); err != nil {
				if errors.IsAlreadyExists(err) {
					// the read from the local cache didn't catch that the resource was created
					// already; simply requeue until the cache is up-to-date
					return ctrl.Result{Requeue: true}, nil
				} else {
					logger.Error(err, "Failed to create Reaper Service")
					return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
				}
			}
			logger.Info("Reaper Service created successfully")
			return ctrl.Result{}, nil
		} else {
			logger.Error(err, "Failed to get Reaper Service")
			return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
		}
	}
	if !annotations.CompareHashAnnotations(actualService, desiredService) {
		logger.Info("Updating Reaper Service")
		updatedService := actualService.DeepCopy()
		desiredService.DeepCopyInto(updatedService)
		updatedService.SetResourceVersion(actualService.GetResourceVersion())
		updatedService.Spec.ClusterIP = actualService.Spec.ClusterIP
		updatedService.Spec.ClusterIPs = actualService.Spec.ClusterIPs
		if err := controllerutil.SetControllerReference(actualReaper, updatedService, r.Scheme); err != nil {
			logger.Error(err, "Failed to set controller reference on updated Reaper Service")
			return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
		} else if err := r.Update(ctx, updatedService); err != nil {
			logger.Error(err, "Failed to update Reaper Service")
			return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
		} else {
			logger.Info("Reaper Service updated successfully")
			return ctrl.Result{}, nil
		}
	}
	logger.Info("Reaper Service is ready")
	return ctrl.Result{}, nil
}

func (r *ReaperReconciler) configureReaper(ctx context.Context, actualReaper *reaperapi.Reaper, actualDc *cassdcapi.CassandraDatacenter, logger logr.Logger) (ctrl.Result, error) {
	manager := r.NewManager()
	if err := manager.Connect(actualReaper); err != nil {
		logger.Error(err, "failed to connect to reaper instance")
		return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
	} else if found, err := manager.VerifyClusterIsConfigured(ctx, actualDc); err != nil {
		logger.Error(err, "failed to verify the cluster is registered with reaper")
		return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
	} else if !found {
		logger.Info("registering cluster with reaper")
		if err = manager.AddClusterToReaper(ctx, actualDc); err != nil {
			logger.Error(err, "failed to register cluster with reaper")
			return ctrl.Result{RequeueAfter: r.DefaultDelay}, err
		}
	}
	return ctrl.Result{}, nil
}

func (r *ReaperReconciler) collectAuthVars(ctx context.Context, actualReaper *reaperapi.Reaper, logger logr.Logger) ([]*corev1.EnvVar, error) {
	cqlVars, err := r.collectCqlAuthVars(ctx, actualReaper, logger)
	if err != nil {
		return nil, err
	}
	jmxVars, err := r.collectJmxAuthVars(ctx, actualReaper, logger)
	if err != nil {
		return nil, err
	}
	return append(cqlVars, jmxVars...), nil
}

func (r *ReaperReconciler) collectCqlAuthVars(ctx context.Context, actualReaper *reaperapi.Reaper, logger logr.Logger) ([]*corev1.EnvVar, error) {
	if len(actualReaper.Spec.CassandraUserSecretRef) > 0 {
		secretKey := types.NamespacedName{Namespace: actualReaper.Namespace, Name: actualReaper.Spec.CassandraUserSecretRef}
		if secret, err := r.getSecret(ctx, secretKey); err != nil {
			logger.Error(err, "Failed to get Cassandra authentication secret", "CassandraUserSecretName", secretKey)
			return nil, err
		} else if usernameEnvVar, passwordEnvVar, err := reaper.GetCassandraAuthEnvironmentVars(secret); err != nil {
			logger.Error(err, "Failed to get Cassandra authentication env vars", "CassandraUserSecretName", secretKey)
			return nil, err
		} else {
			return []*corev1.EnvVar{usernameEnvVar, passwordEnvVar, reaper.EnableCassAuthVar}, nil
		}
	}
	return nil, nil
}

func (r *ReaperReconciler) collectJmxAuthVars(ctx context.Context, actualReaper *reaperapi.Reaper, logger logr.Logger) ([]*corev1.EnvVar, error) {
	if len(actualReaper.Spec.JmxUserSecretRef) > 0 {
		secretKey := types.NamespacedName{Namespace: actualReaper.Namespace, Name: actualReaper.Spec.JmxUserSecretRef}
		if secret, err := r.getSecret(ctx, secretKey); err != nil {
			logger.Error(err, "Failed to get JMX authentication secret", "JmxUserSecretName", secretKey)
			return nil, err
		} else if usernameEnvVar, passwordEnvVar, err := reaper.GetJmxAuthEnvironmentVars(secret); err != nil {
			logger.Error(err, "Failed to get JMX authentication env vars", "JmxUserSecretName", secretKey)
			return nil, err
		} else {
			return []*corev1.EnvVar{usernameEnvVar, passwordEnvVar}, nil
		}
	}
	return nil, nil
}

func (r *ReaperReconciler) getSecret(ctx context.Context, secretKey types.NamespacedName) (*corev1.Secret, error) {
	secret := &corev1.Secret{}
	err := r.Get(ctx, secretKey, secret)
	return secret, err
}

func (r *ReaperReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&reaperapi.Reaper{}, builder.WithPredicates(predicate.GenerationChangedPredicate{})).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Complete(r)
}
