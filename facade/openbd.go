package facade

import (
	"bytes"
	"context"
	"io"

	"github.com/spiegel-im-spiegel/books-data/api/openbd"
	"github.com/spiegel-im-spiegel/books-data/entity"
)

func searchOpenBD(ctx context.Context, id string, rawFlag bool) (io.Reader, error) {
	obdapi := openbd.New(ctx)
	if rawFlag {
		return obdapi.LookupRawData(id)
	}
	book, err := obdapi.LookupBook(id)
	if err != nil {
		return nil, err
	}
	b, err := book.Format(tmpltPath)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func findOpenBD(ctx context.Context, id string) (*entity.Book, error) {
	return openbd.New(ctx).LookupBook(id)
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
