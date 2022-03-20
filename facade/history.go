package facade

import (
	"github.com/goark/books-data/api"
	"github.com/goark/books-data/ecode"
	"github.com/goark/books-data/review"
	"github.com/goark/books-data/review/logger"
	"github.com/goark/errs"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//newSearchCmd returns cobra.Command instance for show sub-command
func newHistroyCmd(ui *rwi.RWI) *cobra.Command {
	historyCmd := &cobra.Command{
		Use:   "history [flags]",
		Short: "Lookup review data from history log",
		Long:  "Lookup review data from history log",
		RunE: func(cmd *cobra.Command, args []string) error {
			logf := viper.GetString("review-log")
			if len(logf) == 0 {
				return debugPrint(ui, errs.New("error in history command", errs.WithCause(ecode.ErrNoData)))
			}
			revs, err := logger.ImportJSONFile(logf)
			if err != nil {
				return debugPrint(ui, err)
			}
			if len(revs) == 0 {
				return debugPrint(ui, errs.New("error in history command", errs.WithCause(ecode.ErrNoData)))
			}

			var rev *review.Review = nil
			//Search by ASIN code
			if len(asin) > 0 {
				rev = revs.Find(api.TypePAAPI.String(), asin)
			}
			//Search by ISBN code
			if rev == nil && len(isbn) > 0 {
				rev = revs.Find(api.TypeOpenBD.String(), isbn)
			}
			//Search by Aozora-bunko card no.
			if rev == nil && len(card) > 0 {
				rev = revs.Find(api.TypeAozoraAPI.String(), card)
			}
			if rev == nil {
				return debugPrint(ui, errs.New("error in history command", errs.WithCause(ecode.ErrNoData)))
			}

			b, err := rev.Format(tmpltPath)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.Output(string(b)))
		},
	}
	//options

	return historyCmd
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
