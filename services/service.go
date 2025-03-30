package services

import (
	"context"
	"database/sql"
	"fmt"
)

var DBConnK = struct{}{}


func dbConn(ctx context.Context) (*sql.Conn, error) {
	value := ctx.Value(DBConnK) 
	if value != nil {
		conn, ok := value.(*sql.Conn)
		if ok {
			return conn, nil
		}
		return nil, fmt.Errorf("service: Inavlid DB connection object: %T", conn)
	}
	return nil, fmt.Errorf("service: DB Connection not present in context")
}
