package database

import (
	"database/sql"
	"fmt"

	"github.com/asyauqi1511/go-workshop-service/internal/entity"
)

const (
	// dataSourceFormat <username>:<password>@tcp(<hostname>:<port>)/<database_name> .
	dataSourceFormat = "%s:%s@tcp(%s:%d)/%s"

	// driverName is kind of sql we will use.
	driverName = "mysql"
)

func ConnectDB(config entity.DBConfig) (*sql.DB, error) {
	dataSource := fmt.Sprintf(dataSourceFormat, config.Username, config.Password, config.Hostname, config.Port, config.DatabaseName)

	db, err := sql.Open(driverName, dataSource)
	return db, err
}
