package cmd

import (
  "fmt"
  
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "ffind",
  Version: 1.0,
  Short: "a tool to find interesting files for a pentester",
  Long: `ffind is tool that looks for configuration files or files 
        that commonly contain sensitive information within a specified directory`,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
  }
}
