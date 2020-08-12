package facade

import (
	"context"
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
	"github.com/spiegel-im-spiegel/gocli/signal"
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
				return errs.New("--pipe", errs.WithCause(err))
			}
			//Ratins
			rating, err := cmd.Flags().GetInt("rating")
			if err != nil {
				return errs.New("--rating", errs.WithCause(err))
			}
			//Date of review
			dt, err := cmd.Flags().GetString("review-date")
			if err != nil {
				return errs.New("--review-date", errs.WithCause(err))
			}
			//URL of book page
			bookpage, err := cmd.Flags().GetString("bookpage-url")
			if err != nil {
				return errs.New("--bookpage-url", errs.WithCause(err))
			}
			//URL of book cover image
			bookcover, err := cmd.Flags().GetString("image-url")
			if err != nil {
				return errs.New("--image-url", errs.WithCause(err))
			}
			//Description
			desc := ""
			if pipeFlag {
				w := &strings.Builder{}
				if _, err := io.Copy(w, ui.Reader()); err != nil {
					return debugPrint(ui, errs.New("Cannot read Stdin", errs.WithCause(err)))
				}
				desc = w.String()
			} else if len(args) > 1 {
				return errs.Wrap(os.ErrInvalid, errs.WithContext("args", strings.Join(args, ",")))
			} else if len(args) == 1 {
				desc = args[0]
			}

			//Create context
			ctx := signal.Context(context.Background(), os.Interrupt)

			err = nil
			var bk *entity.Book = nil
			//Search by ASIN code
			if len(asin) > 0 {
				p, err := getPaapiParams()
				if err != nil {
					return debugPrint(ui, err)
				}
				bk, err = findPAAPI(ctx, asin, p)
				if err != nil {
					if !checkError(err) {
						return debugPrint(ui, err)
					}
				}
			}
			if bk == nil && len(isbn) > 0 {
				//by openBD
				bk, err = findOpenBD(ctx, isbn)
				if err != nil {
					if !checkError(err) {
						return debugPrint(ui, err)
					}
				}
			}
			if bk == nil && len(card) > 0 {
				//by Aozora-API
				bk, err = findAozoraAPI(ctx, card)
				if err != nil {
					if !checkError(err) {
						return debugPrint(ui, err)
					}
				}
			}
			if err == nil && bk == nil {
				err = errs.New("error in review command", errs.WithCause(ecode.ErrNoData))
			}
			if err != nil {
				return debugPrint(ui, err)
			}

			rev := review.New(
				bk,
				review.WithDate(dt),
				review.WithRating(rating),
				review.WithDescription(desc),
				review.WithBookPage(bookpage),
				review.WithBookCover(bookcover),
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
	reviewCmd.Flags().StringP("bookpage-url", "", "", "URL of book page")
	reviewCmd.Flags().StringP("image-url", "", "", "URL of book cover image")
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
		return errs.Wrap(err)
	}
	revs = revs.Append(rev)
	if err := revs.ExportJSONFile(path); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

/* Copyright 2019,2020 Spiegel
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
