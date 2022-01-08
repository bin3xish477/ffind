package cmd

import (
  "fmt"
  "filepath"
)

func fileMatch(p string, f os.FileInfo, err error) error {
  for _, ext := range fileExtentions {
    if strings.HasSuffix(p, ext) {
      fmt.Printf("[•] found: %s", p)
    }
  }
}

var scan := &cobra.Command{
  Use:   "scan",
	Short: "ffind scan",
	Run: func(cmd *cobra.Command, args []string) {
    if err := os.filepath.Walk(path, fileMatch); err != nil {
      fmt.Printf"error walking directory: %s", err.Error())
    }
	},
}

func init() {
  root.AddCommand(scan)
  
  scan.Flag().StringVarP(
    &path, "path", "p", ".", "the directory to traverse for interesting files",
  )
  
  scan.MarkFlagRequired("path")
}
