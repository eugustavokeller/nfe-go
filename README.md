# NFe-Go

Sistema em Go para emissão de Notas Fiscais Eletrônicas (NFe) e integração com a SEFAZ utilizando SOAP, assinatura digital e validação de XML.

## Sumário

- [Descrição](#descrição)
- [Funcionalidades](#funcionalidades)
- [Pré-requisitos](#pré-requisitos)
- [Configuração do Projeto](#configuração-do-projeto)
- [Como Usar](#como-usar)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Contribuições](#contribuições)
- [Licença](#licença)

## Descrição

Este projeto tem como objetivo facilitar a emissão de Notas Fiscais Eletrônicas (NFe) e integração com os serviços da SEFAZ. Ele abrange a geração de XMLs, assinatura digital, envio via SOAP, e consulta do status de processamento das notas.

## Funcionalidades

- Geração de XML para NFe no formato 4.00.
- Assinatura digital dos XMLs utilizando certificados digitais no formato `.pfx`.
- Envio assíncrono de notas para a SEFAZ via SOAP.
- Consulta do status do protocolo de processamento.
- Configuração de ambiente para homologação e produção.

## Pré-requisitos

- Go 1.20 ou superior
- Certificado digital no formato `.pfx`
- Ambiente de homologação/produção configurado
- Conexão com a internet para comunicação com os serviços da SEFAZ

## Configuração do Projeto

1. Clone o repositório:

```bash
git clone https://github.com/eugustavokeller/nfe-go.git
cd nfe-go
```

2. Clone o repositório:

```bash
go mod tidy
```

3. Configure as variáveis de ambiente no arquivo .env:

CERTIFICATE_PATH=/caminho/para/seu_certificado.pfx
CERTIFICATE_PASSWORD=sua_senha
SEFAZ_URL_HOMOLOGACAO=https://homnfe.sefaz.am.gov.br/services2/services/NfeAutorizacao4
SEFAZ_URL=https://nfe.sefaz.am.gov.br/services2/services/NfeAutorizacao4
SEFAZ_URL_CONSULTA_HOMOLOGACAO=https://homnfe.sefaz.am.gov.br/services2/services/NfeConsultaProtocolo4
SEFAZ_URL_CONSULTA=https://nfe.sefaz.am.gov.br/services2/services/NfeConsultaProtocolo4
AMBIENTE=homologacao

4. Compile o projeto:

```bash
go build -o nfe-go
```

5. Execute o programa:

```bash
./nfe-go
```

## Como Usar

1. Certifique-se de que o arquivo .env está devidamente configurado.
2. Ao executar o programa, ele:
   • Gera um XML de teste.
   • Assina digitalmente o XML utilizando o certificado fornecido.
   • Envia o XML para a SEFAZ e retorna o protocolo de recebimento.
   • Realiza consultas periódicas ao status do protocolo.
3. O status final da nota será exibido no terminal.

## Estrutura do Projeto

nfe-go/
├── main.go # Entrada principal do projeto
├── services/
│ ├── cert.go # Funções para carregar e validar certificados digitais
│ ├── xml.go # Geração e manipulação de XMLs
│ ├── soap.go # Envio e consulta via SOAP
├── myservice/
│ └── myservice.go # Código gerado pela biblioteca gowsdl para comunicação com a SEFAZ
├── .env # Configurações do ambiente (não incluído no repositório)
├── go.mod # Gerenciador de dependências do Go
├── go.sum # Checksum das dependências
└── README.md # Documentação do projeto

## Tecnologias Utilizadas

    •	Go: Linguagem de programação.
    •	gowsdl: Biblioteca para integração SOAP.
    •	go-pkcs12: Biblioteca para manipulação de certificados digitais.
    •	encoding/xml: Para geração de XML.
    •	dotenv: Carregamento de variáveis de ambiente.

## Contribuições

Contribuições são bem-vindas! Por favor, abra uma issue ou envie um pull request para melhorias, correções de bugs ou novas funcionalidades.

## Licença

Este projeto é licenciado sob a licença MIT. Consulte o arquivo LICENSE para mais detalhes.

### Observações

Certifique-se de ajustar o texto e as seções conforme necessário para refletir o estado atual do seu projeto.