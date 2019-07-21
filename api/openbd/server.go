package openbd

import (
	"net/http"
	"net/url"
)

//CommandType is Type of openBD commands
type CommandType int

const (
	GET      CommandType = iota + 1 //GET command
	POST                            //POST command
	COVERAGE                        //COVERAGE command
)

var commandMap = map[CommandType]string{
	GET:      "https://api.openbd.jp/v1/get",
	POST:     "https://api.openbd.jp/v1/post",
	COVERAGE: "https://api.openbd.jp/v1/coverage",
}

func (c CommandType) String() string {
	if s, ok := commandMap[c]; ok {
		return s
	}
	return ""
}

//Server is informations of OpenPGP key server
type Server struct {
	cmd      CommandType //Type of openBD commands
	proxyURL string      //URL of proxy server
}

func (s *Server) CreateClient() *Client {
	if len(s.proxyURL) > 0 {
		if proxyUrl, err := url.Parse(s.proxyURL); err == nil {
			return &Client{cmd: s.cmd, client: &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}}
		}
	}
	return &Client{cmd: s.cmd, client: &http.Client{}}
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
