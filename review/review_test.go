package review

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReview(t *testing.T) {
	res := `{"Book":` + testBookResp + `,"Date":"2019-01-01","Rating":4,"Star":[true,true,true,true,false],"Description":"実はちゃんと読んでない（笑） 学生時代に読んでおけばよかった。"}`
	tc := NewTestAPI()
	book, err := tc.LookupBook("card56642")
	if err != nil {
		t.Errorf("testAPI.LookupBook() error is \"%v\", want nil", err)
		fmt.Printf("Info: %+v\n", err)
		return
	}
	rev := New(
		book,
		WithDate("2019-01-01"),
		WithRating(4),
		WithDescription("実はちゃんと読んでない（笑） 学生時代に読んでおけばよかった。"),
	)
	str := rev.String()
	if str != res {
		t.Errorf("review.New() = \"%v\", want \"%v\"", str, res)
	}
}

var testTemplate = `<div class="hreview">
  <div class="photo">{{ if .Book.URL }}<a class="item url" href="{{ .Book.URL }}">{{ end }}<img src="{{ .Book.Image.URL }}" alt="photo">{{ if .Book.URL }}</a>{{ end }}</div>
  <dl class="fn">
    <dt>{{ if .Book.URL }}<a href="{{ .Book.URL }}">{{ end }}{{ .Book.Title }}{{ if .Book.URL }}</a>{{ end }}</dt>{{ if .Book.Authors }}
	<dd>{{ range $i, $v := .Book.Authors }}{{ if ne $i 0 }}, {{ end }}{{ $v }}{{ end }}</dd>{{ end }}
    <dd>{{ .Book.Publisher }}{{ with .Book.PublicationDate }} {{ . }}{{ end }}{{ with .Book.LastRelease }} (Release {{ . }}){{ end }}</dd>
    <dd>{{ .Book.ProductType }}</dd>{{ if .Book.Codes }}
	<dd>{{ range $i, $v := .Book.Codes }}{{ if ne $i 0 }}, {{ end }}{{ $v }}{{ end }}</dd>{{ end }}{{ if gt .Rating 0 }}
    <dd>評価:<abbr class="rating fa-sm" title="{{ .Rating }}">{{ range .Star }}&nbsp;{{ if . }}<i class="fas fa-star"></i>{{ else }}<i class="far fa-star"></i>{{ end }}{{ end }}</abbr></dd>{{ end }}
  </dl>
  <p class="description">{{ .Description }}</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>reviewer</a> on <abbr class="dtreviewed" title="{{ .Date }}">{{ .Date }}</abbr> (powered by <a href="{{ .Book.Service.URL }}" >{{ .Book.Service.Name }}</a>)</p>
</div>`

var formattedText = `<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.aozora.gr.jp/cards/001383/card56642.html"><img src="https://text.baldanders.info/images/aozora/card56642.svg" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.aozora.gr.jp/cards/001383/card56642.html">陰翳礼讃</a></dt>
	<dd>谷崎 潤一郎</dd>
    <dd>青空文庫 2016-06-10 (Release 2019-02-24)</dd>
    <dd>Book</dd>
	<dd>card56642 (青空文庫)</dd>
    <dd>評価:<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">実はちゃんと読んでない（笑） 学生時代に読んでおけばよかった。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>reviewer</a> on <abbr class="dtreviewed" title="2019-01-01">2019-01-01</abbr> (powered by <a href="https://www.aozora.gr.jp/" >青空文庫</a>)</p>
</div>`

func TestTemplate(t *testing.T) {
	tc := NewTestAPI()
	book, err := tc.LookupBook("card56642")
	if err != nil {
		t.Errorf("testAPI.LookupBook() error is \"%v\", want nil", err)
		fmt.Printf("Info: %+v\n", err)
		return
	}
	rev := New(
		book,
		WithDate("2019-01-01"),
		WithRating(4),
		WithDescription("実はちゃんと読んでない（笑） 学生時代に読んでおけばよかった。"),
	)
	b, err := rev.Format(bytes.NewBufferString(testTemplate))
	if err != nil {
		t.Errorf("review.Review.Format() error is \"%v\", want nil", err)
		fmt.Printf("Info: %+v\n", err)
		return
	}
	if string(b) != formattedText {
		t.Errorf("review.New() = \"%v\", want \"%v\"", string(b), formattedText)
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
