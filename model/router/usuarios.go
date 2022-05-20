package router

import (
	"ec2/model/rest"
	"net/http"
)

var rotaConfig = []Rotas{
	{

		URI:          "/usuario",
		Metodo:       http.MethodPost,
		Handlerfunc:  rest.InserirUsuario,
		Autenticacao: false,
	}, {

		URI:          "/usuario/{id}",
		Metodo:       http.MethodPut,
		Handlerfunc:  rest.PutUsuario,
		Autenticacao: true,
	}, {

		URI:          "/usuario/{id}",
		Metodo:       http.MethodGet,
		Handlerfunc:  rest.BuscaUsuarioById,
		Autenticacao: false,
	}, {
		URI:          "/usuario",
		Metodo:       http.MethodGet,
		Handlerfunc:  rest.BuscaUsuario,
		Autenticacao: false,
	}, {
		URI:          "/login",
		Metodo:       http.MethodPost,
		Handlerfunc:  rest.Login,
		Autenticacao: false,
	},
}
