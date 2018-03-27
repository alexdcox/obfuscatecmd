package main

import (
	"fmt"
	"os/exec"
	"strings"
	"bytes"
	"github.com/urfave/cli"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "obfuscatecmd"
	app.Usage = "Make an otherwise readable bash command completely illegible and many times the character length."
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "Alex Cox",
			Email: "noreply@alexdcox.co.uk",
		},
	}
	app.UsageText = "obfuscatecmd [command options] '[bash commands]'"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name: "iterations, i",
			Usage: "Determines the `NUMBER` of times to obfuscate and pack the provided command.",
			Value: 5,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			return cli.ShowCommandHelp(c, "")
		}

		payload := c.Args().Get(c.NArg() - 1)
		iterations := c.Int("iterations")

		if payload == "" {
			return cli.ShowCommandHelp(c, "")
		}

		for i := 1; i <= iterations; i++ {
			payload = packWithDecoder(obfuscate(payload))
		}

		fmt.Println(payload)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func obfuscate(s string) string {
	basecmd := exec.Command("base64")
	basecmd.Stdin = strings.NewReader(s)
	var baseout bytes.Buffer
	basecmd.Stdout = &baseout
	if err := basecmd.Run(); err != nil {
		panic(err)
	}

	revcmd := exec.Command("rev")
	revcmd.Stdin = strings.NewReader(baseout.String())
	var revout bytes.Buffer
	revcmd.Stdout = &revout
	if err := revcmd.Run(); err != nil {
		panic(err)
	}

	return revout.String()
}

func packWithDecoder(encoded string) string {
	return fmt.Sprintf(
		"a='%s'; %s $a | %s | %s | %s;",
		encoded,
		substringCommand(encoded, "echo"),
		substringCommand(encoded, "rev"),
		substringCommand(encoded, "base64 -D"),
		substringCommand(encoded, "sh"),
	)
}

func substringCommand(full, part string) (ans string) {
next:
	for _, bc := range part {
		for si, sc := range full {
			if sc == bc {
				ans += fmt.Sprintf("${a:%d:1}", si)
				continue next
			}
		}

		ans += string(bc)
	}

	return
}
