module github.com/dressc-go/decryptorservice/pkg/server

go 1.16

replace github.com/dressc-go/decryptorservice/pkg/config => ../../pkg/config

replace github.com/dressc-go/decryptorservice/pkg/cryptkey => ../../pkg/cryptkey

replace github.com/dressc-go/decryptors/base64OeapSha256 => github.com/dressc-go/decryptors/base64OeapSha256 v0.0.0-20231004111351-e4120e1a4872

replace github.com/dressc-go/decryptors/common => github.com/dressc-go/decryptors/common v0.0.0-20231004111351-e4120e1a4872

require (
	github.com/dressc-go/decryptors/base64OeapSha1 v0.0.0-20230514215302-511fca39e0a8
	github.com/dressc-go/decryptors/base64OeapSha256 v0.0.0-00010101000000-000000000000
	github.com/dressc-go/decryptorservice/pkg/config v0.0.0-00010101000000-000000000000
	github.com/dressc-go/decryptorservice/pkg/cryptkey v0.0.0-00010101000000-000000000000
	github.com/dressc-go/zlogger v0.0.0-20230514215442-66582730ec37
)
