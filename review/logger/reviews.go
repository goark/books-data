package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/books-data/review"
	"github.com/spiegel-im-spiegel/errs"
)

//Reviews is list of review.Review
type Reviews []review.Review

//ImportJSONFile import Reviews data from file with JSON format
func ImportJSONFile(path string) (Reviews, error) {
	file, err := os.Open(path)
	if err != nil {
		return Reviews{}, nil
	}
	defer file.Close()

	return ImportJSON(file)
}

//ImportJSONFile import Reviews data with JSON format
func ImportJSON(r io.Reader) (Reviews, error) {
	dec := json.NewDecoder(r)
	revs := Reviews{}
	err := dec.Decode(&revs)
	return revs, errs.Wrap(err, "error in logger.ImportJSON() function")
}

//exportJSON returns byte string with JSON format
func (revs Reviews) exportJSON() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	err := enc.Encode(revs)
	return buf, errs.Wrap(err, "error in logger.ReviewLog.exportJSON() function")
}

//ExportJSONFile outputs file with JSON format
func (revs Reviews) ExportJSONFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errs.Wrap(err, "error in logger.ReviewLog.ExportJSONFile() function")
	}
	defer file.Close()

	b, err := revs.exportJSON()
	if err != nil {
		return err
	}
	_, err = io.Copy(file, b)
	return errs.Wrap(err, "error in logger.ReviewLog.ExportJSONFile() function")
}

func (revs Reviews) String() string {
	b, err := revs.exportJSON()
	if err != nil {
		return ""
	}
	return b.String()
}

//Find returns review.Review
func (revs Reviews) Find(svcType, id string) *review.Review {
	for i := 0; i < len(revs); i++ {
		rev := &revs[i]
		if rev.Book.Type == svcType && rev.Book.ID == id {
			return rev
		}
	}
	return nil
}

//Append adds new review.Review data in Reviews list
func (revs Reviews) Append(rev *review.Review) Reviews {
	if r := revs.Find(rev.Book.Type, rev.Book.ID); r != nil {
		r.CopyFrom(rev)
	} else {
		revs = append(revs, *rev)
	}
	return revs
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
