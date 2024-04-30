package main

import (
	app "github/Rarik88/go_final_project"
	con "github/Rarik88/go_final_project/pkg/const"
	hl "github/Rarik88/go_final_project/pkg/handler"
	"github/Rarik88/go_final_project/pkg/hub"
	api "github/Rarik88/go_final_project/pkg/service"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// добавлена обработка ощибок в Дб
func main() {
	db := hub.Sqlite(con.DB_NAME_SET)
	if db == nil {
		log.Fatalf("Не удалось открыть соединение с базой данных")
	}
	defer db.Close()
	repo := hub.NewDB(db)
	service := api.NewApi(repo)
	handler := hl.NewHandler(service)
	serv := new(app.Server)

	err := serv.Run("7540", handler.Init())
	if err != nil {
		log.Fatal(err)
	}
}
