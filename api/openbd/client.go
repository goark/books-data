package openbd

import (
	"net/http"
	"net/url"

	"github.com/spiegel-im-spiegel/errs"
)

//Client is http.Client for openBD RESTful API
type Client struct {
	cmd    CommandType
	client *http.Client
}

func (c *Client) Get(v url.Values) (*http.Response, error) {
	url := c.cmd.String() + "?" + v.Encode()
	resp, err := c.client.Get(url)
	if err != nil {
		return resp, errs.Wrapf(err, "error in Client.Get(\"%v\") function", url)
	}
	return resp, nil
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
