/*
Copyright 2023 The Kubernetes Authors.

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

// Package heartbeat_with_lease contains the node heartbeat with lease for kwok.
package heartbeat_with_lease //nolint:revive

import (
	_ "embed"
)

var (
	// DefaultNodeHeartbeatWithLease is the default node heartbeat yaml.
	//go:embed node-heartbeat-with-lease.yaml
	DefaultNodeHeartbeatWithLease string
)
