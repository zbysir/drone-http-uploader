package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	c := cli.NewApp()
	c.Name = "qiniu plugin"
	c.Usage = ""
	c.Version = ""

	c.Action = func(ctx *cli.Context) (err error) {
		p := Plugin{
			AccessKey: ctx.String("access-key"),
			SercetKey: ctx.String("secret-key"),
			Bucket:    ctx.String("bucket"),
			Zone:      ctx.String("zone"),
			Prefix:    ctx.String("prefix"),
			Path:      ctx.String("path"),
		}

		return p.Exec()
	}

	c.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "access-key",
			Usage:  "qiniu access key",
			EnvVar: "ACCESS_KEY,PLUGIN_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "secret-key",
			Usage:  "qiniu secret key",
			EnvVar: "SECRET_KEY,PLUGIN_SECRET_KEY",
		},
		cli.StringFlag{
			Name:   "zone",
			Usage:  "bucket zone",
			EnvVar: "ZONE,PLUGIN_ZONE",
			Value:  "huadong",
		},
		cli.StringFlag{
			Name:   "bucket",
			Usage:  "bucket name",
			EnvVar: "BUCKET,PLUGIN_BUCKET",
		},
		cli.StringFlag{
			Name:   "prefix",
			Usage:  "upload key prefix",
			EnvVar: "PREFIX,PLUGIN_PREFIX",
		},
		cli.StringFlag{
			Name:   "path",
			Usage:  "file path",
			EnvVar: "PATH,PLUGIN_PATH",
		},
	}

	err := c.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
