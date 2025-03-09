package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/file" method:"post" mime:"multipart/form-data" tags:"工具" sm:"上传文件"` //mime: 接口的mime类型，例如multiparty/form-data一般是全局设置，默认为application、json
	File   *ghttp.UploadFile                                                           `json:"file" type:"file" dc:"选择上传文件"` //File:上传单个文件，Files上传多个文件
}

type FileUploadRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url" dc:"图片地址"`
}
