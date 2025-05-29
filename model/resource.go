/*
SPDX-License-Identifier: GPL-3.0-or-later

Copyright (C) 2025 Aaron Mathis aaron.mathis@gmail.com

This file is part of GoSight.

GoSight is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GoSight is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with GoSight. If not, see https://www.gnu.org/licenses/.
*/

// Package model provides core data structures for the GoSight observability platform.
// This package defines the resource registry schema and related types used for
// tracking infrastructure components, applications, and telemetry sources.
package model

import (
	"time"
)

// Resource represents a trackable entity in the GoSight observability platform.
// Resources form a hierarchical structure representing the relationship between
// agents, hosts, containers, applications, and external telemetry sources.
//
// The resource hierarchy follows this pattern:
//   - Agent (root) → Host → Container → App
//   - Agent (root) → Host → App
//   - Syslog (standalone external sources)
//
// Resources are identified by unique IDs generated from their kind and
// identifying characteristics (labels). This ensures consistent identification
// across the platform regardless of discovery order.
//
// Labels vs Tags vs Annotations:
//   - Labels: System-generated, immutable identifiers used for joins and scoping
//   - Tags: User-defined, mutable metadata for filtering and organization
//   - Annotations: Extended metadata for UI and operational purposes
type Resource struct {
	// ID is the unique identifier for this resource, generated from kind and labels.
	// Format: <kind>-<short-sha1-of-identifying-labels>
	// Example: "host-a1b2c3d4", "container-e5f6g7h8"
	ID string `json:"id" db:"id"`

	// Kind specifies the type of resource (agent, host, container, app, syslog).
	// This determines the resource's position in the hierarchy and its capabilities.
	Kind string `json:"kind" db:"kind"`

	// Name is the primary identifier for the resource, typically derived from
	// system metadata like hostname, container name, or service name.
	Name string `json:"name" db:"name"`

	// DisplayName is a human-readable name for UI display purposes.
	// May be the same as Name or a more descriptive variant.
	DisplayName string `json:"display_name" db:"display_name"`

	// Group provides logical grouping of resources, typically based on
	// environment, team, or deployment context.
	Group string `json:"group" db:"group_name"`

	// ParentID references the parent resource in the hierarchy.
	// Empty for root resources (agents and standalone syslog sources).
	ParentID string `json:"parent_id" db:"parent_id"`

	// Labels contains system-generated, immutable key-value pairs used for
	// resource identification and telemetry correlation. These are set by
	// agents and discovery processes and should not be modified by users.
	//
	// Common labels include:
	//   - agent_id: Unique agent identifier
	//   - host_id: Unique host identifier
	//   - container_id: Container runtime identifier
	//   - service: Service or application name
	Labels map[string]string `json:"labels"`

	// Tags contains user-defined, mutable key-value pairs for organization,
	// filtering, and operational metadata. These can be modified through
	// the API and UI for resource management.
	//
	// Common tags include:
	//   - env: Environment (prod, staging, dev)
	//   - team: Owning team
	//   - critical: Criticality level
	//   - version: Application version
	Tags map[string]string `json:"tags"`

	// Status indicates the current operational state of the resource.
	// This is independent of telemetry heartbeat and represents the
	// actual condition of the resource.
	Status string `json:"status" db:"status"`

	// LastSeen is the timestamp of the most recent telemetry or heartbeat
	// from this resource. Used for staleness detection and health monitoring.
	LastSeen time.Time `json:"last_seen" db:"last_seen"`

	// FirstSeen is the timestamp when this resource was first discovered
	// and registered in the system. Immutable after creation.
	FirstSeen time.Time `json:"first_seen" db:"first_seen"`

	// CreatedAt is the timestamp when this resource record was created
	// in the database. Used for auditing and lifecycle tracking.
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// UpdatedAt is the timestamp of the most recent update to this
	// resource record. Updated on any field modification.
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Location specifies the physical or logical location of the resource.
	// Examples: datacenter, region, availability zone, rack.
	Location string `json:"location" db:"location"`

	// Environment indicates the deployment environment context.
	// Examples: production, staging, development, test.
	Environment string `json:"environment" db:"environment"`

	// Owner identifies the team, user, or service responsible for this resource.
	// Used for accountability and access control.
	Owner string `json:"owner" db:"owner"`

	// Platform specifies the underlying platform or cloud provider.
	// Examples: aws, gcp, azure, kubernetes, bare-metal.
	Platform string `json:"platform" db:"platform"`

	// Runtime indicates the container runtime or execution environment.
	// Examples: docker, containerd, cri-o, runc.
	Runtime string `json:"runtime" db:"runtime"`

	// Version contains version information for the resource, such as
	// agent version, OS version, or application version.
	Version string `json:"version" db:"version"`

	// OS specifies the operating system running on the resource.
	// Examples: linux, windows, darwin.
	OS string `json:"os" db:"os"`

	// Arch indicates the processor architecture of the resource.
	// Examples: amd64, arm64, x86.
	Arch string `json:"arch" db:"arch"`

	// IPAddress contains the primary IP address associated with the resource.
	// Used for network-based correlation and connectivity information.
	IPAddress string `json:"ip_address" db:"ip_address"`

	// ResourceType provides additional classification within a kind.
	// Examples: "vm", "bare-metal", "serverless" for hosts.
	ResourceType string `json:"resource_type" db:"resource_type"`

	// Cluster identifies the cluster or orchestration group this resource
	// belongs to. Commonly used in Kubernetes and container environments.
	Cluster string `json:"cluster" db:"cluster"`

	// Namespace provides logical isolation within a cluster or platform.
	// Examples: Kubernetes namespace, Docker compose project.
	Namespace string `json:"namespace" db:"namespace"`

	// Annotations contains extended metadata for UI display, automation,
	// and integration purposes. Unlike tags, these are typically managed
	// by the system and external tools rather than users.
	Annotations map[string]string `json:"annotations"`

	// Updated is a transient flag indicating whether this resource has
	// been modified and needs to be persisted. Not stored in database.
	Updated bool `json:"-" db:"-"`
}

// ResourceFilter defines criteria for filtering resources in list and search operations.
// Multiple filter criteria are combined with AND logic, while multiple values
// within a single criterion use OR logic.
//
// Example usage:
//
//	filter := &ResourceFilter{
//	    Kinds: []string{"host", "container"},
//	    Environment: []string{"production"},
//	    Labels: map[string]string{"team": "platform"},
//	}
type ResourceFilter struct {
	// Kinds filters resources by their type (agent, host, container, app, syslog).
	// If empty, all kinds are included.
	Kinds []string `json:"kinds"`

	// Groups filters resources by their logical grouping.
	// If empty, all groups are included.
	Groups []string `json:"groups"`

	// Status filters resources by their operational state.
	// If empty, all statuses are included.
	Status []string `json:"status"`

	// Labels filters resources that have all specified label key-value pairs.
	// All specified labels must match (AND logic).
	Labels map[string]string `json:"labels"`

	// Tags filters resources that have all specified tag key-value pairs.
	// All specified tags must match (AND logic).
	Tags map[string]string `json:"tags"`

	// Environment filters resources by their deployment environment.
	// If empty, all environments are included.
	Environment []string `json:"environment"`

	// Owner filters resources by their assigned owner.
	// If empty, all owners are included.
	Owner []string `json:"owner"`

	// LastSeenSince filters resources that have been seen since the specified time.
	// Used for staleness detection and active resource queries.
	LastSeenSince *time.Time `json:"last_seen_since"`
}

// ResourceSearchQuery defines parameters for full-text search across resources.
// Combines text search with filtering capabilities for comprehensive resource discovery.
//
// The search query supports various formats:
//   - Simple text: "web-server"
//   - Label queries: "env:production"
//   - Complex expressions: "web AND (prod OR staging)"
type ResourceSearchQuery struct {
	// Query is the search text to match against resource names, labels, and tags.
	// Supports full-text search syntax and boolean operators.
	Query string `json:"query"`

	// Kinds restricts search to specific resource types.
	// If empty, all kinds are searched.
	Kinds []string `json:"kinds"`

	// Groups restricts search to specific resource groups.
	// If empty, all groups are searched.
	Groups []string `json:"groups"`

	// Status restricts search to resources with specific operational states.
	// If empty, all statuses are searched.
	Status []string `json:"status"`

	// Labels restricts search to resources with specific labels.
	// Applied as additional filters on top of the text query.
	Labels map[string]string `json:"labels"`

	// Tags restricts search to resources with specific tags.
	// Applied as additional filters on top of the text query.
	Tags map[string]string `json:"tags"`

	// Limit specifies the maximum number of results to return.
	// Default and maximum values are enforced by the implementation.
	Limit int `json:"limit"`

	// Offset specifies the number of results to skip for pagination.
	// Used in combination with Limit for result paging.
	Offset int `json:"offset"`
}

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
