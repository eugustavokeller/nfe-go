package services

import (
	"encoding/xml"
	"fmt"
)

// Estruturas do XML da NFe
type Ide struct {
	CUF      string `xml:"cUF"`
	CNF      string `xml:"cNF"`
	NatOp    string `xml:"natOp"`
	Mod      string `xml:"mod"`
	Serie    string `xml:"serie"`
	NNF      string `xml:"nNF"`
	DhEmi    string `xml:"dhEmi"`
	TpNF     string `xml:"tpNF"`
	IdDest   string `xml:"idDest"`
	CMunFG   string `xml:"cMunFG"`
	TpImp    string `xml:"tpImp"`
	TpEmis   string `xml:"tpEmis"`
	CDV      string `xml:"cDV"`
	TpAmb    string `xml:"tpAmb"`
	FinNFe   string `xml:"finNFe"`
	IndFinal string `xml:"indFinal"`
	IndPres  string `xml:"indPres"`
	ProcEmi  string `xml:"procEmi"`
	VerProc  string `xml:"verProc"`
}

type Emit struct {
	CNPJ      string    `xml:"CNPJ"`
	XNome     string    `xml:"xNome"`
	XFant     string    `xml:"xFant"`
	EnderEmit EnderEmit `xml:"enderEmit"`
	IE        string    `xml:"IE"`
	CRT       string    `xml:"CRT"`
}

type EnderEmit struct {
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XCpl    string `xml:"xCpl,omitempty"`
	XBairro string `xml:"xBairro"`
	CMun    string `xml:"cMun"`
	XMun    string `xml:"xMun"`
	UF      string `xml:"UF"`
	CEP     string `xml:"CEP"`
	CPais   string `xml:"cPais"`
	XPais   string `xml:"xPais"`
	Fone    string `xml:"fone"`
}

type Dest struct {
	CNPJ      string    `xml:"CNPJ"`
	XNome     string    `xml:"xNome"`
	EnderDest EnderDest `xml:"enderDest"`
	IndIEDest string    `xml:"indIEDest"`
	IE        string    `xml:"IE"`
	Email     string    `xml:"email,omitempty"`
}

type EnderDest struct {
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XBairro string `xml:"xBairro"`
	CMun    string `xml:"cMun"`
	XMun    string `xml:"xMun"`
	UF      string `xml:"UF"`
	CEP     string `xml:"CEP"`
	CPais   string `xml:"cPais"`
	XPais   string `xml:"xPais"`
	Fone    string `xml:"fone"`
}

type Det struct {
	NItem string `xml:"nItem,attr"`
	Prod  Prod   `xml:"prod"`
}

type Prod struct {
	CProd    string  `xml:"cProd"`
	XProd    string  `xml:"xProd"`
	CEAN     string  `xml:"cEAN"`
	CEANTrib string  `xml:"cEANTrib"`
	NCM      string  `xml:"NCM"`
	CFOP     string  `xml:"CFOP"`
	UCom     string  `xml:"uCom"`
	QCom     float64 `xml:"qCom"`
	VUnCom   float64 `xml:"vUnCom"`
	VProd    float64 `xml:"vProd"`
	UTrib    string  `xml:"uTrib"`
	QTrib    float64 `xml:"qTrib"`
	VUnTrib  float64 `xml:"vUnTrib"`
	IndTot   string  `xml:"indTot"`
}

type Total struct {
	ICMSTot ICMSTot `xml:"ICMSTot"`
}

type ICMSTot struct {
	VBC   string `xml:"vBC"`
	VICMS string `xml:"vICMS"`
	VProd string `xml:"vProd"`
	VNF   string `xml:"vNF"`
}

type InfNFe struct {
	XMLName xml.Name `xml:"infNFe"`
	Id      string   `xml:"Id,attr"`
	Versao  string   `xml:"versao,attr"`
	Ide     Ide      `xml:"ide"`
	Emit    Emit     `xml:"emit"`
	Dest    Dest     `xml:"dest"`
	Det     []Det    `xml:"det"`
	Total   Total    `xml:"total"`
}

type NFe struct {
	XMLName xml.Name `xml:"NFe"`
	Xmlns   string   `xml:"xmlns,attr"`
	InfNFe  InfNFe   `xml:"infNFe"`
}

// Função para gerar XML da NFe
func GerarXMLNFe() (string, error) {
	nfe := NFe{
		Xmlns: "http://www.portalfiscal.inf.br/nfe",
		InfNFe: InfNFe{
			Id:     "NFe41240706101244000490550010000067271091023595",
			Versao: "4.00",
			Ide: Ide{
				CUF:      "41",
				CNF:      "09102359",
				NatOp:    "Remessa em bonificação, doação ou brinde",
				Mod:      "55",
				Serie:    "1",
				NNF:      "6727",
				DhEmi:    "2024-07-01T16:56:48-03:00",
				TpNF:     "1",
				IdDest:   "1",
				CMunFG:   "4108809",
				TpImp:    "1",
				TpEmis:   "1",
				CDV:      "5",
				TpAmb:    "1",
				FinNFe:   "1",
				IndFinal: "0",
				IndPres:  "0",
				ProcEmi:  "0",
				VerProc:  "1.0.0",
			},
			Emit: Emit{
				CNPJ:  "06101244000490",
				XNome: "INKOR INDUSTRIA CATARINENSE DE COLAS E REJUNTES",
				XFant: "INKOR INDUSTRIA CATARINENSE DE COLAS E REJUNTES",
				EnderEmit: EnderEmit{
					XLgr:    "RUA MINISTRO GABRIEL PASSOS",
					Nro:     "470",
					XCpl:    "GLP 02",
					XBairro: "CENTRO",
					CMun:    "4108809",
					XMun:    "Guaira",
					UF:      "PR",
					CEP:     "85980000",
					CPais:   "1058",
					XPais:   "BRASIL",
					Fone:    "05136633247",
				},
				IE:  "9089733135",
				CRT: "3",
			},
			Dest: Dest{
				CNPJ:  "41521606000150",
				XNome: "RARO COMERCIO DE ACABAMENTOS LTDA",
				EnderDest: EnderDest{
					XLgr:    "AV JOSE MARIA DE BRITO",
					Nro:     "642",
					XBairro: "JARDIM DAS NACOES",
					CMun:    "4108304",
					XMun:    "Foz do Iguacu",
					UF:      "PR",
					CEP:     "85864320",
					CPais:   "1058",
					XPais:   "BRASIL",
					Fone:    "4535221316",
				},
				IndIEDest: "1",
				IE:        "9088804802",
				Email:     "natalinofonseca@hotmaio.com",
			},
			Det: []Det{
				{
					NItem: "1",
					Prod: Prod{
						CProd:    "2413",
						XProd:    "ARGAMASSA SUPERKOLA ACIII 20 KG",
						CEAN:     "7898476909292",
						CEANTrib: "7898476909292",
						NCM:      "32149000",
						CFOP:     "5910",
						UCom:     "SC",
						QCom:     47.0,
						VUnCom:   18.25,
						VProd:    857.75,
						UTrib:    "SC",
						QTrib:    47.0,
						VUnTrib:  18.25,
						IndTot:   "1",
					},
				},
			},
			Total: Total{
				ICMSTot: ICMSTot{
					VBC:   "857.75",
					VICMS: "102.93",
					VProd: "857.75",
					VNF:   "1008.97",
				},
			},
		},
	}

	xmlBytes, err := xml.MarshalIndent(nfe, "", "  ")
	if err != nil {
		return "", fmt.Errorf("erro ao gerar o XML NFe: %w", err)
	}

	xmlString := `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + string(xmlBytes)
	return xmlString, nil
}

func GerarXMLNFCe() (string, error) {
	nfce := NFe{
		InfNFe: InfNFe{
			Versao: "4.00",
			Ide: Ide{
				CUF:      "35",
				NatOp:    "Venda",
				Mod:      "65",
				Serie:    "1",
				NNF:      "12345",
				DhEmi:    "2024-12-28T10:00:00-03:00",
				TpNF:     "1",
				IdDest:   "1",
				CMunFG:   "3550308",
				TpImp:    "4",
				TpEmis:   "1",
				CDV:      "7",
				TpAmb:    "2",
				FinNFe:   "1",
				IndFinal: "1",
				IndPres:  "1",
				ProcEmi:  "0",
				VerProc:  "1.00",
			},
			Emit: Emit{
				CNPJ:  "12345678901234",
				XNome: "Empresa Teste LTDA",
				XFant: "Empresa Teste",
				EnderEmit: EnderEmit{
					XLgr:    "Rua Exemplo",
					Nro:     "123",
					XBairro: "Bairro Exemplo",
					CMun:    "3550308",
					XMun:    "São Paulo",
					UF:      "SP",
					CEP:     "01001000",
					CPais:   "1058",
					XPais:   "Brasil",
					Fone:    "1123456789",
				},
				IE:  "123456789",
				CRT: "3",
			},
			Dest: Dest{
				CNPJ:  "98765432100012",
				XNome: "Cliente Teste LTDA",
				EnderDest: EnderDest{
					XLgr:    "Rua Cliente",
					Nro:     "456",
					XBairro: "Bairro Cliente",
					CMun:    "3550308",
					XMun:    "São Paulo",
					UF:      "SP",
					CEP:     "01002000",
					CPais:   "1058",
					XPais:   "Brasil",
					Fone:    "1123456789",
				},
				IndIEDest: "1",
			},
			Total: Total{
				ICMSTot: ICMSTot{
					VBC:   "100.00",
					VICMS: "18.00",
					VProd: "100.00",
					VNF:   "118.00",
					// VFrete: "0.00",
					// VSeg:   "0.00",
					// VDesc:  "0.00",
					// VOutro: "0.00",
				},
			},
		},
	}

	xmlBytes, err := xml.MarshalIndent(nfce, "", "  ")
	if err != nil {
		return "", fmt.Errorf("erro ao gerar o XML NFCe: %w", err)
	}

	xmlString := `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + string(xmlBytes)
	return xmlString, nil
}
