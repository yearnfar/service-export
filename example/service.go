package example

import "context"

type Service interface {
	GetExample(ctx context.Context, req *GetExampleRequest) (resp *GetExampleResponse, err error)
}

type GetExampleRequest struct{}
type GetExampleResponse struct{}
