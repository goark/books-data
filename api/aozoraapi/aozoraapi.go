package aozoraapi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/goark/aozora-api"
	"github.com/goark/books-data/api"
	"github.com/goark/books-data/entity"
	"github.com/goark/books-data/entity/values"
	"github.com/goark/errs"
)

//OpenBD is a api.API class for openBD API
type AozoraAPI struct {
	svcType api.ServiceType //Service Type
	server  *aozora.Server  //server info.
}

var _ api.API = (*AozoraAPI)(nil) //AozoraAPI is compatible with api.API interface

//New returns OpenBD instance
func New() *AozoraAPI {
	return &AozoraAPI{svcType: api.TypeAozoraAPI, server: aozora.New()}
}

//Name returns name of API
func (a *AozoraAPI) Name() string {
	return a.svcType.String()
}

//LookupRawData returns openBD raw data
func (a *AozoraAPI) LookupRawData(ctx context.Context, id string) (io.Reader, error) {
	bookId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errs.New(
			fmt.Sprintf("invalid book id: %v", id),
			errs.WithContext("id", id),
		)
	}
	b, err := a.server.CreateClient().LookupBookRawContext(ctx, bookId)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("id", id))
	}
	return bytes.NewReader(b), nil
}

//LookupBook returns Book data from openBD
func (a *AozoraAPI) LookupBook(ctx context.Context, id string) (*entity.Book, error) {
	bookId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errs.New(
			fmt.Sprintf("invalid book id: %v", id),
			errs.WithCause(err),
			errs.WithContext("id", id),
		)
	}
	bk, err := a.server.CreateClient().LookupBookContext(ctx, bookId)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("id", id))
	}

	book := &entity.Book{
		Type:            a.Name(),
		ID:              strconv.Itoa(bk.BookID),
		Title:           bk.Title,
		SubTitle:        bk.Subtitle,
		OriginalTitle:   bk.OriginalTitle,
		URL:             bk.CardURL,
		ProductType:     "青空文庫",
		Codes:           []entity.Code{{Name: "図書カードNo.", Value: strconv.Itoa(bk.BookID)}},
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

/* Copyright 2019-2021 Spiegel
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
