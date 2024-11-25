package response

import (
	"net/http"
	"reflect"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Response(r *http.Request, w http.ResponseWriter, data any, err error) {
	if err != nil {
		httpx.WriteJson(w, http.StatusOK, Error(err))
		return
	}
	httpx.WriteJson(w, http.StatusOK, Success(data))
}

func Success(data interface{}) *Body {
	resp := &Body{
		Code: 200,
		Msg:  "success",
	}

	if data == nil {
		resp.Data = nil
		return resp
	}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			resp.Data = nil
			return resp
		}
		// decode from interface
		data = val.Elem().Interface()
	}
	resp.Data = data
	return resp
}

func Error(err error) *Body {
	return &Body{
		Code: 500,
		Msg:  err.Error(),
		Data: nil,
	}
}

//
//作者：uccs
//链接：https://juejin.cn/post/7371311251434471474
//来源：稀土掘金
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
