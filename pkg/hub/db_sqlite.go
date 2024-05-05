package hub

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"

	con "github/Rarik88/go_final_project/pkg/const"
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

func Sqlite(dbname string) (*sqlx.DB, error) {
	dbfile, err := CheckDb(dbname)
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Open(SqlDriver, dbfile)
	if err != nil {
		return nil, err
	}
	return db, err
}

func CheckDb(dbName string) (string, error) {
	appPath, err := os.Executable()
	if err != nil {
		return "", err
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
