module github.com/go-http-test

go 1.13

replace github.com/go-http-test/injector => ./injector

replace github.com/go-http-test/model => ./domain/model

replace github.com/go-http-test/repository => ./domain/repository

replace github.com/go-http-test/infrastructure => ./infrastructure

replace github.com/go-http-test/usecase => ./usecase

require (
	github.com/go-http-test/infrastructure v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-http-test/injector v0.0.0-00010101000000-000000000000
	github.com/go-http-test/model v0.0.0-00010101000000-000000000000
	github.com/go-http-test/repository v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-http-test/usecase v0.0.0-00010101000000-000000000000
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/labstack/echo v3.3.10+incompatible
	github.com/stretchr/testify v1.7.0 // indirect
)
