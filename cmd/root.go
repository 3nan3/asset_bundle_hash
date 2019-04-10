package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/3nan3/asset_bundle_hash/pkg"
)

var (
	rootCmd = &cobra.Command {
		Use:   "asset_bundle_hash [manifest_file]",
		Short: "generate asset bundle hash from origin assets",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := pkg.CalculateHash(args[0]); err != nil {
				cmd.Println(err)
				os.Exit(1)
			}
		},
	}
)

func Execute() {
	rootCmd.Execute()
}
