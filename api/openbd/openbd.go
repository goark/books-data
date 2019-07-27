package openbd

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"

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

//ServerOptFunc is self-referential function for functional options pattern
type ServerOptFunc func(*Server)

//New returns OpenBD instance
func New(cmd CommandType, opts ...ServerOptFunc) api.API {
	sv := &Server{cmd: cmd}
	for _, opt := range opts {
		opt(sv)
	}
	return &OpenBD{server: sv}
}

//lookupJSON returns JSON data from openBD
func (api *OpenBD) lookupJSON(id string) ([]byte, error) {
	v := url.Values{
		"isbn": {id},
	}
	resp, err := api.server.CreateClient().Get(v)
	if err != nil {
		return nil, errs.Wrap(err, "error in OpenBD.lookupJSON() function")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, errs.Wrap(err, "error in OpenBD.lookupJSON() function")
	}
	return body, nil
}

//unmarshalJSON returns unmarshalled JSON data
func unmarshalJSON(jsondata []byte) (*obd.Book, error) {
	books := []obd.Book{}
	if err := json.Unmarshal(jsondata, &books); err != nil {
		return nil, errs.Wrap(err, "error in OpenBD.unmarshalJSON() function")
	}
	if len(books) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData, "error in OpenBD.unmarshalJSON() function")
	}
	return &books[0], nil
}

///LookupRawData returns openBD raw data
func (api *OpenBD) LookupRawData(id string) (io.Reader, error) {
	res, err := api.lookupJSON(id)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(res), nil
}

///LookupBook returns Book data from openBD
func (api *OpenBD) LookupBook(id string) (*entity.Book, error) {
	data, err := api.lookupJSON(id)
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
		ID:          bd.GetISBN(),
		Title:       bd.GetTitle(),
		SeriesTitle: bd.GetSeries(),
		Image: struct {
			URL    string
			Height uint16
			Width  uint16
		}{
			URL: bd.GetImageLink(),
		},
		ProductType: "Book",
		Codes:       []entity.Code{entity.Code{Name: "ISBN", Value: bd.GetISBN()}},
		Authors:     []string{bd.GetAuthor()},
		Publisher:   bd.GetPublisher(),
		Service: struct {
			Name string
			URL  string
		}{Name: "openBD", URL: "https://openbd.jp/"},
	}
	if tm, err := bd.GetPubdate(); err == nil {
		book.PublicationDate = values.NewDate(tm)
	}

	return book, nil
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
