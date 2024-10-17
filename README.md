
# Learning Notes

A brief description of what this project does and who it's for

## Go Commands

```
## Enable dependency tracking
go mod init example/hello

## Run code
go run .

## Add module requirements
go mod tidy

## Module replace
go mod edit -replace example.com/greetings=../greetings
go mod tidy

## Compile to executable
go build
go install
```

## Compile and Install
After running `go build` and `go install` the executable file will generated in `~/go/bin`