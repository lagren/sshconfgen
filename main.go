package main

import "github.com/lagren/sshconfgen/cmd"

var version string
var commit string
var date string

func main() {
	cmd.Execute(version, commit, date)
}
