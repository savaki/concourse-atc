package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

var flags = []cli.Flag{
	cli.BoolFlag{"dev", "dev mode; lax security", "ATC_DEV"},
	cli.StringFlag{"callbacksURL", "http://127.0.0.1:8080", "URL used for callbacks to reach the ATCD (excluding basic auth)", "ATC_CALLBACK_URL"},
	cli.StringFlag{"checkInterval", "1m0s", "interval on which to poll for new versions of resources", "ATC_CHECK_INTERVAL"},
	cli.StringFlag{"httpUsername", "", "basic auth username for the server", "ATC_USERNAME"},
	cli.StringFlag{"httpPassword", "", "basic auth password for the server", "ATC_PASSWORD"},
	cli.StringFlag{"sqlDataSource", "postgres://127.0.0.1:5432/atc?sslmode=disable", "database/sql data source configuration string", "ATC_SQL_DATASOURCE"},
	cli.StringFlag{"sqlDriver", "postgres", "database/sql driver name", "ATC_SQL_DRIVER"},
	cli.StringFlag{"public", "web/public", "path to directory containing public resources (javascript, css, etc.)", "ATC_PUBLIC"},
	cli.StringFlag{"templates", "web/templates", "path to directory containing the html templates", "ATC_TEMPLATES"},
	cli.IntFlag{"webListenPort", 8080, "port for the web server to listen on", "ATC_PORT"},
	cli.StringFlag{"atc", "atc", "path to atc command", "ATC_ATC"},
}

func main() {
	app := cli.NewApp()
	app.Flags = flags
	app.Action = Run
	app.Run(os.Args)
}

func Run(c *cli.Context) {
	args := []string{}
	for _, flag := range flags {
		name := ""
		switch v := flag.(type) {
		case cli.StringFlag:
			name = v.Name
		case cli.IntFlag:
			name = v.Name
		case cli.BoolFlag:
			name = v.Name
		}
		if name == "atc" {
			continue
		}
		if value := c.String(name); value != "" {
			args = append(args, fmt.Sprintf("-%s", name))
			args = append(args, value)
		}
	}

	atc := c.String("atc")
	cmd := exec.Command(atc, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
