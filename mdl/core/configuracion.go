package core

import (
	"database/sql"
	"encoding/json"

	"github.com/ipostelcore/ipostel/sys"
	"github.com/ipostelcore/ipostel/util"
)

type Configuracion struct {
}

type ReporteVentas struct {
	Codigo         string `json:"codigo"`
	Descripcion    string `json:"descripcion"`
	Fecha          string `json:"fecha"`
	Concepto       string `json:"concepto"`
	CodigoServicio string `json:"codigoservicio"`
	Servicio       string `json:"servicio"`
	Total          string `json:"total"`
}

//VentasPorSemanas Reporte de ventas por semanas
func (Cnf *Configuracion) VentasPorSemanas() (jSon []byte, err error) {
	var lst []ReporteVentas
	s := `SELECT e.codopt, 
			o2.descripcion, 
			e.FECHA, 
			e.CONCEPTO, 
			e.CODSERV, 
			m2.descripcion AS servicio, 
			SUM(e.total) AS total_dia
		FROM estadisticas e 
		LEFT OUTER join oficinas o2 ON o2.codoficina = e.codopt 
		LEFT OUTER join maestroservicios m2 ON m2.codserv = e.codserv
			GROUP BY e.codopt, o2.descripcion , e.fecha, e.concepto, e.codserv,m2.descripcion `
	sq, err := sys.PuntoPostalPostgres.Query(s)
	util.Error(err)

	for sq.Next() {
		var rvnt ReporteVentas
		var cod, desc, fecha, conc, codserv, serv, total sql.NullString

		err := sq.Scan(&cod, &desc, &fecha, &conc, &codserv, &serv, &total)
		util.Error(err)
		rvnt.Codigo = util.ValidarNullString(cod)

		lst = append(lst, rvnt)

	}
	jSon, err = json.Marshal(lst)
	return

}
