/*
Copyright AppsCode Inc. and Contributors

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

package phase

import (
	esapi "kubedb.dev/apimachinery/apis/elasticsearch/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1"
	olddbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

	kmapi "kmodules.xyz/client-go/api/v1"
	cutil "kmodules.xyz/client-go/conditions"
)

func DashboardPhaseFromCondition(conditions []kmapi.Condition) esapi.DashboardPhase {
	if !cutil.IsConditionTrue(conditions, string(esapi.DashboardConditionProvisioned)) {
		return esapi.DashboardPhaseProvisioning
	}

	if !cutil.IsConditionTrue(conditions, string(esapi.DashboardConditionAcceptingConnection)) {
		return esapi.DashboardPhaseNotReady
	}

	// TODO: implement deployment watcher to handle replica ready

	if cutil.HasCondition(conditions, string(esapi.DashboardConditionServerHealthy)) {

		if !cutil.IsConditionTrue(conditions, string(esapi.DashboardConditionServerHealthy)) {

			_, cond := cutil.GetCondition(conditions, string(esapi.DashboardConditionServerHealthy))

			if cond.Reason == esapi.DashboardStateRed {
				return esapi.DashboardPhaseNotReady
			} else {
				return esapi.DashboardPhaseCritical
			}
		}

		return esapi.DashboardPhaseReady
	}

	return esapi.DashboardPhaseNotReady
}

func PhaseFromCondition(conditions []kmapi.Condition) olddbapi.DatabasePhase {
	// Generally, the conditions should maintain the following chronological order
	// For normal restore process:
	//   ProvisioningStarted --> ReplicaReady --> AcceptingConnection --> DataRestoreStarted --> DataRestored --> Ready --> Provisioned
	// For restoring the volumes (PerconaXtraDB):
	//	 ProvisioningStarted --> DataRestoreStarted --> DataRestored --> ReplicaReady --> AcceptingConnection --> Ready --> Provisioned

	// These are transitional conditions. They can update any time. So, their order may vary:
	// 1. ReplicaReady
	// 2. AcceptingConnection
	// 3. DataRestoreStarted
	// 4. DataRestored
	// 5. Ready
	// 6. Paused
	// 7. HealthCheckPaused

	var phase olddbapi.DatabasePhase

	// ================================= Handling "HealthCheckPaused" condition ==========================
	// If the condition is present and its "true", then the phase should be "Unknown".
	// Skip if the database isn't provisioned yet.
	if cutil.IsConditionTrue(conditions, kubedb.DatabaseHealthCheckPaused) {
		return olddbapi.DatabasePhaseUnknown
	}

	// ==================================  Handling "ProvisioningStarted" condition  ========================
	// If the condition is present and its "true", then the phase should be "Provisioning".
	if cutil.IsConditionTrue(conditions, kubedb.DatabaseProvisioningStarted) {
		phase = olddbapi.DatabasePhaseProvisioning
	}

	// ================================== Handling "Halted" condition =======================================
	// The "Halted" condition has higher priority, that's why it is placed at the top.
	// If the condition is present and its "true", then the phase should be "Halted".
	if cutil.IsConditionTrue(conditions, kubedb.DatabaseHalted) {
		return olddbapi.DatabasePhaseHalted
	}

	// =================================== Handling "DataRestoreStarted" and "DataRestored" conditions  ==================================================
	// For data restoring, there could be the following scenarios:
	// 1. if condition["DataRestoreStarted"] = true, the phase should be "Restoring".
	//		And there will be no "false" status for "DataRestoreStarted" type.
	// 2. if condition["DataRestored"] = false, the phase should be "NotReady".
	//		if the status is "true", the phase should depend on the rest of checks.
	if cutil.IsConditionTrue(conditions, kubedb.DatabaseDataRestoreStarted) {
		// TODO:
		// 		- remove these conditions.
		//		- It is here for backward compatibility.
		//		- Just return "Restoring" in future.
		if cutil.HasCondition(conditions, kubedb.DatabaseDataRestored) {
			if cutil.IsConditionFalse(conditions, kubedb.DatabaseDataRestored) {
				return olddbapi.DatabasePhaseNotReady
			}
		} else {
			return olddbapi.DatabasePhaseDataRestoring
		}
	}
	if cutil.HasCondition(conditions, kubedb.DatabaseDataRestored) && cutil.IsConditionFalse(conditions, kubedb.DatabaseDataRestored) {
		return olddbapi.DatabasePhaseNotReady
	}

	// ================================= Handling "AcceptingConnection" condition ==========================
	// If the condition is present and its "false", then the phase should be "NotReady".
	// Skip if the database isn't provisioned yet.
	if cutil.IsConditionFalse(conditions, kubedb.DatabaseAcceptingConnection) && cutil.IsConditionTrue(conditions, kubedb.DatabaseProvisioned) {
		return olddbapi.DatabasePhaseNotReady
	}

	// ================================= Handling "ReplicaReady" condition ==========================
	// If the condition is present and its "false", then the phase should be "Critical".
	// Skip if the database isn't provisioned yet.
	if cutil.IsConditionFalse(conditions, kubedb.DatabaseReplicaReady) && cutil.IsConditionTrue(conditions, kubedb.DatabaseProvisioned) {
		return olddbapi.DatabasePhaseCritical
	}

	// ================================= Handling "Ready" condition ==========================
	// Skip if the database isn't provisioned yet.
	if cutil.IsConditionFalse(conditions, kubedb.DatabaseReady) && cutil.IsConditionTrue(conditions, kubedb.DatabaseProvisioned) {
		return olddbapi.DatabasePhaseCritical
	}
	// Ready, if the database is provisioned and readinessProbe passed.
	if cutil.IsConditionTrue(conditions, kubedb.DatabaseReady) && cutil.IsConditionTrue(conditions, kubedb.DatabaseProvisioned) {
		return olddbapi.DatabasePhaseReady
	}

	// ================================= Handling "Provisioned" and "Paused" conditions ==========================
	// These conditions does not have any effect on the database phase. They are only for internal usage.
	// So, we don't have to do anything for them.
	return phase
}

func PhaseFromConditionV1(conditions []kmapi.Condition) dbapi.DatabasePhase {
	// Generally, the conditions should maintain the following chronological order
	// For normal restore process:
	//   ProvisioningStarted --> ReplicaReady --> AcceptingConnection --> DataRestoreStarted --> DataRestored --> Ready --> Provisioned
	// For restoring the volumes (PerconaXtraDB):
	//	 ProvisioningStarted --> DataRestoreStarted --> DataRestored --> ReplicaReady --> AcceptingConnection --> Ready --> Provisioned

	// These are transitional conditions. They can update any time. So, their order may vary:
	// 1. ReplicaReady
	// 2. AcceptingConnection
	// 3. DataRestoreStarted
	// 4. DataRestored
	// 5. Ready
	// 6. Paused
	// 7. HealthCheckPaused

	var phase dbapi.DatabasePhase

	// ================================= Handling "HealthCheckPaused" condition ==========================
	// If the condition is present and its "true", then the phase should be "Unknown".
	// Skip if the database isn't provisioned yet.
	if cutil.IsConditionTrue(conditions, kubedb.DatabaseHealthCheckPaused) {
		return dbapi.DatabasePhaseUnknown
	}

	// ==================================  Handling "ProvisioningStarted" condition  ========================
	// If the condition is present and its "true", then the phase should be "Provisioning".
	if cutil.IsConditionTrue(conditions, kubedb.DatabaseProvisioningStarted) {
		phase = dbapi.DatabasePhaseProvisioning
	}

	// ================================== Handling "Halted" condition =======================================
	// The "Halted" condition has higher priority, that's why it is placed at the top.
	// If the condition is present and its "true", then the phase should be "Halted".
	if cutil.IsConditionTrue(conditions, kubedb.DatabaseHalted) {
		return dbapi.DatabasePhaseHalted
	}

	// =================================== Handling "DataRestoreStarted" and "DataRestored" conditions  ==================================================
	// For data restoring, there could be the following scenarios:
	// 1. if condition["DataRestoreStarted"] = true, the phase should be "Restoring".
	//		And there will be no "false" status for "DataRestoreStarted" type.
	// 2. if condition["DataRestored"] = false, the phase should be "NotReady".
	//		if the status is "true", the phase should depend on the rest of checks.
	if cutil.IsConditionTrue(conditions, kubedb.DatabaseDataRestoreStarted) {
		// TODO:
		// 		- remove these conditions.
		//		- It is here for backward compatibility.
		//		- Just return "Restoring" in future.
		if cutil.HasCondition(conditions, kubedb.DatabaseDataRestored) {
			if cutil.IsConditionFalse(conditions, kubedb.DatabaseDataRestored) {
				return dbapi.DatabasePhaseNotReady
			}
		} else {
			return dbapi.DatabasePhaseDataRestoring
		}
	}
	if cutil.HasCondition(conditions, kubedb.DatabaseDataRestored) && cutil.IsConditionFalse(conditions, kubedb.DatabaseDataRestored) {
		return dbapi.DatabasePhaseNotReady
	}

	// ================================= Handling "AcceptingConnection" condition ==========================
	// If the condition is present and its "false", then the phase should be "NotReady".
	// Skip if the database isn't provisioned yet.
	if cutil.IsConditionFalse(conditions, kubedb.DatabaseAcceptingConnection) && cutil.IsConditionTrue(conditions, kubedb.DatabaseProvisioned) {
		return dbapi.DatabasePhaseNotReady
	}

	// ================================= Handling "ReplicaReady" condition ==========================
	// If the condition is present and its "false", then the phase should be "Critical".
	// Skip if the database isn't provisioned yet.
	if cutil.IsConditionFalse(conditions, kubedb.DatabaseReplicaReady) && cutil.IsConditionTrue(conditions, kubedb.DatabaseProvisioned) {
		return dbapi.DatabasePhaseCritical
	}

	// ================================= Handling "Ready" condition ==========================
	// Skip if the database isn't provisioned yet.
	if cutil.IsConditionFalse(conditions, kubedb.DatabaseReady) && cutil.IsConditionTrue(conditions, kubedb.DatabaseProvisioned) {
		return dbapi.DatabasePhaseCritical
	}
	// Ready, if the database is provisioned and readinessProbe passed.
	if cutil.IsConditionTrue(conditions, kubedb.DatabaseReady) && cutil.IsConditionTrue(conditions, kubedb.DatabaseProvisioned) {
		return dbapi.DatabasePhaseReady
	}

	// ================================= Handling "Provisioned" and "Paused" conditions ==========================
	// These conditions does not have any effect on the database phase. They are only for internal usage.
	// So, we don't have to do anything for them.
	return phase
}

// compareLastTransactionTime compare two condition's "LastTransactionTime" and return an integer based on the followings:
// 1. If both conditions does not exist, then return 0
// 2. If cond1 exist but cond2 does not, then return 1
// 3. If cond1 does not exist but cond2 exist, then return -1
// 3. If cond1.LastTransactionTime > cond2.LastTransactionTime, then return 1
// 4. If cond1.LastTransactionTime = cond2.LastTransactionTime, then return 0
// 5. If cond1.LastTransactionTime < cond2.LastTransactionTime, then return -1
func compareLastTransactionTime(conditions []kmapi.Condition, type1, type2 string) int32 {
	idx1, cond1 := cutil.GetCondition(conditions, type1)
	idx2, cond2 := cutil.GetCondition(conditions, type2)
	// both condition does not exist
	if idx1 == -1 && idx2 == -1 {
		return 0
	}
	// cond1 exist but cond2 does not
	if idx1 != -1 && idx2 == -1 {
		return 1
	}
	// cond2 does not exist but cond2 exist
	if idx1 == -1 && idx2 != -1 {
		return -1
	}

	if cond1.LastTransitionTime.After(cond2.LastTransitionTime.Time) {
		// cond1 is newer than cond2
		return 1
	} else if cond2.LastTransitionTime.After(cond1.LastTransitionTime.Time) {
		// cond1 is older than cond2
		return -1
	}
	return 0
}
