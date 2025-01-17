package main

import (
	"encoding/base64"
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
	"github.com/oklog/ulid/v2"
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
	rootCmd.AddCommand(&cobra.Command{Use: "ulid", Run: func(cmd *cobra.Command, args []string) {
		log.Println(ulid.Make())
	}})
	rootCmd.AddCommand(&cobra.Command{Use: "lasthour", Run: func(cmd *cobra.Command, args []string) {
		t := time.Now()
		log.Printf("lasthour: %s", t.Add(-time.Hour).Format(time.RFC1123))
		log.Printf("now: %s", t)
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

	rootCmd.AddCommand(&cobra.Command{Use: "to_yaml", Run: func(cmd *cobra.Command, args []string) {
		log.Printf("%s ->\n[\"%s\"]", args, strings.Join(args, `", "`))
	}})

	rootCmd.AddCommand(&cobra.Command{Use: "to_string", Run: func(cmd *cobra.Command, args []string) {
		// log.Printf("%s ->\n", args, strings.Trim(strings.Join(args)))
		for i, v := range args {
			log.Println(i, v)
		}
	}})

	rootCmd.AddCommand(&cobra.Command{Use: "kt", Run: func(cmd *cobra.Command, args []string) {
		for _, raw := range args {
			tok := strings.Split(raw, "-")
			t := base36.Decode(tok[0])
			log.Printf("%s -> %s", raw, time.Unix(int64(t/1000), 0))
		}
	}})

	base64CMD := &cobra.Command{Use: "base64"}
	base64CMD.AddCommand(&cobra.Command{Use: "encode", Run: func(cmd *cobra.Command, args []string) {
		for _, raw := range args {
			log.Printf("%s -> %s", raw, base64.StdEncoding.EncodeToString([]byte(raw)))
		}
	}})
	base64CMD.AddCommand(&cobra.Command{Use: "decode", Run: func(cmd *cobra.Command, args []string) {
		for _, raw := range args {
			res, err := base64.StdEncoding.DecodeString(raw)
			if err != nil {
				log.Printf("%s -> error: %v", raw, err)
			}
			log.Printf("%s -> %s", raw, res)
		}
	}})

	rootCmd.AddCommand(base64CMD)

	rootCmd.AddCommand(&cobra.Command{Use: "charcount", Run: func(cmd *cobra.Command, args []string) {
		for _, raw := range args {
			log.Printf("%s -> %d", raw, len(raw))
		}
	}})

	rootCmd.AddCommand(&cobra.Command{Use: "sqlGo2Cli", Run: func(cmd *cobra.Command, args []string) {
		for _, raw := range args {
			u, err := url.Parse(raw)
			fatal(err)

			log.Println(u.User)
			// log.Printf("raw -> %d", len(raw))
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
