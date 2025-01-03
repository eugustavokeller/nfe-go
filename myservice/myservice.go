// Code generated by gowsdl DO NOT EDIT.

package myservice

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/hooklift/gowsdl/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type NfeAutorizacaoLoteRequest struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4 nfeAutorizacaoLoteRequest"`

	NfeDadosMsg string `xml:"nfeDadosMsg,omitempty" json:"nfeDadosMsg,omitempty"`
}

type NfeAutorizacaoLoteResponse struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4 nfeAutorizacaoLoteResponse"`

	RetNFeAutorizacaoLote string `xml:"retNFeAutorizacaoLote,omitempty" json:"retNFeAutorizacaoLote,omitempty"`
}

type NFeAutorizacao4PortType interface {
	NfeAutorizacaoLote(request *NfeAutorizacaoLoteRequest) (*NfeAutorizacaoLoteResponse, error)

	NfeAutorizacaoLoteContext(ctx context.Context, request *NfeAutorizacaoLoteRequest) (*NfeAutorizacaoLoteResponse, error)
}

type nFeAutorizacao4PortType struct {
	client *soap.Client
}

func NewNFeAutorizacao4PortType(client *soap.Client) NFeAutorizacao4PortType {
	return &nFeAutorizacao4PortType{
		client: client,
	}
}

func (service *nFeAutorizacao4PortType) NfeAutorizacaoLoteContext(ctx context.Context, request *NfeAutorizacaoLoteRequest) (*NfeAutorizacaoLoteResponse, error) {
	response := new(NfeAutorizacaoLoteResponse)
	err := service.client.CallContext(ctx, "http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4/nfeAutorizacaoLote", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *nFeAutorizacao4PortType) NfeAutorizacaoLote(request *NfeAutorizacaoLoteRequest) (*NfeAutorizacaoLoteResponse, error) {
	return service.NfeAutorizacaoLoteContext(
		context.Background(),
		request,
	)
}
