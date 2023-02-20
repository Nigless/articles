// This is custom goose binary with sqlite3 support only.

package main

import (
	"flag"
	"log"
	"os"
	"rest-api/config"
	"rest-api/pkg/postgres"

	"github.com/pressly/goose"
)

var (
	flags   = flag.NewFlagSet("migrate", flag.ExitOnError)
	command = flags.String("command", "up", "up or down")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	if len(args) > 1 {
		log.Fatalf("to mach arguments")
		return
	}

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to read config: %v\n", err)
	}

	db, err := postgres.New(cfg.Postgres)

	if err != nil {
		log.Fatalf("failed to open postgres: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("failed to close postgres: %v\n", err)
		}
	}()

	command := args[0]
	if err := goose.Run(command, db, "./pkg/postgres/migrations"); err != nil {
		log.Fatalf("%v: %v", command, err)
	}
}
