package facade

import (
	"bytes"
	"testing"

	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
)

func TestHistory(t *testing.T) {
	testCases := []struct {
		args   []string
		out    string
		outErr string
	}{
		{args: []string{"history", "-i", "9784152098702", "-l", "../testdata/review-log.json"}, out: `{"Book":{"Type":"openbd","ID":"9784152098702","Title":"三体","Image":{"URL":"https://cover.openbd.jp/9784152098702.jpg"},"ProductType":"Book","Creators":[{"Name":"劉慈欣／著 大森望／翻訳 光吉さくら／翻訳 ワンチャイ／翻訳 立原透耶／監修"}],"Publisher":"早川書房","Codes":[{"Name":"ISBN","Value":"9784152098702"}],"PublicationDate":"2019-07-04","LastRelease":"","Service":{"Name":"openBD","URL":"https://openbd.jp/"}},"Date":"2019-07-14","Rating":4,"Star":[true,true,true,true,false],"Description":"流行ってるらしいので買ってみた。 Kindle の肥やしにならないことを祈ろう（読めって！）。"}`, outErr: ""},
	}

	for _, tc := range testCases {
		out := new(bytes.Buffer)
		errOut := new(bytes.Buffer)
		ui := rwi.New(
			rwi.WithWriter(out),
			rwi.WithErrorWriter(errOut),
		)
		exit := Execute(ui, tc.args)
		if exit != exitcode.Normal {
			t.Errorf("Execute() err = \"%v\", want \"%v\".", exit, exitcode.Normal)
		}
		if out.String() != tc.out {
			t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", out.String(), tc.out)
		}
		if errOut.String() != tc.outErr {
			t.Errorf("Execute() Stderr = \"%v\", want \"%v\".", errOut.String(), tc.outErr)
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
