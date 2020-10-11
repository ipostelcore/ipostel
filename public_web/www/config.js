let _url = new URL(document.URL);

class Conexion{
    constructor(){
        this.IP = _url.hostname;
        this.Puerto = ":8012";
        this.PuertoSSL = ":2286";
        this.API = "/ipostel/api/";
        this.URL = "https://" + this.IP + this.PuertoSSL + this.API;
        this.URLIMG = "/imagenes/";
        this.URLTEMP = _url.hostname + "/ipostel/temp/";
        this.URLSEC = "http://" + this.IP + this.Puerto;
    }
}
    
function CargarAPI(options){
    var xhttp = new XMLHttpRequest();
    xhttp.open(options.metodo, options.sURL);
    xhttp.setRequestHeader("Authorization", "Bearer " + sessionStorage.getItem('ipostelToken'));
    var promise = new Promise(function(resolve, reject) {
        xhttp.addEventListener('readystatechange', function() {
            if ( xhttp.readyState === 4 && xhttp.status === 200) {
                if(options.Objeto != undefined){
                    options.Objeto = JSON.parse(xhttp.responseText);
                }
                resolve(xhttp);
            }
            if( xhttp.status === 401){
                if ( xhttp.responseText != "" ) {
                    respuesta = JSON.parse(xhttp.responseText);
                    $.notify(respuesta.msj);
                }
            }
        });
        
        xhttp.addEventListener('error', function() {
            if ( xhttp.responseText != "" ) {
                respuesta = JSON.parse(xhttp.responseText);
                if (respuesta.tipo != 0){
                    $.notify("Se ha Insertado correctamente", "success");
                }else{
                    alert(xhttp.responseText);
                }
            }
            reject(xhttp);
        });
    });
    
    if(options.valores != undefined){
        xhttp.send(JSON.stringify(options.valores));
    }else{
        xhttp.send();
    }

    return promise;
}
    

/**
 * 
 * @param {string} id  
 * @param {string} url | HTML 
 */
function CargarUrl(id, url){
    var xhttp = new XMLHttpRequest();
    xhttp.open('GET', url + '.html');
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            $('#'+id).html(xhttp.responseText);
        }
    };
    xhttp.onerror = function() {
        if (this.readyState == 4 && this.status == 0) {
            $.notify("El archivo no ha sido encontrado");
        }

    };
    xhttp.send();
}