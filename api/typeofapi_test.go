package api

import "testing"

func TestServiceType(t *testing.T) {
	testCases := []struct {
		t   ServiceType
		str string
	}{
		{t: ServiceType(0), str: "unknown"},
		{t: TypePAAPI, str: "paapi"},
		{t: TypeOpenBD, str: "openbd"},
		{t: TypeOthers, str: "others"},
		{t: ServiceType(4), str: "unknown"},
	}

	for _, tc := range testCases {
		s := tc.t.String()
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
