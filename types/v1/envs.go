/*
Copyright 2017 The OpenEBS Authors.

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

package v1

import (
	"os"
	"strings"

	"github.com/openebs/maya/pkg/util"
)

// ENVKey is a typed variable that holds all environment
// variables
type ENVKey string

const (
	// VolumeTypeENVK is the ENV key to fetch the volume type
	VolumeTypeENVK ENVKey = "OPENEBS_IO_VOLUME_TYPE"

	// OrchProviderENVK is the ENV key to fetch volume's
	// orchestration provider
	OrchProviderENVK ENVKey = "OPENEBS_IO_ORCH_PROVIDER"

	// K8sStorageClassENVK is the ENV key to fetch volume's storage class
	// This is applicable only when K8s is volume's orchestration
	// provider
	K8sStorageClassENVK ENVKey = "K8S_IO_STORAGE_CLASS"

	// NamespaceENVK is the ENV key to fetch the
	// namespace where volume operations will be executed
	NamespaceENVK ENVKey = "OPENEBS_IO_NAMESPACE"

	// K8sOutClusterENVK is the ENV key to fetch outside
	// K8s cluster information. This K8s cluster is different
	// from the current K8s cluster where this binary will
	// run.
	K8sOutClusterENVK ENVKey = "K8S_IO_OUT_CLUSTER"

	// CapacityENVK is the ENV key to fetch volume's
	// capacity value
	CapacityENVK ENVKey = "OPENEBS_IO_CAPACITY"

	// JivaReplicasENVK is the ENV key to fetch jiva replica
	// count
	JivaReplicasENVK ENVKey = "OPENEBS_IO_JIVA_REPLICA_COUNT"

	// JivaControllersENVK is the ENV key to fetch jiva controller
	// count
	JivaControllersENVK ENVKey = "OPENEBS_IO_JIVA_CONTROLLER_COUNT"

	// JivaReplicaImageENVK is the ENV key to fetch jiva replica
	// image
	JivaReplicaImageENVK ENVKey = "OPENEBS_IO_JIVA_REPLICA_IMAGE"

	// JivaControllerImageENVK is the ENV key to fetch jiva controller
	// image
	JivaControllerImageENVK ENVKey = "OPENEBS_IO_JIVA_CONTROLLER_IMAGE"
)

// VolumeTypeENV will fetch the value of volume type
// from ENV variable if present
func VolumeTypeENV() VolumeType {
	val := getEnv(VolumeTypeENVK)
	return VolumeType(val)
}

// OrchProviderENV will fetch the value of volume's orchestrator
// from ENV variable if present
func OrchProviderENV() OrchProvider {
	val := getEnv(OrchProviderENVK)
	return OrchProvider(val)
}

func K8sStorageClassENV() string {
	val := getEnv(K8sStorageClassENVK)
	return val
}

func NamespaceENV() string {
	val := getEnv(NamespaceENVK)
	return val
}

func K8sOutClusterENV() string {
	val := getEnv(K8sOutClusterENVK)
	return val
}

func CapacityENV() string {
	val := getEnv(CapacityENVK)
	return val
}

func JivaReplicasENV() *int32 {
	val := util.StrToInt32(getEnv(JivaReplicasENVK))
	return val
}

func JivaReplicaImageENV() string {
	val := getEnv(JivaReplicaImageENVK)
	return val
}

func JivaControllersENV() *int32 {
	val := util.StrToInt32(getEnv(JivaControllersENVK))
	return val
}

func JivaControllerImageENV() string {
	val := getEnv(JivaControllerImageENVK)
	return val
}

// getEnv fetches the environment variable value from the machine's
// environment
func getEnv(envKey ENVKey) string {
	return strings.TrimSpace(os.Getenv(string(envKey)))
}
