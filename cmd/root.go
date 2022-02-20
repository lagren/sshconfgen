package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use: "sshconfgen",
}

func Execute(version, commit, date string) {
	if version == "" {
		rootCmd.Version = "development build"
	} else if strings.HasSuffix(version, "-next") {
		rootCmd.Version = fmt.Sprintf("%s %s (commit %s)", version, date, commit)
	} else {
		rootCmd.Version = fmt.Sprintf("%s (%s)", version, date)
	}

	if err := rootCmd.Execute(); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}
