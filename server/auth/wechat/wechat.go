package wechat

import (
	"fmt"

	"github.com/medivhzhan/weapp/v2"
)

type Service struct {
	AppId  string
	Secret string
}

func (s *Service) Resolve(code string) (string, error) {
	res, err := weapp.Login(s.AppId, s.Secret, code)
	if err != nil {
		// 处理一般错误信息
		return "", fmt.Errorf("登陆错误:%v", err)
	}
	if err := res.GetResponseError(); err != nil {
		// 处理微信返回错误信息
		return "", fmt.Errorf("返回信息错误:%v", err)
	}
	return res.OpenID, nil
}
