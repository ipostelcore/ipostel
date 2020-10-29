package api

import (
	"net/http"

	"github.com/ipostelcore/ipostel/mdl/core"
	"github.com/ipostelcore/ipostel/sys/seguridad"
)

//UsuarioConectado Seguridad Informatica
var UsuarioConectado seguridad.Usuario

//API estructuras generales
type API struct {
	OID    int    `json:"oid"`
	Source string `json:"source"`
	Urls   string `json:"urls"`
	Base   string `json:"base"`
	Query  string `json:"query"`
	Autor  string `json:"autor"`
}

//Agregar una interfaz blanca
func (a *API) Agregar(w http.ResponseWriter, r *http.Request) {

}

//Consultar una interfaz blanca
func (a *API) Consultar(w http.ResponseWriter, r *http.Request) {

}

//Oficinas Listar todas las oficinas
func (a *API) Oficinas(w http.ResponseWriter, r *http.Request) {

	var c core.Core
	Cabecera(w, r)
	w.WriteHeader(http.StatusOK)
	j, _ := c.Oficinas()
	w.Write(j)
	return
}
