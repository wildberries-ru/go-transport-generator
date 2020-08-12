package main

import (
	"flag"
	"fmt"
	"os"
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
	httpErrors           = "http-errors"
	instrumentingService = "metrics"
	logService           = "log"
	mockService          = "mock"
	swagger              = "swagger"

	requestAPIPathSuffix           = "api-path"
	requestContentTypeSuffix       = "content-type"
	requestMultipartValueTagSuffix = "value-tag"
	requestMultipartFileTagSuffix  = "file-tag"
	requestJsonTagSuffix           = "json-tag"
	requestHeaderSuffix            = "header"
	requestMethodSuffix            = "method"
	requestQuerySuffix             = "query"
	requestURIPathSuffix           = "uri-path"
	requestErrorsSuffix            = "errors"
	responseContentTypeSuffix      = "response-content-type"
	responseContentEncodingSuffix  = "response-content-encoding"
	responseJsonTagSuffix          = "response-json-tag"
	responseBodySuffix             = "response-body"
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

	version = "v0.1.1"
)

func main() {
	var (
		info api.GenerationInfo
	)

	inFile := flag.String("in", "./pkg/service", "relative path to dir with services")
	swaggerFile := flag.String("swagger", ".", "relative path to swagger file generate")

	info.SwaggerToJson = flag.Bool("json", false, "use json swagger output file")
	info.SwaggerToYaml = flag.Bool("yaml", false, "use yaml swagger output file")
	info.SwaggerInfo.Description = flag.String("desc", "", "swagger description")
	info.SwaggerInfo.Title = flag.String("title", "", "swagger title")
	info.SwaggerInfo.Version = flag.String("version", "", "swagger version")
	servers := flag.String("servers", "", `swagger servers in format: http://some.url = some url description\r\nhttp://another.url = another url description`)

	printVersion := flag.Bool("v", false, "print version and exit")
	flag.Parse()
	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
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

	if *info.SwaggerToJson == false && *info.SwaggerToYaml == false {
		info.SwaggerToYaml = &yaml
	}

	tagsParser := request.NewErrorProcessor(httpServer, requestErrorsSuffix,
		request.NewURIPath(httpServer, requestURIPathSuffix,
			request.NewQuery(httpServer, requestQuerySuffix,
				request.NewMethod(httpServer, requestMethodSuffix,
					request.NewHeader(httpServer, requestHeaderSuffix,
						request.NewContentType(httpServer, requestContentTypeSuffix,
							request.NewMultipartFileTag(httpServer, requestMultipartFileTagSuffix,
								request.NewMultipartValueTag(httpServer, requestMultipartValueTagSuffix,
									request.NewJsonTag(httpServer, requestJsonTagSuffix,
										request.NewAPIPath(httpServer, requestAPIPathSuffix, &request2.Term{}))))))))))
	tagsParser = response.NewStatus(httpServer, responseStatusSuffix,
		response.NewHeader(httpServer, responseHeaderSuffix,
			response.NewContentType(httpServer, responseContentTypeSuffix,
				response.NewEncodingType(httpServer, responseContentEncodingSuffix,
					response.NewJsonTag(httpServer, responseJsonTagSuffix,
						response.NewFile(httpServer, responseFileSuffix,
							response.NewBody(httpServer, responseBodySuffix, tagsParser)))))))
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
	t.Funcs(template.FuncMap{"notin": func(s []string, f string) bool {
		for _, v := range s {
			if v == f {
				return false
			}
		}
		return true
	}})

	imp := imports.NewImports()

	httpServerRender := httpserver.NewServer(t, httpServerPkgName, httpServerFilePath, imp)
	httpServerTransportRender := httpserver.NewTransport(t, httpServerPkgName, httpServerTransportFilePath, imp)
	httpServerBuilderRender := httpserver.NewBuilder(t, httpServerPkgName, httpServerBuilderFilePath, imp)
	httpClientRender := httpclient.NewClient(t, httpClientPkgName, httpClientFilePath, imp)
	httpClientTransportRender := httpclient.NewTransport(t, httpClientPkgName, httpClientTransportFilePath, imp)
	httpClientBuilderRender := httpclient.NewBuilder(t, httpClientPkgName, httpClientBuilderFilePath, imp)
	httpUIErrorsRender := httperrors.NewUI(t, httpErrorsPkgName, httpUIErrorsFilePath, imp)
	httpClientErrorsRender := httperrors.NewClient(t, httpErrorsPkgName, httpClientErrorsFilePath, imp)
	instrumentingRender := service.NewInstrumenting(t, serviceInstrumentingFilePath, imp)
	loggingRender := service.NewLogging(t, serviceLoggingFilePath, imp)
	mockRender := service.NewMock(t, httpClientPkgName, serviceMockFilePath, imp)
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
			httpClientRender,
			httpClientTransportRender,
			httpClientBuilderRender,
		),
		httpsClient: processor.NewHTTPClient(
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

	servicesProcessor := processor.NewServices(tagMark, processors, httpMethodProcessor)
	servicePreProcessor := preprocessor.NewService(servicesProcessor, goGeneratedAutomaticallyPrefix, swaggerRender)

	info.SwaggerAbsOutputPath = *swaggerFile
	err := servicePreProcessor.Process(*inFile, *inFile, &info)
	if err != nil {
		fmt.Printf("servicePreProcessor.Process error: %v\n", err)
	}

	os.Exit(0)
}
