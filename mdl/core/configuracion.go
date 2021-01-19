package core

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
