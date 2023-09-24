/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kopoze/kpz/pkg/hosts"
	"github.com/spf13/cobra"
)

var subdomainInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Add main domain from configuration file inside hosts file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts.InitDomain()
	},
}

var subdomainAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add subdomain inside hosts file",
	Long: `The 'add' command allows you to add subdomains to your system's hosts file. The hosts file is used to map hostnames to IP addresses on your local machine, making it useful for customizing domain-to-IP mappings for development or testing purposes.

	Usage:
	  kpz subdomain add <subdomain1> [<subdomain2> ...]

	Examples:
	  1. Add a single subdomain:
		 $ sudo kpz subdomain add mysubdomain

	  2. Add multiple subdomains at once:
		 $ sudo kpz subdomain add sub1 sub2 sub3

	This command takes one or more subdomain names as arguments and appends them to your system's hosts file. Please note that this command may require administrative privileges to modify the hosts file.

	Note: Be cautious when modifying your hosts file, as it can affect how your system resolves domain names. Always make sure you have the necessary permissions and understand the implications of the changes you make. If something goes wrong, you can restore your previous hosts file from the backup located under the /etc/kopoze/hosts/backup folder.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			hosts.AddSubdomain(arg)
		}
	},
}

var subdomainRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove subdomain inside hosts file",
	Long: `The 'remove' command allows you to remove previously added subdomains from your system's hosts file. The hosts file is used to map hostnames to IP addresses on your local machine, making it useful for customizing domain-to-IP mappings for development or testing purposes.

Usage:
  kpz subdomain remove <subdomain1> [<subdomain2> ...]

Examples:
  1. Remove a single subdomain:
     $ sudo kpz subdomain remove mysubdomain

  2. Remove multiple subdomains at once:
     $ sudo kpz subdomain remove sub1 sub2 sub3

This command takes one or more subdomain names as arguments and removes them from your system's hosts file. Please note that this command may require administrative privileges to modify the hosts file.

Note: Be cautious when using the 'remove' command, as it can affect how your system resolves domain names. Always make sure you have the necessary permissions and understand the implications of the changes you make. If something goes wrong, you can restore your previous hosts file from the backup located under the /etc/kopoze/hosts/backup folder.
`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			hosts.RemoveSubdomain(arg)
		}
	},
}

// subdomainCmd represents the subdomain command
var subdomainCmd = &cobra.Command{
	Use:   "subdomain",
	Short: "Manage subdomain inside hosts file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subdomain called")
	},
}

func init() {
	rootCmd.AddCommand(subdomainCmd)
	subdomainCmd.AddCommand(subdomainInitCmd)
	subdomainCmd.AddCommand(subdomainAddCmd)
	subdomainCmd.AddCommand(subdomainRemoveCmd)
	// TODO: Add rename subdomain command

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subdomainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subdomainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
