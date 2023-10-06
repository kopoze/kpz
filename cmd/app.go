/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/kopoze/kpz/pkg/client"
	"github.com/spf13/cobra"
)

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Manage app inside database",
	Long:  `Manage app inside database.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var appAddCmd = &cobra.Command{
	Use:   "add [flags] app_name app_subdomain app_port",
	Short: "Add app to the database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			cmd.Help()
			os.Exit(0)
		}
		client.Create(args[0], args[1], args[2])
	},
}

var appRemoveCmd = &cobra.Command{
	Use:   "remove [flags] app_id",
	Short: "Remove app from the database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(0)
		}
		client.Delete(args[0])
	},
}

var appListCmd = &cobra.Command{
	Use:   "list",
	Short: "List existing app from the database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.List()
	},
}

var appUpdateCmd = &cobra.Command{
	Use:   "update [flags] app_id app_field app_value",
	Short: "Update existing app in the database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			cmd.Help()
			os.Exit(0)
		}
		client.Update(args[0], args[1], args[2])
	},
}

func init() {
	rootCmd.AddCommand(appCmd)
	appCmd.AddCommand(appAddCmd)
	appCmd.AddCommand(appUpdateCmd)
	appCmd.AddCommand(appRemoveCmd)
	appCmd.AddCommand(appListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
