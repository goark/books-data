package facade

import (
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/books-data/entity"
	"github.com/spiegel-im-spiegel/books-data/review"
	"github.com/spiegel-im-spiegel/books-data/review/logger"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newSearchCmd returns cobra.Command instance for show sub-command
func newReviewCmd(ui *rwi.RWI) *cobra.Command {
	reviewCmd := &cobra.Command{
		Use:   "review [flags] [description]",
		Short: "Make review data",
		Long:  "Make review data",
		RunE: func(cmd *cobra.Command, args []string) error {
			//pipe flag
			pipeFlag, err := cmd.Flags().GetBool("pipe")
			if err != nil {
				return errs.Wrap(err, "--pipe")
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

			err = nil
			var bk *entity.Book = nil
			//Search by ASIN code
			if len(asin) > 0 {
				bk, err = findPAAPI(asin, false)
				if err != nil {
					if !checkError(err) {
						return debugPrint(ui, err)
					}
				}
			}
			//Search by ISBN code
			if bk == nil && len(isbn) > 0 {
				//by PA-API
				bk, err = findPAAPI(isbn, true)
				if err != nil {
					if !checkError(err) {
						return debugPrint(ui, err)
					}
				}
			}
			if bk == nil && len(isbn) > 0 {
				//by openBD
				bk, err = findOpenBD(isbn)
				if err != nil {
					if !checkError(err) {
						return debugPrint(ui, err)
					}
				}
			}
			if err == nil && bk == nil {
				err = errs.Wrap(ecode.ErrNoData, "error in facade.reviewCmd")
			}
			if err != nil {
				return debugPrint(ui, err)
			}

			rev := review.New(
				bk,
				review.WithDate(dt),
				review.WithRating(rating),
				review.WithDescription(desc),
			)
			if err := updateReviewLog(rev); err != nil {
				return debugPrint(ui, err)
			}
			b, err := rev.Format(tmpltPath)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.Output(string(b)))
		},
	}
	//options
	reviewCmd.Flags().IntP("rating", "r", 0, "Rating of product")
	reviewCmd.Flags().StringP("review-date", "", "", "Date of review")
	reviewCmd.Flags().BoolP("pipe", "", false, "Import description from Stdin")

	return reviewCmd
}

func updateReviewLog(rev *review.Review) error {
	path := viper.GetString("review-log")
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
