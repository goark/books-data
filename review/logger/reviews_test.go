package logger

import (
	"fmt"
	"strings"
	"testing"
)

var (
	testBookdata = `{"Type":"test","ID":"card56642","Title":"陰翳礼讃","URL":"https://www.aozora.gr.jp/cards/001383/card56642.html","Image":{"URL":"https://text.baldanders.info/images/aozora/card56642.svg","Height":227,"Width":321},"ProductType":"Book","Authors":["谷崎 潤一郎"],"Publisher":"青空文庫","Codes":[{"Name":"青空文庫","Value":"card56642"}],"PublicationDate":"2016-06-10","LastRelease":"2019-02-24","Service":{"Name":"青空文庫","URL":"https://www.aozora.gr.jp/"}}`
	testRevsdata = `[{"Book":` + testBookdata + `,"Date":"2019-01-01","Rating":4,"Star":[true,true,true,true,false],"Description":"実はちゃんと読んでない（笑） 学生時代に読んでおけばよかった。"}]`
	testRes      = `[
  {
    "Book": {
      "Type": "test",
      "ID": "card56642",
      "Title": "陰翳礼讃",
      "URL": "https://www.aozora.gr.jp/cards/001383/card56642.html",
      "Image": {
        "URL": "https://text.baldanders.info/images/aozora/card56642.svg",
        "Height": 227,
        "Width": 321
      },
      "ProductType": "Book",
      "Authors": [
        "谷崎 潤一郎"
      ],
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
    },
    "Date": "2019-01-01",
    "Rating": 4,
    "Star": [
      true,
      true,
      true,
      true,
      false
    ],
    "Description": "実はちゃんと読んでない（笑） 学生時代に読んでおけばよかった。"
  }
]
`
)

func TestReviews(t *testing.T) {
	testCases := []struct {
		revs string
		str  string
	}{
		{revs: testRevsdata, str: testRes},
	}

	for _, tc := range testCases {
		revs, err := ImportJSON(strings.NewReader(tc.revs))
		if err != nil {
			t.Errorf("ImportJSON() is \"%v\", want nil", err)
			fmt.Printf("Info: %+v", err)
			continue
		}
		s := revs.String()
		if s != tc.str {
			t.Errorf("\"%v\" != \"%v\"", s, tc.str)
		}
		revs = (&Reviews{}).Append(&revs[0])
		s = revs.String()
		if s != tc.str {
			t.Errorf("\"%v\" != \"%v\"", s, tc.str)
		}
		revs = revs.Append(&revs[0])
		s = revs.String()
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
