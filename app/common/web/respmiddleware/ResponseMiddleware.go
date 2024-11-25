package respmiddleware

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	response "business-platform/app/common/web/resp"
)

type ResponseMiddleware struct{}

func NewResponseMiddleware() *ResponseMiddleware {
	return &ResponseMiddleware{}
}

func (m *ResponseMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		httpx.SetErrorHandler(func(err error) (int, any) {
			return http.StatusOK, response.Error(err)
		})

		httpx.SetOkHandler(func(ctx context.Context, data any) any {
			return response.Success(data)
		})

		next(w, r)
	}
}
