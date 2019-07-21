package values

import (
	"encoding/json"
	"fmt"
	"testing"
)

type ForTestStruct struct {
	DateTaken    Date `json:"date_taken"`
	DateImported Date `json:"date_imported"`
}

func TestUnmarshal(t *testing.T) {
	data := `
{
	"date_taken": "2005-03-26",
	"date_imported": "2005-03-25T20:10:13+00:00"
}
`
	res := `{"date_taken":"2005-03-26","date_imported":"2005-03-25"}`
	tst := &ForTestStruct{}
	if err := json.Unmarshal([]byte(data), tst); err != nil {
		t.Errorf("Unmarshal() = \"%v\", want nil.", err)
		return
	}
	if b, err := json.Marshal(tst); err != nil {
		t.Errorf("json.Marshal() = \"%v\", want nil.", err)
	} else if string(b) != res {
		t.Errorf("JSON of ForTestStruct = \"%v\", want \"%v\".", string(b), res)
	}
}

func TestUnmarshalErr(t *testing.T) {
	data := `
{
	"date_taken": "2005/03/26"
}
`
	tst := &ForTestStruct{}
	if err := json.Unmarshal([]byte(data), tst); err == nil {
		t.Error("Unmarshal() error = nil, not want nil.")
	} else {
		fmt.Printf("Info: %+v\n", err)
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
