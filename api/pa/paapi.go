package pa

import (
	"encoding/xml"
	"io"
	"time"

	amazonproduct "github.com/DDRBoxman/go-amazon-product-api"
	"github.com/spiegel-im-spiegel/books-data/api"
	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/books-data/entity"
	"github.com/spiegel-im-spiegel/books-data/entity/values"
	"github.com/spiegel-im-spiegel/errs"
)

//PaAPI is class of PA-API
type PaAPI struct {
	svcType api.ServiceType //Service Type
	server  *Server         //server info.
}

//PaAPIOptFunc is self-referential function for functional options pattern
type PaAPIOptFunc func(*Server)

//New returns PaAPI instance
func New(opts ...PaAPIOptFunc) api.API {
	server := &Server{}
	for _, opt := range opts {
		opt(server)
	}
	return &PaAPI{svcType: api.TypePAAPI, server: server}
}

//WithMarketplace returns function for setting Marketplace
func WithMarketplace(mp string) PaAPIOptFunc {
	return func(s *Server) {
		if s != nil {
			s.marketplace = mp
		}
	}
}

//WithAssociateTag returns function for setting Associate Tag
func WithAssociateTag(tag string) PaAPIOptFunc {
	return func(s *Server) {
		if s != nil {
			s.associateTag = tag
		}
	}
}

//WithAccessKey returns function for setting Access Key
func WithAccessKey(key string) PaAPIOptFunc {
	return func(s *Server) {
		if s != nil {
			s.accessKey = key
		}
	}
}

//WithSecretKey returns function for setting Secret Access Key
func WithSecretKey(key string) PaAPIOptFunc {
	return func(s *Server) {
		if s != nil {
			s.secretKey = key
		}
	}
}

//WithSecretKey returns function for setting Secret Access Key
func WithEnableISBN(isbn bool) PaAPIOptFunc {
	return func(s *Server) {
		if s != nil {
			s.enableISBN = isbn
		}
	}
}

//Name returns name of API
func (a *PaAPI) Name() string {
	return a.svcType.String()
}

///LookupRawData returns PA-API raw data
func (a *PaAPI) LookupRawData(id string) (io.Reader, error) {
	return a.server.CreateClient().LookupXML(id)
}

///LookupBook returns Book data from PA-API
func (a *PaAPI) LookupBook(id string) (*entity.Book, error) {
	data, err := a.LookupRawData(id)
	if err != nil {
		return nil, errs.Wrap(err, "error in PaAPI.LookupBook() function")
	}
	res, err := unmarshalXML(data)
	if err != nil {
		return nil, errs.Wrap(err, "error in PaAPI.LookupBook() function")
	}
	if !res.Items.Request.IsValid {
		return nil, errs.Wrap(ecode.ErrInvalidAPIResponse, "error in PaAPI.LookupBook() function")
	}
	if len(res.Items.Item) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData, "error in PaAPI.LookupBook() function")
	}
	item := res.Items.Item[0]
	book := &entity.Book{
		Type:  a.Name(),
		ID:    item.ASIN,
		Title: item.ItemAttributes.Title,
		URL:   item.DetailPageURL,
		Image: entity.BookCover{
			URL:    item.MediumImage.URL,
			Height: item.MediumImage.Height,
			Width:  item.MediumImage.Width,
		},
		ProductType: item.ItemAttributes.Binding,
		Codes:       []entity.Code{entity.Code{Name: "ASIN", Value: item.ASIN}},
		Authors:     item.ItemAttributes.Author,
		Publisher:   item.ItemAttributes.Publisher,
		Service:     entity.Service{Name: "PA-API", URL: "https://affiliate.amazon.co.jp/assoc_credentials/home"},
	}
	if len(item.ItemAttributes.Creator) > 0 {
		book.Creators = []entity.Creator{}
		for _, c := range item.ItemAttributes.Creator {
			book.Creators = append(book.Creators, entity.Creator{Name: c.Value, Role: c.Role})
		}
	}
	if len(item.ItemAttributes.EAN) > 0 {
		book.Codes = append(book.Codes, entity.Code{Name: "EAN", Value: item.ItemAttributes.EAN})
	}
	if len(item.ItemAttributes.EISBN) > 0 {
		book.Codes = append(book.Codes, entity.Code{Name: "EISBN", Value: item.ItemAttributes.EISBN})
	}
	if len(item.ItemAttributes.PublicationDate) > 0 {
		if tm, err := time.Parse("2006-01-02", item.ItemAttributes.PublicationDate); err == nil {
			book.PublicationDate = values.NewDate(tm)
		}
	}
	if len(item.ItemAttributes.ReleaseDate) > 0 {
		if tm, err := time.Parse("2006-01-02", item.ItemAttributes.ReleaseDate); err == nil {
			book.LastRelease = values.NewDate(tm)
		}
	}
	return book, nil
}

//unmarshalXML returns unmarshalled XML data
func unmarshalXML(xmldata io.Reader) (*amazonproduct.ItemLookupResponse, error) {
	res := &amazonproduct.ItemLookupResponse{}
	if err := xml.NewDecoder(xmldata).Decode(res); err != nil {
		return nil, errs.Wrap(err, "error in unmarshalXML() function")
	}
	return res, nil
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
