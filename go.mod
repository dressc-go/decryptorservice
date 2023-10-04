module github.com/dressc-go/decryptorservice

go 1.16

replace github.com/dressc-go/decryptorservice/pkg/config => github.com/dressc-go/decryptorservice/pkg/config v0.0.0-20231004123244-ad8d1dfc077b

replace github.com/dressc-go/decryptorservice/pkg/server => github.com/dressc-go/decryptorservice/pkg/server v0.0.0-20231004123244-ad8d1dfc077b

replace github.com/dressc-go/decryptors/base64OeapSha256 => github.com/dressc-go/decryptors/base64OeapSha256 v0.0.0-20231004111351-e4120e1a4872

replace github.com/dressc-go/decryptors/base64OeapSha1 => github.com/dressc-go/decryptors/base64OeapSha1 v0.0.0-20231004111351-e4120e1a4872

replace github.com/dressc-go/decryptorservice/pkg/cryptkey => github.com/dressc-go/decryptorservice/pkg/cryptkey v0.0.0-20231004124920-edb0a8e10faf

replace github.com/dressc-go/decryptors/common => github.com/dressc-go/decryptors/common v0.0.0-20231004111351-e4120e1a4872

require (
	github.com/dressc-go/decryptors/common v0.0.0-20231004111351-e4120e1a4872 // indirect
	github.com/dressc-go/decryptorservice/pkg/config v0.0.0-20231004123244-ad8d1dfc077b
	github.com/dressc-go/decryptorservice/pkg/server v0.0.0-00010101000000-000000000000
	github.com/dressc-go/zlogger v0.0.0-20230514215442-66582730ec37
	github.com/pkg/errors v0.9.1
)
