package main

import (
	"flag"
	"github.com/fydy/issue-blog"
	_ "github.com/fydy/issue-blog/issues"
	//"github.com/gohugoio/hugo/commands"
)

func main() {
	var config string
	flag.StringVar(&config, "config", "./mirror.yaml", "the file path of mirror config")
	flag.Parse()
	mirror.SyncWithConfig(config)
	//commands.Execute()
}
