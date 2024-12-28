package services

import (
	"encoding/xml"
	"fmt"
)

type Ide struct {
	CUF   string `xml:"cUF"`
	NatOp string `xml:"natOp"`
	Mod   string `xml:"mod"`
	Serie string `xml:"serie"`
	NNF   string `xml:"nNF"`
}

type Emit struct {
	CNPJ  string `xml:"CNPJ"`
	XNome string `xml:"xNome"`
}

type InfNFe struct {
	XMLName xml.Name `xml:"infNFe"`
	Versao  string   `xml:"versao,attr"`
	Ide     Ide      `xml:"ide"`
	Emit    Emit     `xml:"emit"`
}

type NFe struct {
	XMLName xml.Name `xml:"nfe"`
	InfNFe  InfNFe   `xml:"infNFe"`
}

func GerarXMLNFe() (string, error) {
	nfe := NFe{
		InfNFe: InfNFe{
			Versao: "4.00",
			Ide: Ide{
				CUF:   "35",
				NatOp: "Venda",
				Mod:   "55",
				Serie: "1",
				NNF:   "12345",
			},
			Emit: Emit{
				CNPJ:  "12345678901234",
				XNome: "Empresa Teste LTDA",
			},
		},
	}

	xmlBytes, err := xml.MarshalIndent(nfe, "", "  ")
	if err != nil {
		return "", fmt.Errorf("erro ao gerar o XML: %w", err)
	}

	xmlString := `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + string(xmlBytes)
	return xmlString, nil
}
