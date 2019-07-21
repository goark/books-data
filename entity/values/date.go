package values

import (
	"fmt"
	"strings"
	"time"

	"github.com/spiegel-im-spiegel/books-data/errs"
)

//Time is wrapper class of time.Time
type Date struct {
	time.Time
}

//NewDate returns Time instance
func NewDate(tm time.Time) Date {
	return Date{tm}
}

//UnmarshalJSON returns result of Unmarshal for json.Unmarshal()
func (t *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	tm, err := time.Parse(time.RFC3339, s+"T09:00:00Z")
	if err != nil {
		tm, err = time.Parse(time.RFC3339, s)
		if err != nil {
			return errs.Wrap(err, "error in unmarshaling JSON to Date type")
		}
	}
	*t = Date{tm}
	return nil
}

//MarshalJSON returns time string with RFC3339 format
func (t *Date) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", t)), nil
}

func (t Date) String() string {
	return t.Format("2006-01-02")
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
