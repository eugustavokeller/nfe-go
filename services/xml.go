package services

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
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
	XMLName xml.Name `xml:"total"`
	ICMSTot ICMSTot  `xml:"ICMSTot"`
}

type ICMSTot struct {
	XMLName    xml.Name `xml:"ICMSTot"`
	VBC        string   `xml:"vBC"`
	VICMS      string   `xml:"vICMS"`
	VICMSDeson string   `xml:"vICMSDeson,omitempty"`
	VFCP       string   `xml:"vFCP,omitempty"`
	VBSCST     string   `xml:"vBCST,omitempty"`
	VST        string   `xml:"vST,omitempty"`
	VFCPST     string   `xml:"vFCPST,omitempty"`
	VFCPSTRet  string   `xml:"vFCPSTRet,omitempty"`
	VProd      string   `xml:"vProd"`
	VFrete     string   `xml:"vFrete,omitempty"`
	VSeg       string   `xml:"vSeg,omitempty"`
	VDesc      string   `xml:"vDesc,omitempty"`
	VII        string   `xml:"vII,omitempty"`
	VIPI       string   `xml:"vIPI,omitempty"`
	VIPIDevol  string   `xml:"vIPIDevol,omitempty"`
	VPIS       string   `xml:"vPIS,omitempty"`
	VCOFINS    string   `xml:"vCOFINS,omitempty"`
	VOutro     string   `xml:"vOutro,omitempty"`
	VNF        string   `xml:"vNF"`
	VTotTrib   string   `xml:"vTotTrib,omitempty"`
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

type AutXML struct {
	XMLName xml.Name `xml:"autXML"`
	CNPJ    string   `xml:"CNPJ"`
}

type Transp struct {
	XMLName    xml.Name   `xml:"transp"`
	ModFrete   string     `xml:"modFrete"`
	Transporta Transporta `xml:"transporta"`
	Vol        Vol        `xml:"vol"`
}

type Transporta struct {
	XMLName xml.Name `xml:"transporta"`
	CNPJ    string   `xml:"CNPJ"`
	XNome   string   `xml:"xNome"`
	XEnder  string   `xml:"xEnder"`
	XMun    string   `xml:"xMun"`
	UF      string   `xml:"UF"`
}

type Vol struct {
	XMLName xml.Name `xml:"vol"`
	QVol    string   `xml:"qVol"`
	Esp     string   `xml:"esp"`
	PesoL   string   `xml:"pesoL"`
	PesoB   string   `xml:"pesoB"`
}

type Cobr struct {
	XMLName xml.Name `xml:"cobr"`
	Fat     Fat      `xml:"fat"`
	Dup     []Dup    `xml:"dup"`
}

type Fat struct {
	XMLName xml.Name `xml:"fat"`
	NFat    string   `xml:"nFat"`
	VOrig   string   `xml:"vOrig"`
	VDesc   string   `xml:"vDesc"`
	VLiq    string   `xml:"vLiq"`
}

type Dup struct {
	XMLName xml.Name `xml:"dup"`
	NDup    string   `xml:"nDup"`
	DVenc   string   `xml:"dVenc"`
	VDup    string   `xml:"vDup"`
}

type Pag struct {
	XMLName xml.Name `xml:"pag"`
	DetPag  DetPag   `xml:"detPag"`
}

type DetPag struct {
	XMLName xml.Name `xml:"detPag"`
	IndPag  string   `xml:"indPag"`
	TPag    string   `xml:"tPag"`
	VPag    string   `xml:"vPag"`
}

type InfAdic struct {
	XMLName xml.Name `xml:"infAdic"`
	InfCpl  string   `xml:"infCpl"`
}

type InfRespTec struct {
	XMLName  xml.Name `xml:"infRespTec"`
	CNPJ     string   `xml:"CNPJ"`
	XContato string   `xml:"xContato"`
	Email    string   `xml:"email"`
	Fone     string   `xml:"fone"`
}

// Função para gerar a chave de acesso da NFe
func GerarChaveAcesso(ide Ide, emit Emit, nNF string, tpEmis string) (string, error) {
	// Validar e formatar a data de emissão
	// Data no formato AAAAMM (ano e mês)
	anoMes, err := time.Parse("2006-01-02", ide.DhEmi[:10])
	if err != nil {
		return "", fmt.Errorf("data de emissão inválida: %v", err)
	}
	anoMesFormatado := anoMes.Format("200601")
	// Código aleatório de 8 dígitos (cNF)
	cNF := ide.CNF
	if len(cNF) != 8 {
		return "", fmt.Errorf("o campo cNF deve ter 8 dígitos")
	}
	// Montar a chave de acesso sem o dígito verificador
	chaveSemDV := fmt.Sprintf(
		"%02s%s%s%s%02s%03s%09s%08s%s",
		ide.CUF,         // Código da UF
		anoMesFormatado, // Ano e mês (AAAAMM)
		emit.CNPJ,       // CNPJ do emitente
		ide.Mod,         // Modelo (55 ou 65)
		ide.Serie,       // Série da nota
		nNF,             // Número da nota fiscal
		tpEmis,          // Tipo de emissão
		cNF,             // Código numérico
		"0",             // Placeholder para o dígito verificador
	)
	// Calcular o dígito verificador (DV) da chave
	dv, err := calcularDV(chaveSemDV[:43])
	if err != nil {
		return "", fmt.Errorf("erro ao calcular dígito verificador: %v", err)
	}
	// Adicionar o dígito verificador à chave
	chaveAcesso := chaveSemDV[:43] + strconv.Itoa(dv)
	return chaveAcesso, nil
}

// Função para calcular o Dígito Verificador (DV) da chave de acesso
func calcularDV(chave string) (int, error) {
	if len(chave) != 43 {
		return 0, fmt.Errorf("a chave de acesso deve ter exatamente 43 caracteres")
	}
	pesos := []int{4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	total := 0
	for i, r := range chave {
		valor, err := strconv.Atoi(string(r))
		if err != nil {
			return 0, fmt.Errorf("erro ao converter caractere para número: %v", err)
		}
		total += valor * pesos[i%len(pesos)]
	}
	resto := total % 11
	if resto == 0 || resto == 1 {
		return 0, nil
	}
	return 11 - resto, nil
}
