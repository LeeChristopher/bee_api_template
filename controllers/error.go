package controllers

type ErrorController struct {
	baseController
}

func (m *ErrorController) Initialization() {
}

func (m *ErrorController) Error401() {
	m.SetResponse(401, `Unauthorized!`, nil)
}

func (m *ErrorController) Error403() {
	m.SetResponse(403, `Forbidden!`, nil)
}

func (m *ErrorController) Error404() {
	m.SetResponse(404, `Not Found!`, nil)
}

func (m *ErrorController) Error500() {
	m.SetResponse(500, `Internal Server Error!`, nil)
}

func (m *ErrorController) Error503() {
	m.SetResponse(503, `Service Unavailable!`, nil)
}
