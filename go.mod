module github.com/dressc-go/decryptorservice

go 1.16

replace github.com/dressc-go/decryptorservice/pkg/config => github.com/dressc-go/decryptorservice/pkg/config v0.0.0-20231004123244-ad8d1dfc077b

replace github.com/dressc-go/decryptorservice/pkg/server => github.com/dressc-go/decryptorservice/pkg/server v0.0.0-20231004123244-ad8d1dfc077b

require (
	github.com/dressc-go/decryptorservice/pkg/config v0.0.0-20231004123244-ad8d1dfc077b // indirect
	github.com/dressc-go/zlogger v0.0.0-20230514215442-66582730ec37
	github.com/pkg/errors v0.9.1
)
