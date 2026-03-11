package main

import (
	"Study/postgres/simple_connection"
	"Study/postgres/simple_sql"
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	if err := simple_sql.InsertRow(
		ctx,
		conn,
		"Обед2",
		"Покушать",
		false,
		time.Now(),
	); err != nil {
		panic(err)
	}

	fmt.Println("Успешно!")
}
