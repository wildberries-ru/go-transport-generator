package unmarshalling

import (
	"context"
	"github.com/prometheus/common/log"
)

// @gtg http-server http-client http-errors
type Service interface {
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/testeasyjson
	// @gtg http-server-query param1={param1}
	// @gtg http-server-response-json-tag field1 field1
	// @gtg http-server-response-json-tag field2 field2
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	TestEasyJson(ctx context.Context, param1 int) (field1 , field2 string, err error)
}

type TestService struct {

}

const FieldValue1 = "hello \"world"
const FieldValue2 = "test"

func (s *TestService ) TestEasyJson(ctx context.Context, param1 int) (field1 , field2 string, err error){

	switch param1 {
	case 0:
		field1 = FieldValue1
		field2 = FieldValue2
	case 1:
		field1 = FieldValue2
		field2 = FieldValue1
	default:
		log.Errorf("unknown param value: %v", param1)
	}

	return
}

func NewService()  *TestService {

	return &TestService{}

}