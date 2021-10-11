package main

import (
	"context"
	"database/sql"
	"fmt"
)

var database *sql.DB

var (
	port = 1433
	password = "iamnwani01"
	user = "SA"
)

var dbContext = context.Background()

var connectionString = fmt.Sprintf("user id=%s;password=%s;port=%d", user, password, port)
