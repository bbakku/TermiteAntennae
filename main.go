package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var configDir = "conf/"

type DatabaseConf struct {
	DriverName, DataSourceName string
}

func main() {
	dbConfigFile, err := os.Open(configDir + "database.json")
	if err != nil {
		panic(err)
	}

	var dbConf DatabaseConf
	dbConfigDecoder := json.NewDecoder(dbConfigFile)
	err = dbConfigDecoder.Decode(&dbConf)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(dbConf.DriverName, dbConf.DataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
