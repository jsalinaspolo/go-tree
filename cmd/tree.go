package cmd

import (
	"github.com/jsalinaspolo/yaml-tree/pkg/tree"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func treeCmd() *cobra.Command {
	fileFlag := "file"

	cmd := &cobra.Command{
		Use:   "tree",
		Short: "Prints a tree of a yaml file",
		PreRun: func(cmd *cobra.Command, args []string) {
			_ = viper.BindPFlag(fileFlag, cmd.Flags().Lookup(fileFlag))

		},

		RunE: func(cmd *cobra.Command, args []string) error {
			file := viper.Get(fileFlag).(string)
			err := tree.Generate(file)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringP(fileFlag, "f", "", "yaml file")
	_ = cmd.MarkFlagRequired(fileFlag)

	return cmd
}

func init() {
	rootCmd.AddCommand(treeCmd())
}
