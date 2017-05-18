package xmlstructures

import (
	"encoding/xml"
	"fmt"
)

/****************************************************************************************************************************************
*
*
* Estructura general de un CFD
*
*
****************************************************************************************************************************************/

// Comprobante Comprobante Estándar de Comprobante Fiscal Digital por Internet XML
type Comprobante struct {
	XMLName           xml.Name         `xml:"cfdi:Comprobante"`
	Version           string           `xml:"Version,attr"`           // Atributo requerido con valor prefijado a 3.3 que indica la versión del estándar bajo el que se encuentra expresado el comprobante. Default: "3.3" Req.
	Serie             string           `xml:"Serie,attr"`             // Atributo opcional para precisar la serie para control interno del contribuyente. Este atributo acepta una cadena de caracteres. Pattern ([A-Z]|[a-z]|[0-9]| |Ñ|ñ|!|&quot;|%|&amp;|&apos;| ́|- |:|;|&gt;|=|&lt;|@|_|,|\{|\}|`|~|á|é|í|ó|ú|Á|É|Í|Ó|Ú|ü| Ü){1,25} Opc.
	Folio             string           `xml:"Folio,attr"`             // Atributo opcional para control interno del contribuyente que expresa el folio del comprobante, acepta una cadena de caracteres. Pattern ([A-Z]|[a-z]|[0-9]| |Ñ|ñ|!|&quot;|%|&amp;|&apos;| ́|- |:|;|&gt;|=|&lt;|@|_|,|\{|\}|`|~|á|é|í|ó|ú|Á|É|Í|Ó|Ú|ü |Ü){1,40} Opc.
	Fecha             string           `xml:"Fecha,attr"`             // Atributo requerido para la expresión de la fecha y hora de expedición del Comprobante Fiscal Digital por Internet. Se expresa en la forma AAAA- MM-DDThh:mm:ss y debe corresponder con la hora local donde se expide el comprobante. tdCFDI:t_FechaH Req.
	Sello             string           `xml:"Sello,attr"`             // Atributo requerido para contener el sello digital del comprobante fiscal, al que hacen referencia las reglas de resolución miscelánea vigente. El sello debe ser expresado como una cadena de texto en formato Base 64. Req.
	FormaPago         string           `xml:"FormaPago,attr"`         // Atributo condicional para expresar la clave de la forma de pago de los bienes o servicios amparados por el comprobante. Si no se conoce la forma de pago este atributo se debe omitir. catCFDI:c_FormaPago Opc.
	NoCertificado     string           `xml:"NoCertificado,attr"`     // Atributo requerido para expresar el número de serie del certificado de sello digital que ampara al comprobante, de acuerdo con el acuse correspondiente a 20 posiciones otorgado por el sistema del SAT. Pattern [0-9]{20} Req.
	Certificado       string           `xml:"Certificado,attr"`       // Atributo requerido que sirve para incorporar el certificado de sello digital que ampara al comprobante, como texto en formato base 64. Req.
	CondicionesDePago string           `xml:"CondicionesDePago,attr"` // Atributo condicional para expresar las condiciones comerciales aplicables para el pago del comprobante fiscal digital por Internet. Este atributo puede ser condicionado mediante atributos o complementos. Pattern ([A-Z]|[a-z]|[0-9]| |Ñ|ñ|!|&quot;|%|&amp;|&apos;| ́|- |:|;|&gt;|=|&lt;|@|_|,|\{|\}|`|~|á|é|í|ó|ú|Á|É|Í|Ó|Ú|ü |Ü){1,1000} Opc.
	SubTotal          float64          `xml:"SubTotal,attr"`          // Atributo requerido para representar la suma de los importes de los conceptos antes de descuentos e impuesto. No se permiten valores negativos. Req.
	Descuento         float64          `xml:"Descuento,attr"`         // Atributo condicional para representar el importe total de los descuentos aplicables antes de impuestos. No se permiten valores negativos. Se debe registrar cuando existan conceptos con descuento. Opc.
	Moneda            string           `xml:"Moneda,attr"`            // Atributo requerido para identificar la clave de la moneda utilizada para expresar los montos, cuando se usa moneda nacional se registra MXN. Conforme con la especificación ISO 4217. catCFDI:c_Moneda Req.
	TipoCambio        string           `xml:"TipoCambio,attr"`        // Atributo condicional para representar el tipo de cambio conforme con la moneda usada. Es requerido cuando la clave de moneda es distinta de MXN y de XXX. Opc.
	Total             float64          `xml:"Total,attr"`             // Atributo requerido para representar la suma del subtotal, menos los descuentos aplicables, más las contribuciones recibidas. Req.
	TipoDeComprobante string           `xml:"TipoDeComprobante,attr"` // Atributo requerido para expresar la clave del efecto del comprobante fiscal para el contribuyente emisor. Req
	MetodoPago        string           `xml:"MetodoPago,attr"`        // Atributo condicional para precisar la clave del método de pago que aplica para este comprobante fiscal digital por Internet, conforme al Artículo 29-A fracción VII incisos a y b del CFF.Opc.
	LugarExpedicion   string           `xml:"LugarExpedicion,attr"`   // Atributo requerido para incorporar el código postal del lugar de expedición del comprobante (domicilio de la matriz o de la sucursal). Req.
	Confirmacion      string           `xml:"Confirmacion,attr"`      // Atributo condicional para registrar la clave de confirmación que entregue el PAC para expedir el comprobante con importes grandes, con un tipo de cambio fuera del rango establecido o con ambos casos. Es requerido cuando se registra un tipo de cambio o un total fuera del rango establecido. Pattern [0-9a-zA-Z]{5}. Opc.
	Relacionados      CFDIRelacionados `xml:"cfdi:CfdiRelacionados"`
	Emisor            CFDIEmisor       `xml:"cfdi:Emisor"`
	Receptor          CFDIReceptor     `xml:"cfdi:Receptor"`
	Conceptos         CFDIConceptos    `xml:"cfdi:Conceptos"`
	// CFDIImpuestos
}

/*****************************************************************************************************************************************
*
*	Seccion referente a los CFD relacionados del CFD
*
*
*****************************************************************************************************************************************/

// CFDIRelacionados Nodo opcional para precisar la información de los comprobantes relacionados.
type CFDIRelacionados struct {
	XMLName         xml.Name        `xml:"cfdi:CfdiRelacionados"`
	TipoRelacion    string          `xml:"tipoRelacion,attr"`    // Atributo requerido para indicar la clave de la relación que existe entre éste que se esta generando y el o los CFDI previos. catCFDI:c_TipoRelacion Req.
	CfdiRelacionado CFDIRelacionado `xml:"cfdi:CfdiRelacionado"` // Nodo requerido para precisar la información de los comprobantes relacionados.
}

// CFDIRelacionado Nodo opcional para precisar la información de los comprobantes relacionados.
type CFDIRelacionado struct {
	XMLName xml.Name `xml:"cfdi:CfdiRelacionado"`
	UUID    string   `xml:"UUID,attr"` // Atributo opcional para registrar el folio fiscal (UUID) de un CFDI relacionado con el presente comprobante, por ejemplo: Si el CFDI relacionado es un comprobante de traslado que sirve para registrar el movimiento de la mercancía. Si este comprobante se usa como nota de crédito o nota de débito del comprobante relacionado. Si este comprobante es una devolución sobre el comprobante relacionado. Si éste sustituye a una factura cancelada. Opc.
}

/*****************************************************************************************************************************************
*
*	Seccion referente al emisor del CFD
*
*
****************************************************************************************************************************************/

// CFDIEmisor Nodo requerido para expresar la información del contribuyente emisor del comprobante.
type CFDIEmisor struct {
	XMLName       xml.Name `xml:"cfdi:Emisor"`
	RFC           string   `xml:"Rfc,attr"`           // Atributo requerido para registrar la Clave del Registro Federal de Contribuyentes correspondiente al contribuyente emisor del comprobante. Pattern [a-zA-Z]{3-4}[0,9]{6}[a-zA-Z0-9]{3} Req.
	Nombre        string   `xml:"Nombre,attr"`        // Atributo opcional para registrar el nombre, denominación o razón social del contribuyente emisor del comprobante. Pattern  ([A-Z]|[a-z]|[0-9]| |Ñ|ñ|!|&quot;|%|&amp;|&apos;| ́|- |:|;|&gt;|=|&lt;|@|_|,|\{|\}|`|~|á|é|í|ó|ú|Á|É|Í|Ó|Ú|ü|Ü){1,254} Opc.
	RegimenFiscal string   `xml:"RegimenFiscal,attr"` // Atributo requerido para incorporar la clave del régimen del contribuyente emisor al que aplicará el efecto fiscal de este comprobante.
}

/*****************************************************************************************************************************************
*
*	Seccion referente al receptor del CFD
*
*
****************************************************************************************************************************************/

// CFDIReceptor Nodo requerido para precisar la información del contribuyente receptor del comprobante.
type CFDIReceptor struct {
	XMLName          xml.Name `xml:"cfdi:Receptor"`
	RFC              string   `xml:"Rfc,attr"`              // Atributo requerido para precisar la Clave del Registro Federal de Contribuyentes correspondiente al contribuyente receptor del comprobante. Req.
	Nombre           string   `xml:"Nombre,attr"`           // Atributo opcional para precisar el nombre, denominación o razón social del contribuyente receptor del comprobante. Opc.
	ResidenciaFiscal string   `xml:"ResidenciaFiscal,attr"` // Atributo condicional para registrar la clave del país de residencia para efectos fiscales del receptor del comprobante, cuando se trate de un extranjero, y que es conforme con la especificación ISO 3166-1 alpha-3. Es requerido cuando se incluya el complemento de comercio exterior o se registre el atributo NumRegIdTrib. c_Pais. Opc.
	NumRegIDTrib     string   `xml:"NumRegIdTrib,attr"`     // Atributo condicional para expresar el número de registro de identidad fiscal del receptor cuando sea residente en el extranjero. Es requerido cuando se incluya el complemento de comercio exterior. Opc.
	UsoCFDI          string   `xml:"UsoCFDI,attr"`          // Atributo requerido para expresar la clave del uso que dará a esta factura el receptor del CFDI. c_UsoCFDI. Req.
}

/*****************************************************************************************************************************************
*
*	Seccion referente a la lista de conceptos del CFD
*
*
****************************************************************************************************************************************/

// CFDIConceptos Nodo requerido para listar los conceptos cubiertos por el comprobante.
type CFDIConceptos struct {
	XMLName   xml.Name     `xml:"cfdi:Conceptos"`
	Conceptos CFDIConcepto // Lista de conceptos
}

// CFDIConcepto Nodo requerido para registrar la información detallada de un bien o servicio amparado en el comprobante.
type CFDIConcepto struct {
	XMLName          xml.Name           `xml:"cfdi:Concepto"`
	ClaveProdServ    string             `xml:"ClaveProdServ,attr"`    // Atributo requerido para expresar la clave del producto o del servicio amparado por el presente concepto. Es requerido y deben utilizar las claves del catálogo de productos y servicios, cuando los conceptos que registren por sus actividades correspondan con dichos conceptos. c_ClaveProdServ Req.
	NoIdentificacion string             `xml:"NoIdentificacion,attr"` // Atributo opcional para expresar el número de parte, identificador del producto o del servicio, la clave de producto o servicio, SKU o equivalente, propia de la operación del emisor, amparado por el presente concepto. Opcionalmente se puede utilizar claves del estándar GTIN. Pattern ([A-Z]|[a-z]|[0-9]| |Ñ|ñ|!|&quot;|%|&amp;|&apos;| ́|- |:|;|&gt;|=|&lt;|@|_|,|\{|\}|`|~|á|é|í|ó|ú|Á|É|Í|Ó|Ú|ü|Ü){1,100}. Opc.
	Cantidad         float64            `xml:"Cantidad,attr"`         // Atributo requerido para precisar la cantidad de bienes o servicios del tipo particular definido por el presente concepto. decimales (6) Req.
	ClaveUnidad      string             `xml:"ClaveUnidad,attr"`      // Atributo requerido para precisar la clave de unidad de medida estandarizada aplicable para la cantidad expresada en el concepto. La unidad debe corresponder con la descripción del concepto. catCFDI:c_ClaveUnidad Req.
	Unidad           string             `xml:"Unidad,attr"`           // Atributo opcional para precisar la unidad de medida propia de la operación del emisor, aplicable para la cantidad expresada en el concepto. La unidad debe corresponder con la descripción del concepto. Pattern ([A-Z]|[a-z]|[0-9]| |Ñ|ñ|!|&quot;|%|&amp;|&apos;| ́|- |:|;|&gt;|=|&lt;|@|_|,|\{|\}|`|~|á|é|í|ó|ú|Á|É|Í|Ó|Ú|ü|Ü){1,20}.Opc.
	Descripcion      string             `xml:"Descripcion,attr"`      // Atributo requerido para precisar la descripción del bien o servicio cubierto por el presente concepto. Pattern ([A-Z]|[a-z]|[0-9]| |Ñ|ñ|!|&quot;|%|&amp;|&apos;| ́|- |:|;|&gt;|=|&lt;|@|_|,|\{|\}|`|~|á|é|í|ó|ú|Á|É|Í|Ó|Ú|ü|Ü){1,1000} Opc.
	ValorUnitario    float64            `xml:"ValorUnitario,attr"`    // Atributo requerido para precisar el valor o precio unitario del bien o servicio cubierto por el presente concepto. tdCFDI:t_Importe Req.
	Importe          float64            `xml:"Importe,attr"`          //Atributo requerido para precisar el importe total de los bienes o servicios del presente concepto. Debe ser equivalente al resultado de multiplicar la cantidad por el valor unitario expresado en el concepto. No se permiten valores negativos. tdCFDI:t_Importe Req.
	Descuento        float64            `xml:"Descuento,attr"`        // Atributo opcional para representar el importe de los descuentos aplicables al concepto. No se permiten valores negativos. tdCFDI:t_Importe Opc.
	Impuestos        CFDIImpuestosInner `xml:"cfdi:Impuestos"`
}

// CFDIImpuestosInner Nodo opcional para capturar los impuestos aplicables al presente concepto. Cuando un concepto no registra un impuesto, implica que no es objeto del mismo.
type CFDIImpuestosInner struct {
	XMLName     xml.Name                      `xml:"cfdi:Impuestos"`
	Traslados   CFDIImpuestosTrasladosInner   `xml:"cfdi:Traslados"`   //Nodo opcional para asentar los impuestos trasladados aplicables al presente concepto.
	Retenciones CFDIImpuestosRetencionesInner `xml:"cfdi:Retenciones"` //Nodo opcional para asentar los impuestos trasladados aplicables al presente concepto.
}

// CFDIImpuestosTrasladosInner Nodo opcional para asentar los impuestos trasladados aplicables al presente concepto.
type CFDIImpuestosTrasladosInner struct {
	XMLName   xml.Name                   `xml:"cfdi:Traslados"`
	Traslados CFDIImpuestosTrasladoInner `xml:"cfdi:Traslado"`
}

// CFDIImpuestosTrasladoInner Nodo requerido para asentar la información detallada de un traslado de impuestos aplicable al presente concepto.
type CFDIImpuestosTrasladoInner struct {
	XMLName    xml.Name `xml:"cfdi:Traslado"`
	Base       float64  `xml:"Base,attr"`       // Atributo requerido para señalar la base para el cálculo del impuesto, la determinación de la base se realiza de acuerdo con las disposiciones fiscales vigentes. No se permiten valores negativos.
	Impuesto   string   `xml:"Impuesto,attr"`   // Atributo requerido para señalar la clave del tipo de impuesto trasladado aplicable al concepto.
	TipoFactor string   `xml:"TipoFactor,attr"` // Atributo requerido para señalar la clave del tipo de factor que se aplica a la base del impuesto.
	TasaOCuota float64  `xml:"TasaOCuota,attr"` // Atributo condicional para señalar el valor de la tasa o cuota del impuesto que se traslada para el presente concepto. Es requerido cuando el atributo TipoFactor tenga un valor que corresponda a Tasa o Cuota.
	Importe    float64  `xml:"Importe,attr"`    // Atributo condicional para señalar el importe del impuesto trasladado que aplica al concepto. No se permiten valores negativos. Es requerido cuando TipoFactor sea Tasa o Cuota
}

// CFDIImpuestosRetencionesInner Nodo opcional para asentar los impuestos retenidos aplicables al presente concepto.
type CFDIImpuestosRetencionesInner struct {
	XMLName     xml.Name                    `xml:"cfdi:Retenciones"`
	Retenciones CFDIImpuestosRetencionInner `xml:"cfdi:Retencion"`
}

// CFDIImpuestosRetencionInner Nodo requerido para asentar la información detallada de una retención de impuestos aplicable al presente concepto.
type CFDIImpuestosRetencionInner struct {
	XMLName    xml.Name `xml:"cfdi:Retencion"`
	Base       float64  `xml:"Base,attr"`       // Atributo requerido para señalar la base para el cálculo del impuesto, la determinación de la base se realiza de acuerdo con las disposiciones fiscales vigentes. No se permiten valores negativos.
	Impuesto   string   `xml:"Impuesto,attr"`   // Atributo requerido para señalar la clave del tipo de impuesto trasladado aplicable al concepto.
	TipoFactor string   `xml:"TipoFactor,attr"` // Atributo requerido para señalar la clave del tipo de factor que se aplica a la base del impuesto.
	TasaOCuota float64  `xml:"TasaOCuota,attr"` // Atributo condicional para señalar el valor de la tasa o cuota del impuesto que se traslada para el presente concepto. Es requerido cuando el atributo TipoFactor tenga un valor que corresponda a Tasa o Cuota.
	Importe    float64  `xml:"Importe,attr"`    // Atributo condicional para señalar el importe del impuesto trasladado que aplica al concepto. No se permiten valores negativos. Es requerido cuando TipoFactor sea Tasa o Cuota
}

// // CFDIInformacionAduanera Nodo opcional para introducir la información aduanera aplicable cuando se trate de ventas de primera mano de mercancías importadas o se trate de operaciones de comercio exterior con bienes o servicios.
// type CFDIInformacionAduanera struct {
// 	NumeroPedimento string // Atributo requerido para expresar el número del pedimento que ampara la importación del bien que se expresa en el siguiente formato: últimos 2 dígitos del año de validación seguidos por dos espacios, 2 dígitos de la aduana de despacho seguidos por dos espacios, 4 dígitos del número de la patente seguidos por dos espacios, 1 dígito que corresponde al último dígito del año en curso, salvo que se trate de un pedimento consolidado iniciado en el año inmediato anterior o del pedimento original de una rectificación, seguido de 6 dígitos de la numeración progresiva por aduana. Pattern [0-9]{2} [0-9]{2} [0-9]{4} [0-9]{7}
// }

// // CFDICuentaPredial Nodo opcional para asentar el número de cuenta predial con el que fue registrado el inmueble, en el sistema catastral de la entidad federativa de que trate, o bien para incorporar los datos de identificación del certificado de participación inmobiliaria no amortizable.
// type CFDICuentaPredial struct {
// 	Numero string // Atributo requerido para precisar el número de la cuenta predial del inmueble cubierto por el presente concepto, o bien para incorporar los datos de identificación del certificado de participación inmobiliaria no amortizable, tratándose de arrendamiento Pattern [0-9]{1,150}
// }

// // ComplementoConcepto Nodo opcional donde se incluyen los nodos complementarios de extensión al concepto definidos por el SAT, de acuerdo con las disposiciones particulares para un sector o actividad específica.
// type ComplementoConcepto struct {
// 	Complemento interface{}
// }

// /*****************************************************************************************************************************************
// *
// * Seccion relacionada con el nodo Impuestos del CFD
// *
// *
// ****************************************************************************************************************************************/

// // CFDIImpuestos Nodo condicional para expresar el resumen de los impuestos aplicables.
// type CFDIImpuestos struct {
// 	TotalImpuestosRetenidos   float64 // Atributo condicional para expresar el total de los impuestos retenidos que se desprenden de los conceptos expresados en el comprobante fiscal digital por Internet. No se permiten valores negativos. Es requerido cuando en los conceptos se registren impuestos retenidos
// 	TotalImpuestosTrasladados float64 // Atributo condicional para expresar el total de los impuestos trasladados que se desprenden de los conceptos expresados en el comprobante fiscal digital por Internet. No se permiten valores negativos. Es requerido cuando en los conceptos se registren impuestos trasladados.
// 	CFDIRetenciones
// 	CFDITraslados
// }

// // CFDIRetenciones Nodo condicional para capturar los impuestos retenidos aplicables. Es requerido cuando en los conceptos se registre algún impuesto retenido.
// type CFDIRetenciones struct {
// 	Retenciones []CFDIRetencion
// }

// // CFDIRetencion Nodo requerido para la información detallada de una retención de impuesto específico
// type CFDIRetencion struct {
// 	Impuesto string  // Atributo requerido para señalar la clave del tipo de impuesto retenido
// 	Importe  float64 // Atributo requerido para señalar el monto del impuesto retenido. No se permiten valores negativos.
// }

// // CFDITraslados Nodo condicional para capturar los impuestos trasladados aplicables. Es requerido cuando en los conceptos se registre un impuesto trasladado.
// type CFDITraslados struct {
// 	Traslados []CFDITraslado
// }

// // CFDITraslado Nodo requerido para la información detallada de un traslado de impuesto específico.
// type CFDITraslado struct {
// 	Impuesto   string  // Atributo requerido para señalar la clave del tipo de impuesto trasladado.
// 	TipoFactor string  // Atributo requerido para señalar la clave del tipo de factor que se aplica a la base del impuesto.
// 	TasaOCuota float64 // Atributo requerido para señalar el valor de la tasa o cuota del impuesto que se traslada por los conceptos amparados en el comprobante.
// 	Importe    float64 // Atributo requerido para señalar la suma del importe del impuesto trasladado, agrupado por impuesto, TipoFactor y TasaOCuota. No se permiten valores negativos.
// }

// /*
//  */

// // CFDIComplemento Nodo opcional donde se incluye el complemento Timbre Fiscal Digital de manera obligatoria y los nodos complementarios determinados por el SAT, de acuerdo con las disposiciones particulares para un sector o actividad específica.
// type CFDIComplemento struct {
// 	elemento interface{}
// }

// /*
//  */

// // CFDIAddenda Nodo opcional para recibir las extensiones al presente formato que sean de utilidad al contribuyente. Para las reglas de uso del mismo, referirse al formato origen.
// type CFDIAddenda struct {
// 	elemento interface{}
// }

//MarshallData2XML Transformar Estructura a XML
func MarshallData2XML(comprobante Comprobante) {

	output, err := xml.MarshalIndent(comprobante, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(string(output))

}
