package con

import (
	"log"
	"os"

	"github/Rarik88/go_final_project/config"
)

// Объявление констант, используемых для настройки веб-сервера и базы данных
const (
	// WebDir = "web" // Директория с веб-ресурсами
	WebDir                                = "D:/World/go_final_project/web"
	ENV_PORT                              = "TODO_PORT"                               // Переменная окружения для порта
	API_NEXTDATE                          = "/api/nextdate"                           // URL-путь для API
	API_TASK                              = "/api/task"                               // URL-путь для API
	DATE_EVENT                            = "date"                                    // Константа для даты
	REPEAT_EVENT                          = "repeat"                                  // Константа для повторения
	INFO_GETTING_PORT_FROM_ENVIRONMENT    = "Получаем порт из окружения..."           // Сообщение о получении порта из переменной окружения
	INFO_GETTING_DB_NAME_FROM_ENVIRONMENT = "Получаем имя БД из окружения..."         // Сообщение о получении имени базы данных из переменной окружения
	INFO_USING_DEFAULT_PORT               = "Порт не задан. Будем использовать 7540"  // Сообщение о использовании порта по умолчанию
	PORT_SET                              = "Порт задан - "                           // Сообщение о задании порта
	DB_NAME_SET                           = "./db/scheduler.db"                       // Сообщение о задании имени базы данных
	SQL_DRIVER                            = "sqlite3"                                 // Драйвер для работы с базой данных SQLite
	SQL_CREATE_TABLES                     = "CREATE TABLE IF NOT EXISTS scheduler " + // SQL-запрос для создания таблицы в базе данных
		"(id INTEGER PRIMARY KEY AUTOINCREMENT, " + // Определение первичного ключа
		"date TEXT, " + // Поле для хранения даты
		"title TEXT, " + // Поле для заголовка задачи
		"comment TEXT, " + // Поле для комментария к задаче
		"repeat VARCHAR(128));" // Поле для правил повторения задачи
	SQL_CREATE_INDEX = "CREATE INDEX IF NOT EXISTS idx_date ON scheduler (date)"
)

// EnvPORT получает порт из переменной окружения или использует значение по умолчанию
func EnvPORT(key string) string {
	log.Println(INFO_GETTING_PORT_FROM_ENVIRONMENT) // Вывод сообщения о получении порта из переменной окружения
	port := os.Getenv(key)                          // Получение значения порта из переменной окружения
	if len(port) == 0 {                             // Если порт не задан
		log.Println(INFO_USING_DEFAULT_PORT) // Вывод сообщения о использовании порта по умолчанию
		port = config.Port                   // Использование значения порта по умолчанию из пакета config
	} else {
		log.Println(PORT_SET + port) // Вывод сообщения о задании порта
	}
	return ":" + port // Возвращение порта в формате ":<port>"
}
