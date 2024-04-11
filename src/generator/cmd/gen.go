package cmd

import (
	"github.com/kube-stack/multicloud_service/src/generator/gen"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen -f [config file path]",
	Short: "gen cloud code based on the config file",
	Long:  "gen cloud code based on the config file",
	//Args:  cobra.ExactArgs(1),
	RunE: GenCloudCode,
}

// flags
var configFile string

func GenCloudCode(cmd *cobra.Command, args []string) error {
	config := gen.LoadCloudConfig(configFile)
	generator := gen.NewCloudAPIGenerator(config.CloudType)
	err := generator.DoGen(config)
	if err != nil {
		return err
	}
	return nil
}
