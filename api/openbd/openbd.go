package openbd

import (
	"bytes"
	"encoding/json"
	"io"

	obd "github.com/seihmd/openbd"
	"github.com/spiegel-im-spiegel/books-data/api"
	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/books-data/entity"
	"github.com/spiegel-im-spiegel/books-data/entity/values"
	"github.com/spiegel-im-spiegel/errs"
)

//OpenBD is a api.API class for openBD API
type OpenBD struct {
	server *Server //server info.
}

//New returns OpenBD instance
func New(cmd CommandType) api.API {
	return &OpenBD{server: &Server{svcType: api.TypeOpenBD, cmd: cmd}}
}

//Name returns name of API
func (a *OpenBD) Name() string {
	return a.server.svcType.String()
}

///LookupRawData returns openBD raw data
func (a *OpenBD) LookupRawData(id string) (io.Reader, error) {
	res, err := a.server.CreateClient().LookupJSON(id)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(res), nil
}

///LookupBook returns Book data from openBD
func (a *OpenBD) LookupBook(id string) (*entity.Book, error) {
	data, err := a.LookupRawData(id)
	if err != nil {
		return nil, errs.Wrap(err, "error in OpenBD.LookupBook() function")
	}
	bd, err := unmarshalJSON(data)
	if err != nil {
		return nil, errs.Wrap(err, "error in OpenBD.LookupBook() function")
	}
	if !bd.IsValidData() {
		return nil, errs.Wrap(ecode.ErrInvalidAPIResponse, "error in OpenBD.LookupBook() function")
	}

	book := &entity.Book{
		Type:        a.Name(),
		ID:          bd.GetISBN(),
		Title:       bd.GetTitle(),
		SeriesTitle: bd.GetSeries(),
		Image: entity.BookCover{
			URL: bd.GetImageLink(),
		},
		ProductType: "Book",
		Codes:       []entity.Code{entity.Code{Name: "ISBN", Value: bd.GetISBN()}},
		Authors:     []string{bd.GetAuthor()},
		Publisher:   bd.GetPublisher(),
		Service:     entity.Service{Name: "openBD", URL: "https://openbd.jp/"},
	}
	if tm, err := bd.GetPubdate(); err == nil {
		book.PublicationDate = values.NewDate(tm)
	}

	return book, nil
}

//unmarshalJSON returns unmarshalled JSON data
func unmarshalJSON(jsondata io.Reader) (*obd.Book, error) {
	books := []obd.Book{}
	if err := json.NewDecoder(jsondata).Decode(&books); err != nil {
		return nil, errs.Wrap(err, "error in OpenBD.unmarshalJSON() function")
	}
	if len(books) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData, "error in OpenBD.unmarshalJSON() function")
	}
	return &books[0], nil
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
