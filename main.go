package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"software.sslmate.com/src/go-pkcs12"
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
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return
	}

	// Obter caminho e senha do certificado das variáveis de ambiente
	caminhoCert := os.Getenv("CERTIFICATE_PATH")
	senhaCert := os.Getenv("CERTIFICATE_PASSWORD")

	if caminhoCert == "" || senhaCert == "" {
		fmt.Println("Caminho do certificado ou senha não definidos no .env")
		return
	}

	privateKey, cert, err := carregarCertificado(caminhoCert, senhaCert)
	if err != nil {
		fmt.Println("Erro ao carregar o certificado:", err)
		return
	}

	// Gerar um XML básico para teste
	xmlTeste := `<nfe>
		<infNFe versao="4.00">
			<ide>
			<cUF>35</cUF>
			<natOp>Venda</natOp>
			<mod>55</mod>
			<serie>1</serie>
			<nNF>12345</nNF>
			</ide>
			<emit>
			<CNPJ>12345678901234</CNPJ>
			<xNome>Empresa Teste LTDA</xNome>
			</emit>
		</infNFe>
	</nfe>`

	// Assinar o XML
	assinatura, err := assinarXML(xmlTeste, privateKey)
	if err != nil {
		fmt.Println("Erro ao assinar o XML:", err)
		return
	}

	fmt.Println("Assinatura gerada com sucesso!")
	fmt.Println("Assinatura Base64:", assinatura)

	fmt.Println("Certificado carregado com sucesso!")
	fmt.Printf("Nome do certificado: %s\n", cert.Subject.CommonName)

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

// Função para carregar o certificado do arquivo .pfx
func carregarCertificado(caminhoCert string, senha string) (*rsa.PrivateKey, *x509.Certificate, error) {
	// Carregar o arquivo .pfx
	pfxData, err := os.ReadFile(caminhoCert)
	if err != nil {
		return nil, nil, fmt.Errorf("erro ao carregar o arquivo .pfx: %w", err)
	}

	// Extrair a chave privada e o certificado
	privateKey, cert, err := pkcs12.Decode(pfxData, senha)
	if err != nil {
		return nil, nil, fmt.Errorf("erro ao decodificar o arquivo .pfx: %w", err)
	}

	// Verificar se a chave é do tipo RSA
	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, fmt.Errorf("a chave privada não é do tipo RSA")
	}

	return rsaKey, cert, nil
}

// Função para assinar o XML
func assinarXML(xmlString string, privateKey *rsa.PrivateKey) (string, error) {
	// Criar o hash SHA-1 do XML
	hasher := sha1.New()
	_, err := hasher.Write([]byte(xmlString))
	if err != nil {
		return "", fmt.Errorf("erro ao calcular hash do XML: %w", err)
	}
	hash := hasher.Sum(nil)

	// Assinar o hash com a chave privada
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hash)
	if err != nil {
		return "", fmt.Errorf("erro ao assinar o hash: %w", err)
	}

	// Codificar a assinatura em Base64
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	return signatureBase64, nil
}
