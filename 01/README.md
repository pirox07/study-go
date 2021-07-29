# 01

## やること

- OneLogin API を使ってユーザ情報を取得する
- API にアクセスするための Token を取得する

```shell
git clone xxx
cd 01
# API Credential
export CLIENT_ID=< Client ID >
export CLIENT_SECRET=< Client Secret >

go build main.go
./main
```

## 参考
- [http - pkg.go.dev](https://pkg.go.dev/net/http@go1.16.6)
- [Generate Tokens - developers.onelogin.com](https://developers.onelogin.com/api-docs/2/oauth20-tokens/generate-tokens-2)
- [Get User - developers.onelogin.com](https://developers.onelogin.com/api-docs/2/users/get-user)
- [JSON-to-Go](https://mholt.github.io/json-to-go/)
