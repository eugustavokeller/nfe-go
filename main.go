package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Estruturas para geração do XML
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

func gerarXMLNFe() (string, error) {
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

func salvarXMLNoArquivo(conteudo, nomeArquivo string) error {
	file, err := os.Create(nomeArquivo)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(conteudo)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %w", err)
	}
	return nil
}

func main() {
	xml, err := gerarXMLNFe()
	if err != nil {
		fmt.Println("Erro ao gerar XML:", err)
		return
	}

	fmt.Println("XML Gerado:\n", xml)

	// Salvar XML em um arquivo
	arquivo := "nfe.xml"
	err = salvarXMLNoArquivo(xml, arquivo)
	if err != nil {
		fmt.Println("Erro ao salvar XML no arquivo:", err)
		return
	}

	fmt.Printf("XML salvo com sucesso no arquivo: %s\n", arquivo)
}
