package main

import (
	"log"
	"runtime"
	"strings"

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

	err := rootCmd.Execute()
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		_, f, l, _ := runtime.Caller(1)
		log.Fatalf("\n!!!!!  FATAL ERROR  !!!!!\n%s:%d: %s\n¡¡¡¡¡  FATAL ERROR  ¡¡¡¡¡", f, l, err)
	}
}
