package openbd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spiegel-im-spiegel/books-data/api"
	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/books-data/entity"
	"github.com/spiegel-im-spiegel/books-data/entity/values"
	"github.com/spiegel-im-spiegel/errs"
	obd "github.com/spiegel-im-spiegel/openbd-api"
)

//OpenBD is a api.API class for openBD API
type OpenBD struct {
	svcType api.ServiceType //Service Type
	server  *obd.Server     //server info.
	ctx     context.Context //context
}

var _ api.API = (*OpenBD)(nil) //OpenBD is compatible with api.API interface

//New returns OpenBD instance
func New(ctx context.Context) *OpenBD {
	return &OpenBD{svcType: api.TypeOpenBD, server: obd.New(), ctx: ctx}
}

//Name returns name of API
func (a *OpenBD) Name() string {
	return a.svcType.String()
}

///LookupRawData returns openBD raw data
func (a *OpenBD) LookupRawData(id string) (io.Reader, error) {
	res, err := a.server.CreateClient(obd.WithContext(a.ctx), obd.WithHttpClient(&http.Client{})).LookupBooksRaw([]string{id})
	if err != nil {
		return nil, errs.Wrap(
			err,
			fmt.Sprintf("invalid book id: %v", id),
			errs.WithContext("id", id),
		)
	}
	return bytes.NewReader(res), nil
}

///LookupBook returns Book data from openBD
func (a *OpenBD) LookupBook(id string) (*entity.Book, error) {
	data, err := a.LookupRawData(id)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("id", id))
	}
	bd, err := unmarshalJSON(data)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("id", id))
	}
	if !bd.IsValid() {
		return nil, errs.Wrap(ecode.ErrInvalidAPIResponse, "", errs.WithContext("id", id))
	}

	book := &entity.Book{
		Type:        a.Name(),
		ID:          bd.Id(),
		Title:       bd.Title(),
		SeriesTitle: bd.SeriesTitle(),
		Image: entity.BookCover{
			URL: bd.ImageURL(),
		},
		ProductType:     "Book",
		Codes:           []entity.Code{entity.Code{Name: "ISBN", Value: bd.ISBN()}},
		Creators:        getCreators(bd),
		Publisher:       bd.Publisher(),
		PublicationDate: values.NewDate(bd.PublicationDate().Time),
		Service:         entity.Service{Name: "openBD", URL: "https://openbd.jp/"},
	}
	return book, nil
}

//getCreators returns list of creators
func getCreators(bd *obd.Book) []entity.Creator {
	creators := []entity.Creator{}
	if bd == nil {
		return creators
	}
	for _, author := range bd.Authors() {
		creators = append(creators, entity.Creator{Name: author})
	}
	return creators
}

//unmarshalJSON returns unmarshalled JSON data
func unmarshalJSON(jsondata io.Reader) (*obd.Book, error) {
	books := []obd.Book{}
	if err := json.NewDecoder(jsondata).Decode(&books); err != nil {
		return nil, errs.Wrap(err, "")
	}
	if len(books) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData, "")
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
