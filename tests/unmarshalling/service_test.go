package unmarshalling

import (
	"context"
	"github.com/buaazp/fasthttprouter"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"github.com/wildberries-ru/go-transport-generator/tests/unmarshalling/httpclient"
	"github.com/wildberries-ru/go-transport-generator/tests/unmarshalling/httpserver"
	"sync"
	"testing"
	"time"
)

func TestEasyJson(t *testing.T) {

	svc := NewService()

	router := fasthttprouter.New()
	httpserver.New(router, svc, nil, nil, nil, nil)

	fasthttpServer := &fasthttp.Server{
		Handler: router.Handler,
	}

	stop := make(chan struct {})
	poll := make(chan struct {}, 100)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := fasthttpServer.ListenAndServe(":8080" ); err != nil {
			t.Fatalf("%v", err)
		}
	}()

	time.Sleep(1 * time.Second)

	wg.Add(1)
	go func() {
		defer wg.Done()
		<- stop
		_ = fasthttpServer.Shutdown()
		t.Log("fasthttpServer stopped")
	}()

	client, err := httpclient.New("http://localhost:8080", 100,nil,map[interface{}]httpclient.Option{})
	assert.Nilf(t,err, "failed to create client")
	var wg2 sync.WaitGroup

	var result []struct {
		val1 string
		val2 string
	}

	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func(val int) {
			poll <- struct{}{}
			defer wg2.Done()
			field1 , field2 , err := client.TestEasyJson(context.Background(), val %2)
			<- poll
			assert.Nilf(t,err, "TestEasyJson failed")
			result = append(result, struct {
				val1 string
				val2 string
			}{val1:field1 , val2: field2})

		}(i)
	}
	wg2.Wait()

	for _, val := range result {
		if val.val1 != FieldValue1 && val.val1 != FieldValue2 {
			assert.FailNowf(t, "field1 has wrong value", "%v is not expected", val.val1)
		}
		if val.val2 != FieldValue1 && val.val2 != FieldValue2 {
			assert.FailNowf(t, "field2 has wrong value", "%v is not expected", val.val2)
		}
	}


	stop <- struct{}{}
	wg.Done()
	t.Log("done")
}

