package facade

import (
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newSearchCmd returns cobra.Command instance for show sub-command
func newSearchCmd(ui *rwi.RWI) *cobra.Command {
	searchCmd := &cobra.Command{
		Use:   "search [flags]",
		Short: "Search for books data",
		Long:  "Search for books data",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Raw flag (PA-API)
			rawFlag, err := cmd.Flags().GetBool("raw")
			if err != nil {
				return errs.Wrap(err, "--raw")
			}

			var lastError error
			//Search by ASIN code
			if len(asin) > 0 {
				r, err := searchPAAPI(asin, false, rawFlag)
				if err == nil {
					return debugPrint(ui, ui.WriteFrom(r))
				}
				if !checkError(err) {
					return debugPrint(ui, err)
				}
				lastError = err
			}
			//Search by ISBN code
			if len(isbn) > 0 {
				//by PA-API
				r, err := searchPAAPI(isbn, true, rawFlag)
				if err == nil {
					return debugPrint(ui, ui.WriteFrom(r))
				}
				if !checkError(err) {
					return debugPrint(ui, err)
				}
				//by openBD
				r, err = searchOpenBD(isbn, rawFlag)
				if err == nil {
					return debugPrint(ui, ui.WriteFrom(r))
				}
				if !checkError(err) {
					return debugPrint(ui, err)
				}
				lastError = err
			}
			if len(card) > 0 {
				//by Aozora-API
				r, err := searchAozoraAPI(card, rawFlag)
				if err == nil {
					return debugPrint(ui, ui.WriteFrom(r))
				}
				if !checkError(err) {
					return debugPrint(ui, err)
				}
				lastError = err
			}
			return debugPrint(ui, lastError)
		},
	}
	//options
	searchCmd.Flags().BoolP("raw", "", false, "Output raw data from API")

	return searchCmd
}

func checkError(err error) bool {
	switch true {
	case errs.Is(err, ecode.ErrInvalidAPIParameter):
	case errs.Is(err, ecode.ErrInvalidAPIResponse):
	case errs.Is(err, ecode.ErrNoData):
	default:
		return false
	}
	return true
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
