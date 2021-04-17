module github.com/go-http-test/infrastructure

go 1.13

replace github.com/go-http-test/model => ../model

replace github.com/go-http-test/repository => ../model

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
)
