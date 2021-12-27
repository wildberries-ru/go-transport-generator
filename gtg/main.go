package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/vetcher/go-astra/types"

	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
	"github.com/wildberries-ru/go-transport-generator/pkg/api"
	"github.com/wildberries-ru/go-transport-generator/pkg/imports"
	"github.com/wildberries-ru/go-transport-generator/pkg/mod"
	request2 "github.com/wildberries-ru/go-transport-generator/pkg/parser"
	"github.com/wildberries-ru/go-transport-generator/pkg/parser/httpserver/log"
	"github.com/wildberries-ru/go-transport-generator/pkg/parser/httpserver/request"
	"github.com/wildberries-ru/go-transport-generator/pkg/parser/httpserver/response"
	swagger2 "github.com/wildberries-ru/go-transport-generator/pkg/parser/swagger"
	"github.com/wildberries-ru/go-transport-generator/pkg/preprocessor"
	"github.com/wildberries-ru/go-transport-generator/pkg/processor"
	"github.com/wildberries-ru/go-transport-generator/pkg/render/httpclient"
	"github.com/wildberries-ru/go-transport-generator/pkg/render/httperrors"
	"github.com/wildberries-ru/go-transport-generator/pkg/render/httpserver"
	"github.com/wildberries-ru/go-transport-generator/pkg/render/service"
)

const (
	httpServer           = "http-server"
	httpClient           = "http-client"
	httpsClient          = "https-client"
	httpsClientInsecure  = "https-client-insecure"
	httpErrors           = "http-errors"
	instrumentingService = "metrics"
	logService           = "log"
	mockService          = "mock"
	swagger              = "swagger"

	requestAPIPathSuffix           = "api-path"
	requestContentTypeSuffix       = "content-type"
	requestMultipartValueTagSuffix = "value-tag"
	requestMultipartFileTagSuffix  = "file-tag"
	requestFormUrlencodedTagSuffix = "form-urlencoded"
	requestJSONTagSuffix           = "json-tag"
	requestPlainObjectSuffix       = "plain-object"
	requestHeaderSuffix            = "header"
	requestCookieSuffix            = "cookie"
	requestMethodSuffix            = "method"
	requestQuerySuffix             = "query"
	requestURIPathSuffix           = "uri-path"
	requestErrorsSuffix            = "errors"
	responseContentTypeSuffix      = "response-content-type"
	responseContentEncodingSuffix  = "response-content-encoding"
	responseJSONTagSuffix          = "response-json-tag"
	responseBodySuffix             = "response-body"
	responseCookieSuffix           = "response-cookie"
	responseHeaderSuffix           = "response-header"
	responseFileSuffix             = "response-file"
	responseStatusSuffix           = "response-status"
	swaggerDescriptionSuffix       = "description"
	swaggerServersSuffix           = "servers"
	swaggerSummarySuffix           = "summary"
	swaggerTitleSuffix             = "title"
	swaggerVersionSuffix           = "version"
	ignoreSuffix                   = "ignore"

	httpServerPkgName = "httpserver"
	httpClientPkgName = "httpclient"
	httpErrorsPkgName = "httperrors"

	swaggerFilename = "swagger"

	tagMark = "@gtg"
)

var (
	httpServerFilePath           = []string{"httpserver", "server.go"}
	httpServerTransportFilePath  = []string{"httpserver", "transport.go"}
	httpServerBuilderFilePath    = []string{"httpserver", "builder.go"}
	httpClientFilePath           = []string{"httpclient", "client.go"}
	httpClientTransportFilePath  = []string{"httpclient", "transport.go"}
	httpClientBuilderFilePath    = []string{"httpclient", "builder.go"}
	httpUIErrorsFilePath         = []string{"httperrors", "ui.go"}
	httpClientErrorsFilePath     = []string{"httperrors", "client.go"}
	serviceInstrumentingFilePath = []string{"instrumenting.go"}
	serviceLoggingFilePath       = []string{"logging.go"}
	serviceMockFilePath          = []string{"httpclient", "service_mock.go"}

	goGeneratedAutomaticallyPrefix = []byte("// CODE GENERATED AUTOMATICALLY")
	yaml                           = true

	version = "v0.0.1"
)

func main() {
	var (
		info api.GenerationInfo
	)
	inFile := flag.String("in", "./pkg/service", "relative path to dir with services")
	// swagger
	swaggerFile := flag.String("swagger", ".", "relative path to swagger file generate")
	info.SwaggerToJSON = flag.Bool("json", false, "use json swagger output file")
	info.SwaggerToYaml = flag.Bool("yaml", false, "use yaml swagger output file")
	info.SwaggerInfo.Description = flag.String("desc", "", "swagger description")
	info.SwaggerInfo.Title = flag.String("title", "", "swagger title")
	info.SwaggerInfo.Version = flag.String("version", "", "swagger version")
	servers := flag.String("servers", "", `swagger servers in format: http://some.url = some url description\r\nhttp://another.url = another url description`)
	//templates
	// http client
	httpClientBuilderFile := flag.String("http-client-builder", "./templates/httpclient/builder", "")
	httpClientBuilderTestFile := flag.String("http-client-builder-test", "./templates/httpclient/builder_test", "")
	httpClientClientFile := flag.String("http-client-client", "./templates/httpclient/client", "")
	httpClientClientTestFile := flag.String("http-client-client-test", "./templates/httpclient/client_test", "")
	httpClientTransportFile := flag.String("http-client-transport", "./templates/httpclient/transport", "")
	// http errors
	httpErrorsClientFile := flag.String("http-errors-client", "./templates/httperrors/client", "")
	httpErrorsUIFile := flag.String("http-errors-ui", "./templates/httperrors/ui", "")
	// http server
	httpServerBuilderFile := flag.String("http-server-builder", "./templates/httpserver/builder", "")
	httpServerServerFile := flag.String("http-server-server", "./templates/httpserver/server", "")
	httpServerTransportFile := flag.String("http-server-transport", "./templates/httpserver/transport", "")
	// service
	serviceInstrumentingFile := flag.String("service-instrumenting", "./templates/service/instrumenting", "")
	serviceLoggingFile := flag.String("service-logging", "./templates/service/logging", "")
	serviceMockFile := flag.String("service-mock", "./templates/service/mock", "")
	// version
	printVersion := flag.Bool("v", false, "print version and exit")
	flag.Parse()
	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}
	// read templates
	var (
		fp  string
		err error
	)
	fp, err = filepath.Abs(*httpClientBuilderFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpClientBuilder, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*httpClientBuilderTestFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpClientBuilderTest, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*httpClientClientFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpClientClient, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*httpClientClientTestFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpClientClientTest, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*httpClientTransportFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpClientTransport, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*httpErrorsClientFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpErrorsClient, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*httpErrorsUIFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpErrorsUI, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*httpServerBuilderFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpServerBuilder, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*httpServerServerFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpServerServer, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*httpServerTransportFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpServerTransport, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*serviceInstrumentingFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	serviceInstrumenting, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*serviceLoggingFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	serviceLogging, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fp, err = filepath.Abs(*serviceMockFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	serviceMock, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *servers != "" {
		for _, srv := range strings.Split(*servers, `\r\n`) {
			srvs := strings.Split(srv, " = ")
			if len(srvs) != 2 {
				fmt.Println("wrong servers parameters, try use `http://some.url = some url description\r\nhttp://another.url = another url description`")
				os.Exit(0)
			}
			info.SwaggerInfo.Servers = append(info.SwaggerInfo.Servers, v1.Server{
				URL:         srvs[0],
				Description: srvs[1],
			})
		}
	}

	if !*info.SwaggerToJSON && !*info.SwaggerToYaml {
		info.SwaggerToYaml = &yaml
	}

	tagsParser := request.NewErrorProcessor(httpServer, requestErrorsSuffix,
		request.NewURIPath(httpServer, requestURIPathSuffix,
			request.NewQuery(httpServer, requestQuerySuffix,
				request.NewMethod(httpServer, requestMethodSuffix,
					request.NewHeader(httpServer, requestHeaderSuffix,
						request.NewCookie(httpServer, requestCookieSuffix,
							request.NewContentType(httpServer, requestContentTypeSuffix,
								request.NewMultipartFileTag(httpServer, requestMultipartFileTagSuffix,
									request.NewMultipartValueTag(httpServer, requestMultipartValueTagSuffix,
										request.NewJSONTag(httpServer, requestJSONTagSuffix,
											request.NewPlainObjectTag(httpServer, requestPlainObjectSuffix,
												request.NewAPIPath(httpServer, requestAPIPathSuffix,
													request.NewFormUrlencodedTag(httpServer, requestFormUrlencodedTagSuffix, &request2.Term{})))))))))))))
	tagsParser = response.NewStatus(httpServer, responseStatusSuffix,
		response.NewCookies(httpServer, responseCookieSuffix,
			response.NewHeader(httpServer, responseHeaderSuffix,
				response.NewContentType(httpServer, responseContentTypeSuffix,
					response.NewEncodingType(httpServer, responseContentEncodingSuffix,
						response.NewJSONTag(httpServer, responseJSONTagSuffix,
							response.NewFile(httpServer, responseFileSuffix,
								response.NewBody(httpServer, responseBodySuffix, tagsParser))))))))
	tagsParser = log.NewLogIgnore(logService, ignoreSuffix, tagsParser)
	swaggerMethodTagParser := swagger2.NewVersion(swagger, swaggerVersionSuffix,
		swagger2.NewTitle(swagger, swaggerTitleSuffix,
			swagger2.NewSummary(swagger, swaggerSummarySuffix,
				swagger2.NewServers(swagger, swaggerServersSuffix,
					swagger2.NewDescription(swagger, swaggerDescriptionSuffix, &swagger2.Term{})))))

	t := template.New("")
	t.Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }})
	t.Funcs(template.FuncMap{"len": func(s []string) int { return len(s) }})
	t.Funcs(template.FuncMap{"lenVariables": func(s []types.Variable) int { return len(s) }})
	t.Funcs(template.FuncMap{"lenMap": func(s map[string]string) int { return len(s) }})
	t.Funcs(template.FuncMap{"length": func(s string) int { return len(s) }})
	t.Funcs(template.FuncMap{"popFirst": func(s []types.Variable) []types.Variable { return s[1:] }})
	t.Funcs(template.FuncMap{"popLast": func(s []types.Variable) []types.Variable { return s[:len(s)-1] }})
	t.Funcs(template.FuncMap{"joinFullVariables": func(s []types.Variable, c string, any ...string) string {
		t := make([]string, len(s))
		for i, m := range s {
			t[i] = m.String()
		}
		if len(any) > 0 {
			t = append(t, any...)
		}
		return strings.Join(t, c)
	}})
	t.Funcs(template.FuncMap{"joinVariableNames": func(s []types.Variable, c string, any ...string) string {
		t := make([]string, len(s))
		for i, m := range s {
			t[i] = m.Name
		}
		if len(any) > 0 {
			t = append(t, any...)
		}
		return strings.Join(t, c)
	}})
	t.Funcs(template.FuncMap{"joinVariableNamesWithEllipsis": func(s []types.Variable, c string, any ...string) string {
		t := make([]string, len(s))
		for i, m := range s {
			switch m.Type.(type) {
			case types.TEllipsis:
				t[i] = m.Name + "..."
			default:
				t[i] = m.Name
			}
		}
		if len(any) > 0 {
			t = append(t, any...)
		}
		return strings.Join(t, c)
	}})
	t.Funcs(template.FuncMap{"low": func(s string) string {
		a := []rune(s)
		a[0] = unicode.ToLower(a[0])
		return string(a)
	}})
	t.Funcs(template.FuncMap{"join": func(s []string, c string) string {
		if len(s) > 0 {
			return strings.Join(s, c)
		}
		return ""
	}})
	t.Funcs(template.FuncMap{"concat": func(s string, c string) string {
		return s + c
	}})
	t.Funcs(template.FuncMap{"up": func(s string) string {
		if s == "id" {
			return "ID"
		}
		a := []rune(s)
		a[0] = unicode.ToUpper(a[0])
		return string(a)
	}})
	t.Funcs(template.FuncMap{"stripType": func(s string) string {
		parts := strings.Split(s, ".")
		return parts[len(parts)-1]
	}})
	t.Funcs(template.FuncMap{"isSliceType": func(s string) bool {
		return len(s) > 2 && s[0] == '['
	}})
	t.Funcs(template.FuncMap{"isMapType": func(s string) bool {
		return len(s) > 3 && s[0] == 'm' && s[1] == 'a' && s[2] == 'p'
	}})
	t.Funcs(template.FuncMap{"isError": func(t types.Type) bool {
		return strings.EqualFold(t.String(), "error")
	}})
	t.Funcs(template.FuncMap{"isFileHeader": func(t types.Type) bool {
		return strings.HasSuffix(t.String(), "multipart.FileHeader")
	}})
	t.Funcs(template.FuncMap{"isFileHeaderPlaceholder": func(t string) bool {
		return strings.HasSuffix(t, "multipart.FileHeader")
	}})
	t.Funcs(template.FuncMap{"isBytesPlaceholder": func(t string) bool {
		return strings.HasSuffix(t, "[]byte")
	}})
	t.Funcs(template.FuncMap{"notin": func(s []string, f string) bool {
		for _, v := range s {
			if v == f {
				return false
			}
		}
		return true
	}})
	t.Funcs(template.FuncMap{"in": func(s map[string]*api.MetricsPlaceholder, f string) bool {
		for _, v := range s {
			if v.Name == f {
				return true
			}
		}
		return false
	}})
	t.Funcs(template.FuncMap{"contains": func(s string, substr string) bool {
		return strings.Contains(s, substr)
	}})

	imp := imports.NewImports()

	httpServerRender := httpserver.NewServer(t, httpServerPkgName, httpServerFilePath, imp, string(httpServerServer))
	httpServerTransportRender := httpserver.NewTransport(t, httpServerPkgName, httpServerTransportFilePath, imp, string(httpServerTransport))
	httpServerBuilderRender := httpserver.NewBuilder(t, httpServerPkgName, httpServerBuilderFilePath, imp, string(httpServerBuilder))
	httpClientRender := httpclient.NewClient(t, httpClientPkgName, httpClientFilePath, imp, string(httpClientClient), string(httpClientClientTest))
	httpClientTransportRender := httpclient.NewTransport(t, httpClientPkgName, httpClientTransportFilePath, imp, string(httpClientTransport))
	httpClientBuilderRender := httpclient.NewBuilder(t, httpClientPkgName, httpClientBuilderFilePath, imp, string(httpClientBuilder), string(httpClientBuilderTest))
	httpUIErrorsRender := httperrors.NewUI(t, httpErrorsPkgName, httpUIErrorsFilePath, imp, string(httpErrorsUI))
	httpClientErrorsRender := httperrors.NewClient(t, httpErrorsPkgName, httpClientErrorsFilePath, imp, string(httpErrorsClient))
	instrumentingRender := service.NewInstrumenting(t, serviceInstrumentingFilePath, imp, string(serviceInstrumenting))
	loggingRender := service.NewLogging(t, serviceLoggingFilePath, imp, string(serviceLogging))
	mockRender := service.NewMock(t, httpClientPkgName, serviceMockFilePath, imp, string(serviceMock))
	swaggerRender := httpserver.NewSwagger(swaggerFilename)

	httpMethodProcessor := processor.NewHTTPMethod(tagMark, tagsParser)
	processors := map[string]processor.Processor{
		httpServer: processor.NewHTTPServer(
			httpServerRender,
			httpServerTransportRender,
			httpServerBuilderRender,
		),
		httpClient: processor.NewHTTPClient(
			false,
			false,
			httpClientRender,
			httpClientTransportRender,
			httpClientBuilderRender,
		),
		httpsClient: processor.NewHTTPClient(
			true,
			false,
			httpClientRender,
			httpClientTransportRender,
			httpClientBuilderRender,
		),
		httpsClientInsecure: processor.NewHTTPClient(
			true,
			true,
			httpClientRender,
			httpClientTransportRender,
			httpClientBuilderRender,
		),
		httpErrors:           processor.NewErrors(tagMark, httpUIErrorsRender, httpClientErrorsRender),
		instrumentingService: processor.NewInstrumenting(instrumentingRender),
		logService:           processor.NewLogging(loggingRender),
		mockService:          processor.NewMock(mockRender),
		swagger:              processor.NewSwagger(tagMark, httpMethodProcessor, swaggerMethodTagParser, mod.NewMod(), goGeneratedAutomaticallyPrefix),
	}

	servicesProcessor := processor.NewServices(tagMark, processors, httpMethodProcessor, instrumentingService)
	servicePreProcessor := preprocessor.NewService(servicesProcessor, goGeneratedAutomaticallyPrefix, swaggerRender)

	info.SwaggerAbsOutputPath = *swaggerFile
	err = servicePreProcessor.Process(*inFile, *inFile, &info)
	if err != nil {
		fmt.Printf("servicePreProcessor.Process error: %v\n", err)
	}

	os.Exit(0)
}
