package entity

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/books-data/entity/values"
	"github.com/spiegel-im-spiegel/books-data/format"
	"github.com/spiegel-im-spiegel/errs"
)

//Code is entity class of book code
type Code struct {
	Name  string
	Value string
}

func (c Code) String() string {
	if len(c.Value) > 0 && len(c.Name) > 0 {
		return fmt.Sprintf("%v (%v)", c.Value, c.Name)
	}
	return c.Value
}

//Creator is entity class of creator info.
type Creator struct {
	Name string
	Role string `json:",omitempty"`
}

func (c Creator) String() string {
	if len(c.Role) > 0 && len(c.Name) > 0 {
		return fmt.Sprintf("%v (%v)", c.Name, c.Role)
	}
	return c.Name
}

//BookCover is entity class of book cover image info.
type BookCover struct {
	URL    string
	Height uint16 `json:",omitempty"`
	Width  uint16 `json:",omitempty"`
}

//Service is entity class of API service info.
type Service struct {
	Name string
	URL  string
}

//Book is entity class of information for book
type Book struct {
	Type            string
	ID              string
	Title           string
	SubTitle        string `json:",omitempty"`
	SeriesTitle     string `json:",omitempty"`
	OriginalTitle   string `json:",omitempty"`
	URL             string `json:",omitempty"`
	Image           BookCover
	ProductType     string    `json:",omitempty"`
	Creators        []Creator `json:",omitempty"`
	Publisher       string    `json:",omitempty"`
	Codes           []Code
	PublicationDate values.Date
	LastRelease     values.Date
	PublicDomain    bool   `json:",omitempty"`
	FirstAppearance string `json:",omitempty"`
	Service         Service
}

//ImportBookFromJSON import Reviews data with JSON format
func ImportBookFromJSON(r io.Reader) (*Book, error) {
	dec := json.NewDecoder(r)
	bk := Book{}
	err := dec.Decode(&bk)
	return &bk, errs.Wrap(err)
}

func (b *Book) Format(tmpltPath string) ([]byte, error) {
	if b == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer, errs.WithContext("tmpltPath", tmpltPath))
	}
	if len(tmpltPath) == 0 {
		b, err := json.Marshal(b)
		if err != nil {
			return nil, errs.Wrap(err, errs.WithContext("tmpltPath", tmpltPath))
		}
		return b, nil
	}
	buf, err := format.ByTemplateFile(b, tmpltPath)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("tmpltPath", tmpltPath))
	}
	return buf.Bytes(), nil
}

func (b *Book) String() string {
	res, err := b.Format("")
	if err != nil {
		return ""
	}
	return string(res)
}

//CopyFrom copy from Book instance
func (b *Book) CopyFrom(src *Book) *Book {
	if src == nil {
		return nil
	}
	if b == nil {
		return src
	}
	*b = *src
	return b
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
