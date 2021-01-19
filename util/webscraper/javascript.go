package webscraper

func codigoClase(cf CodigoFuente) string {

	return `
	/*!
	* ESBEpicTechnologies v1.0.0 (https://esbepictechnologies.io)
	* Copyright 2020-2021 CodeEpicTechnologies <http://codeepictechnologies.com>
	* Licensed under MIT (https://github.com/gesaodin/esbepic/blob/master/LICENSE)
	*/\n
	class ESB` + cf.Nombre + `{
	/** ` + cf.Goblal + ` **/
	constructor (){
	` + cf.Constructor + `
	}
	
	//Función permite obtener los datos del formulario
	Obtener(){
	` + cf.Obtener + `
		return this;
	}

	//Cargar un documento mediante API de conexion
	Cargar(req){
		var apic = new ApiCore();
        apic.funcion = "ObtenerOficinas";
        var promesa =  CargarAPI({
			metodo : "POST",
			sURL: conn.URL + "crud",
			valores :  apic.Obtener()
		});
	
		promesa.then(function (xhRequest) {
			json = xhRequest.responseText;
			
		});

	}

	//Salvar documento mediante API de conexion
	Salvar(){
		var obj = this.Obtener();
		var apic = new ApiCore();
		apic.funcion = "ESB` + cf.Nombre + `";
		apic.valores = obj;
        var promesa =  CargarAPI({
			metodo : "POST",
			sURL: conn.URL + "crud",
			valores :  apic.Obtener()
		});
	
		promesa.then(function (xhRequest) {
			json = xhRequest.responseText;
			
		});

		
	}
	};
	let ` + cf.Nombre + ` = new  ESB` + cf.Nombre + `();`
}
