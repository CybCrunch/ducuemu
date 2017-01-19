package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"../config"
)

type DBHandler struct {

	cfg	*config.DucuemuConfig
	dbconn	string
}

func NewDBHandler(config *config.DucuemuConfig) (*DBHandler) {
	dbh := &DBHandler{config, ""}

	dbuser, dbpass, dbname, dbhost, dbtype := dbh.cfg.DB()
	if dbtype == "tcp"{
		dbh.dbconn = string(dbuser+":"+dbpass+"@tcp("+dbhost+")/"+dbname)
		fmt.Println("DB Config - User("+dbuser+") Pass("+dbpass+") DB("+dbname+") Hostname("+dbhost+")")
	} else {
		dbh.dbconn = string(dbuser + ":" + dbpass + "@/" + dbname)
		fmt.Println("DB Config - User("+dbuser+") Pass("+dbpass+") DB("+dbname+")")
	}
	return dbh
}

func (dbh *DBHandler) Setup() error {
	return dbh.createTables()
}


func (dbh *DBHandler) CheckCount(rows *sql.Rows) (int, error) {

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil{
			return 0, err
		}
	}
	return count, nil
}


func (dbh *DBHandler) Query(query string) (*sql.Rows, error){

	// Create the database handle, confirm driver is present
	db, _ := sql.Open("mysql", dbh.dbconn)
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil

}
