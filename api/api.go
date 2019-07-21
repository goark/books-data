package api

import (
	"io"

	"github.com/spiegel-im-spiegel/books-data/entity"
)

type API interface {
	LookupBook(id string) (*entity.Book, error) //Lookup book data by API
	LookupRawData(id string) (io.Reader, error) //Lookup raw data by API
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
