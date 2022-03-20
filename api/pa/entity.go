package pa

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goark/books-data/ecode"
	"github.com/goark/books-data/entity"
	"github.com/goark/books-data/entity/values"
	"github.com/goark/errs"
)

type GenInfo struct {
	DisplayValue string
	Label        string `json:",omitempty"`
	Locale       string `json:",omitempty"`
}

type GenInfoInt struct {
	DisplayValue int
	Label        string `json:",omitempty"`
	Locale       string `json:",omitempty"`
}

type GenInfoFloat struct {
	DisplayValue float64
	Label        string `json:",omitempty"`
	Locale       string `json:",omitempty"`
}

type GenInfoTime struct {
	DisplayValue values.Date
	Label        string `json:",omitempty"`
	Locale       string `json:",omitempty"`
}

type IdInfo struct {
	DisplayValues []string
	Label         string `json:",omitempty"`
	Locale        string `json:",omitempty"`
}

type Price struct {
	DisplayAmount string
	Amount        float64
	Currency      string
}

type Item struct {
	ASIN          string
	DetailPageURL string
	Images        *struct {
		Primary *struct {
			Medium *entity.BookCover `json:",omitempty"`
		} `json:",omitempty"`
	} `json:",omitempty"`
	ItemInfo *struct {
		ByLineInfo *struct {
			Brand        *GenInfo `json:",omitempty"`
			Manufacturer *GenInfo `json:",omitempty"`
			Contributors []struct {
				Name   string
				Locale string
				Role   string
			}
		} `json:",omitempty"`
		Classifications *struct {
			Binding      GenInfo
			ProductGroup GenInfo
		} `json:",omitempty"`
		ContentInfo *struct {
			Edition   *GenInfo `json:",omitempty"`
			Languages struct {
				DisplayValues []struct {
					DisplayValue string
					Type         string
				}
				Label  string
				Locale string
			}
			PagesCount struct {
				DisplayValue int
				Label        string
				Locale       string
			}
			PublicationDate GenInfoTime
		} `json:",omitempty"`
		ExternalIds *struct {
			EANs  *IdInfo `json:",omitempty"`
			ISBNs *IdInfo `json:",omitempty"`
			UPCs  *IdInfo `json:",omitempty"`
		} `json:",omitempty"`
		ProductInfo *struct {
			Color          *GenInfo `json:",omitempty"`
			IsAdultProduct struct {
				DisplayValue bool
				Label        string
				Locale       string
			}
			ItemDimensions *struct {
				Height *GenInfoFloat `json:",omitempty"`
				Length *GenInfoFloat `json:",omitempty"`
				Weight *GenInfoFloat `json:",omitempty"`
				Width  *GenInfoFloat `json:",omitempty"`
			} `json:",omitempty"`
			ReleaseDate *GenInfoTime `json:",omitempty"`
			Size        *GenInfo     `json:",omitempty"`
			UnitCount   *GenInfoInt  `json:",omitempty"`
		} `json:",omitempty"`
		Title *GenInfo `json:",omitempty"`
	}
}

func (i *Item) Title() string {
	if i.ItemInfo != nil && i.ItemInfo.Title != nil {
		return i.ItemInfo.Title.DisplayValue
	}
	return ""
}

func (i *Item) BookCover() entity.BookCover {
	if i.Images != nil && i.Images.Primary != nil && i.Images.Primary.Medium != nil {
		return *i.Images.Primary.Medium
	}
	return entity.BookCover{}
}

func (i *Item) ProductType() string {
	if i.ItemInfo != nil && i.ItemInfo.Classifications != nil {
		return i.ItemInfo.Classifications.Binding.DisplayValue
	}
	return ""
}

func (i *Item) Codes() []entity.Code {
	codes := []entity.Code{{Name: "ASIN", Value: i.ASIN}}
	if i.ItemInfo != nil && i.ItemInfo.ExternalIds != nil {
		ids := i.ItemInfo.ExternalIds.EANs
		if ids != nil {
			for _, id := range ids.DisplayValues {
				codes = append(codes, entity.Code{Name: ids.Label, Value: id})
			}
		}
		ids = i.ItemInfo.ExternalIds.ISBNs
		if ids != nil {
			for _, id := range ids.DisplayValues {
				codes = append(codes, entity.Code{Name: ids.Label, Value: id})
			}
		}
		ids = i.ItemInfo.ExternalIds.UPCs
		if ids != nil {
			for _, id := range ids.DisplayValues {
				codes = append(codes, entity.Code{Name: ids.Label, Value: id})
			}
		}
	}
	return codes
}

func (i *Item) Creators() []entity.Creator {
	creators := []entity.Creator{}
	if i.ItemInfo != nil && i.ItemInfo.ByLineInfo != nil {
		for _, contributor := range i.ItemInfo.ByLineInfo.Contributors {
			creators = append(creators, entity.Creator{Name: contributor.Name, Role: contributor.Role})
		}
	}
	return creators
}

func (i *Item) Publisher() string {
	if i.ItemInfo != nil && i.ItemInfo.ByLineInfo != nil && i.ItemInfo.ByLineInfo.Manufacturer != nil {
		return i.ItemInfo.ByLineInfo.Manufacturer.DisplayValue
	}
	return ""
}

func (i *Item) PublicationDate() values.Date {
	if i.ItemInfo != nil && i.ItemInfo.ContentInfo != nil {
		return i.ItemInfo.ContentInfo.PublicationDate.DisplayValue
	}
	return values.Date{}
}

func (i *Item) LastRelease() values.Date {
	if i.ItemInfo != nil && i.ItemInfo.ProductInfo != nil && i.ItemInfo.ProductInfo.ReleaseDate != nil {
		return i.ItemInfo.ProductInfo.ReleaseDate.DisplayValue
	}
	return values.Date{}
}

type Response struct {
	Errors []struct {
		Code    string
		Message string
	} `json:",omitempty"`
	ItemsResult struct {
		Items []Item `json:",omitempty"`
	}
}

func (r *Response) CheckError() error {
	if len(r.Errors) > 0 {
		msg := r.Errors[0].Message
		es := []string{}
		for _, e := range r.Errors {
			es = append(es, fmt.Sprintf("%s (%s)", e.Message, e.Code))
		}
		return errs.New(msg, errs.WithContext("detail", strings.Join(es, "\n")))
	}
	return nil
}

func (r *Response) Output(typeName string) (*entity.Book, error) {
	if r == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer)
	}
	if err := r.CheckError(); err != nil {
		return nil, errs.Wrap(err)
	}
	if len(r.ItemsResult.Items) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	item := r.ItemsResult.Items[0]
	book := &entity.Book{
		Type:            typeName,
		ID:              item.ASIN,
		Title:           item.Title(),
		URL:             item.DetailPageURL,
		Image:           item.BookCover(),
		ProductType:     item.ProductType(),
		Codes:           item.Codes(),
		Creators:        item.Creators(),
		Publisher:       item.Publisher(),
		PublicationDate: item.PublicationDate(),
		LastRelease:     item.LastRelease(),
		Service:         entity.Service{Name: "PA-APIv5", URL: "https://affiliate.amazon.co.jp/assoc_credentials/home"},
	}

	return book, nil
}

//JSON returns JSON data from Response instance
func (r *Response) JSON() ([]byte, error) {
	b, err := json.Marshal(r)
	return b, errs.Wrap(err)
}

//Stringer
func (r *Response) String() string {
	b, err := r.JSON()
	if err != nil {
		return ""
	}
	return string(b)
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
