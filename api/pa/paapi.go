package pa

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
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
	paapi *amazonproduct.AmazonProductAPI
}

//PaAPIOptFunc is self-referential function for functional options pattern
type PaAPIOptFunc func(*amazonproduct.AmazonProductAPI)

//New returns PaAPI instance
func New(opts ...PaAPIOptFunc) api.API {
	paapi := &amazonproduct.AmazonProductAPI{Client: &http.Client{}}
	for _, opt := range opts {
		opt(paapi)
	}
	return &PaAPI{paapi: paapi}
}

//WithMarketplace returns function for setting Marketplace
func WithMarketplace(mp string) PaAPIOptFunc {
	return func(api *amazonproduct.AmazonProductAPI) {
		if api != nil {
			api.Host = mp
		}
	}
}

//WithAssociateTag returns function for setting Associate Tag
func WithAssociateTag(tag string) PaAPIOptFunc {
	return func(api *amazonproduct.AmazonProductAPI) {
		if api != nil {
			api.AssociateTag = tag
		}
	}
}

//WithAccessKey returns function for setting Access Key
func WithAccessKey(key string) PaAPIOptFunc {
	return func(api *amazonproduct.AmazonProductAPI) {
		if api != nil {
			api.AccessKey = key
		}
	}
}

//WithSecretKey returns function for setting Secret Access Key
func WithSecretKey(key string) PaAPIOptFunc {
	return func(api *amazonproduct.AmazonProductAPI) {
		if api != nil {
			api.SecretKey = key
		}
	}
}

//lookupXML returns XML data from PA-API
func (api *PaAPI) lookupXML(id string) (io.Reader, error) {
	params := map[string]string{
		"IdType":        "ASIN",
		"ResponseGroup": "Images,ItemAttributes,Small",
		"ItemId":        id,
	}
	xml, err := api.paapi.ItemLookupWithParams(params)
	if err != nil {
		return nil, errs.Wrap(err, "error in PaAPI.lookupXML() function")
	}
	return bytes.NewBufferString(xml), nil
}

//unmarshalXML returns unmarshalled XML data
func unmarshalXML(xmldata io.Reader) (*amazonproduct.ItemLookupResponse, error) {
	res := &amazonproduct.ItemLookupResponse{}
	if err := xml.NewDecoder(xmldata).Decode(res); err != nil {
		return nil, errs.Wrap(err, "error in unmarshalXML() function")
	}
	return res, nil
}

///LookupRawData returns PA-API raw data
func (api *PaAPI) LookupRawData(id string) (io.Reader, error) {
	return api.lookupXML(id)
}

///LookupBook returns Book data from PA-API
func (api *PaAPI) LookupBook(id string) (*entity.Book, error) {
	data, err := api.lookupXML(id)
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
		ID:    item.ASIN,
		Title: item.ItemAttributes.Title,
		URL:   item.DetailPageURL,
		Image: struct {
			URL    string
			Height uint16
			Width  uint16
		}{
			URL:    item.MediumImage.URL,
			Height: item.MediumImage.Height,
			Width:  item.MediumImage.Width,
		},
		ProductType: item.ItemAttributes.Binding,
		Codes:       []entity.Code{entity.Code{Name: "ASIN", Value: item.ASIN}},
		Authors:     item.ItemAttributes.Author,
		Publisher:   item.ItemAttributes.Publisher,
		Service: struct {
			Name string
			URL  string
		}{Name: "PA-API", URL: "https://affiliate.amazon.co.jp/assoc_credentials/home"},
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
		if tm, err := time.Parse(time.RFC3339, item.ItemAttributes.PublicationDate+"T09:00:00Z"); err == nil {
			book.PublicationDate = values.NewDate(tm)
		}
	}
	if len(item.ItemAttributes.ReleaseDate) > 0 {
		if tm, err := time.Parse(time.RFC3339, item.ItemAttributes.ReleaseDate+"T09:00:00Z"); err == nil {
			book.LastRelease = values.NewDate(tm)
		}
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
