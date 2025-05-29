// Copyright 2024 GoSight Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package model defines core data structures and constants used throughout
// the GoSight observability platform.
package model

// Resource kinds define the types of resources that GoSight tracks and monitors.
// These constants represent the hierarchical relationship between different
// infrastructure components and telemetry sources.
//
// The resource hierarchy is:
//   - Agent (root): The GoSight agent process that collects telemetry
//   - Host: Physical or virtual machines where agents run
//   - Container: Containerized workloads (Docker, containerd, etc.)
//   - App: Application processes running on hosts or in containers
//   - Syslog: External syslog sources and network devices
const (
	// ResourceKindAgent represents a GoSight agent process.
	// Agents are the root of the resource hierarchy and are responsible
	// for collecting and forwarding telemetry data.
	ResourceKindAgent = "agent"

	// ResourceKindHost represents a physical or virtual machine.
	// Hosts are managed by agents and can contain multiple containers
	// and applications. Examples: servers, VMs, cloud instances.
	ResourceKindHost = "host"

	// ResourceKindContainer represents a containerized workload.
	// Containers run on hosts and are managed by container runtimes
	// like Docker, containerd, or CRI-O. They can host applications.
	ResourceKindContainer = "container"

	// ResourceKindApp represents an application or service process.
	// Applications can run directly on hosts or within containers.
	// Examples: web servers, databases, microservices.
	ResourceKindApp = "app"

	// ResourceKindSyslog represents external syslog sources.
	// These are typically network devices, appliances, or systems
	// that send syslog messages but don't run GoSight agents.
	// Examples: routers, switches, firewalls, IoT devices.
	ResourceKindSyslog = "syslog"
)

// Resource statuses define the operational state of tracked resources.
// These statuses are independent of telemetry heartbeat and represent
// the actual operational condition of the resource.
const (
	// ResourceStatusOnline indicates the resource is running and operational.
	// This is the default status for resources actively sending telemetry.
	ResourceStatusOnline = "online"

	// ResourceStatusOffline indicates the resource is not running or unreachable.
	// This may be set when a resource stops sending telemetry or when
	// explicit status information indicates the resource is down.
	ResourceStatusOffline = "offline"

	// ResourceStatusIdle indicates the resource is running but not actively processing.
	// This is commonly used for containers that are paused or applications
	// that are in a standby state.
	ResourceStatusIdle = "idle"

	// ResourceStatusUnknown indicates the resource status cannot be determined.
	// This is the initial status for newly discovered resources or when
	// status information is unavailable.
	ResourceStatusUnknown = "unknown"
)
