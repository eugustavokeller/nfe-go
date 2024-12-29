package main

import (
	"fmt"
	"os"

	"github.com/eugustavokeller/nfe-go/sefaz"
	"github.com/eugustavokeller/nfe-go/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return
	}
	// Inicializar configurações e ferramentas SEFAZ
	config := sefaz.Configuracoes{
		CertificadoPath:  os.Getenv("CERTIFICATE_PATH"),
		CertificadoSenha: os.Getenv("CERTIFICATE_PASSWORD"),
		Ambiente:         os.Getenv("AMBIENTE"),
		SiglaUF:          "SC",
	}
	tools, err := sefaz.NewSefazTools(config)
	if err != nil {
		fmt.Println("Erro ao inicializar ferramentas SEFAZ:", err)
		return
	}
	// Gerar XML
	xmlString, err := services.GerarXMLNFe()
	if err != nil {
		fmt.Println("Erro ao gerar XML:", err)
		return
	}
	// Assinar XML
	xmlAssinado, err := tools.AssinarXML(xmlString)
	if err != nil {
		fmt.Println("Erro ao assinar XML:", err)
		return
	}
	// Enviar Lote
	response, err := tools.EnviarLote([]sefaz.NotaFiscal{{XML: xmlAssinado}}, "123456", 0)
	if err != nil {
		fmt.Println("Erro ao enviar lote:", err)
		return
	}
	fmt.Println("Lote enviado com sucesso!")
	fmt.Printf("Resposta SEFAZ: %+v\n", response)
}
