package router

import (
	"net/http"
	"webapp/src/rest"
)

var configRotas = []Rotas{
	{
		URI:          "/",
		Metodo:       http.MethodGet,
		Handlerfunc:  rest.TelaLogin,
		Autenticacao: false,
	}, {
		URI:          "/login",
		Metodo:       http.MethodGet,
		Handlerfunc:  rest.TelaLogin,
		Autenticacao: false,
	},
	{
		URI:          "/criar_usuario",
		Metodo:       http.MethodGet,
		Handlerfunc:  rest.CarregarTelaCriarUsuario,
		Autenticacao: false,
	},
}
