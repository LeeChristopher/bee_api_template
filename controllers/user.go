package controllers

type UserController struct {
	baseController
}

func (m *UserController) Initialization() {
}

func (m *UserController) Index() {
	m.SetResponse(200, `Index!`, nil)
}

func (m *UserController) Add() {
	m.SetResponse(200, `Add!`, nil)
}

func (m *UserController) Info() {
	m.SetResponse(200, `Info!`, nil)
}

func (m *UserController) Update() {
	m.SetResponse(200, `Update!`, nil)
}

func (m *UserController) Delete() {
	m.SetResponse(200, `Delete!`, nil)
}
