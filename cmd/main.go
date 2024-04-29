package main

import (
	_ "github.com/mattn/go-sqlite3"
	app "github/Rarik88/go_final_project"
	"github/Rarik88/go_final_project/pkg/api"
	con "github/Rarik88/go_final_project/pkg/const"
	hl "github/Rarik88/go_final_project/pkg/handler"
	"github/Rarik88/go_final_project/pkg/hub"
	"log"
)

func main() {
	db := hub.Sqlite(con.DB_NAME_SET)
	repo := hub.NewDB(db)
	service := api.NewApi(repo)
	handler := hl.NewHandler(service)
	serv := new(app.Server)

	err := serv.Run("7540", handler.Init())
	if err != nil {
		log.Fatal(err)
	}
}
