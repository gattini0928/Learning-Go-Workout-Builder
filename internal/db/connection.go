package db

import (
	"log"
	"database/sql"
	"fmt"
	"github.com/gattini0928/Learning-Go-Workout-Builder/internal/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error){
	conf := configs.GetDb()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	log.Println(conf)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return conn, err
}