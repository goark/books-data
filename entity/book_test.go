package entity

import (
	"strings"
	"testing"

	"github.com/goark/books-data/entity/values"
)

func TestCode(t *testing.T) {
	testCases := []struct {
		c   Code
		str string
	}{
		{c: Code{Name: "Name", Value: "Value"}, str: "Value (Name)"},
		{c: Code{Name: "", Value: "Value"}, str: "Value"},
		{c: Code{Name: "Name", Value: ""}, str: ""},
		{c: Code{Name: "", Value: ""}, str: ""},
	}

	for _, tc := range testCases {
		s := tc.c.String()
		if s != tc.str {
			t.Errorf("\"%v\" != \"%v\"", s, tc.str)
		}
	}
}

func TestCreator(t *testing.T) {
	testCases := []struct {
		c   Creator
		str string
	}{
		{c: Creator{Name: "Name", Role: "Role"}, str: "Name (Role)"},
		{c: Creator{Name: "Name", Role: ""}, str: "Name"},
		{c: Creator{Name: "", Role: "Role"}, str: ""},
		{c: Creator{Name: "", Role: ""}, str: ""},
	}

	for _, tc := range testCases {
		s := tc.c.String()
		if s != tc.str {
			t.Errorf("\"%v\" != \"%v\"", s, tc.str)
		}
	}
}

var bookTestCases = []struct {
	b   Book
	str string
}{
	{
		b: Book{
			Type:        "Type",
			ID:          "ID",
			Title:       "Title",
			SubTitle:    "SubTitle",
			SeriesTitle: "SeriesTitle",
			URL:         "URL",
			Image: struct {
				URL    string
				Height uint16 `json:",omitempty"`
				Width  uint16 `json:",omitempty"`
			}{
				URL:    "url",
				Height: 0,
				Width:  100,
			},
			ProductType: "ProductType",
			Creators: []Creator{
				Creator{Name: "Name1", Role: "Role1"},
				Creator{Name: "Name2", Role: "Role2"},
			},
			Publisher: "Publisher",
			Codes: []Code{
				Code{Value: "Value1", Name: "Name1"},
				Code{Value: "Value2", Name: "Name2"},
			},
			PublicationDate: values.Date{},
			LastRelease:     values.Date{},
			Service: struct {
				Name string
				URL  string
			}{
				Name: "Name",
				URL:  "URL",
			},
		},
		str: `{"Type":"Type","ID":"ID","Title":"Title","SubTitle":"SubTitle","SeriesTitle":"SeriesTitle","URL":"URL","Image":{"URL":"url","Width":100},"ProductType":"ProductType","Creators":[{"Name":"Name1","Role":"Role1"},{"Name":"Name2","Role":"Role2"}],"Publisher":"Publisher","Codes":[{"Name":"Name1","Value":"Value1"},{"Name":"Name2","Value":"Value2"}],"PublicationDate":"","LastRelease":"","Service":{"Name":"Name","URL":"URL"}}`,
	},
}

func TestBook(t *testing.T) {
	for _, tc := range bookTestCases {
		s := tc.b.String()
		if s != tc.str {
			t.Errorf("\"%v\" != \"%v\"", s, tc.str)
		}

	}
}

func TestBookCopyFrom(t *testing.T) {
	for _, tc := range bookTestCases {
		dst := &Book{}
		s := dst.CopyFrom(&tc.b).String()
		if s != tc.str {
			t.Errorf("\"%v\" != \"%v\"", s, tc.str)
		}

	}
}

var testBookData = `{"Type":"test","ID":"card56642","Title":"陰翳礼讃","URL":"https://www.aozora.gr.jp/cards/001383/card56642.html","Image":{"URL":"https://text.baldanders.info/images/aozora/card56642.svg","Height":227,"Width":321},"ProductType":"Book","Creators":[{"Name":"谷崎 潤一郎","Role":"著"}],"Publisher":"青空文庫","Codes":[{"Name":"青空文庫","Value":"card56642"}],"PublicationDate":"2016-06-10","LastRelease":"2019-02-24","Service":{"Name":"青空文庫","URL":"https://www.aozora.gr.jp/"}}`
var testBookRes = `@BOOK{Book:card56642,
    TITLE = "陰翳礼讃",
    AUTHOR = "谷崎 潤一郎 (著)",
    PUBLISHER = {青空文庫},
    YEAR = 2016
}
`

func TestFormat(t *testing.T) {
	testCases := []struct {
		data   string
		result string
	}{
		{data: testBookData, result: testBookRes},
	}

	for _, tc := range testCases {
		book, err := ImportBookFromJSON(strings.NewReader(tc.data))
		if err != nil {
			t.Errorf("ImportBookFromJSON() error: %+v", err)
			continue
		}

		b, err := book.Format("../testdata/book-template/template.bib.txt")
		if err != nil {
			t.Errorf("ByTemplateFile() error: %+v", err)
			continue
		}
		s := string(b)
		if s != tc.result {
			t.Errorf("ByTemplateFile() = \"%v\", watnt \"%v\"", s, tc.result)
		}
	}
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
