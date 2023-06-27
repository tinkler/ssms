package main

import (
	"flag"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

var (
	down = flag.Bool("down", false, "migrate up")
	p    = flag.String("p", "service/ssms_auth", "migrate project name")
	dev  = flag.Bool("dev", false, "migrate dev")
)

// Migrate database
func main() {
	flag.Parse()
	if *dev {
		os.Chdir("../../" + *p)
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:" + err.Error())
	}
	var m *migrate.Migrate
	switch *p {
	case "service/ssms_auth":
		m, err = migrate.New(
			"file://./db",
			os.Getenv("MIGRATE_DSN"),
		)
	default:
		log.Fatal("Unknown project name")
	}
	if err != nil {
		log.Fatal(err)
	}
	if *down {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	}
}
