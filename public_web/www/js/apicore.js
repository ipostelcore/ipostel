

class ApiCore{
    constructor(){
        this.prioriodad = "alta";
        this.concurrencia = true;
        this.ruta = "";
        this.funcion = "";
        this.parametros = "";
        this.retorna = true;
        this.migrar = false;
        this.modulo = "";
    }
    Obtener(){
        return this;
    }
    GenerarListado(){
        var promesa =  CargarAPI({
            metodo : "GET",
            sURL: conn.URL + "listar"
        });
        promesa.then(function (xhRequest) {
            var json = JSON.parse( xhRequest.responseText);
            console.log(json);
            CrearGrid(json)
        });
    }
    
    
}

function CrearGrid(db){
    $("#jsGridApi").jsGrid({
        width: "100%",
        height: "600px",
        filtering: false,
        editing: false,
        sorting: true,
        paging: true,
        data: db,
        fields: [
            {  width: 10,
                itemTemplate: function(_, item) {
                    if(item.estatus == false){
                        return $("<i style='color:#A40F0F' class='fas fa-screwdriver'>")
                    }else{
                        return $("<i style='color:#6ABE0B' class='fas fa-check-circle'>")
                    }
              	}
            },
            { name: "protocolo", type: "text", title: "Protocolo",  width: 35,},
            { name: "puerto", type: "text", title: "Puerto",  width: 25,},
            { name: "metodo", type: "text", title: "Metodo",  width: 30,},
            { name: "funcion", type: "text", title: "Funcion" },
            { name: "entorno", type: "text", title: "Ruta",
                itemTemplate: function(_, item) {
                    return item.entorno=="produccion"?"/ipostel/api/crud":"/devel/api/crud"
                }
            },
            { name: "driver", type: "text", title: "Driver" },
            {  width: 10,
                itemTemplate: function(_, item) {
                    return $("<i style='color:#1D9D36' class='fas fa-play'>")
                    	.on("click", function() {
                        	EjecutarAPI(item);
                    	});
              	}
            }
        ]
    });
}

function EjecutarAPI(item){
    console.log(item);
    var xAPI = new ApiCore();
    var api = item.entorno=="produccion"?"/ipostel/api/crud":"/devel/api/crud"
    var ruta = item.protocolo.toLowerCase() + "://" + conn.IP + ":" + item.puerto + api;

    xAPI.funcion = item.funcion;
    xAPI.modulo = item.modulo;
    xAPI.metodo = item.metodo;
    
    console.log(ruta);
    console.log(xAPI.Obtener());

    var promesa =  CargarAPI({
        metodo : item.metodo,
        valores: xAPI.Obtener(),
        sURL: ruta
    });
    promesa.then(function (xhRequest) {
        var json =  xhRequest.responseText;
        console.log(json);
       
    });
}