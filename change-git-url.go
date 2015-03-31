package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/yelinaung/git-url"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"

	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) {
		remote := c.Args()[0]
		if strings.Contains(remote, "@") {
			fmt.Println(giturl.Ssh2Http(remote))
		} else {
			fmt.Println(giturl.Http2Ssh(remote))
		}
	}

	app.Run(os.Args)
}
