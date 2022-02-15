package main

import (
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
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

	err := rootCmd.Execute()
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		_, f, l, _ := runtime.Caller(1)
		log.Fatalf("\n!!!!!  FATAL ERROR  !!!!!\n%s:%d: %s\n¡¡¡¡¡  FATAL ERROR  ¡¡¡¡¡", f, l, err)
	}
}
