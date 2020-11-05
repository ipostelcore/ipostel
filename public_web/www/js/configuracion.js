

class Usuarios{
    constructor(){}
    Crear(){}
    Obtener(){}
}

class Servicios{
    constructor(){
        this.codigo = "";
        this.nombre = "";
        this.descripcion = "";
        this.estatus = false;
    }

    Crear(){

    }
    Obtener(){
        this.nombre = $("#txtnombre").val();
        this.descripcion = $("#txtdescripcion").val();
        this.estatus = $("#cmbestatus option:selected").val();
        return this;
    }
}

class Roles{
    constructor(){}
    Crear(){}
    Obtener(){}
}

function _mdlAgregarRoles(){
    $("#_mdlTitulo").html("Agregar Roles");
    $("#mdlGeneral").modal("show");
}

class Oficinas{
    constructor(){}
    Crear(req){
        console.log(req);
    }
    Obtener(){}
}

function _mdlAgregarServicio(){
    $("#_mdlTitulo").html("Agregar Servicio");
    $("#mdlGeneral").modal("show");
}

function SalvarServicio(){
    var srv = new Servicios();
    console.info(srv.Obtener());
    var opciones = {
        valores : srv.Obtener(),
        metodo : "POST",
        sUrl : "ruta",
        Objeto : srv
    }
    CargarAPI(opciones);

}

function ListarOficinas(){
    var ofc = new Oficinas();
    var promesa =  CargarAPI({
        metodo : "GET",
        sURL: conn.URL + "oficina/listar",
        Objeto: ofc
    });
    promesa.then(function (xhRequest) {
        json = JSON.parse(xhRequest.responseText);
        json.forEach(e => {
            $("#cmboficinas").append(`<option value="${e.id}">${e.descripcion}</option>`);
        });
    });

    console.log("Imprimiendo");

}
