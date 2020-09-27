package controllers

import "bee_api_template/filters"

var authFilter *filters.AuthFilter

type AuthController struct {
	baseController
}

func (m *AuthController) Initialization() {
	authFilter = filters.NewAuthFilter(m.Ctx.Request)
}

/**
 *登录
 */
func (m *AuthController) Login() {
	res, err := authFilter.Login()
	if err != nil {
		m.SetResponse(0, err.Error(), nil)
		return
	}

	m.SetResponse(200, "ok", res)
}
