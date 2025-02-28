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
  "time"
)

const NodeStatusUnknown = "unknown"
const NodeStatusOffline = "offline"
const NodeStatusOnline = "online"
const NodeStatusOk = "ok"
const NodeStatusError = "error"

type Node struct {
  Id                      string    `json:"id"`
  OnlineTime              time.Time `json:"onlineTime"`
  UpdateTime              time.Time `json:"updateTime"`
  EpochTime               time.Time `json:"epochTime"`
  UptimeSeconds           int       `json:"uptimeSeconds"`
  Description             string    `json:"description"`
  Address                 string    `json:"address"`
  Role                    string    `json:"role"`
  Status                  string    `json:"status"`
  Version                 string    `json:"version"`
}

func NewNode(id string) *Node {
  return &Node{
    Id: id,
    Status: NodeStatusUnknown,
    OnlineTime: time.Now(),
    UpdateTime: time.Now(),
  }
}
