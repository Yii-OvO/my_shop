package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"my_shop/api/backend"
	"my_shop/internal/model"
	"my_shop/internal/service"
)

var File = cFile{}

type cFile struct{}

func (c *cFile) Upload(ctx context.Context, req *backend.FileUploadReq) (res *backend.FileUploadRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCodef(gcode.CodeMissingParameter, "请上传文件") // gerror.NewCodef:自定义错误码和错误信息
	}
	upload, err := service.File().Upload(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	return &backend.FileUploadRes{
		Name: upload.Name,
		Url:  upload.Url,
	}, nil
}
