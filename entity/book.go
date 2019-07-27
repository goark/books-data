package entity

import (
	"encoding/json"
	"fmt"

	"github.com/spiegel-im-spiegel/books-data/entity/values"
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
	return c.Name
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

//Book is entity class of information for book
type Book struct {
	ID          string
	Title       string
	SubTitle    string `json:",omitempty"`
	SeriesTitle string `json:",omitempty"`
	URL         string `json:",omitempty"`
	Image       struct {
		URL    string
		Height uint16
		Width  uint16
	}
	ProductType     string `json:",omitempty"`
	Authors         []string
	Creators        []Creator `json:",omitempty"`
	Publisher       string
	Codes           []Code
	PublicationDate values.Date
	LastRelease     values.Date
	Service         struct {
		Name string
		URL  string
	}
}

func (b *Book) String() string {
	res, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(res)
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
