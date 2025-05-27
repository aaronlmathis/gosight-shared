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

package model

import (
	"time"
)

type Resource struct {
    ID          string            `json:"id" db:"id"`
    Kind        string            `json:"kind" db:"kind"`
    Name        string            `json:"name" db:"name"`
    DisplayName string            `json:"display_name" db:"display_name"`
    Group       string            `json:"group" db:"group_name"`
    ParentID    string            `json:"parent_id" db:"parent_id"`

    Labels      map[string]string `json:"labels"`
    Tags        map[string]string `json:"tags"`

    Status      string            `json:"status" db:"status"`
    LastSeen    time.Time         `json:"last_seen" db:"last_seen"`
    FirstSeen   time.Time         `json:"first_seen" db:"first_seen"`
    CreatedAt   time.Time         `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time         `json:"updated_at" db:"updated_at"`

    Location     string           `json:"location" db:"location"`
    Environment  string           `json:"environment" db:"environment"`
    Owner        string           `json:"owner" db:"owner"`
    Platform     string           `json:"platform" db:"platform"`
    Runtime      string           `json:"runtime" db:"runtime"`
    Version      string           `json:"version" db:"version"`
    OS           string           `json:"os" db:"os"`
    Arch         string           `json:"arch" db:"arch"`
    IPAddress    string           `json:"ip_address" db:"ip_address"`

    ResourceType string           `json:"resource_type" db:"resource_type"`
    Cluster      string           `json:"cluster" db:"cluster"`
    Namespace    string           `json:"namespace" db:"namespace"`
    Annotations  map[string]string `json:"annotations"`

    Updated     bool              `json:"-" db:"-"`
}

type ResourceFilter struct {
    Kinds       []string          `json:"kinds"`
    Groups      []string          `json:"groups"`
    Status      []string          `json:"status"`
    Labels      map[string]string `json:"labels"`
    Tags        map[string]string `json:"tags"`
    Environment []string          `json:"environment"`
    Owner       []string          `json:"owner"`
    LastSeenSince *time.Time      `json:"last_seen_since"`
}

const (
    ResourceKindHost      = "host"
    ResourceKindContainer = "container"
    ResourceKindApp       = "app"
    ResourceKindDevice    = "device"
    ResourceKindSyslog    = "syslog"
    ResourceKindOtel      = "otel"
    
    ResourceStatusOnline  = "online"
    ResourceStatusOffline = "offline"
    ResourceStatusIdle    = "idle"
    ResourceStatusUnknown = "unknown"
)