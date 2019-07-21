package review

import (
	"bytes"
	"text/template"

	"github.com/spiegel-im-spiegel/books-data/errs"
)

//Format returns Formatted text by template
func Format(obj interface{}, tmplt string) (*bytes.Buffer, error) {
	t, err := template.New("Formatting").Parse(tmplt)
	if err != nil {
		return nil, errs.Wrap(err, "error in review.Format() function")
	}
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, obj); err != nil {
		return buf, errs.Wrap(err, "error in review.Format() function")
	}
	return buf, nil
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
