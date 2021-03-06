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
	"github.com/openebs/maya/pkg/util"
)

// These are a set of defaults
const (
	// DefaultVolumeType contains the default volume type
	DefaultVolumeType VolumeType = JivaVolumeType

	// DefaultOrchProvider contains the default orchestrator
	DefaultOrchProvider OrchProvider = K8sOrchProvider

	// DefaultNamespace contains the default namespace where
	// volume operations will be executed
	DefaultNamespace string = "default"

	// DefaultCapacity contains the default volume capacity
	DefaultCapacity string = "5G"

	// DefaultJivaControllerImage contains the default jiva controller
	// image
	DefaultJivaControllerImage string = "openebs/jiva:0.4.0"

	// DefaultJivaReplicaImage contains the default jiva replica image
	DefaultJivaReplicaImage string = "openebs/jiva:0.4.0"
)

var (
	// DefaultJivaReplicas contains the default jiva replica count
	DefaultJivaReplicas *int32 = util.StrToInt32("2")

	// DefaultJivaControllers contains the default jiva controller
	// count
	DefaultJivaControllers *int32 = util.StrToInt32("1")
)
