## Golang transport generator

### Pre generation
For example `pkg/service/service.go` contains Service interface
```
// Service ...
// @gtg http-server log metrics http-client
type Service interface {
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/v1/namespaces/{namespace}/details/{detail}
	// @gtg http-server-query file={fileID}&some={someID}
	// @gtg http-server-header X-Auth-Token {token}
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-header X-Auth-ID {id}
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error)
	// @gtg http-server-method PUT
	// @gtg http-server-uri-path /api/v1/namespaces/{namespace}/details/{detail}
	// @gtg http-server-query test={testID}&bla={blaID}
	// @gtg http-server-header X-Auth-Token {token}
	// @gtg http-server-json-tag pretty ThePretty
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-header X-Auth-ID {id}
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
}
```
Interface definition requirements:
* First parameter should be `ctx context.Context`
* Last result parameter should be `err error`
* All result parameters should be named

[NOTE]
* URI path parameters should be strings
* Headers parameter should be pointers on string

#### Service interface Tags
**http-server** Generates http server
 
**http-client** Generates http client 

**https-client** Generates https client
 
**logs** Generates log wrapper for service. It uses interface `Logger` from this repository and has adapters for common loggers in `log/adapters` directory. You can use any logging package.
 
**metrics** Generates metrics wrapper for service

By default, only method name and error are presented in metrics. You can add additional parameters to metrics:

metrics(fistParameter,secondParameter)

This parameters must be in all yours methods. Support only strings, ints(int, int32. int64, etc) and pointers on them.

**swagger** Generates swagger for http server

#### Service methods Tags
**http-server-method** one of GET POST PUT PATCH DELETE

**http-server-uri-path** URI path with placeholders 

for example `/api/v1/namespaces/{namespace}/details/{detail}`

fill from URI path to variables `namespace` and `detail` (should be strings)

**http-server-query** Query string with placeholders

for example `file={fileID}&some={someID}`

fill from query string to variables `fileID` and `someID`

**http-server-header** Header string with placeholders

for example `X-Auth-Token {token}`

fill from header `X-Auth-Token` to variable `token` (should be pointers to string)

**http-server-content-type** Header Content-Type 

**http-server-response-header** Response Header string with placeholders

for example `X-Auth-Token {token}`

fill header `X-Auth-Token` from variable `token` (should be pointers to string)

**http-server-response-status** HTTP Response Status Code

**http-server-response-content-type** Response Header Content-Type

**http-server-response-content-encoding** Response Header Content-Type Encoding

for example `utf-8`

**http-server-response-json-tag** Response JSON tag 

for example `nothing TheNothing` 

**http-server-response-file** Allows you to return a file as a set of bytes

**swagger-summary** Method summary for a swagger

for example `Return list of items`

**swagger-description** Method description for a swagger

for example `size - set page size for pagination, offset - set offset for pagination`

### Generation
Build binary and put it into binary path.
```
go build -o ~/go/bin/gtg ./cmd/gtg
```
Run gtg binary from root of your project repository.
You should use `--in` flag to set service folder, by default it's `./pkg/service/`.
```
gtg --in ./example/
```
Run walks all folders from `--in` and find all service interfaces. Very usefull for multiple domains. 
### Deploy
Swagger generation
```
gtg --in="./example" --desc="service description" --title="some title for service" --version="v0.0.10" --servers="http://some.url = some url description\r\nhttp://another.url = another url description"
```

### Docker running example
```
docker run --rm -v $(pwd):/app -w /app/  frankegoesdown/gtg:v0.1.1 gtg --in example/service/
```
