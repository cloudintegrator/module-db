package database

import (
	"context"
	"fmt"
)

type DBConnector struct {
}

func (db *DBConnector) Process(ctx *context.Context) {
	fmt.Println("Processing data...")
}
