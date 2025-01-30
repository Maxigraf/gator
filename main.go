package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/maxigraf/gator/internal/command"
	"github.com/maxigraf/gator/internal/config"
	"github.com/maxigraf/gator/internal/database"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		fmt.Printf("Config reading failed: %v\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)

	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
	}

	queries := database.New(db)

	state := command.State{
		Database: queries,
		Config:   &cfg,
	}

	commands := command.Commands{}
	commands.Register("login", command.HandlerLogin)
	commands.Register("register", command.HandlerRegister)
	commands.Register("reset", command.HandlerReset)
	commands.Register("users", command.HandlerUsers)
	commands.Register("agg", command.HandlerAggregate)
	commands.Register("addfeed", command.HandlerAddFeed)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	command := command.Command{
		Name: args[1],
		Args: args[2:],
	}

	err = commands.Run(&state, command)

	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		os.Exit(1)
	}
}
