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

// Package model contains the data structures used in GoSight.
package model

// Meta represents metadata about the host, agent, and environment.
// This is used to provide context for the data being collected and sent to the server.
// ------------------------------------
// Combined Meta (for metrics, logs, traces)
// ------------------------------------
type Meta struct {
    // --- Agent / Host / Endpoint context ---
    AgentID      string `json:"agent_id"`      // Unique ID for the agent
    AgentVersion string `json:"agent_version"` // Version of the agent

    HostID     string `json:"host_id"`     // Unique ID for the host
    EndpointID string `json:"endpoint_id"` // Unique ID for the endpoint
    ResourceID string `json:"resource_id"` // Unique ID for the resource
    Kind       string `json:"kind"`        // e.g. "host", "container", "service"

    Hostname           string `json:"hostname"`
    IPAddress          string `json:"ip_address"`
    OS                 string `json:"os,omitempty"`
    OSVersion          string `json:"os_version,omitempty"`
    Platform           string `json:"platform,omitempty"`
    PlatformFamily     string `json:"platform_family,omitempty"`
    PlatformVersion    string `json:"platform_version,omitempty"`
    KernelArchitecture string `json:"kernel_architecture,omitempty"`
    KernelVersion      string `json:"kernel_version,omitempty"`
    Architecture       string `json:"architecture,omitempty"`

    VirtualizationSystem string `json:"virtualization_system,omitempty"`
    VirtualizationRole   string `json:"virtualization_role,omitempty"`

    // --- Cloud / IaaS context ---
    CloudProvider    string `json:"cloud_provider,omitempty"`    // AWS, Azure, GCP
    Region           string `json:"region,omitempty"`
    AvailabilityZone string `json:"availability_zone,omitempty"` // or Zone
    InstanceID       string `json:"instance_id,omitempty"`
    InstanceType     string `json:"instance_type,omitempty"`
    AccountID        string `json:"account_id,omitempty"`
    ProjectID        string `json:"project_id,omitempty"`     // GCP
    ResourceGroup    string `json:"resource_group,omitempty"` // Azure
    VPCID            string `json:"vpc_id,omitempty"`         // AWS, GCP
    SubnetID         string `json:"subnet_id,omitempty"`      // AWS, GCP, Azure
    ImageID          string `json:"image_id,omitempty"`       // AMI, Image, etc.
    ServiceID        string `json:"service_id,omitempty"`     // if a managed service is the source

    // --- Kubernetes / Container context ---
    ContainerID        string            `json:"container_id,omitempty"`
    ContainerName      string            `json:"container_name,omitempty"`
    ContainerImageID   string            `json:"container_image_id,omitempty"`
    ContainerImageName string            `json:"container_image_name,omitempty"`
    PodName            string            `json:"pod_name,omitempty"`
    PodUID             string            `json:"pod_uid,omitempty"`
    PodLabels          map[string]string `json:"pod_labels,omitempty"`
    PodAnnotations     map[string]string `json:"pod_annotations,omitempty"`
    DeploymentName     string            `json:"deployment_name,omitempty"`
    OwnerKind          string            `json:"owner_kind,omitempty"` // e.g. "Deployment", "ReplicaSet"
    OwnerName          string            `json:"owner_name,omitempty"`
    Namespace          string            `json:"namespace,omitempty"`       
    NamespaceUID       string            `json:"namespace_uid,omitempty"`
    ClusterName        string            `json:"cluster_name,omitempty"`
    ClusterUID         string            `json:"cluster_uid,omitempty"`
    NodeName           string            `json:"node_name,omitempty"`
    NodeLabels         map[string]string `json:"node_labels,omitempty"`
    ServiceAccount     string            `json:"service_account,omitempty"`

    // --- Application / OTel Resource / Service context ---
    ServiceName               string `json:"service_name,omitempty"`               // OTel: service.name
    ServiceNamespace          string `json:"service_namespace,omitempty"`          // OTel: service.namespace
    ServiceInstanceID         string `json:"service_instance_id,omitempty"`        // OTel: service.instance.id
    ServiceVersion            string `json:"service_version,omitempty"`            // OTel: service.version
    TelemetrySDKName          string `json:"telemetry_sdk_name,omitempty"`         // OTel: telemetry.sdk.name
    TelemetrySDKVersion       string `json:"telemetry_sdk_version,omitempty"`      // OTel: telemetry.sdk.version
    TelemetrySDKLanguage      string `json:"telemetry_sdk_language,omitempty"`     // OTel: telemetry.sdk.language
    InstrumentationLibrary    string `json:"instrumentation_library,omitempty"`
    InstrumentationLibVersion string `json:"instrumentation_lib_version,omitempty"`
	
	Application  string `json:"application,omitempty"`
	Environment  string `json:"environment,omitempty"` // dev, staging, prod
	Service      string `json:"service,omitempty"`     // if a microservice
	Version      string `json:"version,omitempty"`


    // --- Process / Runtime context ---
    ProcessID      int    `json:"process_id,omitempty"`
    ProcessName    string `json:"process_name,omitempty"`
    RuntimeName    string `json:"runtime_name,omitempty"`    // e.g. "go", "java"
    RuntimeVersion string `json:"runtime_version,omitempty"` // e.g. "go1.20"

    // --- Networking / Mesh / Security context ---
    PublicIP         string `json:"public_ip,omitempty"`
    PrivateIP        string `json:"private_ip,omitempty"`
    MACAddress       string `json:"mac_address,omitempty"`
    NetworkInterface string `json:"network_interface,omitempty"`
    MeshPeerVersion  string `json:"mesh_peer_version,omitempty"`
    MTLSEnabled      bool   `json:"mtls_enabled,omitempty"`
    TLSVersion       string `json:"tls_version,omitempty"`
    CipherSuite      string `json:"cipher_suite,omitempty"`
    AuthMethod       string `json:"auth_method,omitempty"`          // e.g. "oauth2", "api_key", "mtls"
    User             string `json:"user,omitempty"`                 // OS‐level user or app‐level user
    JWTClaims        map[string]interface{} `json:"jwt_claims,omitempty"`

    // --- Deployment / CI-CD context ---
    DeploymentID   string `json:"deployment_id,omitempty"`
    GitCommitHash  string `json:"git_commit,omitempty"`
    BuildTimestamp string `json:"build_timestamp,omitempty"`

    // --- Log‐specific fields (formerly in LogMeta) ---
    AppName       string `json:"app_name,omitempty"`        // e.g. "nginx", "mysvc"
    AppVersion    string `json:"app_version,omitempty"`     // if known (e.g. "1.4.2")
    Unit          string `json:"unit,omitempty"`            // systemd unit name (for journald)
    EventID       string `json:"event_id,omitempty"`        // Windows Event ID or similar
    Executable    string `json:"exe,omitempty"`             // path to binary, if available
    Path          string `json:"path,omitempty"`            // original source log path (e.g. "/var/log/nginx/error.log")
    Extra         map[string]string `json:"extra,omitempty"` // any collector‐specific fields you previously put in LogMeta

    // --- Custom / User-defined tags & labels ---
    Labels map[string]string `json:"labels,omitempty"` // System-generated labels
    Tags   map[string]string `json:"tags,omitempty"`   // User-defined tags
}