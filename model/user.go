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

type User struct {
  Id												string									`json:"id"`
  CreateTime                time.Time 							`json:"createTime"`
  UpdateTime	              time.Time 							`json:"updateTime"`
  Email											string									`json:"email"`
  FirstName                 string    							`json:"firstName"`
  LastName	                string    							`json:"lastName"`
  Role    	                string    							`json:"role"`
  Status                    string                  `json:"status"`
}

func NewUser() *User {
  return &User{
		CreateTime: time.Now(),
    Email: "",
    FirstName: "",
    LastName: "",
    Status: "",
  }
}
