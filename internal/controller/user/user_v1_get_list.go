package user

import (
	"context"

	"demo/api/user/v1"
	"demo/internal/dao"
	"demo/internal/model/do"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	err = dao.User.Ctx(ctx).Where(do.User{
		Age:    req.Age,
		Status: req.Status,
	}).Scan(&res.List)
	return
}
