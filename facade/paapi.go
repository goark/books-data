package facade

import (
	"io"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/books-data/api/pa"
	"github.com/spiegel-im-spiegel/books-data/review"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	cfgFile string //config file
)

//newpaapiCmd returns cobra.Command instance for show sub-command
func newPaApiCmd(ui *rwi.RWI) *cobra.Command {
	paapiCmd := &cobra.Command{
		Use:   "paapi [flags] [description]",
		Short: "Search for books data by PA-API",
		Long:  "Search for books data by PA-API",
		RunE: func(cmd *cobra.Command, args []string) error {
			//ASIN code
			id, err := cmd.Flags().GetString("asin")
			if err != nil {
				return errs.Wrap(err, "--asin")
			}
			if len(id) == 0 {
				return errs.Wrap(os.ErrInvalid, "No ASIN code")
			}
			//Ratins
			rating, err := cmd.Flags().GetInt("rating")
			if err != nil {
				return errs.Wrap(err, "--rating")
			}
			//Date of review
			dt, err := cmd.Flags().GetString("review-date")
			if err != nil {
				return errs.Wrap(err, "--review-date")
			}
			//Template data
			tf, err := cmd.Flags().GetString("template")
			if err != nil {
				return errs.Wrap(err, "--template")
			}
			var tr io.Reader
			if len(tf) > 0 {
				file, err := os.Open(tf)
				if err != nil {
					return err
				}
				defer file.Close()
				tr = file
			}
			//Description
			desc := ""
			if pipeFlag {
				w := &strings.Builder{}
				if _, err := io.Copy(w, ui.Reader()); err != nil {
					return debugPrint(ui, errs.Wrap(err, "Cannot read Stdin"))
				}
				desc = w.String()
			} else if len(args) > 1 {
				return errs.Wrap(os.ErrInvalid, strings.Join(args, " "))
			} else if len(args) == 1 {
				desc = args[0]
			}

			//Creatte API
			paapi := pa.New(
				pa.WithMarketplace(viper.GetString("marketplace")),
				pa.WithAssociateTag(viper.GetString("associate-tag")),
				pa.WithAccessKey(viper.GetString("access-key")),
				pa.WithSecretKey(viper.GetString("secret-key")),
			)
			if rawFlag {
				res, err := paapi.LookupRawData(id)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.WriteFrom(res))
			}
			book, err := paapi.LookupBook(id)
			if err != nil {
				return debugPrint(ui, err)
			}
			b, err := review.New(
				book,
				review.WithDate(dt),
				review.WithRating(rating),
				review.WithDescription(desc),
			).Format(tr)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.Output(string(b)))
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

	//parameters for review
	paapiCmd.Flags().StringP("asin", "a", "", "Amazon ASIN code")
	paapiCmd.Flags().IntP("rating", "r", 0, "Rating of product")
	paapiCmd.Flags().StringP("review-date", "", "", "Date of review")
	paapiCmd.Flags().StringP("template", "t", "", "Template file")

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
