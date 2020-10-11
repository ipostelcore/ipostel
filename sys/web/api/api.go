package api

import (
	"github.com/ipostelcore/ipostel/sys/seguridad"
)

//UsuarioConectado Seguridad Informatica
var UsuarioConectado seguridad.Usuario

//API estructuras generales
type API struct {
	OID    int    `json:"oid"`
	Source string `json:"source"`
	Url    string `json:"url"`
	Base   string `json:"base"`
	Query  string `json:"query"`
	Autor  string `json:"autor"`
}
