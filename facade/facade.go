package facade

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/gocli/config"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	//Name is applicatin name
	Name = "books-data"
	//Version is version for applicatin
	Version = "developer version"
)
var (
	debugFlag         bool   //debug flag
	cfgFile           string //config file
	asin              string //ASIN code
	isbn              string //ISBN code
	card              string //Aozora-bunko card no.
	tmpltPath         string //template file path
	configFile        = "config"
	defaultConfigPath = config.Path(Name, configFile+".yaml")
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   Name,
		Short: "Search for books data",
		Long:  "Search for books data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugPrint(ui, ecode.ErrNoCommand)
		},
	}
	rootCmd.SetArgs(args)               //arguments of command-line
	rootCmd.SetIn(ui.Reader())          //Stdin
	rootCmd.SetOutput(ui.ErrorWriter()) //Stdout and Stderr
	rootCmd.AddCommand(newVersionCmd(ui))
	rootCmd.AddCommand(newSearchCmd(ui))
	rootCmd.AddCommand(newReviewCmd(ui))
	rootCmd.AddCommand(newHistroyCmd(ui))

	//global options
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("Config file (default %v)", defaultConfigPath))
	rootCmd.PersistentFlags().StringP("review-log", "l", "", "Config: Review log file (JSON format)")
	rootCmd.PersistentFlags().StringP("marketplace", "", "webservices.amazon.co.jp", "Config: PA-API Marketplace")
	rootCmd.PersistentFlags().StringP("associate-tag", "", "", "Config: PA-API Associate Tag")
	rootCmd.PersistentFlags().StringP("access-key", "", "", "Config: PA-API Access Key ID")
	rootCmd.PersistentFlags().StringP("secret-key", "", "", "Config: PA-API Secret Access Key")

	//Bind config file
	_ = viper.BindPFlag("review-log", rootCmd.PersistentFlags().Lookup("review-log"))
	_ = viper.BindPFlag("marketplace", rootCmd.PersistentFlags().Lookup("marketplace"))
	_ = viper.BindPFlag("associate-tag", rootCmd.PersistentFlags().Lookup("associate-tag"))
	_ = viper.BindPFlag("access-key", rootCmd.PersistentFlags().Lookup("access-key"))
	_ = viper.BindPFlag("secret-key", rootCmd.PersistentFlags().Lookup("secret-key"))
	cobra.OnInitialize(initConfig)

	//global options (others)
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "for debug")
	rootCmd.PersistentFlags().StringVarP(&asin, "asin", "a", "", "Amazon ASIN code")
	rootCmd.PersistentFlags().StringVarP(&isbn, "isbn", "i", "", "ISBN code")
	rootCmd.PersistentFlags().StringVarP(&card, "aozora-card", "c", "", "Aozora-bunko card no.")
	rootCmd.PersistentFlags().StringVarP(&tmpltPath, "template-file", "t", "", "Template file for formatted output")

	return rootCmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find config directory.
		confDir := config.Dir(Name)
		if len(confDir) == 0 {
			confDir = "." //current directory
		}
		// Search config in home directory with name ".books-data.yaml" (without extension).
		viper.AddConfigPath(confDir)
		viper.SetConfigName(configFile)
	}
	viper.AutomaticEnv()     // read in environment variables that match
	_ = viper.ReadInConfig() // If a config file is found, read it in.
}

func debugPrint(ui *rwi.RWI, err error) error {
	if debugFlag && err != nil {
		fmt.Fprintf(ui.ErrorWriter(), "%+v\n", err)
		return nil
	}
	return err
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
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
