package services

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/beevik/etree"
	dsig "github.com/russellhaering/goxmldsig"
	"software.sslmate.com/src/go-pkcs12"
)

// AssinarXML assina o XML da NFe usando RSA-SHA1
func AssinarXML(xmlContent string, privateKey *rsa.PrivateKey, certificate *x509.Certificate) (string, error) {
	// Parse do XML
	fmt.Println("XML recebido:", xmlContent)
	doc := etree.NewDocument()
	err := doc.ReadFromString(xmlContent)
	if err != nil {
		return "", fmt.Errorf("erro ao parsear o XML: %v", err)
	}

	// Encontrar o elemento `infNFe` com o atributo Id
	element := doc.FindElement("//*[@Id]")
	if element == nil {
		return "", fmt.Errorf("elemento com atributo `Id` não encontrado no XML")
	}

	// Canonicalizar o elemento `infNFe`
	canonicalizer := dsig.MakeC14N10ExclusiveCanonicalizerWithPrefixList("")
	canonicalized, err := canonicalizer.Canonicalize(element)
	if err != nil {
		return "", fmt.Errorf("erro ao canonicalizar o elemento: %v", err)
	}

	// Calcular o hash SHA-1 do conteúdo canonicalizado
	hash := sha1.Sum(canonicalized)

	// Gerar a assinatura digital usando RSA-SHA1
	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA1, hash[:])
	if err != nil {
		return "", fmt.Errorf("erro ao gerar assinatura digital: %v", err)
	}

	// Codificar a assinatura em Base64
	signatureValue := base64.StdEncoding.EncodeToString(signature)

	// Gerar o elemento <Signature>
	signatureElement := gerarElementoSignature(signatureValue, certificate, element.SelectAttrValue("Id", ""))

	// Adicionar o elemento <Signature> ao documento
	root := doc.Root()
	root.AddChild(signatureElement)

	// Gerar o XML final assinado
	finalXML, err := doc.WriteToString()
	if err != nil {
		return "", fmt.Errorf("erro ao gerar o XML final: %v", err)
	}
	fmt.Println("XML assinado:", finalXML)
	return finalXML, nil
}

func CarregarCertificado(caminhoCert string, senha string) (*rsa.PrivateKey, *x509.Certificate, error) {
	// Carregar certificado
	pfxData, err := os.ReadFile(caminhoCert)
	if err != nil {
		return nil, nil, fmt.Errorf("erro ao ler o arquivo do certificado: %v", err)
	}
	privateKey, cert, err := pkcs12.Decode(pfxData, senha)
	if err != nil {
		return nil, nil, fmt.Errorf("erro ao decodificar o certificado: %v", err)
	}
	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, fmt.Errorf("chave privada não é RSA")
	}
	return rsaKey, cert, nil
}

func gerarElementoSignature(signatureValue string, certificate *x509.Certificate, referenceID string) *etree.Element {
	// Criar elemento <Signature>
	signature := etree.NewElement("Signature")
	signature.CreateAttr("xmlns", "http://www.w3.org/2000/09/xmldsig#")

	// Criar <SignedInfo>
	signedInfo := etree.NewElement("SignedInfo")

	// Adicionar <CanonicalizationMethod>
	canonicalizationMethod := etree.NewElement("CanonicalizationMethod")
	canonicalizationMethod.CreateAttr("Algorithm", "http://www.w3.org/TR/2001/REC-xml-c14n-20010315")
	signedInfo.AddChild(canonicalizationMethod)

	// Adicionar <SignatureMethod>
	signatureMethod := etree.NewElement("SignatureMethod")
	signatureMethod.CreateAttr("Algorithm", "http://www.w3.org/2000/09/xmldsig#rsa-sha1")
	signedInfo.AddChild(signatureMethod)

	// Adicionar <Reference>
	reference := etree.NewElement("Reference")
	reference.CreateAttr("URI", "#"+referenceID)

	// Adicionar <Transforms>
	transforms := etree.NewElement("Transforms")
	transform1 := etree.NewElement("Transform")
	transform1.CreateAttr("Algorithm", "http://www.w3.org/2000/09/xmldsig#enveloped-signature")
	transforms.AddChild(transform1)

	transform2 := etree.NewElement("Transform")
	transform2.CreateAttr("Algorithm", "http://www.w3.org/TR/2001/REC-xml-c14n-20010315")
	transforms.AddChild(transform2)

	reference.AddChild(transforms)

	// Adicionar <DigestMethod>
	digestMethod := etree.NewElement("DigestMethod")
	digestMethod.CreateAttr("Algorithm", "http://www.w3.org/2000/09/xmldsig#sha1")
	reference.AddChild(digestMethod)

	// Adicionar <DigestValue>
	digestValue := etree.NewElement("DigestValue")
	digestValue.SetText(signatureValue) // DigestValue real deve ser calculado
	reference.AddChild(digestValue)

	signedInfo.AddChild(reference)

	// Adicionar <SignedInfo> ao <Signature>
	signature.AddChild(signedInfo)

	// Adicionar <SignatureValue>
	signatureValueElement := etree.NewElement("SignatureValue")
	signatureValueElement.SetText(signatureValue)
	signature.AddChild(signatureValueElement)

	// Adicionar <KeyInfo>
	keyInfo := etree.NewElement("KeyInfo")
	x509Data := etree.NewElement("X509Data")
	x509Certificate := etree.NewElement("X509Certificate")
	x509Certificate.SetText(base64.StdEncoding.EncodeToString(certificate.Raw))
	x509Data.AddChild(x509Certificate)
	keyInfo.AddChild(x509Data)
	signature.AddChild(keyInfo)

	return signature
}
