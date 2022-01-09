package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
)

var args struct {
	Path string `arg:"-p,--path" default:"." help:"the directory to search"`
}

func byteCountFromDecimal(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func findMatch(path string, f os.FileInfo, err error) error {
	for ext, description := range fileExtentions {
		if strings.HasSuffix(path, ext) {
			fmt.Printf("[%s%sM%s] %s%spath%s: %s, %s%stype%s: %s, %s%ssize%s: %s\n",
				yellow, bold, end, red, bold, end, path,
				green, bold, end, description,
				blue, bold, end, byteCountFromDecimal(f.Size()))
		}
	}

	return nil
}

func walk(path string) {
	if err := filepath.Walk(path, findMatch); err != nil {
		fmt.Printf("%serror%s: %s\n", red, end, err.Error())
	}
}

func main() {
	arg.MustParse(&args)

	now := time.Now()

	path := filepath.Clean(args.Path)

	walk(path)

	end := time.Since(now)
	fmt.Printf("\nTime elapsed: %s\n", end)
}
