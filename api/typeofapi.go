package api

//ECode is error codes for books-data
type ServiceType int

const (
	TypePAAPI ServiceType = iota + 1
	TypeOpenBD
)

var strTypes = map[ServiceType]string{
	TypePAAPI:  "paapi",
	TypeOpenBD: "openbd",
}

func (t ServiceType) String() string {
	if s, ok := strTypes[t]; ok {
		return s
	}
	return "unknown"
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
