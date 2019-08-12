package facade

import (
	"bytes"
	"io"

	"github.com/spiegel-im-spiegel/books-data/api/aozoraapi"
	"github.com/spiegel-im-spiegel/books-data/entity"
)

func searchAozoraAPI(id string, rawFlag bool) (io.Reader, error) {
	aozora := aozoraapi.New()
	if rawFlag {
		return aozora.LookupRawData(id)
	}
	book, err := aozora.LookupBook(id)
	if err != nil {
		return nil, err
	}
	b, err := book.Format(tmpltPath)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func findAozoraAPI(id string) (*entity.Book, error) {
	return aozoraapi.New().LookupBook(id)
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
