package main

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"os"
)

/**
 此处定义服务业务逻辑部分
 */
type OperationService interface {
	Add(left int, right int) int
}

/**
服务接口实现
 */
type operationService struct {

}

func (operationService) Add(left int, right int) int {
	return 	left + right
}

/**
定义请求参数以及响应参数
 */
type addRequest struct {
	Left int `json:"left"`
	Right int `json:"right"`
}

type addResponse struct {
	Result string `json:"result"`
}

func makeAddEndpoint(svc OperationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(addRequest)
		result := svc.Add(req.Left, req.Right)
		return addResponse{fmt.Sprintf("%d and %d is %d", req.Left, req.Right, result)}, nil
	}
}

func decodeAddRequest(_ context.Context, r *http.Request) (interface{}, error)  {
	var reuest addRequest
	if err := json.NewDecoder(r.Body).Decode(&reuest);err !=nil {
		return nil, err
	}
	return reuest, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc OperationService
	svc = operationService{}
	svc = LogginMiddleware{logger, svc}
	addHandler := httptransport.NewServer(
		makeAddEndpoint(svc),
		decodeAddRequest,
		encodeResponse,)

	http.Handle("/add", addHandler)

	logger.Log("err", http.ListenAndServe("192.168.1.32:8080", nil))
}

