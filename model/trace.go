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

import "time"

// TraceSpan represents a single span in a distributed trace.
// It contains metadata about the span, including its ID, parent span ID,
// name, service name, start and end times, duration, status, attributes,
// and any events associated with the span.
// It is designed to be compatible with OpenTelemetry's span format.
type TraceSpan struct {
	TraceID      string `json:"trace_id"`
	SpanID       string `json:"span_id"`
	ParentSpanID string `json:"parent_span_id,omitempty"`
	Name         string `json:"name"`
	ServiceName  string `json:"service_name,omitempty"`
	EndpointID   string `json:"endpoint_id,omitempty"`
	AgentID      string `json:"agent_id,omitempty"`
	HostID       string `json:"host_id,omitempty"`

	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	DurationMs float64   `json:"duration_ms"`

	StatusCode    string `json:"status_code,omitempty"` // "OK", "ERROR", etc.
	StatusMessage string `json:"status_message,omitempty"`

	Attributes map[string]string `json:"attributes,omitempty"` // OpenTelemetry key-value pairs

	Events        []SpanEvent       `json:"events,omitempty"`
	ResourceAttrs map[string]string `json:"resource_attrs,omitempty"` // service.name, host.name, etc.
}

// SpanEvent represents an event that occurred during the execution of a span.
type SpanEvent struct {
	Name       string            `json:"name"`
	Timestamp  time.Time         `json:"timestamp"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

// TracePayload represents a collection of trace spans sent from an agent
type TracePayload struct {
    Meta   *Meta       `json:"meta"`
    Traces []TraceSpan `json:"traces"`
}
