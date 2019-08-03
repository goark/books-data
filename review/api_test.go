package review

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/spiegel-im-spiegel/books-data/api"
	"github.com/spiegel-im-spiegel/books-data/entity"
)

type testAPI struct {
	output string
}

func NewTestAPI() api.API {
	return testAPI{testBookData}
}

func (ta testAPI) Name() string {
	return "test"
}

func (ta testAPI) LookupRawData(id string) (io.Reader, error) {
	return bytes.NewBufferString(ta.output), nil
}

func (ta testAPI) LookupBook(id string) (*entity.Book, error) {
	r, _ := ta.LookupRawData(id)
	dec := json.NewDecoder(r)
	books := &entity.Book{}
	if err := dec.Decode(books); err != nil {
		return nil, err
	}
	books.Codes[0].Value = id
	return books, nil
}

var (
	testBookData = `
{
    "Type": "test",
    "ID": "card56642",
    "Title": "陰翳礼讃",
    "SubTitle": "",
    "SeriesTitle": "",
    "URL": "https://www.aozora.gr.jp/cards/001383/card56642.html",
      "Image": {
		  "URL": "https://text.baldanders.info/images/aozora/card56642.svg",
		  "Height": 227,
		  "Width": 321
	  },
	  "ProductType": "Book",
	  "Authors": [ "谷崎 潤一郎" ],
	  "Publisher": "青空文庫",
	  "Codes": [
	  	{
			"Name": "青空文庫",
			"Value": "card56642"
		}
	  ],
	  "PublicationDate": "2016-06-10",
	  "LastRelease": "2019-02-24",
	  "Service": {
		  "Name": "青空文庫",
		  "URL": "https://www.aozora.gr.jp/"
	  }
}
`
	testBookResp = `{"Type":"test","ID":"card56642","Title":"陰翳礼讃","URL":"https://www.aozora.gr.jp/cards/001383/card56642.html","Image":{"URL":"https://text.baldanders.info/images/aozora/card56642.svg","Height":227,"Width":321},"ProductType":"Book","Authors":["谷崎 潤一郎"],"Publisher":"青空文庫","Codes":[{"Name":"青空文庫","Value":"card56642"}],"PublicationDate":"2016-06-10","LastRelease":"2019-02-24","Service":{"Name":"青空文庫","URL":"https://www.aozora.gr.jp/"}}`
)

func TestAPI(t *testing.T) {
	tc := NewTestAPI()
	book, err := tc.LookupBook("card56642")
	if err != nil {
		t.Errorf("testAPI.LookupBook() error is \"%v\", want nil", err)
		fmt.Printf("Info: %+v\n", err)
		return
	}
	str := book.String()
	if str != testBookResp {
		t.Errorf("testAPI.LookupBook() = \"%v\", want \"%v\"", str, testBookResp)
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
