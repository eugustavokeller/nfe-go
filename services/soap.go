package services

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/eugustavokeller/nfe-go/myservice"
	"github.com/hooklift/gowsdl/soap"
)

func EnviarSOAP(url string, xmlContent string) (string, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	soapEnvelope := fmt.Sprintf(`
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:nfe="http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4">
			<soapenv:Header/>
			<soapenv:Body>
				<nfe:nfeDadosMsg>
					%s
				</nfe:nfeDadosMsg>
			</soapenv:Body>
		</soapenv:Envelope>
	`, xmlContent)
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
	return string(body), nil
}

func createSoapClientWithCertificate(sefazURL string) (*soap.Client, error) {
	// Caminho e senha do certificado digital do cliente
	certPath := os.Getenv("CERTIFICATE_PATH")         // Caminho para o .pfx
	certPassword := os.Getenv("CERTIFICATE_PASSWORD") // Senha do certificado
	// Caminho do certificado raiz da SEFAZ
	sefazRootCertPath := os.Getenv("CERTIFICATE_ROOT_SEFAZ")
	// Carregar o certificado digital do cliente
	privateKey, cert, err := CarregarCertificado(certPath, certPassword)
	if err != nil {
		return nil, fmt.Errorf("falha ao carregar certificado do cliente: %w", err)
	}
	// Configurar o certificado TLS
	tlsCert := tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  privateKey,
	}
	// Carregar o certificado raiz da SEFAZ
	rootCert, err := os.ReadFile(sefazRootCertPath)
	if err != nil {
		return nil, fmt.Errorf("falha ao carregar o certificado raiz da SEFAZ: %w", err)
	}
	// Adicionar o certificado raiz ao pool de certificados confiáveis
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(rootCert) {
		return nil, fmt.Errorf("falha ao adicionar o certificado raiz ao pool")
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		RootCAs:      certPool,
	}
	// Criar cliente SOAP configurado com opções
	soapClient := soap.NewClient(sefazURL, soap.WithTLS(tlsConfig))
	return soapClient, nil
}

func ConsultarStatusProtocolo(protocolo string) (string, error) {
	sefazURL := os.Getenv("SEFAZ_URL_CONSULTA_HOMOLOGACAO")
	if os.Getenv("AMBIENTE") == "producao" {
		sefazURL = os.Getenv("SEFAZ_URL_CONSULTA")
	}
	soapClient := soap.NewClient(sefazURL)
	service := myservice.NewNFeAutorizacao4PortType(soapClient)
	request := &myservice.NfeAutorizacaoLoteRequest{
		NfeDadosMsg: protocolo,
	}
	for i := 0; i < 5; i++ { // Limite de 5 tentativas
		response, err := service.NfeAutorizacaoLoteContext(context.Background(), request)
		if err != nil {
			return "", fmt.Errorf("erro ao consultar o status do protocolo: %w", err)
		}
		status := response.RetNFeAutorizacaoLote
		fmt.Println("Status do protocolo:", status)
		if status == "Autorizado" || status == "Rejeitado" {
			return status, nil
		}
		fmt.Println("Nota em processamento. Tentando novamente em 5 segundos...")
		time.Sleep(5 * time.Second)
	}
	return "", fmt.Errorf("tempo limite atingido para consulta do protocolo")
}
