package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	c := cli.NewApp()
	c.Name = "http uploader plugin"
	c.Usage = ""
	c.Version = ""

	c.Action = func(ctx *cli.Context) (err error) {
		p := Plugin{
			Debug: ctx.Bool("debug"),
			Url:   ctx.String("url"),
			Path:  ctx.String("path"),
		}

		return p.Exec()
	}

	c.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "enable log",
			EnvVar: "PLUGIN_DEBUG,DEBUG",
		},
		cli.StringFlag{
			Name:   "url",
			Usage:  "service upload url",
			EnvVar: "PLUGIN_URL,URL",
		},
		cli.StringFlag{
			Name:   "path",
			Usage:  "file path",
			EnvVar: "PLUGIN_PATH,PATH",
		},
	}

	err := c.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
