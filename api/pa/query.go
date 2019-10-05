package pa

import (
	"encoding/json"

	"github.com/spiegel-im-spiegel/errs"
	paapi5 "github.com/spiegel-im-spiegel/pa-api"
)

//Query is query data class for PA-API v5
type Query struct {
	OpeCode     paapi5.Operation `json:"Operation"`
	Marketplace string
	PartnerTag  string
	PartnerType string
	ItemIds     []string `json:",omitempty"`
	ItemIdType  string   `json:",omitempty"`
	Resources   []string `json:",omitempty"`
}

var _ paapi5.Query = (*Query)(nil) //Query is compatible with paapi5.Query interface

//NewQuery creates new Query instance
func NewQuery(marketplace, partnerTag, partnerType string, itms []string) *Query {
	q := &Query{
		OpeCode:     paapi5.GetItems,
		Marketplace: marketplace,
		PartnerTag:  partnerTag,
		PartnerType: partnerType,
		ItemIds:     itms,
		ItemIdType:  "ASIN",
	}
	return q
}

func (q *Query) Operation() paapi5.Operation {
	if q == nil {
		return paapi5.NullOperation
	}
	return q.OpeCode
}

func (q *Query) Payload() ([]byte, error) {
	if q == nil {
		return nil, errs.Wrap(paapi5.ErrNullPointer, "")
	}
	q.Resources = []string{}
	q.Resources = append(
		q.Resources,
		"Images.Primary.Medium",
		"ItemInfo.ByLineInfo",
		"ItemInfo.ContentInfo",
		"ItemInfo.ContentRating",
		"ItemInfo.Classifications",
		"ItemInfo.ExternalIds",
		"ItemInfo.Features",
		"ItemInfo.ManufactureInfo",
		"ItemInfo.ProductInfo",
		"ItemInfo.TechnicalInfo",
		"ItemInfo.Title",
		"ItemInfo.TradeInInfo",
	)
	b, err := json.Marshal(q)
	return b, errs.Wrap(err, "")
}

func (q *Query) String() string {
	b, err := q.Payload()
	if err != nil {
		return ""
	}
	return string(b)
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
