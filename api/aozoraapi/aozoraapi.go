package aozoraapi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/spiegel-im-spiegel/aozora-api"
	"github.com/spiegel-im-spiegel/books-data/api"
	"github.com/spiegel-im-spiegel/books-data/entity"
	"github.com/spiegel-im-spiegel/books-data/entity/values"
	"github.com/spiegel-im-spiegel/errs"
)

//OpenBD is a api.API class for openBD API
type AozoraAPI struct {
	svcType api.ServiceType //Service Type
	server  *aozora.Server  //server info.
	ctx     context.Context //context
}

var _ api.API = (*AozoraAPI)(nil) //AozoraAPI is compatible with api.API interface

//New returns OpenBD instance
func New(ctx context.Context) *AozoraAPI {
	return &AozoraAPI{svcType: api.TypeAozoraAPI, server: aozora.New(), ctx: ctx}
}

//Name returns name of API
func (a *AozoraAPI) Name() string {
	return a.svcType.String()
}

///LookupRawData returns openBD raw data
func (a *AozoraAPI) LookupRawData(id string) (io.Reader, error) {
	bookId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errs.Wrap(
			err,
			fmt.Sprintf("invalid book id: %v", id),
			errs.WithContext("id", id),
		)
	}
	b, err := a.server.CreateClient(aozora.WithContext(a.ctx), aozora.WithHttpClient(&http.Client{})).LookupBookRaw(bookId)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("id", id))
	}
	return bytes.NewReader(b), nil
}

///LookupBook returns Book data from openBD
func (a *AozoraAPI) LookupBook(id string) (*entity.Book, error) {
	bookId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errs.Wrap(
			err,
			fmt.Sprintf("invalid book id: %v", id),
			errs.WithContext("id", id),
		)
	}
	bk, err := a.server.CreateClient(aozora.WithContext(a.ctx), aozora.WithHttpClient(&http.Client{})).LookupBook(bookId)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("id", id))
	}

	book := &entity.Book{
		Type:            a.Name(),
		ID:              strconv.Itoa(bk.BookID),
		Title:           bk.Title,
		SubTitle:        bk.Subtitle,
		OriginalTitle:   bk.OriginalTitle,
		URL:             bk.CardURL,
		ProductType:     "青空文庫",
		Codes:           []entity.Code{entity.Code{Name: "図書カードNo.", Value: strconv.Itoa(bk.BookID)}},
		Creators:        getCreators(bk),
		PublicationDate: values.NewDate(bk.ReleaseDate.Time),
		LastRelease:     values.NewDate(bk.LastModified.Time),
		PublicDomain:    !bk.Copyright,
		FirstAppearance: bk.FirstAppearance,
		Service:         entity.Service{Name: "aozorahack", URL: "https://aozorahack.org/"},
	}

	return book, nil
}

func getCreators(bk *aozora.Book) []entity.Creator {
	creators := []entity.Creator{}
	if bk == nil {
		return creators
	}
	if len(bk.Authors) > 0 {
		for _, a := range bk.Authors {
			creators = append(creators, entity.Creator{Name: fmt.Sprintf("%v %v", a.LastName, a.FirstName)})
		}
	}
	if len(bk.Translators) > 0 {
		for _, t := range bk.Translators {
			creators = append(creators, entity.Creator{Name: fmt.Sprintf("%v %v", t.LastName, t.FirstName), Role: "翻訳"})
		}
	}
	return creators
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
