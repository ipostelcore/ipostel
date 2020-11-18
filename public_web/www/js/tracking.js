
/**
 * Realizado por: Maikol Echenique 
 * Fecha: 16-11-2020 
 * Titulo: Creación de clases Módulo Tracking
 */


class Consignacion{
    constructor(){
        this.hora = "";
        this.fecha = "";
        this.tipoenvio = "";
        this.sticker = "";
        this.oficinaorigen = "";
        this.peso = "";
        this.tarifa = "";
        this.exonerado = "";
        this.monto = "";
        this.total = "";
        this.cobrado = "";
        this.contado = "";
        this.credito = "";
        this.cedula = "";
        this.nombreremitente = "";
        this.direccion = "";
        this.telefono = "";
        this.correo = "";
        this.nombredestinatario = "";
        this.pais = "";
        this.oficinadestino = "";
        this.direcciondestino = "";
        this.telefonodestino = "";

    }
    Crear(req){}
    Obtener(){
        this.hora = $("#txtHora").val();
        this.fecha = $("#txtFecha").val();
        this.estatus = $("#cmbTipoEnvio option:selected").val();
        this.sticker = $("#txtSticker").val();
        this.oficinaorigen = $("#cmbOficinaOrigen option:selected").val();
        this.peso = $("#txtPeso").val();
        this.tarifa = $("#cmbTarifa option:selected").val();
        this.exonerado = $("#radExonerado").val();
        this.monto = $("#txtMonto").val();
        this.total = $("#txtTotal").val();
        this.cobrado = $("#txtCobrado").val();
        this.contado = $("#txtContado").val();
        this.credito = $("#txtCredito").val();
        this.cedula = $("#txtCedula").val();
        this.nombreremitente = $("#txtNombre").val();
        this.direccion = $("#txtDireccion").val();
        this.correo = $("#txtCorreo").val();
        this.nombredestinatario = $("#txtNombreDestinatario").val();
        this.pais = $("#txtPais").val();
        this.oficnadestino = $("#cmbOficinaDestino option:selected").val();
        this.direcciondestino = $("#txtDireccionDestino").val();
        this.telefonodestino = $("#txtTelefonoDestino").val();

        return this;
    }
}

function GuardarConsignacion(){

}

class Seguimiento{
    constructor(){
        this.hora = "";
        this.fecha = "";
        this.sticker = "";
        this.destinatario = "";
        this.remitente = "";
        this.peso = "";
        this.estatus = "";
        this.observacion = "";

    }
    Crear(req){}
    Obtener(){
        this.hora = $("#txtHora").val();
        this.fecha = $("#txtFecha").val();
        this.sticker = $("#txtsticker").val();
        this.destinatario = $("#txtDestinatario").val();
        this.remitente = $("#txtRemitente").val();
        this.peso = $("#txtPeso").val();
        this.estatus = $("#cmbEstatus option:selected").val();
        this.observacion = $("#txtObservacion").val();
        
    
        return this;
    }
}

function BuscarPaquete(){

}

class Distribucion{
    constructor(){
        this.hora = "";
        this.fecha = "";
        this.sticker = "";
        this.destinatario = "";
        this.remitente = "";
        this.peso = "";
        this.estatus = "";
        this.observacion = "";
    
    }
        Crear(req){}
        Obtener(){
        this.hora = $("#txtHora").val();
        this.fecha = $("#txtFecha").val();
        this.sticker = $("#txtSticker").val();
        this.destinatario = $("#txtDestinatario").val();
        this.remitente = $("#txtRemitente").val();
        this.peso = $("#txtPeso").val();
        this.tipodistribucion = $("#cmbTipoDistibucion option:selected").val();
        this.observacion = $("#txtObservacion").val();
        
    
        return this;
    }

}
    class Distribucion{
        constructor(){
            this.hora = "";
            this.fecha = "";
            this.sticker = "";
            this.destinatario = "";
            this.remitente = "";
            this.peso = "";
            this.estatus = "";
            this.observacion = "";
    
        }
        Crear(req){}
        Obtener(){
        this.hora = $("#txtHora").val();
        this.fecha = $("#txtFecha").val();
        this.sticker = $("#txtSticker").val();
        this.destinatario = $("#txtDestinatario").val();
        this.remitente = $("#txtRemitente").val();
        this.peso = $("#txtPeso").val();
        this.tipodistribucion = $("#cmbTipoDistibucion option:selected").val();
        this.observacion = $("#txtObservacion").val();
        
    
        return this;
    }
}
class Ruta{
    constructor(){
        this.localidad = "";
        this.zona = "";
        this.ruta = "";

    }
    Crear(req){}
    Obtener(){
    this.localidad = $("#txtLocalidad").val();
    this.zona = $("#txtZona").val();
    this.ruta = $("#txtRuta").val();
    

    return this;
    }
}

class Apartados{
    constructor(){
        this.iniciorango = "";
        this.finrango = "";
        this.oficinaperteneciente = "";
        this.zonapostal = "";
        this.ruta = "";

    }
    Crear(req){}
    Obtener(){
    this.iniciorango = $("#txtInicioRango").val();
    this.finrango = $("#txtFinRango").val();
    this.oficinaperteneciente = $("#cmbOficinaPerteneciente option:selected").val();
    this.zonapostal = $("#txtZonaPostal").val();
    this.ruta = $("#txtRuta").val();

    return this;
    }
}


function GuardarApartados(){

}

class Repartidores{
    constructor(){
        this.repatidor = "";
        this.rutamultiple = "";

    }
    Crear(req){}
    Obtener(){
    this.repatidor = $("#txtRepartidor").val();
    this.rutamultiple = $("#cmbRutaMultiple option:selected").val();

    return this;
    }
}


function GuardarRepartidor(){

}


class Incentivos{
    constructor(){
        this.desde = "";
        this.hasta = "";
        this.monto = "";
    
    }
    Crear(req){}
    Obtener(){
    this.desde = $("#txtDesde").val();
    this.hasta = $("#txtHasta").val();
    this.monto = $("#txtMonto").val();
    

    return this;
    }
}


function GuardarRepartidor(){

}





