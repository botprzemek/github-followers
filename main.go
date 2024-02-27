package main

import (
	"github-followers/config"
	"github-followers/utils/api"
	"github-followers/utils/cache"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Default(*cli.Context) error {
	err := cache.Load()
	if err != nil {
		return err
	}

	if err = cache.Save(); err != nil {
		return err
	}

	api.GetUsers()

	return err
}

func main() {
	config.Load()

	app := &cli.App{
		Name:                 "github-followers",
		Usage:                "make an explosive entrance",
		EnableBashCompletion: true,
		Action:               Default,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// https://api.github.com/user/followers
// https://api.github.com/user/following/cherrymui
