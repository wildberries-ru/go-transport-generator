// Package httpclient ...
// CODE GENERATED AUTOMATICALLY
// DO NOT EDIT
package httpclient

import (
	"context"
	"mime/multipart"

	"github.com/stretchr/testify/mock"
	v1 "github.com/wildberries-ru/go-transport-generator/example/api/v1"
)

// MockService ...
type MockService struct {
	mock.Mock
}

// UploadDocument ...
func (s *MockService) UploadDocument(ctx context.Context, token *string, name string, extension string, categoryID string, supplierID *int64, contractID *int64, data multipart.File) (err error) {
	args := s.Called(context.Background(), token, name, extension, categoryID, supplierID, contractID, data)
	return args.Error(0)
}

// GetWarehouses ...
func (s *MockService) GetWarehouses(ctx context.Context) (pets map[string]v1.Detail, err error) {
	args := s.Called(context.Background())
	return args.Get(0).(map[string]v1.Detail), args.Error(1)
}

// GetDetails ...
func (s *MockService) GetDetails(ctx context.Context, namespace string, detail string, fileID uint32, someID *uint64, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error) {
	args := s.Called(context.Background(), namespace, detail, fileID, someID, token)
	return args.Get(0).(v1.Detail), args.Get(1).(v1.Namespace), args.Get(2).(*string), args.Error(3)
}

// GetDetailsEmbedStruct ...
func (s *MockService) GetDetailsEmbedStruct(ctx context.Context, namespace string, detail string) (response v1.GetDetailsEmbedStructResponse, err error) {
	args := s.Called(context.Background(), namespace, detail)
	return args.Get(0).(v1.GetDetailsEmbedStructResponse), args.Error(1)
}

// GetDetailsListEmbedStruct ...
func (s *MockService) GetDetailsListEmbedStruct(ctx context.Context, namespace string, detail string) (details []v1.Detail, err error) {
	args := s.Called(context.Background(), namespace, detail)
	return args.Get(0).([]v1.Detail), args.Error(1)
}

// PutDetails ...
func (s *MockService) PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	args := s.Called(context.Background(), namespace, detail, testID, blaID, token, pretty, yang)
	return args.Get(0).(v1.Detail), args.Get(1).(v1.Namespace), args.Get(2).(*string), args.Error(3)
}

// GetSomeElseDataUtf8 ...
func (s *MockService) GetSomeElseDataUtf8(ctx context.Context) (cool v1.Detail, nothing v1.Namespace, id *string, err error) {
	args := s.Called(context.Background())
	return args.Get(0).(v1.Detail), args.Get(1).(v1.Namespace), args.Get(2).(*string), args.Error(3)
}

// GetFile ...
func (s *MockService) GetFile(ctx context.Context) (data []byte, fileName string, err error) {
	args := s.Called(context.Background())
	return args.Get(0).([]byte), args.Get(1).(string), args.Error(2)
}
