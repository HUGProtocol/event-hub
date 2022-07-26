package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const DBNAME = "hug"

type DBService struct {
	DBUni *sql.DB
	DBVer *sql.DB
}

func Init(password string, url string, verName string) (*DBService, error) {
	dsn := "root:" + password + "@tcp(" + url + ")/" + DBNAME
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	dsnV := "root:" + password + "@tcp(" + url + ")/" + verName
	dbV, err := sql.Open("mysql", dsnV)
	if err != nil {
		return nil, err
	}
	err = dbV.Ping()
	if err != nil {
		return nil, err
	}
	return &DBService{
		DBUni: db,
		DBVer: dbV,
	}, nil

}


