// This file defines interfaces that determines an orchestrator w.r.t maya api
// server. All the features that maya api server wants from an orchestrator is
// defined in these set of interfaces.
package orchprovider

import (
	volProfile "github.com/openebs/maya/lib/profile/volumeprovisioner"
	"github.com/openebs/maya/lib/types/v1"
)

// OrchestrationInterface is an interface abstraction of a real orchestrator.
// It represents an abstraction that maya api server expects from its
// orchestrator.
//
// NOTE:
//  OrchestratorInterface should be the only interface that exposes orchestration
// contracts.
type OrchestratorInterface interface {
	// Label assigned against the orchestration provider
	Label() string

	// Name of the orchestration provider
	Name() string

	// Region where this orchestration provider is running/deployed
	Region() string

	// StorageOps gets the instance that deals with storage related operations.
	// Will return false if not supported.
	StorageOps() (StorageOps, bool)
}

// StorageOps exposes various storage related operations that deals with
// storage placements, scheduling, etc. The low level work is done by the
// orchestrator.
type StorageOps interface {

	// AddStorage will add persistent volume running as containers
	//
	// TODO
	//    Use VSM as the return type
	AddStorage(volProProfile volProfile.VolumeProvisionerProfile) (*v1.PersistentVolume, error)

	// DeleteStorage will remove the persistent volume
	//
	// TODO
	//    Use VSM as the return type
	DeleteStorage(volProProfile volProfile.VolumeProvisionerProfile) error

	// ReadStorage will fetch information about the persistent volume
	//
	// TODO
	//    Use VSM as the return type
	ReadStorage(volProProfile volProfile.VolumeProvisionerProfile) (*v1.PersistentVolume, error)

	// ListStorage will list a collection of VSMs in a given context e.g. namespace
	// if working in a K8s setup, etc.
	ListStorage(volProProfile volProfile.VolumeProvisionerProfile) (*v1.PersistentVolumeList, error)
}