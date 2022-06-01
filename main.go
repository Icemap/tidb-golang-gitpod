// Copyright 2022 PingCAP, Inc.
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

package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:4000)/test?charset=utf8mb4")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT VERSION()")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    if rows.Next() {
        version := ""
        err = rows.Scan(&version)
        if err != nil {
            panic(err)
        }

        fmt.Println(version)
    }
}
