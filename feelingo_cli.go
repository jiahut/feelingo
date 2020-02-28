package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"os/exec"

	"github.com/urfave/cli/v2"
)

// const iflsig = "AINFCbYAAAAAXljHYl_9dMXdKRenz7XGHbZyB37fYHhd"

// var searchPrefix = fmt.Sprintf("https://google.com/search?btnI&iflsig=%s&q=", iflsig)

const searchPrefix = "https://duckduckgo.com/?q=%5C"

func main() {
	app := cli.App{
		Name:  "feelingo",
		Usage: "I am feeling good for search ",
		Action: func(c *cli.Context) error {

			if c.NArg() < 1 {
				fmt.Printf("Must specify search query ")
				os.Exit(1)
			}

			searchStr := strings.Join(c.Args().Slice(), "+")

			searchUrl := searchPrefix + searchStr

			log.Println(searchUrl)
			return runCmd("xdg-open", searchUrl)

		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runCmd(prog string, args ...string) error {
	cmd := exec.Command(prog, args...)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	return cmd.Run()
}
