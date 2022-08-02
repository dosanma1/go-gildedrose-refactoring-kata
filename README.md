# GO Starter

- Run :

```shell
go run texttest_fixture.go [<number-of-days>; default: 2]
```

- Run tests :

```shell
go test ./...
```

- Run tests and coverage :

```shell
go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```

or

```shell
go clean --cache && go test -v -cover github.com/dosanma1/go-gildedrose-refactoring-kata/...
```
