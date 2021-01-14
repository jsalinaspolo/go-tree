package cmd

import (
	"github.com/jsalinaspolo/yaml-tree/pkg/tree"
	"github.com/spf13/cobra"
)

func treeCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "tree",
		Short: "Prints a tree of a structure file",
		PreRun: func(cmd *cobra.Command, args []string) {
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			tree.Generate()

			return nil
		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(treeCmd())
}
