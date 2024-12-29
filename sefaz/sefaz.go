package sefaz

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/beevik/etree"
	"github.com/eugustavokeller/nfe-go/services"
	dsig "github.com/russellhaering/goxmldsig"
)

// Configurações e certificados necessários
type Configuracoes struct {
	EmitenteID       int
	CertificadoPath  string
	CertificadoSenha string
	Ambiente         string // "producao" ou "homologacao"
	SiglaUF          string
}

type NotaFiscal struct {
	ID          int
	Modelo      string
	XML         string
	Status      string
	Recibo      string
	ChaveAcesso string
}

type SefazTools struct {
	Configuracoes Configuracoes
	Certificado   *x509.Certificate
	PrivateKey    *rsa.PrivateKey
	URLPortal     string
}

// Estrutura para resposta do SEFAZ
type SefazResponse struct {
	CStat      int
	Motivo     string
	Recibo     string
	Protocolo  string
	XMLRetorno string
}

func NewSefazTools(config Configuracoes) (*SefazTools, error) {
	// Carregar certificado e chave privada
	privKey, cert, err := services.CarregarCertificado(config.CertificadoPath, config.CertificadoSenha)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar certificado: %v", err)
	}

	urlSefaz := os.Getenv("SEFAZ_URL_HOMOLOGACAO")
	if config.Ambiente == "producao" {
		urlSefaz = os.Getenv("SEFAZ_URL")
	}
	return &SefazTools{
		Configuracoes: config,
		Certificado:   cert,
		PrivateKey:    privKey,
		URLPortal:     urlSefaz,
	}, nil
}

func (t *SefazTools) AssinarXML(xmlContent string) (string, error) {
	if t.Certificado == nil || t.PrivateKey == nil {
		return "", errors.New("certificado ou chave privada não carregados")
	}

	// Assinar o XML com a chave privada
	assinatura, err := AssinarConteudo(xmlContent, t.PrivateKey, "NFe41240706101244000490550010000067271091023595")
	if err != nil {
		return "", fmt.Errorf("erro ao assinar XML: %v", err)
	}

	// Inserir a assinatura no XML
	xmlAssinado := InserirAssinaturaNoXML(xmlContent, assinatura)
	fmt.Println("XML assinado sefaz.go - 89:", xmlAssinado)
	return xmlAssinado, nil
}

func (t *SefazTools) EnviarLote(notasFiscais []NotaFiscal, idLote string, indSinc int) (*SefazResponse, error) {
	if len(notasFiscais) == 0 {
		return nil, errors.New("nenhuma nota fiscal fornecida para envio")
	}

	// Gerar XML do lote
	var loteXML string
	for _, nota := range notasFiscais {
		loteXML += nota.XML
	}
	requestXML := fmt.Sprintf(`
		<enviNFe xmlns="http://www.portalfiscal.inf.br/nfe" versao="4.00">
			<idLote>%s</idLote>
			<indSinc>%d</indSinc>
			%s
		</enviNFe>
	`, idLote, indSinc, loteXML)

	// Enviar para o endpoint SEFAZ
	urlServico := fmt.Sprintf("%s/ws/NfeAutorizacao/NFeAutorizacao4.asmx", t.URLPortal)
	responseXML, err := EnviarSOAP(urlServico, requestXML)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar lote para SEFAZ: %v", err)
	}

	// Processar resposta
	response := &SefazResponse{}
	err = xml.Unmarshal([]byte(responseXML), response)
	if err != nil {
		return nil, fmt.Errorf("erro ao processar resposta da SEFAZ: %v", err)
	}

	return response, nil
}

func ConsultarRecibo(url string, recibo string) (string, error) {
	soapEnvelope := fmt.Sprintf(`
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:nfe="http://www.portalfiscal.inf.br/nfe/wsdl/NFeRetAutorizacao4">
			<soapenv:Header/>
			<soapenv:Body>
				<nfe:consReciNFe>
					<tpAmb>2</tpAmb>
					<nRec>%s</nRec>
				</nfe:consReciNFe>
			</soapenv:Body>
		</soapenv:Envelope>
	`, recibo)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(soapEnvelope))
	if err != nil {
		return "", fmt.Errorf("erro ao criar requisição SOAP: %v", err)
	}

	request.Header.Set("Content-Type", "text/xml; charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("erro ao enviar requisição SOAP: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta SOAP: %v", err)
	}

	// Parse da resposta para verificar status
	var soapResp struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			XMLName  xml.Name `xml:"Body"`
			Response struct {
				XMLName xml.Name `xml:"retConsReciNFe"`
				CStat   string   `xml:"cStat"`
				XMotivo string   `xml:"xMotivo"`
			} `xml:"retConsReciNFe"`
		} `xml:"Body"`
	}

	err = xml.Unmarshal(body, &soapResp)
	if err != nil {
		return "", fmt.Errorf("erro ao parsear resposta SOAP: %v", err)
	}

	return soapResp.Body.Response.XMotivo, nil
}

func EnviarSOAP(url string, xmlContent string) (string, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(xmlContent))
	if err != nil {
		return "", fmt.Errorf("erro ao criar requisição SOAP: %v", err)
	}

	request.Header.Set("Content-Type", "text/xml; charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("erro ao enviar requisição SOAP: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta SOAP: %v", err)
	}

	return string(body), nil
}

// AssinarConteudo realiza a assinatura do XML com base em uma chave privada
func AssinarConteudo(xmlContent string, privateKey *rsa.PrivateKey, elementID string) (string, error) {
	// Parse do XML usando etree
	doc := etree.NewDocument()
	err := doc.ReadFromString(xmlContent)
	if err != nil {
		return "", fmt.Errorf("erro ao parsear o XML: %v", err)
	}

	// Localizar o elemento a ser assinado
	element := doc.FindElement(fmt.Sprintf("//*[@Id='%s']", elementID))
	if element == nil {
		return "", fmt.Errorf("elemento com ID '%s' não encontrado", elementID)
	}

	// Canonicalizar o elemento
	canonicalizer := dsig.MakeC14N10ExclusiveCanonicalizerWithPrefixList("")
	canonicalized, err := canonicalizer.Canonicalize(element)
	if err != nil {
		return "", fmt.Errorf("erro ao canonicalizar o elemento: %v", err)
	}

	// Calcular o hash SHA-256 do elemento canonicalizado
	hash := sha256.Sum256(canonicalized)

	// Criar a assinatura RSA-SHA256 do hash
	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", fmt.Errorf("erro ao assinar o hash: %v", err)
	}

	// Codificar a assinatura em base64
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	return signatureBase64, nil
}

func InserirAssinaturaNoXML(xmlContent, assinatura string) string {
	// Inserir a assinatura no local apropriado do XML
	// Assumimos que a assinatura deve ser inserida imediatamente antes do fechamento da tag `</infNFe>`
	insertPoint := "</infNFe>"
	signedXML := strings.Replace(xmlContent, insertPoint, assinatura+insertPoint, 1)

	return signedXML
}
