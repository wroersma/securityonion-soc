// Copyright 2019 Jason Ertel (jertel). All rights reserved.
// Copyright 2020-2021 Security Onion Solutions, LLC. All rights reserved.
//
// This program is distributed under the terms of version 2 of the
// GNU General Public License.  See LICENSE for further details.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package model

import (
  "github.com/security-onion-solutions/securityonion-soc/config"
)

type Info struct {
  Version        string                   `json:"version"`
  License        string                   `json:"license"`
  Parameters     *config.ClientParameters `json:"parameters"`
  ElasticVersion string                   `json:"elasticVersion"`
  WazuhVersion   string                   `json:"wazuhVersion"`
}
