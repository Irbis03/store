package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"example.com/backend/internal"
)

func main() {
	// открываем соединение с базой данных
	connStr := "user=postgres dbname=mydb sslmode=disable password=password"

	var err error
	internal.Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// проверяем соединение
	err = internal.Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// инициализируем приложение и запускаем приложение
	app := internal.NewApi()

	log.Fatal(app.Listen(":5000"))
}