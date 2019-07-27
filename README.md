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
  openbd      Search for books data by openBD
  paapi       Search for books data by PA-API
  version     Print the version number

Flags:
      --debug   for debug
  -h, --help    help for books-data
      --pipe    Import description from Stdin
      --raw     Output raw data from openBD

Use "books-data [command] --help" for more information about a command.
```

### Search for books data from [openBD]

```
$ books-data openbd -h
Search for books data by openBD

Usage:
  books-data openbd [flags] [description]

Flags:
  -h, --help                 help for openbd
  -i, --isbn string          ISBN code
  -r, --rating int           Rating of product
      --review-date string   Date of review
  -t, --template string      Template file

Global Flags:
      --debug   for debug
      --pipe    Import description from Stdin
      --raw     Output raw data from openBD

$ books-data openbd -i 427406932X -r 5 "This book is Interesting." | jq .
{
  "Book": {
    "ID": "9784274069321",
    "Title": "リーン開発の現場 : カンバンによる大規模プロジェクトの運営",
    "Image": {
      "URL": "https://cover.openbd.jp/9784274069321.jpg",
      "Height": 0,
      "Width": 0
    },
    "ProductType": "Book",
    "Authors": [
      "Kniberg,Henrik／著 オーム社／著 オーム社開発局／著 市谷聡啓／翻訳 ほか"
    ],
    "Publisher": "オーム社",
    "Codes": [
      {
        "Name": "ISBN",
        "Value": "9784274069321"
      }
    ],
    "PublicationDate": "2013-10-01",
    "LastRelease": "0001-01-01",
    "Service": {
      "Name": "openBD",
      "URL": "https://openbd.jp/"
    }
  },
  "Date": "2019-07-27",
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

$ books-data openbd -i 427406932X -r 5 "This book is Interesting." -t template/template.bib.txt 
@BOOK{Book:9784274069321,
    TITLE = "リーン開発の現場 : カンバンによる大規模プロジェクトの運営",
    AUTHOR = "Kniberg,Henrik／著 オーム社／著 オーム社開発局／著 市谷聡啓／翻訳 ほか",
    PUBLISHER = {オーム社},
    YEAR = 2013
}
```

### Search for books data by [PA-API]

```
$ cat ~/.paapi.yaml
marketplace: webservices.amazon.co.jp
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890

$ books-data paapi -h
Search for books data by PA-API

Usage:
  books-data paapi [flags] [description]

Flags:
      --access-key string      PA-API: Access Key ID
  -a, --asin string            Amazon ASIN code
      --associate-tag string   PA-API: Associate Tag
      --config string          config file (default $HOME/.paapi.yaml)
  -h, --help                   help for paapi
      --marketplace string     PA-API: Marketplace (default "webservices.amazon.co.jp")
  -r, --rating int             Rating of product
      --review-date string     Date of review
      --secret-key string      PA-API: Secret Access Key
  -t, --template string        Template file

Global Flags:
      --debug   for debug
      --pipe    Import description from Stdin
      --raw     Output raw data from openBD

$ books-data paapi -a 427406932X -r 5 "This book is Interesting." | jq .
{
  "Book": {
    "ID": "427406932X",
    "Title": "リーン開発の現場 カンバンによる大規模プロジェクトの運営",
    "URL": "https://www.amazon.co.jp/exec/obidos/ASIN/427406932X/mytag-20",
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
    "LastRelease": "0001-01-01",
    "Service": {
      "Name": "PA-API",
      "URL": "https://affiliate.amazon.co.jp/assoc_credentials/home"
    }
  },
  "Date": "2019-07-27",
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

$ books-data paapi -a 427406932X -r 5 "This book is Interesting." -t template/template.bib.txt 
@BOOK{Book:427406932X,
    TITLE = "リーン開発の現場 カンバンによる大規模プロジェクトの運営",
    AUTHOR = "Henrik Kniberg and 角谷 信太郎 (翻訳) and 市谷 聡啓 (翻訳) and 藤原 大 (翻訳)",
    PUBLISHER = {オーム社},
    YEAR = 2013
}
```

## Reference

- [DDRBoxman/go-amazon-product-api: Wrapper for the Amazon Product Advertising API](https://github.com/DDRBoxman/go-amazon-product-api)
- [seihmd/openbd: openBD API written by Go](https://github.com/seihmd/openbd)
    - [openBDのAPIライブラリをGoでつくりました - Qiita](https://qiita.com/seihmd/items/d1f8b3b54cbc93346d78)

[books-data]: https://github.com/spiegel-im-spiegel/books-data "spiegel-im-spiegel/books-data: Search for Books Data"
[openBD]: https://openbd.jp/
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home
