package facade

import (
	"fmt"
	"runtime"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/books-data/review"
	"github.com/spiegel-im-spiegel/books-data/review/logger"
	"github.com/spiegel-im-spiegel/errs"
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
	debugFlag bool //debug flag
	rawFlag   bool //raw flag
	pipeFlag  bool //pipe flag
)
var (
	cfgFile string //config file
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
	rootCmd.AddCommand(newPaApiCmd(ui))
	rootCmd.AddCommand(newOpenBDCmd(ui))

	//global options in config file
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default $HOME/.books-data.yaml)")
	rootCmd.PersistentFlags().StringP("template-file", "t", "", "Template file for formatted output")
	rootCmd.PersistentFlags().StringP("review-log", "l", "", "Review log file (JSON format)")
	_ = viper.BindPFlag("template-file", rootCmd.PersistentFlags().Lookup("template-file"))
	_ = viper.BindPFlag("review-log", rootCmd.PersistentFlags().Lookup("review-log"))
	cobra.OnInitialize(initConfig)

	//global options (others)
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "for debug")
	rootCmd.PersistentFlags().BoolVarP(&rawFlag, "raw", "", false, "Output raw data from openBD")
	rootCmd.PersistentFlags().BoolVarP(&pipeFlag, "pipe", "", false, "Import description from Stdin")

	return rootCmd
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
		// Search config in home directory with name ".books-data.yaml" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".books-data")
	}
	viper.AutomaticEnv()     // read in environment variables that match
	_ = viper.ReadInConfig() // If a config file is found, read it in.
}

func updateReviewLog(rev *review.Review) error {
	path := viper.GetString("review-log")
	fmt.Println("debug", path)
	if len(path) == 0 {
		return nil
	}
	revs, err := logger.ImportJSONFile(path)
	if err != nil {
		return errs.Wrap(err, "error in facade.updateReviewLog() function")
	}
	revs = revs.Append(rev)
	if err := revs.ExportJSONFile(path); err != nil {
		return errs.Wrap(err, "error in facade.updateReviewLog() function")
	}
	return nil
}

func debugPrint(ui *rwi.RWI, err error) error {
	if debugFlag && err != nil {
		fmt.Fprintf(ui.ErrorWriter(), "Error: %+v\n", err)
		return nil
	}
	return errs.Cause(err)
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
