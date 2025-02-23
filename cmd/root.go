/*
Copyright © 2025 Aria Lopez <aria.lopez.dev@proton.me>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aria-afk/redis-clii/client"
	"github.com/aria-afk/redis-clii/gui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "redis-clii",
	Short: "An improved version of the redis-cli",
	Long:  "An improved version of the redis-cli",
	Run:   run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Connection options
	rootCmd.Flags().String("h", "localhost", "IP adress or hostname to connect too. Can be passed via REDISCLI_HOST as well. Default is localhost")
	rootCmd.Flags().String("a", "", "Authentication password. Highly recommended to provide this via REDISCLI_AUTH or TODO: after connection attempt. Default is empty")
	rootCmd.Flags().String("u", "redis://localhost:6379/0", "URI string to connect to redis, format: redis://user:password@host:port/dbnum. Can be passed via REDISCLI_URI, Default is redis://localhost:6379")
	rootCmd.Flags().String("user", "", "DB User, can be passed via REDISCLI_USER, Default is empty")
	rootCmd.Flags().Int("n", 0, "DB number to connect to. Can be passed via REDISCLI_NUM, default is 0")
	rootCmd.Flags().Int("p", 6379, "Port to connect too. Can be passed via REDISCLI_PORT, default is 6379")
	// TODO: TLS/SLS POOLING?
	// TODO: All the other stuff; csv, commands passing, read in from stdin etc
}

func run(cmd *cobra.Command, args []string) {
	cliOpts := client.RedisOptions{}
	envOpts := client.RedisOptions{}

	host, _ := cmd.Flags().GetString("h")
	if cmd.Flags().Changed("h") {
		cliOpts.Host = host
	}
	REDISCLI_HOST := os.Getenv("REDISCLI_HOST")
	if len(REDISCLI_HOST) > 0 {
		envOpts.Host = REDISCLI_HOST
	}

	auth, _ := cmd.Flags().GetString("a")
	if cmd.Flags().Changed("a") {
		cliOpts.Auth = auth
	}
	REDISCLI_AUTH := os.Getenv("REDISCLI_AUTH")
	if len(REDISCLI_AUTH) > 0 {
		envOpts.Auth = REDISCLI_AUTH
	}

	uri, _ := cmd.Flags().GetString("u")
	if cmd.Flags().Changed("u") {
		cliOpts.Uri = uri
	}
	REDISCLI_URI := os.Getenv("REDISCLI_URI")
	if len(REDISCLI_URI) > 0 {
		envOpts.Uri = REDISCLI_URI
	}

	user, _ := cmd.Flags().GetString("user")
	if cmd.Flags().Changed("user") {
		cliOpts.User = user
	}
	REDISCLI_USER := os.Getenv("REDISCLI_USER")
	if len(REDISCLI_USER) > 0 {
		envOpts.User = REDISCLI_USER
	}

	dbnum, _ := cmd.Flags().GetInt("n")
	if cmd.Flags().Changed("n") {
		cliOpts.Number = dbnum
	}
	REDISCLI_NUM := os.Getenv("REDISCLI_NUM")
	if len(REDISCLI_NUM) > 0 {
		parsedNum, err := strconv.ParseInt(REDISCLI_NUM, 0, 64)
		if err != nil {
			// TODO: Bad arg supplied
			fmt.Println(err)
		}
		envOpts.Number = int(parsedNum)
	}

	port, _ := cmd.Flags().GetInt("p")
	if cmd.Flags().Changed("p") {
		cliOpts.Port = port
	}
	REDISCLI_PORT := os.Getenv("REDISCLI_PORT")
	if len(REDISCLI_PORT) > 0 {
		parsedNum, err := strconv.ParseInt(REDISCLI_PORT, 0, 64)
		if err != nil {
			// TODO: Bad arg supplied
			fmt.Println(err)
		}
		envOpts.Port = int(parsedNum)
	}

	defaultOpts := client.RedisOptions{
		Host:   host,
		Auth:   auth,
		User:   user,
		Uri:    uri,
		Number: dbnum,
		Port:   port,
	}

	_, err := client.NewRedis(cmd.Context(), cliOpts, envOpts, defaultOpts)
	if err != nil {
		fmt.Printf("Error connecting to redis:\n%s\ngoodbye.", err)
	}
	ui := gui.NewGUI("localhost:6379")
	ui.Run()
}
