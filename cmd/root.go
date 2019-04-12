package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/3nan3/asset_bundle_hash/pkg"
)

var (
	projectRoot string
	versionFlag bool

	rootCmd = &cobra.Command {
		Use:   "asset_bundle_hash [manifest_file]",
		Short: "generate asset bundle hash from origin assets",
		Args: func(cmd *cobra.Command, args []string) error {
			if versionFlag {
				return nil
			}

			if len(args) != 1 {
				return fmt.Errorf("requires only 1 arg(s), only received %d", len(args))
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if versionFlag {
				pkg.PrintVersion()
				os.Exit(0)
			}

			hash, err := pkg.CalculateHash(args[0], projectRoot)
			if err != nil {
				cmd.Println(err)
				os.Exit(1)
			}

			cmd.Println(hash)
		},
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "print version number")
	rootCmd.Flags().StringVarP(&projectRoot, "project-root", "d", ".", "change unity project root directory")
}
