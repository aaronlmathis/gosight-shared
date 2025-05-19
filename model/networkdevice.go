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

// NetworkDevice represents a syslog-emitting network appliance
type NetworkDevice struct {
	ID        string    `db:"id"`         // UUID primary key
	Name      string    `db:"name"`       // human-readable label
	Vendor    string    `db:"vendor"`     // e.g. "sonicwall", "fortinet"
	Address   string    `db:"address"`    // IP or hostname
	Port      int       `db:"port"`       // syslog port (default 514)
	Protocol  string    `db:"protocol"`   // "udp", "tcp"
	Format    string    `db:"format"`     // "rfc3164", "rfc5424", "cef", etc.
	Facility  string    `db:"facility"`   // syslog facility, e.g. "local0"
	SyslogID  string    `db:"syslog_id"`  // vendor tag or hostname override
	RateLimit int       `db:"rate_limit"` // events/sec throttle (optional)
	CreatedAt time.Time `db:"created_at"` // record creation time
	UpdatedAt time.Time `db:"updated_at"` // record update time
	Status    string    `db:"status"`     // "enabled", "disabled"
}
// NetworkDeviceFilter is used to filter network devices based on various criteria.
type NetworkDeviceFilter struct {
	Name      string
	Vendor    string
	Address   string
	Port      int
	Protocol  string
	Format    string
	Facility  string
	SyslogID  string
	RateLimit int
	Limit  int    
	Order  string 
	Offset int
	Status string
}