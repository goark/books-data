package pa

import (
	"net/http"

	amazonproduct "github.com/DDRBoxman/go-amazon-product-api"
)

//Server is informations of PA-API
type Server struct {
	marketplace  string
	associateTag string
	accessKey    string
	secretKey    string
	enableISBN   bool //enable ISBN code by lookup item
}

func (s *Server) CreateClient() *Client {
	return &Client{
		enableISBN: s.enableISBN,
		paapi: &amazonproduct.AmazonProductAPI{
			Client:       &http.Client{},
			Host:         s.marketplace,
			AssociateTag: s.associateTag,
			AccessKey:    s.accessKey,
			SecretKey:    s.secretKey,
		},
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
