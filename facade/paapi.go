package facade

import (
	"bytes"
	"io"

	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/books-data/api"
	"github.com/spiegel-im-spiegel/books-data/api/pa"
	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/books-data/entity"
	"github.com/spiegel-im-spiegel/errs"
)

func CreatePAAPI(isbnFlag bool) (api.API, error) {
	marketplace := viper.GetString("marketplace")
	if len(marketplace) == 0 {
		return nil, errs.Wrap(ecode.ErrInvalidAPIParameter, "marketplace is empty")
	}
	associateTag := viper.GetString("associate-tag")
	if len(associateTag) == 0 {
		return nil, errs.Wrap(ecode.ErrInvalidAPIParameter, "associate-tag is empty")
	}
	accessKey := viper.GetString("access-key")
	if len(accessKey) == 0 {
		return nil, errs.Wrap(ecode.ErrInvalidAPIParameter, "access-key is empty")
	}
	secretKey := viper.GetString("secret-key")
	if len(secretKey) == 0 {
		return nil, errs.Wrap(ecode.ErrInvalidAPIParameter, "secret-key is empty")
	}
	return pa.New(
		pa.WithMarketplace(marketplace),
		pa.WithAssociateTag(associateTag),
		pa.WithAccessKey(accessKey),
		pa.WithSecretKey(secretKey),
		pa.WithEnableISBN(isbnFlag),
	), nil
}

func searchPAAPI(id string, isbnFlag, rawFlag bool) (io.Reader, error) {
	paapi, err := CreatePAAPI(isbnFlag)
	if err != nil {
		return nil, err
	}
	if rawFlag {
		return paapi.LookupRawData(id)
	}
	book, err := paapi.LookupBook(id)
	if err != nil {
		return nil, err
	}
	b, err := book.Format(tmpltPath)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func findPAAPI(id string, isbnFlag bool) (*entity.Book, error) {
	paapi, err := CreatePAAPI(isbnFlag)
	if err != nil {
		return nil, err
	}
	return paapi.LookupBook(id)
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
