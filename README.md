# [books-data] -- Search for Books Data

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/books-data.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/books-data)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/books-data/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/books-data.svg)](https://github.com/spiegel-im-spiegel/books-data/releases/latest)

## Download and Build

```
$ go get github.com/spiegel-im-spiegel/books-data@latest
```

### Binaries

See [latest release](https://github.com/spiegel-im-spiegel/books-data/releases/latest).

## Usage

```
$ books-data -h
Search for books data

Usage:
  books-data [flags]
  books-data [command]

Available Commands:
  help        Help about any command
  history     Lookup review data from history log
  review      Make review data
  search      Search for books data
  version     Print the version number

Flags:
      --access-key string      Config: PA-API Access Key ID
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default $HOME/.books-data.yaml)
      --debug                  for debug
  -h, --help                   help for books-data
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "webservices.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output

Use "books-data [command] --help" for more information about a command.
```

### Config file

```text
$ cat ~/.paapi.yaml
marketplace: webservices.amazon.co.jp
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890
review-log: /home/username/review-log.json
```

### Search for books data

```
$ books-data search -h
Search for books data

Usage:
  books-data search [flags]

Flags:
  -h, --help   help for search
      --raw    Output raw data from API

Global Flags:
      --access-key string      Config: PA-API Access Key ID
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default $HOME/.books-data.yaml)
      --debug                  for debug
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "webservices.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output

  $ books-data search -a 427406932X | jq .
  {
    "Type": "paapi",
    "ID": "427406932X",
    "Title": "リーン開発の現場 カンバンによる大規模プロジェクトの運営",
    "URL": "https://www.amazon.co.jp/exec/obidos/ASIN/427406932X/baldandersinf-22/",
    "Image": {
      "URL": "https://images-fe.ssl-images-amazon.com/images/I/51llL1uygcL._SL160_.jpg",
      "Height": 160,
      "Width": 116
    },
    "ProductType": "単行本（ソフトカバー）",
    "Authors": [
      "Henrik Kniberg"
    ],
    "Creators": [
      {
        "Name": "角谷 信太郎",
        "Role": "翻訳"
      },
      {
        "Name": "市谷 聡啓",
        "Role": "翻訳"
      },
      {
        "Name": "藤原 大",
        "Role": "翻訳"
      }
    ],
    "Publisher": "オーム社",
    "Codes": [
      {
        "Name": "ASIN",
        "Value": "427406932X"
      },
      {
        "Name": "EAN",
        "Value": "9784274069321"
      }
    ],
    "PublicationDate": "2013-10-26",
    "LastRelease": "",
    "Service": {
      "Name": "PA-API",
      "URL": "https://affiliate.amazon.co.jp/assoc_credentials/home"
    }
  }
```

### Make review data

```text
$ books-data review -h
Make review data

Usage:
  books-data review [flags] [description]

Flags:
  -h, --help                 help for review
      --pipe                 Import description from Stdin
  -r, --rating int           Rating of product
      --review-date string   Date of review

Global Flags:
      --access-key string      Config: PA-API Access Key ID
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default $HOME/.books-data.yaml)
      --debug                  for debug
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "webservices.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output

$ books-data review -i 427406932X -r 5 "This book is Interesting." | jq .
{
  "Book": {
    "Type": "paapi",
    "ID": "427406932X",
    "Title": "リーン開発の現場 カンバンによる大規模プロジェクトの運営",
    "URL": "https://www.amazon.co.jp/exec/obidos/ASIN/427406932X/baldandersinf-22/",
    "Image": {
      "URL": "https://images-fe.ssl-images-amazon.com/images/I/51llL1uygcL._SL160_.jpg",
      "Height": 160,
      "Width": 116
    },
    "ProductType": "単行本（ソフトカバー）",
    "Authors": [
      "Henrik Kniberg"
    ],
    "Creators": [
      {
        "Name": "角谷 信太郎",
        "Role": "翻訳"
      },
      {
        "Name": "市谷 聡啓",
        "Role": "翻訳"
      },
      {
        "Name": "藤原 大",
        "Role": "翻訳"
      }
    ],
    "Publisher": "オーム社",
    "Codes": [
      {
        "Name": "ASIN",
        "Value": "427406932X"
      },
      {
        "Name": "EAN",
        "Value": "9784274069321"
      }
    ],
    "PublicationDate": "2013-10-26",
    "LastRelease": "",
    "Service": {
      "Name": "PA-API",
      "URL": "https://affiliate.amazon.co.jp/assoc_credentials/home"
    }
  },
  "Date": "2019-08-04",
  "Rating": 5,
  "Star": [
    true,
    true,
    true,
    true,
    true
  ],
  "Description": "This book is Interesting."
}

$ cat ~/review-log.json
[
  {
    "Book": {
      "Type": "paapi",
      "ID": "427406932X",
      "Title": "リーン開発の現場 カンバンによる大規模プロジェクトの運営",
      "URL": "https://www.amazon.co.jp/exec/obidos/ASIN/427406932X/baldandersinf-22/",
      "Image": {
        "URL": "https://images-fe.ssl-images-amazon.com/images/I/51llL1uygcL._SL160_.jpg",
        "Height": 160,
        "Width": 116
      },
      "ProductType": "単行本（ソフトカバー）",
      "Authors": [
        "Henrik Kniberg"
      ],
      "Creators": [
        {
          "Name": "角谷 信太郎",
          "Role": "翻訳"
        },
        {
          "Name": "市谷 聡啓",
          "Role": "翻訳"
        },
        {
          "Name": "藤原 大",
          "Role": "翻訳"
        }
      ],
      "Publisher": "オーム社",
      "Codes": [
        {
          "Name": "ASIN",
          "Value": "427406932X"
        },
        {
          "Name": "EAN",
          "Value": "9784274069321"
        }
      ],
      "PublicationDate": "2013-10-26",
      "LastRelease": "",
      "Service": {
        "Name": "PA-API",
        "URL": "https://affiliate.amazon.co.jp/assoc_credentials/home"
      }
    },
    "Date": "2019-08-04",
    "Rating": 5,
    "Star": [
      true,
      true,
      true,
      true,
      true
    ],
    "Description": "This book is Interesting."
  }
]
```

### Lookup review data from history

```text
$ books-data history -h
Lookup review data from history log

Usage:
  books-data history [flags]

Flags:
  -h, --help   help for history

Global Flags:
      --access-key string      Config: PA-API Access Key ID
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default $HOME/.books-data.yaml)
      --debug                  for debug
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "webservices.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output

$ books-data history -a 427406932X | jq .
{
  "Book": {
    "Type": "paapi",
    "ID": "427406932X",
    "Title": "リーン開発の現場 カンバンによる大規模プロジェクトの運営",
    "URL": "https://www.amazon.co.jp/exec/obidos/ASIN/427406932X/baldandersinf-22/",
    "Image": {
      "URL": "https://images-fe.ssl-images-amazon.com/images/I/51llL1uygcL._SL160_.jpg",
      "Height": 160,
      "Width": 116
    },
    "ProductType": "単行本（ソフトカバー）",
    "Authors": [
      "Henrik Kniberg"
    ],
    "Creators": [
      {
        "Name": "角谷 信太郎",
        "Role": "翻訳"
      },
      {
        "Name": "市谷 聡啓",
        "Role": "翻訳"
      },
      {
        "Name": "藤原 大",
        "Role": "翻訳"
      }
    ],
    "Publisher": "オーム社",
    "Codes": [
      {
        "Name": "ASIN",
        "Value": "427406932X"
      },
      {
        "Name": "EAN",
        "Value": "9784274069321"
      }
    ],
    "PublicationDate": "2013-10-26",
    "LastRelease": "",
    "Service": {
      "Name": "PA-API",
      "URL": "https://affiliate.amazon.co.jp/assoc_credentials/home"
    }
  },
  "Date": "2019-08-04",
  "Rating": 5,
  "Star": [
    true,
    true,
    true,
    true,
    true
  ],
  "Description": "This book is Interesting."
}
```

### Formatted output by template file

```text
$ books-data search -a 427406932X -t testdata/book-template/template.bib.txt
@BOOK{Book:427406932X,
    TITLE = "リーン開発の現場 カンバンによる大規模プロジェクトの運営",
    AUTHOR = "Henrik Kniberg and 角谷 信太郎 (翻訳) and 市谷 聡啓 (翻訳) and 藤原 大 (翻訳)",
    PUBLISHER = {オーム社},
    YEAR = 2013
}

$ books-data history -a 427406932X -t testdata/review-template/template.html
<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/427406932X/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51llL1uygcL._SL160_.jpg" width="116" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/427406932X/baldandersinf-22/">リーン開発の現場 カンバンによる大規模プロジェクトの運営</a></dt>
    <dd>Henrik Kniberg</dd>
    <dd>角谷 信太郎 (翻訳), 市谷 聡啓 (翻訳), 藤原 大 (翻訳)</dd>
    <dd>オーム社 2013-10-26</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>427406932X (ASIN), 9784274069321 (EAN)</dd>
    <dd>Rating<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">This book is Interesting.</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-08-04">2019-08-04</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home" >PA-API</a>)</p>
</div>
```

## Reference

- [DDRBoxman/go-amazon-product-api: Wrapper for the Amazon Product Advertising API](https://github.com/DDRBoxman/go-amazon-product-api)
- [seihmd/openbd: openBD API written by Go](https://github.com/seihmd/openbd)
    - [openBDのAPIライブラリをGoでつくりました - Qiita](https://qiita.com/seihmd/items/d1f8b3b54cbc93346d78)

[books-data]: https://github.com/spiegel-im-spiegel/books-data "spiegel-im-spiegel/books-data: Search for Books Data"
[openBD]: https://openbd.jp/
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home
