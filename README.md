# test-jellysmack

### Install dependencies
```
go mod download
```

### Usage
cd to root directory, and run
```
go run cmd/mow-solver/main.go [-f=input_file_path]
```

### Manuel
```
go run cmd/mow-solver/main.go -h
Usage of ./test-jellysmack:
  -f string
        path of input file (default "./input.txt")
```

### Test
```
go test ./...
```

### Algorithm
```
/internal/service/mower.go
```
