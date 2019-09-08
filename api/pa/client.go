package pa

import (
	"bytes"
	"io"

	amazonproduct "github.com/DDRBoxman/go-amazon-product-api"
	"github.com/spiegel-im-spiegel/errs"
)

//Client is http.Client for PA-API
type Client struct {
	enableISBN bool //enable ISBN code by lookup item
	paapi      *amazonproduct.AmazonProductAPI
}

//LookupXML returns XML data from PA-API lookup command
func (c *Client) LookupXML(id string) (io.Reader, error) {
	params := map[string]string{
		"ResponseGroup": "Images,ItemAttributes,Small",
		"ItemId":        id,
	}
	if c.enableISBN {
		params["IdType"] = "ISBN"
		params["SearchIndex"] = "All"
	} else {
		params["IdType"] = "ASIN"
	}
	xml, err := c.paapi.ItemLookupWithParams(params)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithParam("id", id))
	}
	return bytes.NewBufferString(xml), nil
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
