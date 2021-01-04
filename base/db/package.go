package db

import (
	"log"
	"task/model"
)

func Instance(configObj model.Config, db string) Client {
	// configObj := config.Instance()

	dbType := dbType(configObj)

	var dbClientObj Client

	//can implement this for the other db clients
	if dbType == "mysql" {
		dbClientObj = newMySQLClient(configObj, db)
	}

	if dbClientObj != nil {
		return dbClientObj
	}

	log.Fatalln("Invalid db type configured.")
	return nil
}

func dbType(configObj model.Config) string {
	dbType := configObj.DBType

	if dbType == "" {
		dbType = "mysql"
	}

	return dbType
}
