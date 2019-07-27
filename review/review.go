package review

import (
	"encoding/json"
	"io"
	"strings"
	"time"

	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/books-data/entity"
	"github.com/spiegel-im-spiegel/books-data/entity/values"
	"github.com/spiegel-im-spiegel/errs"
)

const MAX_STAR = 5

//Review is entity class for review info.
type Review struct {
	Book        *entity.Book
	Date        values.Date
	Rating      int
	Star        [MAX_STAR]bool
	Description string `json:",omitempty"`
}

//LookupOptFunc is self-referential function for functional options pattern
type ReviewOptFunc func(*Review)

func New(book *entity.Book, opts ...ReviewOptFunc) *Review {
	rev := &Review{Book: book, Date: values.NewDate(time.Now())}
	for _, opt := range opts {
		opt(rev)
	}
	return rev
}

//WithDate returns function for setting IdType
func WithDate(s string) ReviewOptFunc {
	return func(rev *Review) {
		if rev != nil && len(s) > 0 {
			if tm, err := time.Parse(time.RFC3339, s+"T09:00:00Z"); err == nil {
				rev.Date = values.NewDate(tm)
			}
		}
	}
}

//WithRating returns function for setting IdType
func WithRating(r int) ReviewOptFunc {
	return func(rev *Review) {
		if rev != nil {
			if r < 0 {
				r = 0
			} else if r > MAX_STAR {
				r = MAX_STAR
			}
			rev.Rating = r
			for i := 0; i < MAX_STAR; i++ {
				if r > i {
					rev.Star[i] = true
				} else {
					rev.Star[i] = false
				}
			}
		}
	}
}

//WithDate returns function for setting IdType
func WithDescription(s string) ReviewOptFunc {
	return func(rev *Review) {
		if rev != nil {
			rev.Description = s
		}
	}
}

func (r *Review) Format(tr io.Reader) ([]byte, error) {
	if r == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer, "error in review.Review.Format() function")
	}
	if tr == nil {
		b, err := json.Marshal(r)
		if err != nil {
			return nil, errs.Wrap(err, "error in review.Review.Format() function")
		}
		return b, nil
	}
	tmplt := &strings.Builder{}
	if _, err := io.Copy(tmplt, tr); err != nil {
		return nil, errs.Wrap(err, "error in review.Review.Format() function")
	}
	buf, err := Format(r, tmplt.String())
	if err != nil {
		return buf.Bytes(), errs.Wrap(err, "error in review.Review.Format() function")
	}
	return buf.Bytes(), nil
}

func (r *Review) String() string {
	b, err := r.Format(nil)
	if err != nil {
		return ""
	}
	return string(b)
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
