package router

import (
	"ec2/model/rest"
	"net/http"
)

var rotaConfig = []Rotas{
	{

		URI:          "/criausuario",
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
		Autenticacao: true,
	}, {
		URI:          "/usuario",
		Metodo:       http.MethodGet,
		Handlerfunc:  rest.BuscaUsuario,
		Autenticacao: true,
	}, {
		URI:          "/login",
		Metodo:       http.MethodPost,
		Handlerfunc:  rest.Login,
		Autenticacao: false,
	}, {
		URI:          "/usuario_senha/{id}",
		Metodo:       http.MethodPost,
		Handlerfunc:  rest.PutPassword,
		Autenticacao: true,
	},
}
