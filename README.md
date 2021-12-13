# Golang Clean Architech Microservice

A boilerplate for a micro-service, written in GoLang, designed based on clean-architecture.
Default module stack
- [Gin](https://github.com/gin-gonic/gin) - Http router engine
- [Gorm](http://gorm.io/) - ORM
- PostgresSQL
- [Golang-Migrate](https://github.com/golang-migrate/migrate) - DB migration engine
- [Testify](https://github.com/stretchr/testify) - Assertion library, Usecase mock, Repository mock 
- [Go-Mocket](https://github.com/Selvatico/go-mocket) - SQL command mock
- [Migrate](https://github.com/golang-migrate/migrate) - Database migrations [Installation](https://github.com/golang-migrate/migrate/tree/master/cli)
- [GoMock](https://github.com/golang/mock) - Golang Mockgen [Installation](https://github.com/golang/mock)
 
## Getting Started

### Prerequisite
- golang
- gomodule
- docker

### Installing

Pull the code under `$GOPATH`

Install richgo
```
brew tap kyoh86/tap
brew install richgo
```

Install mockgen
```
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen
```

```
Gen mock `make mock-gen module="auth"`

Mockgen example command:<br>
`mockgen -source=./app/repositories/users/init.go -destination=./app/mocks/users/repo.go`
```

#### Run app
```
make dev
```

#### Run test
```
make test
```

### Usage

#### Access Private Repo Lib
```
Git can also be configured to use SSH in place of HTTPS for URLs matching a given prefix. For example, to use SSH for all GitHub access, add these lines to your ~/.gitconfig:

[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

#### Access DB
Go to `http://localhost:5432`
See db credential in docker-compose.yaml

#### Golang Lint
```
https://github.com/golangci/golangci-lint
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
```

## Problems
- Test case is not complete


## Clean Architech
![clean architech](https://raw.githubusercontent.com/athiwatp/assetica/master/clean-arch.png)