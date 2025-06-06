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

import "time"

// Exemplar holds a single exemplar from a metric DataPoint,
// including trace/span IDs for correlation.
type Exemplar struct {
    Value              float64           `json:"value"`                // exemplar value
    Timestamp          time.Time         `json:"timestamp"`            // when exemplar was recorded
    TraceID            string            `json:"trace_id,omitempty"`   // 16-byte hex
    SpanID             string            `json:"span_id,omitempty"`    // 8-byte hex
    FilteredAttributes map[string]string `json:"filtered_attributes,omitempty"`
}

// DataPoint represents one data point of a metric. Depending on DataType,
// only one of Value, Histogram, or Summary will be non-nil.
type DataPoint struct {
    // --- Common fields ---
    Attributes     map[string]string `json:"attributes,omitempty"`     // OTLP attribute map
    StartTimestamp time.Time         `json:"start_timestamp,omitempty"`// for cumulative metrics (Sum/Counter/Histogram)
    Timestamp      time.Time         `json:"timestamp"`                // point timestamp

    // --- Depending on metric type: one of these is non-zero/non-nil ---
    // Gauge or Sum (counter)
    Value float64 `json:"value,omitempty"`

    // Histogram
    Count        uint64    `json:"count,omitempty"`
    Sum          float64   `json:"sum,omitempty"`
    BucketCounts []uint64  `json:"bucket_counts,omitempty"`
    ExplicitBounds []float64 `json:"explicit_bounds,omitempty"`

    // Summary (percentiles)
    QuantileValues []QuantileValue `json:"quantile_values,omitempty"`

    // Exemplars (trace/span IDs for sampled measurements)
    Exemplars []Exemplar `json:"exemplars,omitempty"`
}

// QuantileValue is used when DataType == Summary.
type QuantileValue struct {
    Quantile float64 `json:"quantile"` // e.g., 0.5 for p50
    Value    float64 `json:"value"`
}

// Metric records one OTLP Metric, including its type, unit, and data points.
type Metric struct {
    // Basic identity fields
    Namespace    string            `json:"namespace,omitempty"`     // optional “namespace” (e.g., user‐defined grouping)
    SubNamespace string            `json:"subnamespace,omitempty"`  // optional
    Name         string            `json:"name"`                    // metric name (OTLP: metric.name)
    Description  string            `json:"description,omitempty"`   // OTLP: metric description
    Unit         string            `json:"unit,omitempty"`          // OTLP: metric unit (e.g., "ms", "MiB")
	
	Source string `json:"source,omitempty"`  // gopsutils, prometheus, etc.
	
    // DataType indicates how to interpret the DataPoints slice.
    // Common values: "gauge", "sum", "histogram", "summary", etc.
    DataType string `json:"type,omitempty"`

    // For "sum" (cumulative/counter) metrics, indicates aggregation temporality:
    //   "cumulative" or "delta"
    AggregationTemporality string `json:"aggregation_temporality,omitempty"`

    // Each metric may carry zero or more DataPoints (usually at least one).
    DataPoints []DataPoint `json:"data_points,omitempty"`

    // StorageResolution is your existing retention/rollup hint.
    StorageResolution int `json:"resolution,omitempty"`
}



// StatisticValues represents the minimum, maximum, count, and sum of a metric.
type StatisticValues struct {
	Minimum     float64 `json:"min"`
	Maximum     float64 `json:"max"`
	SampleCount int     `json:"count"`
	Sum         float64 `json:"sum"`
}

type Point struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
}

// Metric represents a single metric data point.
// It includes the namespace, subnamespace, name, timestamp, value, and other data.
// type Metric struct {
// 	Namespace         string            `json:"namespace,omitempty"`
// 	SubNamespace      string            `json:"subnamespace,omitempty"`
// 	Name              string            `json:"name"`
// 	Timestamp         time.Time         `json:"timestamp"`
// 	Value             float64           `json:"value"`
// 	StatisticValues   *StatisticValues  `json:"stats"`
// 	Unit              string            `json:"unit,omitempty"`
// 	Dimensions        map[string]string `json:"dimensions,omitempty"`
// 	StorageResolution int               `json:"resolution,omitempty"`
// 	Type              string            `json:"type,omitempty"`
// }

// MetricPayload represents a collection of metrics to be sent to the server.
// It includes the agent ID, host ID, hostname, endpoint ID, and a list of metrics.
// The timestamp is the time at which the paylaod was packaged.
type MetricPayload struct {
	AgentID    string    `json:"agent_id"`
	HostID     string    `json:"host_id"`
	Hostname   string    `json:"hostname"`
	EndpointID string    `json:"endpoint_id"`
	Metrics    []Metric  `json:"metrics"`
	Meta       *Meta     `json:"meta,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
}

// MetricRow represents a single row of metric data.
type MetricRow struct {
	Value     float64           `json:"value"`
	Labels    map[string]string `json:"tags"`
	Timestamp int64             `json:"timestamp"` // Unix ms
}

// MetricPoint represents a single point in a metric time series.
type MetricPoint struct {
	Timestamp int64   `json:"timestamp"` // Unix ms
	Value     float64 `json:"value"`     // Metric value
}

// MetricSelector is used to select a specific metric for querying.
type MetricSelector struct {
	Name         string
	Namespace    string
	SubNamespace string
	Instant      bool // true = stat card, false = chart or long term
}
