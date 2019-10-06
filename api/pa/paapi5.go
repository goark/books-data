package pa

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/spiegel-im-spiegel/books-data/api"
	"github.com/spiegel-im-spiegel/books-data/entity"
	"github.com/spiegel-im-spiegel/errs"
	paapi5 "github.com/spiegel-im-spiegel/pa-api"
)

//PAAPI5 is a api.API class for PA-API v5
type PAAPI5 struct {
	svcType      api.ServiceType //Service Type
	associateTag string
	accessKey    string
	secretKey    string
	server       *paapi5.Server  //server info.
	ctx          context.Context //context
}

var _ api.API = (*PAAPI5)(nil) //PAAPI5 is compatible with api.API interface

//New returns PaAPI instance
func New(ctx context.Context, marketplace, associateTag, accessKey, secretKey string) *PAAPI5 {
	return &PAAPI5{
		svcType:      api.TypePAAPI,
		associateTag: associateTag,
		accessKey:    accessKey,
		secretKey:    secretKey,
		server:       paapi5.New(paapi5.WithMarketplace(marketplace)),
		ctx:          ctx,
	}
}

//Name returns name of API
func (a *PAAPI5) Name() string {
	return a.svcType.String()
}

///LookupRawData returns PA-API raw data
func (a *PAAPI5) LookupRawData(id string) (io.Reader, error) {
	client := a.server.CreateClient(a.associateTag, a.accessKey, a.secretKey, paapi5.WithContext(a.ctx), paapi5.WithHttpCilent(&http.Client{}))
	q := NewQuery(client.Marketplace(), client.PartnerTag(), client.PartnerType(), []string{id})
	body, err := client.Request(q)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithParam("id", id))
	}

	return bytes.NewReader(body), nil
}

///LookupBook returns Book data from PA-API
func (a *PAAPI5) LookupBook(id string) (*entity.Book, error) {
	r, err := a.LookupRawData(id)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithParam("id", id))
	}
	rsp := Response{}
	if err := json.NewDecoder(r).Decode(&rsp); err != nil {
		return nil, errs.Wrap(err, "", errs.WithParam("id", id))
	}
	book, err := rsp.Output(a.Name())
	return book, errs.Wrap(err, "", errs.WithParam("id", id))
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
