package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/martinlindhe/base36"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "fe", Short: "multipurpose utility functions"}

	rootCmd.AddCommand(&cobra.Command{Use: "upcase", Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			log.Println(strings.ToUpper(v))
		}
	}})
	rootCmd.AddCommand(&cobra.Command{Use: "downcase", Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			log.Println(strings.ToLower(v))
		}
	}})
	rootCmd.AddCommand(&cobra.Command{Use: "uuid", Run: func(cmd *cobra.Command, args []string) {
		log.Println(uuid.New())
	}})
	rootCmd.AddCommand(&cobra.Command{Use: "lasthour", Run: func(cmd *cobra.Command, args []string) {
		t := time.Now()
		log.Printf("%d %d", t.Add(-time.Hour).Unix(), t.Unix())
	}})
	rootCmd.AddCommand(&cobra.Command{Use: "unix", Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			log.Println(time.Now().Unix())
		case 1:
			n, err := strconv.ParseInt(args[0], 10, 64)
			fatal(err)
			log.Println(time.Unix(n, 0))
		default:
			log.Fatal("epoc command accepts 0 or 1 arguments")
		}
	}})
	rootCmd.AddCommand(&cobra.Command{Use: "human", Args: cobra.ExactArgs(1), Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]

		raw, err := strconv.Atoi(arg)
		fatal(err)

		mutated := raw
		for _, prefix := range []string{"", "k", "M", "G", "T", "P"} {
			if mutated < 1024 {
				log.Printf("%d B -> %d %sB", raw, mutated, prefix)
				return
			}
			mutated = mutated / 1024
		}
		log.Fatalf("sorry %d, is out of bounds (we only cover petabytes)", raw)
	}})

	rootCmd.AddCommand(&cobra.Command{Use: "unescape", Run: func(cmd *cobra.Command, args []string) {
		for _, raw := range args {
			escaped, err := url.QueryUnescape(raw)
			if err != nil {
				log.Printf("%s -> %s", raw, err)
			}
			log.Printf("%s -> %s", raw, escaped)

		}
	}})

	rootCmd.AddCommand(&cobra.Command{Use: "escape", Run: func(cmd *cobra.Command, args []string) {
		for _, raw := range args {
			log.Printf("%s -> %s", raw, url.QueryEscape(raw))

		}
	}})

	rootCmd.AddCommand(&cobra.Command{Use: "kt", Run: func(cmd *cobra.Command, args []string) {
		for _, raw := range args {
			tok := strings.Split(raw, "-")
			t := base36.Decode(tok[0])
			log.Printf("%s -> %s", raw, time.Unix(int64(t/1000), 0))
		}
	}})

	err := rootCmd.Execute()
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		_, f, l, _ := runtime.Caller(1)
		fmt.Printf("%s:%d %s", f, l, err)
		os.Exit(-1)
	}
}
