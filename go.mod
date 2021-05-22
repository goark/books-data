module github.com/spiegel-im-spiegel/books-data

go 1.16

require (
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/spiegel-im-spiegel/aozora-api v0.3.0
	github.com/spiegel-im-spiegel/errs v1.0.2
	github.com/spiegel-im-spiegel/gocli v0.10.4
	github.com/spiegel-im-spiegel/openbd-api v0.3.0
	github.com/spiegel-im-spiegel/pa-api v0.9.0
)

replace golang.org/x/sys v0.0.0-20190624142023-c5567b49c5d0 => golang.org/x/sys v0.0.0-20210521090106-6ca3eb03dfc2
