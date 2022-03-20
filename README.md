# [books-data] -- Search for Books Data

[![check vulns](https://github.com/goark/books-data/workflows/vulns/badge.svg)](https://github.com/goark/books-data/actions)
[![lint status](https://github.com/goark/books-data/workflows/lint/badge.svg)](https://github.com/goark/books-data/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/books-data/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/books-data.svg)](https://github.com/goark/books-data/releases/latest)

This package is required Go 1.16 or later.

**Migrated repository to [github.com/goark/books-data][books-data]**

## Build and Install

```
$ go install github.com/goark/books-data@latest
```

### Binaries

See [latest release](https://github.com/goark/books-data/releases/latest).

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
  -c, --aozora-card string     Aozora-bunko card no.
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default /home/username/.config/books-data/config.yaml)
      --debug                  for debug
  -h, --help                   help for books-data
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "www.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output

Use "books-data [command] --help" for more information about a command.
```

### Config file

```text
$ cat ~/.config/books-data/config.yaml
marketplace: www.amazon.co.jp
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
  -c, --aozora-card string     Aozora-bunko card no.
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default /home/username/.config/books-data/config.yaml)
      --debug                  for debug
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "www.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output

$ books-data search -a 427406932X | jq .
{
  "Type": "paapi",
  "ID": "427406932X",
  "Title": "リーン開発の現場 カンバンによる大規模プロジェクトの運営",
  "URL": "https://www.amazon.co.jp/dp/427406932X?tag=mytag-22&linkCode=ogi&th=1&psc=1",
  "Image": {
    "URL": "https://m.media-amazon.com/images/I/51llL1uygcL._SL160_.jpg",
    "Height": 160,
    "Width": 116
  },
  "ProductType": "単行本（ソフトカバー）",
  "Creators": [
    {
      "Name": "Henrik Kniberg",
      "Role": "著"
    },
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
    },
    {
      "Name": "ISBN",
      "Value": "427406932X"
    }
  ],
  "PublicationDate": "2013-10-26",
  "LastRelease": "",
  "Service": {
    "Name": "PA-APIv5",
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
      --bookpage-url string   URL of book page
  -h, --help                  help for review
      --image-url string      URL of book cover image
      --pipe                  Import description from Stdin
  -r, --rating int            Rating of product
      --review-date string    Date of review

Global Flags:
      --access-key string      Config: PA-API Access Key ID
  -c, --aozora-card string     Aozora-bunko card no.
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default /home/username/.config/books-data/config.yaml)
      --debug                  for debug
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "www.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output

$ books-data review -i 427406932X -r 5 "This book is Interesting." | jq .
{
  "Book": {
    "Type": "openbd",
    "ID": "9784274069321",
    "Title": "リーン開発の現場 : カンバンによる大規模プロジェクトの運営",
    "Image": {
      "URL": "https://cover.openbd.jp/9784274069321.jpg"
    },
    "ProductType": "Book",
    "Creators": [
      {
        "Name": "Kniberg, Henrik"
      },
      {
        "Name": "オーム社"
      },
      {
        "Name": "オーム社開発局"
      },
      {
        "Name": "市谷 聡啓"
      },
      {
        "Name": "藤原 大"
      },
      {
        "Name": "角谷 信太郎"
      },
      {
        "Name": "クニバーグ ヘンリック"
      }
    ],
    "Publisher": "オーム社",
    "Codes": [
      {
        "Name": "ISBN",
        "Value": "9784274069321"
      }
    ],
    "PublicationDate": "2013-10-01",
    "LastRelease": "",
    "Service": {
      "Name": "openBD",
      "URL": "https://openbd.jp/"
    }
  },
  "Date": "2019-10-05",
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
      "Type": "openbd",
      "ID": "9784274069321",
      "Title": "リーン開発の現場 : カンバンによる大規模プロジェクトの運営",
      "Image": {
        "URL": "https://cover.openbd.jp/9784274069321.jpg"
      },
      "ProductType": "Book",
      "Creators": [
        {
          "Name": "Kniberg, Henrik"
        },
        {
          "Name": "オーム社"
        },
        {
          "Name": "オーム社開発局"
        },
        {
          "Name": "市谷 聡啓"
        },
        {
          "Name": "藤原 大"
        },
        {
          "Name": "角谷 信太郎"
        },
        {
          "Name": "クニバーグ ヘンリック"
        }
      ],
      "Publisher": "オーム社",
      "Codes": [
        {
          "Name": "ISBN",
          "Value": "9784274069321"
        }
      ],
      "PublicationDate": "2013-10-01",
      "LastRelease": "",
      "Service": {
        "Name": "openBD",
        "URL": "https://openbd.jp/"
      }
    },
    "Date": "2019-10-05",
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
  -c, --aozora-card string     Aozora-bunko card no.
  -a, --asin string            Amazon ASIN code
      --associate-tag string   Config: PA-API Associate Tag
      --config string          Config file (default /home/username/.config/books-data/config.yaml)
      --debug                  for debug
  -i, --isbn string            ISBN code
      --marketplace string     Config: PA-API Marketplace (default "www.amazon.co.jp")
  -l, --review-log string      Config: Review log file (JSON format)
      --secret-key string      Config: PA-API Secret Access Key
  -t, --template-file string   Template file for formatted output

$ books-data history -i 9784274069321 | jq .
{
  "Book": {
    "Type": "openbd",
    "ID": "9784274069321",
    "Title": "リーン開発の現場 : カンバンによる大規模プロジェクトの運営",
    "Image": {
      "URL": "https://cover.openbd.jp/9784274069321.jpg"
    },
    "ProductType": "Book",
    "Creators": [
      {
        "Name": "Kniberg, Henrik"
      },
      {
        "Name": "オーム社"
      },
      {
        "Name": "オーム社開発局"
      },
      {
        "Name": "市谷 聡啓"
      },
      {
        "Name": "藤原 大"
      },
      {
        "Name": "角谷 信太郎"
      },
      {
        "Name": "クニバーグ ヘンリック"
      }
    ],
    "Publisher": "オーム社",
    "Codes": [
      {
        "Name": "ISBN",
        "Value": "9784274069321"
      }
    ],
    "PublicationDate": "2013-10-01",
    "LastRelease": "",
    "Service": {
      "Name": "openBD",
      "URL": "https://openbd.jp/"
    }
  },
  "Date": "2019-10-05",
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
    AUTHOR = "Henrik Kniberg (著) and 角谷 信太郎 (翻訳) and 市谷 聡啓 (翻訳) and 藤原 大 (翻訳)",
    PUBLISHER = {オーム社},
    YEAR = 2013
}

$ books-data history -i 9784274069321 -t testdata/review-template/template.html
<div class="hreview">
  <div class="photo"><img src="https://cover.openbd.jp/9784274069321.jpg" alt="photo"></div>
  <dl class="fn">
    <dt>リーン開発の現場 : カンバンによる大規模プロジェクトの運営</dt>
    <dd>Kniberg, Henrik, オーム社, オーム社開発局, 市谷 聡啓, 藤原 大, 角谷 信太郎, クニバーグ ヘンリック</dd>
    <dd>オーム社 2013-10-01</dd>
    <dd>Book</dd>
    <dd>9784274069321 (ISBN)</dd>
    <dd>Rating<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">This book is Interesting.</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-10-05">2019-10-05</abbr> (powered by <a href="https://openbd.jp/" >openBD</a>)</p>
</div>
```

## Modules Requirement Graph

[![dependency.png](./dependency.png)](./dependency.png)

## Reference

- [DDRBoxman/go-amazon-product-api: Wrapper for the Amazon Product Advertising API](https://github.com/DDRBoxman/go-amazon-product-api)
- [seihmd/openbd: openBD API written by Go](https://github.com/seihmd/openbd)
    - [openBDのAPIライブラリをGoでつくりました - Qiita](https://qiita.com/seihmd/items/d1f8b3b54cbc93346d78)
- [goark/aozora-api: APIs for Aozora-bunko RESTful Service by Golang](https://github.com/goark/aozora-api)

[books-data]: https://github.com/goark/books-data "goark/books-data: Search for Books Data"
[openBD]: https://openbd.jp/
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home
