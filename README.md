## install && start

0. clone the repo and enter local repo directory, then execute: `go mod download`
1. prepare mysql connect params
2. generate config file: `cp .env.templ to .env`
3. change your config in `.env`
4. start: `go run main.go`

## packages and document

1. [gin](https://github.com/gin-gonic/gin) and [gin-doc](https://gin-gonic.com/docs/)
2. [sqlx](https://github.com/jmoiron/sqlx) and [sqlx-doc](http://jmoiron.github.io/sqlx/)
3. [validator](https://gopkg.in/go-playground/validator.v9) and [validator-doc](https://godoc.org/gopkg.in/go-playground/validator.v9)