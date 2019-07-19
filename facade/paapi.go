package facade

import (
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	cfgFile string //config file
)

//newpaapiCmd returns cobra.Command instance for show sub-command
func newPaApiCmd(ui *rwi.RWI) *cobra.Command {
	paapiCmd := &cobra.Command{
		Use:   "paapi",
		Short: "Search for books data by PA-API",
		Long:  "Search for books data by PA-API",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ui.OutputErrln(strings.Join(usage, "\n"))
		},
	}
	//config file option
	paapiCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default $HOME/.paapi.yaml)")
	cobra.OnInitialize(initConfig)

	//options for PA-API
	paapiCmd.Flags().StringP("marketplace", "", "webservices.amazon.co.jp", "PA-API: Marketplace")
	paapiCmd.Flags().StringP("associate-tag", "", "", "PA-API: Associate Tag")
	paapiCmd.Flags().StringP("access-key", "", "", "PA-API: Access Key ID")
	paapiCmd.Flags().StringP("secret-key", "", "", "PA-API: Secret Access Key")
	_ = viper.BindPFlag("marketplace", paapiCmd.Flags().Lookup("marketplace"))
	_ = viper.BindPFlag("associate-tag", paapiCmd.Flags().Lookup("associate-tag"))
	_ = viper.BindPFlag("access-key", paapiCmd.Flags().Lookup("access-key"))
	_ = viper.BindPFlag("secret-key", paapiCmd.Flags().Lookup("secret-key"))

	return paapiCmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}
		// Search config in home directory with name ".paapi.yaml" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".paapi")
	}
	viper.AutomaticEnv()     // read in environment variables that match
	_ = viper.ReadInConfig() // If a config file is found, read it in.
}

/* Copyright 2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
