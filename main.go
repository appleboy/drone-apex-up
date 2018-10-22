package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

// Version set at compile-time
var (
	Version  string
	BuildNum string
)

func main() {
	app := cli.NewApp()
	app.Name = "apex-up client plugin"
	app.Usage = "deploying service to Lambda using apex/up tool"
	app.Copyright = "Copyright (c) 2018 Bo-Yi Wu"
	app.Authors = []cli.Author{
		{
			Name:  "Bo-Yi Wu",
			Email: "appleboy.tw@gmail.com",
		},
	}
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "access-key",
			Usage:  "AWS ACCESS KEY",
			EnvVar: "PLUGIN_ACCESS_KEY,AWS_ACCESS_KEY_ID",
		},
		cli.StringFlag{
			Name:   "secret-key",
			Usage:  "AWS SECRET KEY",
			EnvVar: "PLUGIN_SECRET_KEY,AWS_SECRET_ACCESS_KEY",
		},
		cli.StringSliceFlag{
			Name:   "stage",
			Usage:  "Target stage name",
			EnvVar: "PLUGIN_STAGE",
		},
	}

	app.Version = Version

	if BuildNum != "" {
		app.Version = app.Version + "+" + BuildNum
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("can't start app")
	}
}

func run(c *cli.Context) error {
	var version string
	if BuildNum != "" {
		version = Version + "+" + BuildNum
	} else {
		version = Version
	}
	log.Info().Str("revision", version).Msg("Drone apex-up Plugin Version")

	plugin := Plugin{
		Config: Config{
			AccessKey: c.String("access-key"),
			SecretKey: c.String("secret-key"),
			Stage:     c.StringSlice("stage"),
		},
	}

	return plugin.Exec()
}
