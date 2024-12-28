package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/eugustavokeller/nfe-go/myservice"
	"github.com/hooklift/gowsdl/soap"
)

func EnviarXMLSoap(xmlAssinado string) (string, error) {
	sefazURL := os.Getenv("SEFAZ_URL_HOMOLOGACAO")
	if os.Getenv("AMBIENTE") == "producao" {
		sefazURL = os.Getenv("SEFAZ_URL")
	}

	soapClient := soap.NewClient(sefazURL)
	service := myservice.NewNFeAutorizacao4PortType(soapClient)

	request := &myservice.NfeAutorizacaoLoteRequest{
		NfeDadosMsg: xmlAssinado,
	}

	response, err := service.NfeAutorizacaoLoteContext(context.Background(), request)
	if err != nil {
		return "", fmt.Errorf("erro ao enviar XML via SOAP: %w", err)
	}

	// Retornar o protocolo de recebimento
	protocolo := response.RetNFeAutorizacaoLote
	fmt.Println("Protocolo recebido:", protocolo)
	return protocolo, nil
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
