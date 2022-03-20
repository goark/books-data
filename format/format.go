package format

import (
	"bytes"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/goark/errs"
)

func ByTemplateFile(obj interface{}, tmpltPath string) (*bytes.Buffer, error) {
	file, err := os.Open(tmpltPath)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("tmpltPath", tmpltPath))
	}
	defer file.Close()

	tmplt := &strings.Builder{}
	if _, err := io.Copy(tmplt, file); err != nil {
		return nil, errs.Wrap(err, errs.WithContext("tmpltPath", tmpltPath))
	}
	return Do(obj, tmplt.String())
}

//Do returns Formatted text by template
func Do(obj interface{}, tmplt string) (*bytes.Buffer, error) {
	t, err := template.New("Formatting").Parse(tmplt)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("tmplt", tmplt))
	}
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, obj); err != nil {
		return buf, errs.Wrap(err, errs.WithContext("tmplt", tmplt))
	}
	return buf, nil
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
