package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	dialect     = "pgx"
	fmtDBString = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
)

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String("dir", "migrations", "directory with mifgration files")
)

type DB struct {
	Host     string `env:"PG_HOST"`
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	DBName   string `env:"PG_DBNAME"`
	Port     string `env:"PG_PORT"`
}

func GetDNS(c *DB) string {
	return fmt.Sprintf(fmtDBString, c.Host, c.User, c.Password, c.DBName, c.Port)
}

func main() {
	flags.Parse(os.Args[1:])

	args := flags.Args()

	command := args[0]
	var c DB
	configPath := "./.env"
	if err := cleanenv.ReadConfig(configPath, &c); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	dbString := GetDNS(&c)

	db, err := goose.OpenDBWithDriver(dialect, dbString)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	if err := goose.RunContext(context.Background(), command, db, *dir, args[1:]...); err != nil {
		log.Printf("migrate %v: %v", command, err)
	}
}
