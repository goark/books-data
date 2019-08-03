package entity

import (
	"testing"

	"github.com/spiegel-im-spiegel/books-data/entity/values"
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
				Height uint16
				Width  uint16
			}{
				URL:    "url",
				Height: 100,
				Width:  100,
			},
			ProductType: "ProductType",
			Authors:     []string{"Author1", "Author2"},
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
		str: `{"Type":"Type","ID":"ID","Title":"Title","SubTitle":"SubTitle","SeriesTitle":"SeriesTitle","URL":"URL","Image":{"URL":"url","Height":100,"Width":100},"ProductType":"ProductType","Authors":["Author1","Author2"],"Creators":[{"Name":"Name1","Role":"Role1"},{"Name":"Name2","Role":"Role2"}],"Publisher":"Publisher","Codes":[{"Name":"Name1","Value":"Value1"},{"Name":"Name2","Value":"Value2"}],"PublicationDate":"0001-01-01","LastRelease":"0001-01-01","Service":{"Name":"Name","URL":"URL"}}`,
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
