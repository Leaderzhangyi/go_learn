package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func get_pg_dsn(db_name string) string {
	meta := viper.GetStringMapString("postgres")
	dsn := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=disable",
		meta["host"],
		meta["user"],
		meta["port"],
		meta["password"],
		db_name)

	return dsn

}

func GetRawDB() *pgxpool.Pool {

	dsn := get_pg_dsn("go_majiang")
	var err error
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}

	cfg.MaxConns = 100
	cfg.ConnConfig.ConnectTimeout = time.Second * 100
	cfg.ConnConfig.TLSConfig = nil
	db_pg_raw, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		panic(err)
	}
	return db_pg_raw

}

func main() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("can not find config file")
		panic(err)
	}

	db_pg_raw := GetRawDB()
	defer db_pg_raw.Close()

	// _, err,  := db_pg_raw.Exec(context.Background(), "SELECT 1")
	// do something with db_pg_raw

}
