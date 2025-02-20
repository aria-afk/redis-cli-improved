/*
Copyright © 2025 Aria Lopez <aria.lopez.dev@proton.me>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/aria-afk/redis-clii/client"
	"github.com/aria-afk/redis-clii/gui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "redis-clii",
	Short: "An improved version of the redis-cli",
	Long:  "An improved version of the redis-cli",
	Run:   OpenRedis,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var ctx = context.Background()

func init() {
	rootCmd.Flags().String("addr", "localhost:6379", "Address redis is running on, default is localhost:6379")
	rootCmd.Flags().String("pwd", "", "Password for redis instance, default is no password")
	rootCmd.Flags().Int("db", 0, "DB id to connect to, default is 0")
}

func OpenRedis(cmd *cobra.Command, args []string) {
	addr, _ := cmd.Flags().GetString("addr")
	pwd, _ := cmd.Flags().GetString("pwd")
	dbID, _ := cmd.Flags().GetInt("db")

	_, err := client.NewRedis(cmd.Context(), addr, pwd, dbID)
	if err != nil {
		// TODO: Error message
		os.Exit(1)
	}
	// TODO: better connection message
	fmt.Println("Succesfully connected to redis")
	ui := gui.NewGUI(addr)
	ui.Run()
}
