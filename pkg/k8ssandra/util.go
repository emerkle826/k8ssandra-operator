package k8ssandra

import (
	api "github.com/k8ssandra/k8ssandra-operator/apis/k8ssandra/v1alpha1"
)

func GetDatacenterForDecommission(kc *api.K8ssandraCluster) *api.CassandraDatacenterTemplate {
	dcTemplates := make([]api.CassandraDatacenterTemplate, 0)
	for _, dc := range kc.Spec.Cassandra.Datacenters {
		dcTemplates = append(dcTemplates, dc)
	}

	// First look for a status that already has started decommission
	for dcName, status := range kc.Status.Datacenters {
		if idx := cassandraDatacenterTemplatesContain(dcTemplates, dcName); idx >= 0 {
			if status.DecommissionProgress != api.DecommNone {
				return &dcTemplates[idx]
			}
		}
	}

	// No decommissions are in progress. Pick the first one we find.
	for dcName := range kc.Status.Datacenters {
		if idx := cassandraDatacenterTemplatesContain(dcTemplates, dcName); idx >= 0 {
			return &dcTemplates[idx]
		}
	}

	return nil
}

// cassandraDatacenterTemplatesContain returns true the slice of Datacenter templates t contains a template with
// the name dc.
func cassandraDatacenterTemplatesContain(t []api.CassandraDatacenterTemplate, dc string) int {
	for i, n := range t {
		if n.Meta.Name == dc {
			return i
		}
	}
	return -1
}
