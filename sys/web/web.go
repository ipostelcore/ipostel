package web

//Copyright Carlos Peña
//Modulo de negociación WEB
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ipostelcore/ipostel/sys/web/api"
)

//Variables de Control
var (
	Enrutador   = mux.NewRouter()
	WsEnrutador = mux.NewRouter()
)

//Cargar los diferentes modulos del sistema
func Cargar() {
	CargarModulosPanel()
	CargarModulosWeb()
	CargarModulosSeguridad()

	WMAdminLTE()
	CargarModulosWebDevel()
}

//CargarModulosWeb Cargador de modulos web
func CargarModulosWeb() {
	var wUsuario api.WUsuario
	var ap api.API

	Enrutador.HandleFunc("/", Principal)
	Enrutador.HandleFunc("/ipostel/api/crud/{id}", wUsuario.ValidarToken(ap.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipostel/api/oficina/listar", wUsuario.ValidarToken(ap.Oficinas)).Methods("GET")

	// Enrutador.HandleFunc("/ipostel/api/crud", wUsuario.ValidarToken(per.Actualizar)).Methods("PUT")
	// Enrutador.HandleFunc("/ipostel/api/crud", wUsuario.ValidarToken(per.Insertar)).Methods("POST")
	// Enrutador.HandleFunc("/ipostel/api/crud", wUsuario.ValidarToken(per.Eliminar)).Methods("DELETE")
	// Enrutador.HandleFunc("/ipostel/api/crud", wUsuario.ValidarToken(per.Opciones)).Methods("OPTIONS")
	// Enrutador.HandleFunc("/ipostel/api/listado", wUsuario.ValidarToken(per.Listado)).Methods("POST")

}

//CargarModulosPanel Panel de Contencion
func CargarModulosPanel() {

}

//CargarModulosSeguridad Y cifrado
func CargarModulosSeguridad() {
	var wUsuario api.WUsuario
	// Enrutador.HandleFunc("/ipsfa/app/api/wusuario/crud/{id}", wUsuario.Consultar).Methods("GET")
	Enrutador.HandleFunc("/ipostel/app/api/wusuario/login", wUsuario.Login).Methods("POST")
	Enrutador.HandleFunc("/ipostel/app/api/wusuario/login", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipostel/api/wusuario/validar", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("POST")
	Enrutador.HandleFunc("/ipostel/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")

	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.Crear).Methods("POST")
	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.ValidarToken(wUsuario.CambiarClave)).Methods("PUT")
	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.ValidarToken(wUsuario.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.ValidarToken(wUsuario.Crear)).Methods("POST")
	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.ValidarToken(wUsuario.CambiarClave)).Methods("PUT")
	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.ValidarToken(wUsuario.Opciones)).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipostel/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")

	Enrutador.HandleFunc("/ipostel/api/wusuario/validarphp", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("GET")
}

//Principal Página inicial del sistema o bienvenida
func Principal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenidos al Bus Empresarial de Datos")
}

//WMAdminLTE OpenSource tema de panel de control Tecnología Bootstrap3
func WMAdminLTE() {
	fmt.Println("Módulos de AdminLTE V 3.5")
	// var GP = GPanel{}
	// Enrutador.HandleFunc("/sssifanb/{id}", GP.IrA)
	prefix := http.StripPrefix("/ipostel", http.FileServer(http.Dir("public_web/www")))
	Enrutador.PathPrefix("/ipostel/").Handler(prefix)
	// prefixx := http.StripPrefix("/bdse-admin/public/temp", http.FileServer(http.Dir("public/temp")))
	// Enrutador.PathPrefix("/bdse-admin/public/temp/").Handler(prefixx)
}

//CargarModulosWebDevel Cargador de modulos web
func CargarModulosWebDevel() {
	//var wUsuario api.WUsuario
	// Enrutador.HandleFunc("/devel/api/militar/crud/{id}", per.Consultar).Methods("GET")
	// Enrutador.HandleFunc("/devel/api/militar/crud", per.Actualizar).Methods("PUT")
	// Enrutador.HandleFunc("/devel/api/militar/crud", per.Insertar).Methods("POST")
	// Enrutador.HandleFunc("/devel/api/militar/crud", per.Eliminar).Methods("DELETE")
	// Enrutador.HandleFunc("/devel/api/militar/crud", per.Opciones).Methods("OPTIONS")
	// Enrutador.HandleFunc("/devel/api/militar/reportecomponente", per.EstadisticasPorComponente).Methods("POST")
	// Enrutador.HandleFunc("/devel/api/militar/reportegrado", per.EstadisticasPorGrado).Methods("POST")
	// Enrutador.HandleFunc("/devel/api/militar/reportefamiliar", per.EstadisticasFamiliar).Methods("POST")
	// Enrutador.HandleFunc("/ipsfa/api/militar/listado", per.Listado).Methods("POST")
	// Enrutador.HandleFunc("/ipsfa/api/militar/subirarchivos", per.SubirArchivos).Methods("POST")

}
