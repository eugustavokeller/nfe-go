package main

import (
	"fmt"
	"os"

	"github.com/eugustavokeller/nfe-go/services"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return
	}

	caminhoCert := os.Getenv("CERTIFICATE_PATH")
	senhaCert := os.Getenv("CERTIFICATE_PASSWORD")
	if caminhoCert == "" || senhaCert == "" {
		fmt.Println("Caminho do certificado ou senha não definidos no .env")
		return
	}

	privateKey, cert, err := services.CarregarCertificado(caminhoCert, senhaCert)
	if err != nil {
		fmt.Println("Erro ao carregar o certificado:", err)
		return
	}

	fmt.Printf("Certificado carregado: %s\n", cert.Subject.CommonName)

	// Gerar XML
	xmlString, err := services.GerarXMLNFe()
	if err != nil {
		fmt.Println("Erro ao gerar XML:", err)
		return
	}

	// Assinar XML
	xmlAssinado, err := services.AssinarXML(xmlString, privateKey)
	if err != nil {
		fmt.Println("Erro ao assinar XML:", err)
		return
	}

	// Enviar XML via SOAP e receber o protocolo
	protocolo, err := services.EnviarXMLSoap(xmlAssinado)
	if err != nil {
		fmt.Println("Erro ao enviar XML via SOAP:", err)
		return
	}

	// Consultar o status do protocolo
	status, err := services.ConsultarStatusProtocolo(protocolo)
	if err != nil {
		fmt.Println("Erro ao consultar o status do protocolo:", err)
		return
	}

	fmt.Printf("Status final da nota: %s\n", status)
	fmt.Println("Processo concluído com sucesso!")
}
