//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/k8ssandra/cass-operator/apis/cassandra/v1beta1"
	medusav1alpha1 "github.com/k8ssandra/k8ssandra-operator/apis/medusa/v1alpha1"
	reaperv1alpha1 "github.com/k8ssandra/k8ssandra-operator/apis/reaper/v1alpha1"
	stargatev1alpha1 "github.com/k8ssandra/k8ssandra-operator/apis/stargate/v1alpha1"
	telemetryv1alpha1 "github.com/k8ssandra/k8ssandra-operator/apis/telemetry/v1alpha1"
	"github.com/k8ssandra/k8ssandra-operator/pkg/images"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterTemplate) DeepCopyInto(out *CassandraClusterTemplate) {
	*out = *in
	out.SuperuserSecretRef = in.SuperuserSecretRef
	if in.JmxInitContainerImage != nil {
		in, out := &in.JmxInitContainerImage, &out.JmxInitContainerImage
		*out = new(images.Image)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemLoggerResources != nil {
		in, out := &in.SystemLoggerResources, &out.SystemLoggerResources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.CassandraConfig != nil {
		in, out := &in.CassandraConfig, &out.CassandraConfig
		*out = new(CassandraConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageConfig != nil {
		in, out := &in.StorageConfig, &out.StorageConfig
		*out = new(v1beta1.StorageConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Networking != nil {
		in, out := &in.Networking, &out.Networking
		*out = new(v1beta1.NetworkingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Racks != nil {
		in, out := &in.Racks, &out.Racks
		*out = make([]v1beta1.Rack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Datacenters != nil {
		in, out := &in.Datacenters, &out.Datacenters
		*out = make([]CassandraDatacenterTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CassandraTelemetry != nil {
		in, out := &in.CassandraTelemetry, &out.CassandraTelemetry
		*out = new(telemetryv1alpha1.TelemetrySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.MgmtAPIHeap != nil {
		in, out := &in.MgmtAPIHeap, &out.MgmtAPIHeap
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterTemplate.
func (in *CassandraClusterTemplate) DeepCopy() *CassandraClusterTemplate {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraConfig) DeepCopyInto(out *CassandraConfig) {
	*out = *in
	in.CassandraYaml.DeepCopyInto(&out.CassandraYaml)
	in.JvmOptions.DeepCopyInto(&out.JvmOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraConfig.
func (in *CassandraConfig) DeepCopy() *CassandraConfig {
	if in == nil {
		return nil
	}
	out := new(CassandraConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDatacenterTemplate) DeepCopyInto(out *CassandraDatacenterTemplate) {
	*out = *in
	in.Meta.DeepCopyInto(&out.Meta)
	if in.JmxInitContainerImage != nil {
		in, out := &in.JmxInitContainerImage, &out.JmxInitContainerImage
		*out = new(images.Image)
		(*in).DeepCopyInto(*out)
	}
	if in.CassandraConfig != nil {
		in, out := &in.CassandraConfig, &out.CassandraConfig
		*out = new(CassandraConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemLoggerResources != nil {
		in, out := &in.SystemLoggerResources, &out.SystemLoggerResources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Racks != nil {
		in, out := &in.Racks, &out.Racks
		*out = make([]v1beta1.Rack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Networking != nil {
		in, out := &in.Networking, &out.Networking
		*out = new(v1beta1.NetworkingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageConfig != nil {
		in, out := &in.StorageConfig, &out.StorageConfig
		*out = new(v1beta1.StorageConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Stargate != nil {
		in, out := &in.Stargate, &out.Stargate
		*out = new(stargatev1alpha1.StargateDatacenterTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Reaper != nil {
		in, out := &in.Reaper, &out.Reaper
		*out = new(reaperv1alpha1.ReaperDatacenterTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.MgmtAPIHeap != nil {
		in, out := &in.MgmtAPIHeap, &out.MgmtAPIHeap
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.CassandraTelemetry != nil {
		in, out := &in.CassandraTelemetry, &out.CassandraTelemetry
		*out = new(telemetryv1alpha1.TelemetrySpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDatacenterTemplate.
func (in *CassandraDatacenterTemplate) DeepCopy() *CassandraDatacenterTemplate {
	if in == nil {
		return nil
	}
	out := new(CassandraDatacenterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraYaml) DeepCopyInto(out *CassandraYaml) {
	*out = *in
	if in.Authenticator != nil {
		in, out := &in.Authenticator, &out.Authenticator
		*out = new(string)
		**out = **in
	}
	if in.Authorizer != nil {
		in, out := &in.Authorizer, &out.Authorizer
		*out = new(string)
		**out = **in
	}
	if in.RoleManager != nil {
		in, out := &in.RoleManager, &out.RoleManager
		*out = new(string)
		**out = **in
	}
	if in.RolesValidityMillis != nil {
		in, out := &in.RolesValidityMillis, &out.RolesValidityMillis
		*out = new(int64)
		**out = **in
	}
	if in.RolesUpdateIntervalMillis != nil {
		in, out := &in.RolesUpdateIntervalMillis, &out.RolesUpdateIntervalMillis
		*out = new(int64)
		**out = **in
	}
	if in.PermissionsValidityMillis != nil {
		in, out := &in.PermissionsValidityMillis, &out.PermissionsValidityMillis
		*out = new(int64)
		**out = **in
	}
	if in.PermissionsUpdateIntervalMillis != nil {
		in, out := &in.PermissionsUpdateIntervalMillis, &out.PermissionsUpdateIntervalMillis
		*out = new(int64)
		**out = **in
	}
	if in.CredentialsValidityMillis != nil {
		in, out := &in.CredentialsValidityMillis, &out.CredentialsValidityMillis
		*out = new(int64)
		**out = **in
	}
	if in.CredentialsUpdateIntervalMillis != nil {
		in, out := &in.CredentialsUpdateIntervalMillis, &out.CredentialsUpdateIntervalMillis
		*out = new(int64)
		**out = **in
	}
	if in.NumTokens != nil {
		in, out := &in.NumTokens, &out.NumTokens
		*out = new(int)
		**out = **in
	}
	if in.AllocateTokensForLocalReplicationFactor != nil {
		in, out := &in.AllocateTokensForLocalReplicationFactor, &out.AllocateTokensForLocalReplicationFactor
		*out = new(int)
		**out = **in
	}
	if in.ConcurrentReads != nil {
		in, out := &in.ConcurrentReads, &out.ConcurrentReads
		*out = new(int)
		**out = **in
	}
	if in.ConcurrentWrites != nil {
		in, out := &in.ConcurrentWrites, &out.ConcurrentWrites
		*out = new(int)
		**out = **in
	}
	if in.ConcurrentCounterWrites != nil {
		in, out := &in.ConcurrentCounterWrites, &out.ConcurrentCounterWrites
		*out = new(int)
		**out = **in
	}
	if in.AutoSnapshot != nil {
		in, out := &in.AutoSnapshot, &out.AutoSnapshot
		*out = new(bool)
		**out = **in
	}
	if in.MemtableFlushWriters != nil {
		in, out := &in.MemtableFlushWriters, &out.MemtableFlushWriters
		*out = new(int)
		**out = **in
	}
	if in.CommitLogSegmentSizeMb != nil {
		in, out := &in.CommitLogSegmentSizeMb, &out.CommitLogSegmentSizeMb
		*out = new(int)
		**out = **in
	}
	if in.ConcurrentCompactors != nil {
		in, out := &in.ConcurrentCompactors, &out.ConcurrentCompactors
		*out = new(int)
		**out = **in
	}
	if in.CompactionThroughputMbPerSec != nil {
		in, out := &in.CompactionThroughputMbPerSec, &out.CompactionThroughputMbPerSec
		*out = new(int)
		**out = **in
	}
	if in.SstablePreemptiveOpenIntervalMb != nil {
		in, out := &in.SstablePreemptiveOpenIntervalMb, &out.SstablePreemptiveOpenIntervalMb
		*out = new(int)
		**out = **in
	}
	if in.KeyCacheSizeMb != nil {
		in, out := &in.KeyCacheSizeMb, &out.KeyCacheSizeMb
		*out = new(int)
		**out = **in
	}
	if in.ThriftPreparedStatementCacheSizeMb != nil {
		in, out := &in.ThriftPreparedStatementCacheSizeMb, &out.ThriftPreparedStatementCacheSizeMb
		*out = new(int)
		**out = **in
	}
	if in.PreparedStatementsCacheSizeMb != nil {
		in, out := &in.PreparedStatementsCacheSizeMb, &out.PreparedStatementsCacheSizeMb
		*out = new(int)
		**out = **in
	}
	if in.StartRpc != nil {
		in, out := &in.StartRpc, &out.StartRpc
		*out = new(bool)
		**out = **in
	}
	if in.SlowQueryLogTimeoutMs != nil {
		in, out := &in.SlowQueryLogTimeoutMs, &out.SlowQueryLogTimeoutMs
		*out = new(int)
		**out = **in
	}
	if in.CounterCacheSizeMb != nil {
		in, out := &in.CounterCacheSizeMb, &out.CounterCacheSizeMb
		*out = new(int)
		**out = **in
	}
	if in.FileCacheSizeMb != nil {
		in, out := &in.FileCacheSizeMb, &out.FileCacheSizeMb
		*out = new(int)
		**out = **in
	}
	if in.RowCacheSizeMb != nil {
		in, out := &in.RowCacheSizeMb, &out.RowCacheSizeMb
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraYaml.
func (in *CassandraYaml) DeepCopy() *CassandraYaml {
	if in == nil {
		return nil
	}
	out := new(CassandraYaml)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedObjectMeta) DeepCopyInto(out *EmbeddedObjectMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedObjectMeta.
func (in *EmbeddedObjectMeta) DeepCopy() *EmbeddedObjectMeta {
	if in == nil {
		return nil
	}
	out := new(EmbeddedObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JvmOptions) DeepCopyInto(out *JvmOptions) {
	*out = *in
	if in.HeapSize != nil {
		in, out := &in.HeapSize, &out.HeapSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.HeapNewGenSize != nil {
		in, out := &in.HeapNewGenSize, &out.HeapNewGenSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.AdditionalOptions != nil {
		in, out := &in.AdditionalOptions, &out.AdditionalOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JvmOptions.
func (in *JvmOptions) DeepCopy() *JvmOptions {
	if in == nil {
		return nil
	}
	out := new(JvmOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8ssandraCluster) DeepCopyInto(out *K8ssandraCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8ssandraCluster.
func (in *K8ssandraCluster) DeepCopy() *K8ssandraCluster {
	if in == nil {
		return nil
	}
	out := new(K8ssandraCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K8ssandraCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8ssandraClusterCondition) DeepCopyInto(out *K8ssandraClusterCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8ssandraClusterCondition.
func (in *K8ssandraClusterCondition) DeepCopy() *K8ssandraClusterCondition {
	if in == nil {
		return nil
	}
	out := new(K8ssandraClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8ssandraClusterList) DeepCopyInto(out *K8ssandraClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]K8ssandraCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8ssandraClusterList.
func (in *K8ssandraClusterList) DeepCopy() *K8ssandraClusterList {
	if in == nil {
		return nil
	}
	out := new(K8ssandraClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K8ssandraClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8ssandraClusterSpec) DeepCopyInto(out *K8ssandraClusterSpec) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(bool)
		**out = **in
	}
	if in.Cassandra != nil {
		in, out := &in.Cassandra, &out.Cassandra
		*out = new(CassandraClusterTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Stargate != nil {
		in, out := &in.Stargate, &out.Stargate
		*out = new(stargatev1alpha1.StargateClusterTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Reaper != nil {
		in, out := &in.Reaper, &out.Reaper
		*out = new(reaperv1alpha1.ReaperClusterTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Medusa != nil {
		in, out := &in.Medusa, &out.Medusa
		*out = new(medusav1alpha1.MedusaClusterTemplate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8ssandraClusterSpec.
func (in *K8ssandraClusterSpec) DeepCopy() *K8ssandraClusterSpec {
	if in == nil {
		return nil
	}
	out := new(K8ssandraClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8ssandraClusterStatus) DeepCopyInto(out *K8ssandraClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]K8ssandraClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Datacenters != nil {
		in, out := &in.Datacenters, &out.Datacenters
		*out = make(map[string]K8ssandraStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8ssandraClusterStatus.
func (in *K8ssandraClusterStatus) DeepCopy() *K8ssandraClusterStatus {
	if in == nil {
		return nil
	}
	out := new(K8ssandraClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8ssandraStatus) DeepCopyInto(out *K8ssandraStatus) {
	*out = *in
	if in.Cassandra != nil {
		in, out := &in.Cassandra, &out.Cassandra
		*out = new(v1beta1.CassandraDatacenterStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Stargate != nil {
		in, out := &in.Stargate, &out.Stargate
		*out = new(stargatev1alpha1.StargateStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Reaper != nil {
		in, out := &in.Reaper, &out.Reaper
		*out = new(reaperv1alpha1.ReaperStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8ssandraStatus.
func (in *K8ssandraStatus) DeepCopy() *K8ssandraStatus {
	if in == nil {
		return nil
	}
	out := new(K8ssandraStatus)
	in.DeepCopyInto(out)
	return out
}
