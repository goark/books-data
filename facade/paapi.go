package facade

import (
	"bytes"
	"context"
	"io"

	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/books-data/api"
	"github.com/spiegel-im-spiegel/books-data/api/pa"
	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/books-data/entity"
	"github.com/spiegel-im-spiegel/errs"
)

//paapiParams is parameters for PA-API
type paapiParams struct {
	marketplace  string
	associateTag string
	accessKey    string
	secretKey    string
}

func getPaapiParams() (*paapiParams, error) {
	marketplace := viper.GetString("marketplace")
	if len(marketplace) == 0 {
		return nil, errs.New("marketplace is empty", errs.WithCause(ecode.ErrInvalidAPIParameter))
	}
	associateTag := viper.GetString("associate-tag")
	if len(associateTag) == 0 {
		return nil, errs.New("associate-tag is empty", errs.WithCause(ecode.ErrInvalidAPIParameter))
	}
	accessKey := viper.GetString("access-key")
	if len(accessKey) == 0 {
		return nil, errs.New("access-key is empty", errs.WithCause(ecode.ErrInvalidAPIParameter))
	}
	secretKey := viper.GetString("secret-key")
	if len(secretKey) == 0 {
		return nil, errs.New("secret-key is empty", errs.WithCause(ecode.ErrInvalidAPIParameter))
	}
	return &paapiParams{marketplace: marketplace, associateTag: associateTag, accessKey: accessKey, secretKey: secretKey}, nil
}

func createPAAPI(ctx context.Context, p *paapiParams) api.API {
	return pa.New(
		ctx,
		p.marketplace,
		p.associateTag,
		p.accessKey,
		p.secretKey,
	)
}

func searchPAAPI(ctx context.Context, id string, p *paapiParams, rawFlag bool) (io.Reader, error) {
	if rawFlag {
		return createPAAPI(ctx, p).LookupRawData(id)
	}
	book, err := createPAAPI(ctx, p).LookupBook(id)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("id", id))
	}
	b, err := book.Format(tmpltPath)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("id", id))
	}
	return bytes.NewReader(b), nil
}

func findPAAPI(ctx context.Context, id string, p *paapiParams) (*entity.Book, error) {
	return createPAAPI(ctx, p).LookupBook(id)
}

/* Copyright 2019,2020 Spiegel
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
