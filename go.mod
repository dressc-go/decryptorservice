module github.com/dressc-go/decryptorservice

go 1.16

replace github.com/dressc-go/decryptorservice/pkg/config => ./pkg/config

replace github.com/dressc-go/decryptorservice/pkg/server => ./pkg/server

require (
	github.com/dressc-go/zlogger v0.0.0-20230514215442-66582730ec37
	github.com/pkg/errors v0.9.1
)
