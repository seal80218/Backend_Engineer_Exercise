package main

import (
	"os"
	"reflect"
	"strings"

	"github.com/urfave/cli"
)

type Clawer struct {
	url string
}

var container Clawer

func main() {
	container.url = "https://www.alexa.com/"

	app := cli.NewApp()
	app.Name = "question4"
	app.Usage = "question4"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "action,a",
			Usage: "action",
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) {
	var builder strings.Builder
	builder.WriteString("Question4_")
	builder.WriteString(c.Args().First())
	f := reflect.ValueOf(&container).MethodByName(builder.String())
	params := c.Args().Tail()
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	f.Call(in)
}
