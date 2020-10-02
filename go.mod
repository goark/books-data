module github.com/spiegel-im-spiegel/books-data

go 1.15

require (
	github.com/spf13/cobra v1.0.1-0.20201001152800-40d34bca1bff
	github.com/spf13/viper v1.7.1
	github.com/spiegel-im-spiegel/aozora-api v0.2.6
	github.com/spiegel-im-spiegel/errs v1.0.2
	github.com/spiegel-im-spiegel/gocli v0.10.3
	github.com/spiegel-im-spiegel/openbd-api v0.2.6
	github.com/spiegel-im-spiegel/pa-api v0.7.2
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
)

replace github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
