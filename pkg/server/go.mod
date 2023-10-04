module github.com/dressc-go/decryptorservice/pkg/server

go 1.16

replace github.com/dressc-go/decryptorservice/pkg/config => github.com/dressc-go/decryptorservice/pkg/config v0.0.0-20231004124920-edb0a8e10faf

replace github.com/dressc-go/decryptorservice/pkg/cryptkey => github.com/dressc-go/decryptorservice/pkg/cryptkey v0.0.0-20231004152134-271c0b9b389e

replace github.com/dressc-go/decryptors/base64OeapSha256 => github.com/dressc-go/decryptors/base64OeapSha256 v0.0.0-20231004111351-e4120e1a4872

replace github.com/dressc-go/decryptors/base64OeapSha1 => github.com/dressc-go/decryptors/base64OeapSha1 v0.0.0-20231004111351-e4120e1a4872

replace github.com/dressc-go/decryptors/common => github.com/dressc-go/decryptors/common v0.0.0-20231004111351-e4120e1a4872

require (
	github.com/dressc-go/decryptors/base64OeapSha1 v0.0.0-00010101000000-000000000000
	github.com/dressc-go/decryptors/base64OeapSha256 v0.0.0-00010101000000-000000000000
	github.com/dressc-go/decryptorservice/pkg/config v0.0.0-00010101000000-000000000000
	github.com/dressc-go/decryptorservice/pkg/cryptkey v0.0.0-00010101000000-000000000000
	github.com/dressc-go/zlogger v0.0.0-20230514215442-66582730ec37
)
