/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ClusterState represents the values of the 'cluster_state' enumerated type.
type ClusterState string

const (
	// Error during installation.
	ClusterStateError ClusterState = "error"
	// The cluster is still being installed.
	ClusterStateInstalling ClusterState = "installing"
	// The cluster is waiting to be provisioned.
	ClusterStatePending ClusterState = "pending"
	// Creation of the cluster is waiting for the creation of an account in the cloud provider.
	ClusterStatePendingAccount ClusterState = "pending_account"
	// The cluster is ready to use.
	ClusterStateReady ClusterState = "ready"
	// The cluster is being uninstalled.
	ClusterStateUninstalling ClusterState = "uninstalling"
	// The state of the cluster is unknown.
	ClusterStateUnknown ClusterState = "unknown"
)
