package facade

import (
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/books-data/api/openbd"
	"github.com/spiegel-im-spiegel/books-data/errs"
	"github.com/spiegel-im-spiegel/books-data/review"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newpaapiCmd returns cobra.Command instance for show sub-command
func newOpenBDCmd(ui *rwi.RWI) *cobra.Command {
	openBDCmd := &cobra.Command{
		Use:   "openbd",
		Short: "Search for books data by openBD",
		Long:  "Search for books data by openBD",
		RunE: func(cmd *cobra.Command, args []string) error {
			//ASIN code
			id, err := cmd.Flags().GetString("isbn")
			if err != nil {
				return errs.Wrap(err, "--isbn")
			}
			if len(id) == 0 {
				return errs.Wrap(os.ErrInvalid, "No ISBN code")
			}
			//Ratins
			rating, err := cmd.Flags().GetInt("rating")
			if err != nil {
				return errors.Wrap(err, "--rating")
			}
			//Date of review
			dt, err := cmd.Flags().GetString("review-date")
			if err != nil {
				return errors.Wrap(err, "--review-date")
			}
			//Template data
			tf, err := cmd.Flags().GetString("template")
			if err != nil {
				return errors.Wrap(err, "--template")
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
				return errors.Wrap(os.ErrInvalid, strings.Join(args, " "))
			} else if len(args) == 1 {
				desc = args[0]
			}

			//Creatte API
			obdapi := openbd.New(openbd.GET)
			if rawFlag {
				res, err := obdapi.LookupRawData(id)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.WriteFrom(res))
			}
			book, err := obdapi.LookupBook(id)
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
	//parameters for review
	openBDCmd.Flags().StringP("isbn", "i", "", "ISBN code")
	openBDCmd.Flags().IntP("rating", "r", 0, "Rating of product")
	openBDCmd.Flags().StringP("review-date", "", "", "Date of review")
	openBDCmd.Flags().StringP("template", "t", "", "Template file")

	return openBDCmd
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
