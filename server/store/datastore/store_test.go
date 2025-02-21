// Copyright 2021 Woodpecker Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datastore

import (
	"database/sql"
	"os"
)

func testDriverConfig() (driver, config string) {
	driver = "sqlite3"
	config = ":memory:"

	if os.Getenv("WOODPECKER_DATABASE_DRIVER") != "" {
		driver = os.Getenv("WOODPECKER_DATABASE_DRIVER")
		config = os.Getenv("WOODPECKER_DATABASE_CONFIG")
	}

	return
}

// openTest opens a new database connection for testing purposes.
// The database driver and connection string are provided by
// environment variables, with fallback to in-memory sqlite.
func openTest() *sql.DB {
	db, _ := open(testDriverConfig())
	return db
}

// newTest creates a new database connection for testing purposes.
// The database driver and connection string are provided by
// environment variables, with fallback to in-memory sqlite.
func newTest() *datastore {
	driver, config := testDriverConfig()
	return &datastore{
		DB:     openTest(),
		driver: driver,
		config: config,
	}
}
