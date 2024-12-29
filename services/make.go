package services

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type DynamicElement struct {
	XMLName  xml.Name
	Content  string
	Attrs    []xml.Attr
	Children []DynamicElement
}

// Função genérica para gerar XML a partir de dados dinâmicos
func GenerateDynamicXML(root DynamicElement) (string, error) {
	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	encoder.Indent("", "  ")

	if err := encodeDynamicElement(encoder, root); err != nil {
		return "", fmt.Errorf("erro ao gerar XML: %w", err)
	}

	if err := encoder.Flush(); err != nil {
		return "", fmt.Errorf("erro ao finalizar XML: %w", err)
	}

	return `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + buffer.String(), nil
}

func encodeDynamicElement(encoder *xml.Encoder, element DynamicElement) error {
	startElement := xml.StartElement{Name: element.XMLName, Attr: element.Attrs}

	if err := encoder.EncodeToken(startElement); err != nil {
		return err
	}

	if element.Content != "" {
		if err := encoder.EncodeToken(xml.CharData([]byte(element.Content))); err != nil {
			return err
		}
	}

	for _, child := range element.Children {
		if err := encodeDynamicElement(encoder, child); err != nil {
			return err
		}
	}

	if err := encoder.EncodeToken(startElement.End()); err != nil {
		return err
	}

	return nil
}

func MakeTagIde(ide Ide) DynamicElement {
	return DynamicElement{
		XMLName: xml.Name{Local: "ide"},
		Children: []DynamicElement{
			{XMLName: xml.Name{Local: "cUF"}, Content: ide.CUF},
			{XMLName: xml.Name{Local: "cNF"}, Content: ide.CNF},
			{XMLName: xml.Name{Local: "natOp"}, Content: ide.NatOp},
			{XMLName: xml.Name{Local: "mod"}, Content: ide.Mod},
			{XMLName: xml.Name{Local: "serie"}, Content: ide.Serie},
			{XMLName: xml.Name{Local: "nNF"}, Content: ide.NNF},
			{XMLName: xml.Name{Local: "dhEmi"}, Content: ide.DhEmi},
			{XMLName: xml.Name{Local: "tpNF"}, Content: ide.TpNF},
			{XMLName: xml.Name{Local: "idDest"}, Content: ide.IdDest},
			{XMLName: xml.Name{Local: "cMunFG"}, Content: ide.CMunFG},
			{XMLName: xml.Name{Local: "tpImp"}, Content: ide.TpImp},
			{XMLName: xml.Name{Local: "tpEmis"}, Content: ide.TpEmis},
			{XMLName: xml.Name{Local: "cDV"}, Content: ide.CDV},
			{XMLName: xml.Name{Local: "tpAmb"}, Content: ide.TpAmb},
			{XMLName: xml.Name{Local: "finNFe"}, Content: ide.FinNFe},
			{XMLName: xml.Name{Local: "indFinal"}, Content: ide.IndFinal},
			{XMLName: xml.Name{Local: "indPres"}, Content: ide.IndPres},
			{XMLName: xml.Name{Local: "procEmi"}, Content: ide.ProcEmi},
			{XMLName: xml.Name{Local: "verProc"}, Content: ide.VerProc},
		},
	}
}

func MakeTagEmit(emit Emit) DynamicElement {
	return DynamicElement{
		XMLName: xml.Name{Local: "emit"},
		Children: []DynamicElement{
			{XMLName: xml.Name{Local: "CNPJ"}, Content: emit.CNPJ},
			{XMLName: xml.Name{Local: "xNome"}, Content: emit.XNome},
			{XMLName: xml.Name{Local: "xFant"}, Content: emit.XFant},
			{
				XMLName: xml.Name{Local: "enderEmit"},
				Children: []DynamicElement{
					{XMLName: xml.Name{Local: "xLgr"}, Content: emit.EnderEmit.XLgr},
					{XMLName: xml.Name{Local: "nro"}, Content: emit.EnderEmit.Nro},
					{XMLName: xml.Name{Local: "xCpl"}, Content: emit.EnderEmit.XCpl},
					{XMLName: xml.Name{Local: "xBairro"}, Content: emit.EnderEmit.XBairro},
					{XMLName: xml.Name{Local: "cMun"}, Content: emit.EnderEmit.CMun},
					{XMLName: xml.Name{Local: "xMun"}, Content: emit.EnderEmit.XMun},
					{XMLName: xml.Name{Local: "UF"}, Content: emit.EnderEmit.UF},
					{XMLName: xml.Name{Local: "CEP"}, Content: emit.EnderEmit.CEP},
					{XMLName: xml.Name{Local: "cPais"}, Content: emit.EnderEmit.CPais},
					{XMLName: xml.Name{Local: "xPais"}, Content: emit.EnderEmit.XPais},
					{XMLName: xml.Name{Local: "fone"}, Content: emit.EnderEmit.Fone},
				},
			},
			{XMLName: xml.Name{Local: "IE"}, Content: emit.IE},
			{XMLName: xml.Name{Local: "CRT"}, Content: emit.CRT},
		},
	}
}

func MakeTagDest(dest Dest) DynamicElement {
	return DynamicElement{
		XMLName: xml.Name{Local: "dest"},
		Children: []DynamicElement{
			{XMLName: xml.Name{Local: "CNPJ"}, Content: dest.CNPJ},
			{XMLName: xml.Name{Local: "xNome"}, Content: dest.XNome},
			{
				XMLName: xml.Name{Local: "enderDest"},
				Children: []DynamicElement{
					{XMLName: xml.Name{Local: "xLgr"}, Content: dest.EnderDest.XLgr},
					{XMLName: xml.Name{Local: "nro"}, Content: dest.EnderDest.Nro},
					{XMLName: xml.Name{Local: "xBairro"}, Content: dest.EnderDest.XBairro},
					{XMLName: xml.Name{Local: "cMun"}, Content: dest.EnderDest.CMun},
					{XMLName: xml.Name{Local: "xMun"}, Content: dest.EnderDest.XMun},
					{XMLName: xml.Name{Local: "UF"}, Content: dest.EnderDest.UF},
					{XMLName: xml.Name{Local: "CEP"}, Content: dest.EnderDest.CEP},
					{XMLName: xml.Name{Local: "cPais"}, Content: dest.EnderDest.CPais},
					{XMLName: xml.Name{Local: "xPais"}, Content: dest.EnderDest.XPais},
					{XMLName: xml.Name{Local: "fone"}, Content: dest.EnderDest.Fone},
				},
			},
			{XMLName: xml.Name{Local: "indIEDest"}, Content: dest.IndIEDest},
			{XMLName: xml.Name{Local: "IE"}, Content: dest.IE},
			{XMLName: xml.Name{Local: "email"}, Content: dest.Email},
		},
	}
}

func MakeTagDet(det Det) DynamicElement {
	return DynamicElement{
		XMLName: xml.Name{Local: "det"},
		Attrs: []xml.Attr{
			{Name: xml.Name{Local: "nItem"}, Value: det.NItem},
		},
		Children: []DynamicElement{
			{
				XMLName: xml.Name{Local: "prod"},
				Children: []DynamicElement{
					{XMLName: xml.Name{Local: "cProd"}, Content: det.Prod.CProd},
					{XMLName: xml.Name{Local: "xProd"}, Content: det.Prod.XProd},
					{XMLName: xml.Name{Local: "cEAN"}, Content: det.Prod.CEAN},
					{XMLName: xml.Name{Local: "cEANTrib"}, Content: det.Prod.CEANTrib},
					{XMLName: xml.Name{Local: "NCM"}, Content: det.Prod.NCM},
					{XMLName: xml.Name{Local: "CFOP"}, Content: det.Prod.CFOP},
					{XMLName: xml.Name{Local: "uCom"}, Content: det.Prod.UCom},
					{XMLName: xml.Name{Local: "qCom"}, Content: fmt.Sprintf("%.4f", det.Prod.QCom)},
					{XMLName: xml.Name{Local: "vUnCom"}, Content: fmt.Sprintf("%.10f", det.Prod.VUnCom)},
					{XMLName: xml.Name{Local: "vProd"}, Content: fmt.Sprintf("%.2f", det.Prod.VProd)},
					{XMLName: xml.Name{Local: "uTrib"}, Content: det.Prod.UTrib},
					{XMLName: xml.Name{Local: "qTrib"}, Content: fmt.Sprintf("%.4f", det.Prod.QTrib)},
					{XMLName: xml.Name{Local: "vUnTrib"}, Content: fmt.Sprintf("%.10f", det.Prod.VUnTrib)},
					{XMLName: xml.Name{Local: "indTot"}, Content: det.Prod.IndTot},
				},
			},
		},
	}
}

func MakeTagAutXML(autXML AutXML) DynamicElement {
	return DynamicElement{
		XMLName: xml.Name{Local: "autXML"},
		Children: []DynamicElement{
			{XMLName: xml.Name{Local: "CNPJ"}, Content: autXML.CNPJ},
		},
	}
}

func MakeTagTotal(total Total) DynamicElement {
	return DynamicElement{
		XMLName: xml.Name{Local: "total"},
		Children: []DynamicElement{
			{
				XMLName: xml.Name{Local: "ICMSTot"},
				Children: []DynamicElement{
					{XMLName: xml.Name{Local: "vBC"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VBC)},
					{XMLName: xml.Name{Local: "vICMS"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VICMS)},
					{XMLName: xml.Name{Local: "vICMSDeson"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VICMSDeson)},
					{XMLName: xml.Name{Local: "vFCP"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VFCP)},
					{XMLName: xml.Name{Local: "vBCST"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VBCST)},
					{XMLName: xml.Name{Local: "vST"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VST)},
					{XMLName: xml.Name{Local: "vFCPST"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VFCPST)},
					{XMLName: xml.Name{Local: "vFCPSTRet"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VFCPSTRet)},
					{XMLName: xml.Name{Local: "vProd"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VProd)},
					{XMLName: xml.Name{Local: "vFrete"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VFrete)},
					{XMLName: xml.Name{Local: "vSeg"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VSeg)},
					{XMLName: xml.Name{Local: "vDesc"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VDesc)},
					{XMLName: xml.Name{Local: "vII"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VII)},
					{XMLName: xml.Name{Local: "vIPI"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VIPI)},
					{XMLName: xml.Name{Local: "vIPIDevol"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VIPIDevol)},
					{XMLName: xml.Name{Local: "vPIS"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VPIS)},
					{XMLName: xml.Name{Local: "vCOFINS"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VCOFINS)},
					{XMLName: xml.Name{Local: "vOutro"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VOutro)},
					{XMLName: xml.Name{Local: "vNF"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VNF)},
					{XMLName: xml.Name{Local: "vTotTrib"}, Content: fmt.Sprintf("%.2f", total.ICMSTot.VTotTrib)},
				},
			},
		},
	}
}

func MakeTagTransp(transp Transp) DynamicElement {
	return DynamicElement{
		XMLName: xml.Name{Local: "transp"},
		Children: []DynamicElement{
			{XMLName: xml.Name{Local: "modFrete"}, Content: transp.ModFrete},
			{XMLName: xml.Name{Local: "transporta"},
				Children: []DynamicElement{
					{XMLName: xml.Name{Local: "CNPJ"}, Content: transp.Transporta.CNPJ},
					{XMLName: xml.Name{Local: "xNome"}, Content: transp.Transporta.XNome},
					{XMLName: xml.Name{Local: "xEnder"}, Content: transp.Transporta.XEnder},
					{XMLName: xml.Name{Local: "xMun"}, Content: transp.Transporta.XMun},
					{XMLName: xml.Name{Local: "UF"}, Content: transp.Transporta.UF},
				},
			},
			{XMLName: xml.Name{Local: "vol"},
				Children: []DynamicElement{
					{XMLName: xml.Name{Local: "qVol"}, Content: transp.Vol.QVol},
					{XMLName: xml.Name{Local: "esp"}, Content: transp.Vol.Esp},
					{XMLName: xml.Name{Local: "pesoL"}, Content: transp.Vol.PesoL},
					{XMLName: xml.Name{Local: "pesoB"}, Content: transp.Vol.PesoB},
				},
			},
		},
	}
}

func MakeTagCobr(cobr Cobr) DynamicElement {
	return DynamicElement{
		XMLName: xml.Name{Local: "cobr"},
		Children: []DynamicElement{
			{XMLName: xml.Name{Local: "fat"}, Content: cobr.Fat},
			{XMLName: xml.Name{Local: "dup"}, Content: cobr.Dup},
		},
	}
}

func MakeTagPag(pag Pag) DynamicElement {
	return DynamicElement{
		XMLName: xml.Name{Local: "pag"},
		Children: []DynamicElement{
			{XMLName: xml.Name{Local: "detPag"}, Content: pag.DetPag},
		},
	}
}

func MakeTagInfAdic(infAdic InfAdic) DynamicElement {
	return DynamicElement{
		XMLName: xml.Name{Local: "infAdic"},
		Children: []DynamicElement{
			{XMLName: xml.Name{Local: "infCpl"}, Content: infAdic.InfCpl},
		},
	}
}

func MakeTagInfRespTec(infRespTec InfRespTec) DynamicElement {
	return DynamicElement{
		XMLName: xml.Name{Local: "infRespTec"},
		Children: []DynamicElement{
			{XMLName: xml.Name{Local: "CNPJ"}, Content: infRespTec.CNPJ},
			{XMLName: xml.Name{Local: "xContato"}, Content: infRespTec.XContato},
			{XMLName: xml.Name{Local: "email"}, Content: infRespTec.Email},
			{XMLName: xml.Name{Local: "fone"}, Content: infRespTec.Fone},
		},
	}
}
