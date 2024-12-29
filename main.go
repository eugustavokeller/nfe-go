package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/eugustavokeller/nfe-go/sefaz"
	"github.com/eugustavokeller/nfe-go/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Erro ao carregar arquivo .env: %v", err)
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
		fmt.Errorf("Erro ao inicializar ferramentas SEFAZ: %v", err)
		return
	}

	// Gerar chave de acesso
	chave, err := services.GerarChaveAcesso(ide, emit, nNF, tpEmis) // Dados da NFe
	if err != nil {
		fmt.Println("Erro ao gerar chave de acesso:", err)
		return
	}

	fmt.Println("Chave de Acesso Gerada:", chave)

	// Gerar XML
	infNFe := services.DynamicElement{
		XMLName: xml.Name{Local: "infNFe"},
		Attrs: []xml.Attr{
			{Name: xml.Name{Local: "Id"}, Value: "NFe" + chaveAcesso},
			{Name: xml.Name{Local: "versao"}, Value: "4.00"},
		},
		Children: []services.DynamicElement{
			services.MakeTagIde(services.Ide{ /* Dados aqui */ }),
			services.MakeTagEmit(services.Emit{ /* Dados aqui */ }),
			services.MakeTagDest(services.Dest{ /* Dados aqui */ }),
			services.MakeTagAutXML(services.AutXML{ /* Dados aqui */ }),
			services.MakeTagDet(services.Det{ /* Dados aqui */ }),
			services.MakeTagTotal(services.Total{ /* Dados aqui */ }),
			services.MakeTagTransp(services.Transp{ /* Dados aqui */ }),
			services.MakeTagCobr(services.Cobr{ /* Dados aqui */ }),
			services.MakeTagPag(services.Pag{ /* Dados aqui */ }),
			services.MakeTagInfAdic(services.InfAdic{ /* Dados aqui */ }),
			services.MakeTagInfRespTec(services.InfRespTec{ /* Dados aqui */ }),
		},
	}
	root := services.DynamicElement{
		XMLName: xml.Name{Local: "NFe"},
		Attrs: []xml.Attr{
			{Name: xml.Name{Local: "xmlns"}, Value: "http://www.portalfiscal.inf.br/nfe"},
		},
		Children: []services.DynamicElement{infNFe},
	}
	xmlString, err := services.GenerateDynamicXML(root)
	if err != nil {
		fmt.Errorf("Erro ao gerar XML: %v", err)
		return
	}
	fmt.Println("XML gerado com sucesso! \n", xmlString)

	// Assinar XML
	xmlAssinado, err := tools.AssinarXML(xmlString)
	if err != nil {
		fmt.Errorf("Erro ao assinar XML: %v", err)
		return
	}
	fmt.Println("XML assinado com sucesso! \n", xmlAssinado)
	// Enviar Lote
	response, err := tools.EnviarLote([]sefaz.NotaFiscal{{XML: xmlAssinado}}, "123456", 0)
	if err != nil {
		fmt.Errorf("Erro ao enviar lote: %v", err)
		return
	}
	fmt.Println("Lote enviado com sucesso!")
	fmt.Printf("Resposta SEFAZ: %+v\n", response)
}
