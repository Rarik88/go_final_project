package hub

import (
	"github.com/jmoiron/sqlx"
	"github/Rarik88/go_final_project/pkg/const"
	"log"
	"os"
	"path/filepath"
)

const (
	SqlDriver = "sqlite3"
)

type TaskSQLite struct {
	db *sqlx.DB
}

func NewDB(db *sqlx.DB) *TaskSQLite {
	return &TaskSQLite{db: db}
}

func Sqlite(dbname string) *sqlx.DB {
	dbfile, err := CheckDb(dbname)
	if err != nil {
		return nil
	}
	db, err := sqlx.Open(SqlDriver, dbfile)
	if err != nil {
		return nil
	}
	return db
}

func CheckDb(dbName string) (string, error) {
	appPath, err := os.Executable()
	if err != nil {
	}
	dbFile := filepath.Join(filepath.Dir(appPath), dbName)
	_, err = os.Stat(dbFile)
	if err != nil {
		InstallDB(dbName)
	}
	return dbName, nil
}

func InstallDB(dbName string) {
	db, err := sqlx.Open(con.SQL_DRIVER, dbName)
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных", err)
	}
	defer func() { _ = db.Close() }()

	_, err = db.Exec(con.SQL_CREATE_TABLES)
	if err != nil {
		log.Fatal("Не удалось создать таблицу", err)
	}

	_, err = db.Exec(con.SQL_CREATE_INDEX)
	if err != nil {
		log.Fatal("Не удалось создать индекс", err)
	}
}
