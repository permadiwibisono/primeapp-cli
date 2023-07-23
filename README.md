# primeapp-cli
simple primeapp with go

### Run
```bash
go run .
```

### Run Test
```bash
go test .

// with verbose
go test -v .

// spesific test case
go test -v -run Test_readInput

// with coverage
go test -cover .
go test -coverprofile=coverage.out .
```

### Run coverage
```bash
go tool cover -html=coverage.out
```