package joinville

import (
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "https://nfemwshomologacao.joinville.sc.gov.br"

// NewServicosSoap creates an initializes a ServicosSoap.
func NewServicosSoap(cli *soap.Client) ServicosSoap {
	return &servicosSoap{cli}
}

// ServicosSoap was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ServicosSoap interface {
	// 9.2.4 Cancelamento NFS-e Esse serviço será executado através
	// da chamada ao método CancelarNfse, passando a mensagem XML
	// como parâmetro com a estrutura definida na tabela que segue.
	CancelarNfse(parameters *CancelarNfseEnvio) (*CancelarNfseEnvioResponse, error)

	// 9.2.6 Consulta de Lote de RPS Esse serviço será executado
	// pelo método ConsultarLoteRps, passando a mensagem XML como
	// parâmetro com a estrutura definida na tabela que segue.
	ConsultarLoteRps(parameters *ConsultarLoteRpsEnvio) (*ConsultarLoteRpsEnvioResponse, error)

	// 9.2.7 Consulta de NFS-e por RPS Esse serviço será executado
	// pelo método ConsultarNfsePorRps, passando a mensagem XML como
	// parâmetro com a estrutura definida na tabela que segue.
	ConsultarNfsePorRps(parameters *ConsultarNfseRpsEnvio) (*ConsultarNfseRpsEnvioResponse, error)

	// 9.2.1 Recepção de Lote de RPS Esse serviço será executado,
	// pelo o método RecepcionarLoteRps, passando a mensagem XML como
	// parâmetro com a estrutura definida na tabela que segue.
	RecepcionarLoteRps(parameters *EnviarLoteRpsEnvio) (*EnviarLoteRpsEnvioResponse, error)
}

// DateTime in WSDL format.
type DateTime string

// ArrayOfMensagemRetorno was auto-generated from WSDL.
type ArrayOfMensagemRetorno struct {
	MensagemRetorno []*MensagemRetorno `xml:"MensagemRetorno,omitempty" json:"MensagemRetorno,omitempty" yaml:"MensagemRetorno,omitempty"`
}

// ArrayOfMensagemRetornoLote was auto-generated from WSDL.
type ArrayOfMensagemRetornoLote struct {
	MensagemRetornoLote []*MensagemRetornoLote `xml:"MensagemRetornoLote,omitempty" json:"MensagemRetornoLote,omitempty" yaml:"MensagemRetornoLote,omitempty"`
}

// CancelarNfseEnvio was auto-generated from WSDL.
type CancelarNfseEnvio struct {
	Pedido *interface{} `xml:"Pedido,omitempty" json:"Pedido,omitempty" yaml:"Pedido,omitempty"`
}

// CancelarNfseEnvioResponse was auto-generated from WSDL.
type CancelarNfseEnvioResponse struct {
	CancelarNfseResposta *CancelarNfseResposta `xml:"CancelarNfseResposta,omitempty" json:"CancelarNfseResposta,omitempty" yaml:"CancelarNfseResposta,omitempty"`
}

// CancelarNfseResposta was auto-generated from WSDL.
type CancelarNfseResposta struct {
	RetCancelamento      *RetCancelamento   `xml:"RetCancelamento,omitempty" json:"RetCancelamento,omitempty" yaml:"RetCancelamento,omitempty"`
	ListaMensagemRetorno []*MensagemRetorno `xml:"ListaMensagemRetorno,omitempty" json:"ListaMensagemRetorno,omitempty" yaml:"ListaMensagemRetorno,omitempty"`
}

// CompNfseTc was auto-generated from WSDL.
type CompNfseTc struct {
	Nfse *Nfse `xml:"Nfse,omitempty" json:"Nfse,omitempty" yaml:"Nfse,omitempty"`
}

// ConfirmacaoCancelamento was auto-generated from WSDL.
type ConfirmacaoCancelamento struct {
	Pedido   *PedidoCancelamento `xml:"Pedido,omitempty" json:"Pedido,omitempty" yaml:"Pedido,omitempty"`
	DataHora DateTime            `xml:"DataHora" json:"DataHora" yaml:"DataHora"`
}

// ConsultarLoteRpsEnvio was auto-generated from WSDL.
type ConsultarLoteRpsEnvio struct {
	Prestador *Prestador `xml:"Prestador,omitempty" json:"Prestador,omitempty" yaml:"Prestador,omitempty"`
	Protocolo *string    `xml:"Protocolo,omitempty" json:"Protocolo,omitempty" yaml:"Protocolo,omitempty"`
}

// ConsultarLoteRpsEnvioResponse was auto-generated from WSDL.
type ConsultarLoteRpsEnvioResponse struct {
	ConsultarLoteRpsResposta *ConsultarLoteRpsResposta `xml:"ConsultarLoteRpsResposta,omitempty" json:"ConsultarLoteRpsResposta,omitempty" yaml:"ConsultarLoteRpsResposta,omitempty"`
}

// ConsultarLoteRpsResposta was auto-generated from WSDL.
type ConsultarLoteRpsResposta struct {
	Situacao                 int                         `xml:"Situacao" json:"Situacao" yaml:"Situacao"`
	ListaMensagemRetorno     *ArrayOfMensagemRetorno     `xml:"ListaMensagemRetorno,omitempty" json:"ListaMensagemRetorno,omitempty" yaml:"ListaMensagemRetorno,omitempty"`
	ListaMensagemRetornoLote *ArrayOfMensagemRetornoLote `xml:"ListaMensagemRetornoLote,omitempty" json:"ListaMensagemRetornoLote,omitempty" yaml:"ListaMensagemRetornoLote,omitempty"`
	ListaNfse                *ListaNfse                  `xml:"ListaNfse,omitempty" json:"ListaNfse,omitempty" yaml:"ListaNfse,omitempty"`
}

// ConsultarNfseRpsEnvio was auto-generated from WSDL.
type ConsultarNfseRpsEnvio struct {
	IdentificacaoRps *IdentificacaoRps `xml:"IdentificacaoRps,omitempty" json:"IdentificacaoRps,omitempty" yaml:"IdentificacaoRps,omitempty"`
	Prestador        *Prestador        `xml:"Prestador,omitempty" json:"Prestador,omitempty" yaml:"Prestador,omitempty"`
}

// ConsultarNfseRpsEnvioResponse was auto-generated from WSDL.
type ConsultarNfseRpsEnvioResponse struct {
	ConsultarNfseRpsResposta *ConsultarNfseRpsResposta `xml:"ConsultarNfseRpsResposta,omitempty" json:"ConsultarNfseRpsResposta,omitempty" yaml:"ConsultarNfseRpsResposta,omitempty"`
}

// ConsultarNfseRpsResposta was auto-generated from WSDL.
type ConsultarNfseRpsResposta struct {
	CompNfse             *CompNfseTc             `xml:"CompNfse,omitempty" json:"CompNfse,omitempty" yaml:"CompNfse,omitempty"`
	ListaMensagemRetorno *ArrayOfMensagemRetorno `xml:"ListaMensagemRetorno,omitempty" json:"ListaMensagemRetorno,omitempty" yaml:"ListaMensagemRetorno,omitempty"`
}

// Contato was auto-generated from WSDL.
type Contato struct {
	Telefone *string `xml:"Telefone,omitempty" json:"Telefone,omitempty" yaml:"Telefone,omitempty"`
	Email    *string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`
}

// CpfCnpj was auto-generated from WSDL.
type CpfCnpj struct {
	Cpf  *string `xml:"Cpf,omitempty" json:"Cpf,omitempty" yaml:"Cpf,omitempty"`
	Cnpj *string `xml:"Cnpj,omitempty" json:"Cnpj,omitempty" yaml:"Cnpj,omitempty"`
}

// DadosConstrucaoCivil was auto-generated from WSDL.
type DadosConstrucaoCivil struct {
	CodigoObra *string `xml:"CodigoObra,omitempty" json:"CodigoObra,omitempty" yaml:"CodigoObra,omitempty"`
	Art        *string `xml:"Art,omitempty" json:"Art,omitempty" yaml:"Art,omitempty"`
}

// DadosDeducao was auto-generated from WSDL.
type DadosDeducao struct {
	TipoDeducao                   *string                        `xml:"TipoDeducao,omitempty" json:"TipoDeducao,omitempty" yaml:"TipoDeducao,omitempty"`
	DescricaoDeducao              *string                        `xml:"DescricaoDeducao,omitempty" json:"DescricaoDeducao,omitempty" yaml:"DescricaoDeducao,omitempty"`
	IdentificacaoDocumentoDeducao *IdentificacaoDocumentoDeducao `xml:"IdentificacaoDocumentoDeducao,omitempty" json:"IdentificacaoDocumentoDeducao,omitempty" yaml:"IdentificacaoDocumentoDeducao,omitempty"`
	DadosFornecedor               *DadosFornecedor               `xml:"DadosFornecedor,omitempty" json:"DadosFornecedor,omitempty" yaml:"DadosFornecedor,omitempty"`
	DateEmissao                   DateTime                       `xml:"DateEmissao" json:"DateEmissao" yaml:"DateEmissao"`
	ValorDedutivel                float64                        `xml:"ValorDedutivel" json:"ValorDedutivel" yaml:"ValorDedutivel"`
	ValorUtilizadoDeducao         float64                        `xml:"ValorUtilizadoDeducao" json:"ValorUtilizadoDeducao" yaml:"ValorUtilizadoDeducao"`
}

// DadosEvento was auto-generated from WSDL.
type DadosEvento struct {
	IdentificacaoEvento *string `xml:"IdentificacaoEvento,omitempty" json:"IdentificacaoEvento,omitempty" yaml:"IdentificacaoEvento,omitempty"`
	DescricaoEvento     *string `xml:"DescricaoEvento,omitempty" json:"DescricaoEvento,omitempty" yaml:"DescricaoEvento,omitempty"`
}

// DadosFornecedor was auto-generated from WSDL.
type DadosFornecedor struct {
}

// DadosIntermediario was auto-generated from WSDL.
type DadosIntermediario struct {
	IdentificacaoIntermediario *IdentificacaoPessoaEmpresa `xml:"IdentificacaoIntermediario,omitempty" json:"IdentificacaoIntermediario,omitempty" yaml:"IdentificacaoIntermediario,omitempty"`
	RazaoSocial                *string                     `xml:"RazaoSocial,omitempty" json:"RazaoSocial,omitempty" yaml:"RazaoSocial,omitempty"`
	CodigoMunicipio            *string                     `xml:"CodigoMunicipio,omitempty" json:"CodigoMunicipio,omitempty" yaml:"CodigoMunicipio,omitempty"`
}

// DadosPrestador was auto-generated from WSDL.
type DadosPrestador struct {
	RazaoSocial  *string   `xml:"RazaoSocial,omitempty" json:"RazaoSocial,omitempty" yaml:"RazaoSocial,omitempty"`
	NomeFantasia *string   `xml:"NomeFantasia,omitempty" json:"NomeFantasia,omitempty" yaml:"NomeFantasia,omitempty"`
	Endereco     *Endereco `xml:"Endereco,omitempty" json:"Endereco,omitempty" yaml:"Endereco,omitempty"`
	Contato      *Contato  `xml:"Contato,omitempty" json:"Contato,omitempty" yaml:"Contato,omitempty"`
}

// DeclaracaoPrestacaoServico was auto-generated from WSDL.
type DeclaracaoPrestacaoServico struct {
	InfDeclaracaoPrestacaoServico *InfDeclaracaoPrestacaoServico `xml:"InfDeclaracaoPrestacaoServico,omitempty" json:"InfDeclaracaoPrestacaoServico,omitempty" yaml:"InfDeclaracaoPrestacaoServico,omitempty"`
	Signature                     *string                        `xml:"Signature,omitempty" json:"Signature,omitempty" yaml:"Signature,omitempty"`
}

// Endereco was auto-generated from WSDL.
type Endereco struct {
	Endereco        *string `xml:"Endereco,omitempty" json:"Endereco,omitempty" yaml:"Endereco,omitempty"`
	Numero          *string `xml:"Numero,omitempty" json:"Numero,omitempty" yaml:"Numero,omitempty"`
	Complemento     *string `xml:"Complemento,omitempty" json:"Complemento,omitempty" yaml:"Complemento,omitempty"`
	Bairro          *string `xml:"Bairro,omitempty" json:"Bairro,omitempty" yaml:"Bairro,omitempty"`
	CodigoMunicipio int     `xml:"CodigoMunicipio" json:"CodigoMunicipio" yaml:"CodigoMunicipio"`
	Uf              *string `xml:"Uf,omitempty" json:"Uf,omitempty" yaml:"Uf,omitempty"`
	Cep             *string `xml:"Cep,omitempty" json:"Cep,omitempty" yaml:"Cep,omitempty"`
}

// EnderecoExterior was auto-generated from WSDL.
type EnderecoExterior struct {
	CodigoPais               *string `xml:"CodigoPais,omitempty" json:"CodigoPais,omitempty" yaml:"CodigoPais,omitempty"`
	EnderecoCompletoExterior *string `xml:"EnderecoCompletoExterior,omitempty" json:"EnderecoCompletoExterior,omitempty" yaml:"EnderecoCompletoExterior,omitempty"`
}

// EnviarLoteRpsEnvio was auto-generated from WSDL.
type EnviarLoteRpsEnvio struct {
	LoteRps   *interface{} `xml:"LoteRps,omitempty" json:"LoteRps,omitempty" yaml:"LoteRps,omitempty"`
	Signature *interface{} `xml:"Signature,omitempty" json:"Signature,omitempty" yaml:"Signature,omitempty"`
}

// EnviarLoteRpsEnvioResponse was auto-generated from WSDL.
type EnviarLoteRpsEnvioResponse struct {
	EnviarLoteRpsResposta *EnviarLoteRpsResposta `xml:"EnviarLoteRpsResposta,omitempty" json:"EnviarLoteRpsResposta,omitempty" yaml:"EnviarLoteRpsResposta,omitempty"`
}

// EnviarLoteRpsResposta was auto-generated from WSDL.
type EnviarLoteRpsResposta struct {
	NumeroLote           *int                    `xml:"NumeroLote,omitempty" json:"NumeroLote,omitempty" yaml:"NumeroLote,omitempty"`
	DataRecebimento      *DateTime               `xml:"DataRecebimento,omitempty" json:"DataRecebimento,omitempty" yaml:"DataRecebimento,omitempty"`
	Protocolo            *string                 `xml:"Protocolo,omitempty" json:"Protocolo,omitempty" yaml:"Protocolo,omitempty"`
	ListaMensagemRetorno *ArrayOfMensagemRetorno `xml:"ListaMensagemRetorno,omitempty" json:"ListaMensagemRetorno,omitempty" yaml:"ListaMensagemRetorno,omitempty"`
}

// IdentificacaoDocumentoDeducao was auto-generated from WSDL.
type IdentificacaoDocumentoDeducao struct {
	IdentificacaoNfse     *string `xml:"IdentificacaoNfse,omitempty" json:"IdentificacaoNfse,omitempty" yaml:"IdentificacaoNfse,omitempty"`
	IdentificacaoNfe      *string `xml:"IdentificacaoNfe,omitempty" json:"IdentificacaoNfe,omitempty" yaml:"IdentificacaoNfe,omitempty"`
	OutroDocumentoDeducao *string `xml:"OutroDocumentoDeducao,omitempty" json:"OutroDocumentoDeducao,omitempty" yaml:"OutroDocumentoDeducao,omitempty"`
}

// IdentificacaoNfse was auto-generated from WSDL.
type IdentificacaoNfse struct {
	Numero             int      `xml:"Numero" json:"Numero" yaml:"Numero"`
	CpfCnpj            *CpfCnpj `xml:"CpfCnpj,omitempty" json:"CpfCnpj,omitempty" yaml:"CpfCnpj,omitempty"`
	InscricaoMunicipal *string  `xml:"InscricaoMunicipal,omitempty" json:"InscricaoMunicipal,omitempty" yaml:"InscricaoMunicipal,omitempty"`
	CodigoMunicipio    int      `xml:"CodigoMunicipio" json:"CodigoMunicipio" yaml:"CodigoMunicipio"`
}

// IdentificacaoOrgaoGerador was auto-generated from WSDL.
type IdentificacaoOrgaoGerador struct {
	CodigoMunicipio int     `xml:"CodigoMunicipio" json:"CodigoMunicipio" yaml:"CodigoMunicipio"`
	Uf              *string `xml:"Uf,omitempty" json:"Uf,omitempty" yaml:"Uf,omitempty"`
}

// IdentificacaoPessoaEmpresa was auto-generated from WSDL.
type IdentificacaoPessoaEmpresa struct {
	CpfCnpj            *CpfCnpj `xml:"CpfCnpj,omitempty" json:"CpfCnpj,omitempty" yaml:"CpfCnpj,omitempty"`
	InscricaoMunicipal *string  `xml:"InscricaoMunicipal,omitempty" json:"InscricaoMunicipal,omitempty" yaml:"InscricaoMunicipal,omitempty"`
}

// IdentificacaoRps was auto-generated from WSDL.
type IdentificacaoRps struct {
	Numero int     `xml:"Numero" json:"Numero" yaml:"Numero"`
	Serie  *string `xml:"Serie,omitempty" json:"Serie,omitempty" yaml:"Serie,omitempty"`
	Tipo   int     `xml:"Tipo" json:"Tipo" yaml:"Tipo"`
}

// IdentificacaoTomador was auto-generated from WSDL.
type IdentificacaoTomador struct {
	CpfCnpj            *CpfCnpj `xml:"CpfCnpj,omitempty" json:"CpfCnpj,omitempty" yaml:"CpfCnpj,omitempty"`
	InscricaoMunicipal *string  `xml:"InscricaoMunicipal,omitempty" json:"InscricaoMunicipal,omitempty" yaml:"InscricaoMunicipal,omitempty"`
}

// InfDeclaracaoPrestacaoServico was auto-generated from WSDL.
type InfDeclaracaoPrestacaoServico struct {
	Rps                       *InfRps                     `xml:"Rps,omitempty" json:"Rps,omitempty" yaml:"Rps,omitempty"`
	Competencia               DateTime                    `xml:"Competencia" json:"Competencia" yaml:"Competencia"`
	Servico                   *Servico                    `xml:"Servico,omitempty" json:"Servico,omitempty" yaml:"Servico,omitempty"`
	Prestador                 *IdentificacaoPessoaEmpresa `xml:"Prestador,omitempty" json:"Prestador,omitempty" yaml:"Prestador,omitempty"`
	TomadorServico            *TomadorWs                  `xml:"TomadorServico,omitempty" json:"TomadorServico,omitempty" yaml:"TomadorServico,omitempty"`
	Intermediario             *DadosIntermediario         `xml:"Intermediario,omitempty" json:"Intermediario,omitempty" yaml:"Intermediario,omitempty"`
	ConstrucaoCivil           *DadosConstrucaoCivil       `xml:"ConstrucaoCivil,omitempty" json:"ConstrucaoCivil,omitempty" yaml:"ConstrucaoCivil,omitempty"`
	RegimeEspecialTributacao  *string                     `xml:"RegimeEspecialTributacao,omitempty" json:"RegimeEspecialTributacao,omitempty" yaml:"RegimeEspecialTributacao,omitempty"`
	OptanteSimplesNacional    int                         `xml:"OptanteSimplesNacional" json:"OptanteSimplesNacional" yaml:"OptanteSimplesNacional"`
	IncentivoFiscal           int                         `xml:"IncentivoFiscal" json:"IncentivoFiscal" yaml:"IncentivoFiscal"`
	Evento                    *DadosEvento                `xml:"Evento,omitempty" json:"Evento,omitempty" yaml:"Evento,omitempty"`
	InformacoesComplementares *string                     `xml:"InformacoesComplementares,omitempty" json:"InformacoesComplementares,omitempty" yaml:"InformacoesComplementares,omitempty"`
	Deducao                   *DadosDeducao               `xml:"Deducao,omitempty" json:"Deducao,omitempty" yaml:"Deducao,omitempty"`
	CodigoCei                 *string                     `xml:"CodigoCei,omitempty" json:"CodigoCei,omitempty" yaml:"CodigoCei,omitempty"`
	Id                        string                      `xml:"Id,attr,omitempty" json:"Id,attr,omitempty" yaml:"Id,attr,omitempty"`
}

// InfNfse was auto-generated from WSDL.
type InfNfse struct {
	Numero                            int                         `xml:"Numero" json:"Numero" yaml:"Numero"`
	NumeroRps                         int                         `xml:"NumeroRps" json:"NumeroRps" yaml:"NumeroRps"`
	CodigoVerificacao                 *string                     `xml:"CodigoVerificacao,omitempty" json:"CodigoVerificacao,omitempty" yaml:"CodigoVerificacao,omitempty"`
	DataEmissao                       DateTime                    `xml:"DataEmissao" json:"DataEmissao" yaml:"DataEmissao"`
	NfseSubstituida                   int                         `xml:"NfseSubstituida" json:"NfseSubstituida" yaml:"NfseSubstituida"`
	OutrasInformacoes                 *string                     `xml:"OutrasInformacoes,omitempty" json:"OutrasInformacoes,omitempty" yaml:"OutrasInformacoes,omitempty"`
	ValoresNfse                       *ValoresNfse                `xml:"ValoresNfse,omitempty" json:"ValoresNfse,omitempty" yaml:"ValoresNfse,omitempty"`
	DescricaoCodigoTributacaoMunicPio *string                     `xml:"DescricaoCodigoTributacaoMunicípio,omitempty" json:"DescricaoCodigoTributacaoMunicípio,omitempty" yaml:"DescricaoCodigoTributacaoMunicípio,omitempty"`
	ValorCredito                      float64                     `xml:"ValorCredito" json:"ValorCredito" yaml:"ValorCredito"`
	PrestadorServico                  *DadosPrestador             `xml:"PrestadorServico,omitempty" json:"PrestadorServico,omitempty" yaml:"PrestadorServico,omitempty"`
	OrgaoGerador                      *IdentificacaoOrgaoGerador  `xml:"OrgaoGerador,omitempty" json:"OrgaoGerador,omitempty" yaml:"OrgaoGerador,omitempty"`
	DeclaracaoPrestacaoServico        *DeclaracaoPrestacaoServico `xml:"DeclaracaoPrestacaoServico,omitempty" json:"DeclaracaoPrestacaoServico,omitempty" yaml:"DeclaracaoPrestacaoServico,omitempty"`
	Id                                string                      `xml:"Id,attr,omitempty" json:"Id,attr,omitempty" yaml:"Id,attr,omitempty"`
}

// InfPedidoCancelamento was auto-generated from WSDL.
type InfPedidoCancelamento struct {
	IdentificacaoNfse  *IdentificacaoNfse `xml:"IdentificacaoNfse,omitempty" json:"IdentificacaoNfse,omitempty" yaml:"IdentificacaoNfse,omitempty"`
	CodigoCancelamento int                `xml:"CodigoCancelamento" json:"CodigoCancelamento" yaml:"CodigoCancelamento"`
	Id                 string             `xml:"Id,attr,omitempty" json:"Id,attr,omitempty" yaml:"Id,attr,omitempty"`
}

// InfRps was auto-generated from WSDL.
type InfRps struct {
	IdentificacaoRps *IdentificacaoRps `xml:"IdentificacaoRps,omitempty" json:"IdentificacaoRps,omitempty" yaml:"IdentificacaoRps,omitempty"`
	DataEmissao      DateTime          `xml:"DataEmissao" json:"DataEmissao" yaml:"DataEmissao"`
	Status           int               `xml:"Status" json:"Status" yaml:"Status"`
	RpsSubstituido   *IdentificacaoRps `xml:"RpsSubstituido,omitempty" json:"RpsSubstituido,omitempty" yaml:"RpsSubstituido,omitempty"`
	Id               string            `xml:"Id,attr,omitempty" json:"Id,attr,omitempty" yaml:"Id,attr,omitempty"`
}

// ListaNfse was auto-generated from WSDL.
type ListaNfse struct {
	CompNfse                   []*CompNfseTc      `xml:"CompNfse,omitempty" json:"CompNfse,omitempty" yaml:"CompNfse,omitempty"`
	ListaMensagemAlertaRetorno []*MensagemRetorno `xml:"ListaMensagemAlertaRetorno,omitempty" json:"ListaMensagemAlertaRetorno,omitempty" yaml:"ListaMensagemAlertaRetorno,omitempty"`
}

// MensagemRetorno was auto-generated from WSDL.
type MensagemRetorno struct {
	Codigo   *string `xml:"Codigo,omitempty" json:"Codigo,omitempty" yaml:"Codigo,omitempty"`
	Mensagem *string `xml:"Mensagem,omitempty" json:"Mensagem,omitempty" yaml:"Mensagem,omitempty"`
	Correcao *string `xml:"Correcao,omitempty" json:"Correcao,omitempty" yaml:"Correcao,omitempty"`
}

// MensagemRetornoLote was auto-generated from WSDL.
type MensagemRetornoLote struct {
	IdentificacaoRps *IdentificacaoRps `xml:"IdentificacaoRps,omitempty" json:"IdentificacaoRps,omitempty" yaml:"IdentificacaoRps,omitempty"`
	Codigo           *string           `xml:"Codigo,omitempty" json:"Codigo,omitempty" yaml:"Codigo,omitempty"`
	Mensagem         *string           `xml:"Mensagem,omitempty" json:"Mensagem,omitempty" yaml:"Mensagem,omitempty"`
}

// Nfse was auto-generated from WSDL.
type Nfse struct {
	InfNfse   *InfNfse `xml:"InfNfse,omitempty" json:"InfNfse,omitempty" yaml:"InfNfse,omitempty"`
	Signature *string  `xml:"Signature,omitempty" json:"Signature,omitempty" yaml:"Signature,omitempty"`
	Versao    *string  `xml:"versao,omitempty" json:"versao,omitempty" yaml:"versao,omitempty"`
}

// NfseCancelamento was auto-generated from WSDL.
type NfseCancelamento struct {
	Confirmacao *ConfirmacaoCancelamento `xml:"Confirmacao,omitempty" json:"Confirmacao,omitempty" yaml:"Confirmacao,omitempty"`
	Versao      *string                  `xml:"versao,omitempty" json:"versao,omitempty" yaml:"versao,omitempty"`
}

// PedidoCancelamento was auto-generated from WSDL.
type PedidoCancelamento struct {
	InfPedidoCancelamento *InfPedidoCancelamento `xml:"InfPedidoCancelamento,omitempty" json:"InfPedidoCancelamento,omitempty" yaml:"InfPedidoCancelamento,omitempty"`
	Signature             *string                `xml:"Signature,omitempty" json:"Signature,omitempty" yaml:"Signature,omitempty"`
}

// Prestador was auto-generated from WSDL.
type Prestador struct {
	CpfCnpj            *CpfCnpj `xml:"CpfCnpj,omitempty" json:"CpfCnpj,omitempty" yaml:"CpfCnpj,omitempty"`
	InscricaoMunicipal *string  `xml:"InscricaoMunicipal,omitempty" json:"InscricaoMunicipal,omitempty" yaml:"InscricaoMunicipal,omitempty"`
	TypeAttrXSI        string   `xml:"xsi:type,attr,omitempty"`
	TypeNamespace      string   `xml:"xmlns:objtype,attr,omitempty"`

	OverrideTypeAttrXSI   *string `xml:"-"`
	OverrideTypeNamespace *string `xml:"-"`
}

// SetXMLType was auto-generated from WSDL.
func (t *Prestador) SetXMLType() {
	if t.OverrideTypeAttrXSI != nil {
		t.TypeAttrXSI = *t.OverrideTypeAttrXSI
	} else {
		t.TypeAttrXSI = "objtype:Prestador"
	}
	if t.OverrideTypeNamespace != nil {
		t.TypeNamespace = *t.OverrideTypeNamespace
	} else {
		t.TypeNamespace = "https://nfemwshomologacao.joinville.sc.gov.br"
	}
}

// RetCancelamento was auto-generated from WSDL.
type RetCancelamento struct {
	NfseCancelamento *NfseCancelamento `xml:"NfseCancelamento,omitempty" json:"NfseCancelamento,omitempty" yaml:"NfseCancelamento,omitempty"`
}

// Servico was auto-generated from WSDL.
type Servico struct {
	Valores                 *Valores `xml:"Valores,omitempty" json:"Valores,omitempty" yaml:"Valores,omitempty"`
	IssRetido               int      `xml:"IssRetido" json:"IssRetido" yaml:"IssRetido"`
	ResponsavelRetencao     *string  `xml:"ResponsavelRetencao,omitempty" json:"ResponsavelRetencao,omitempty" yaml:"ResponsavelRetencao,omitempty"`
	ItemListaServico        *string  `xml:"ItemListaServico,omitempty" json:"ItemListaServico,omitempty" yaml:"ItemListaServico,omitempty"`
	CodigoCnae              *string  `xml:"CodigoCnae,omitempty" json:"CodigoCnae,omitempty" yaml:"CodigoCnae,omitempty"`
	Discriminacao           *string  `xml:"Discriminacao,omitempty" json:"Discriminacao,omitempty" yaml:"Discriminacao,omitempty"`
	CodigoMunicipio         int      `xml:"CodigoMunicipio" json:"CodigoMunicipio" yaml:"CodigoMunicipio"`
	CodigoPais              *string  `xml:"CodigoPais,omitempty" json:"CodigoPais,omitempty" yaml:"CodigoPais,omitempty"`
	ExigibilidadeISS        *string  `xml:"ExigibilidadeISS,omitempty" json:"ExigibilidadeISS,omitempty" yaml:"ExigibilidadeISS,omitempty"`
	IdentifNaoExigibilidade int      `xml:"IdentifNaoExigibilidade" json:"IdentifNaoExigibilidade" yaml:"IdentifNaoExigibilidade"`
	MunicipioIncidencia     int      `xml:"MunicipioIncidencia" json:"MunicipioIncidencia" yaml:"MunicipioIncidencia"`
	NumeroProcesso          *int     `xml:"NumeroProcesso,omitempty" json:"NumeroProcesso,omitempty" yaml:"NumeroProcesso,omitempty"`
}

// TomadorWs was auto-generated from WSDL.
type TomadorWs struct {
	IdentificacaoTomador *IdentificacaoTomador `xml:"IdentificacaoTomador,omitempty" json:"IdentificacaoTomador,omitempty" yaml:"IdentificacaoTomador,omitempty"`
	RazaoSocial          *string               `xml:"RazaoSocial,omitempty" json:"RazaoSocial,omitempty" yaml:"RazaoSocial,omitempty"`
	NifTomador           *string               `xml:"NifTomador,omitempty" json:"NifTomador,omitempty" yaml:"NifTomador,omitempty"`
	Endereco             *Endereco             `xml:"Endereco,omitempty" json:"Endereco,omitempty" yaml:"Endereco,omitempty"`
	EnderecoExterior     *EnderecoExterior     `xml:"EnderecoExterior,omitempty" json:"EnderecoExterior,omitempty" yaml:"EnderecoExterior,omitempty"`
	Contato              *Contato              `xml:"Contato,omitempty" json:"Contato,omitempty" yaml:"Contato,omitempty"`
}

// Valores was auto-generated from WSDL.
type Valores struct {
	ValorServicos          float64  `xml:"ValorServicos" json:"ValorServicos" yaml:"ValorServicos"`
	ValorDeducoes          *float64 `xml:"ValorDeducoes,omitempty" json:"ValorDeducoes,omitempty" yaml:"ValorDeducoes,omitempty"`
	ValorPis               *float64 `xml:"ValorPis,omitempty" json:"ValorPis,omitempty" yaml:"ValorPis,omitempty"`
	ValorCofins            *float64 `xml:"ValorCofins,omitempty" json:"ValorCofins,omitempty" yaml:"ValorCofins,omitempty"`
	ValorInss              *float64 `xml:"ValorInss,omitempty" json:"ValorInss,omitempty" yaml:"ValorInss,omitempty"`
	ValorIr                *float64 `xml:"ValorIr,omitempty" json:"ValorIr,omitempty" yaml:"ValorIr,omitempty"`
	ValorCsll              *float64 `xml:"ValorCsll,omitempty" json:"ValorCsll,omitempty" yaml:"ValorCsll,omitempty"`
	AliquotaPis            *float64 `xml:"AliquotaPis,omitempty" json:"AliquotaPis,omitempty" yaml:"AliquotaPis,omitempty"`
	AliquotaCofins         *float64 `xml:"AliquotaCofins,omitempty" json:"AliquotaCofins,omitempty" yaml:"AliquotaCofins,omitempty"`
	AliquotaInss           *float64 `xml:"AliquotaInss,omitempty" json:"AliquotaInss,omitempty" yaml:"AliquotaInss,omitempty"`
	AliquotaIr             *float64 `xml:"AliquotaIr,omitempty" json:"AliquotaIr,omitempty" yaml:"AliquotaIr,omitempty"`
	AliquotaCsll           *float64 `xml:"AliquotaCsll,omitempty" json:"AliquotaCsll,omitempty" yaml:"AliquotaCsll,omitempty"`
	ValTotTributos         *float64 `xml:"ValTotTributos,omitempty" json:"ValTotTributos,omitempty" yaml:"ValTotTributos,omitempty"`
	ValorIss               *float64 `xml:"ValorIss,omitempty" json:"ValorIss,omitempty" yaml:"ValorIss,omitempty"`
	Aliquota               *float64 `xml:"Aliquota,omitempty" json:"Aliquota,omitempty" yaml:"Aliquota,omitempty"`
	DescontoIncondicionado *float64 `xml:"DescontoIncondicionado,omitempty" json:"DescontoIncondicionado,omitempty" yaml:"DescontoIncondicionado,omitempty"`
	DescontoCondicionado   *float64 `xml:"DescontoCondicionado,omitempty" json:"DescontoCondicionado,omitempty" yaml:"DescontoCondicionado,omitempty"`
}

// ValoresNfse was auto-generated from WSDL.
type ValoresNfse struct {
	ValorIss         float64 `xml:"ValorIss" json:"ValorIss" yaml:"ValorIss"`
	Aliquota         float64 `xml:"Aliquota" json:"Aliquota" yaml:"Aliquota"`
	ValorLiquidoNfse float64 `xml:"ValorLiquidoNfse" json:"ValorLiquidoNfse" yaml:"ValorLiquidoNfse"`
	BaseCalculo      float64 `xml:"BaseCalculo" json:"BaseCalculo" yaml:"BaseCalculo"`
	AliquotaIr       float64 `xml:"AliquotaIr" json:"AliquotaIr" yaml:"AliquotaIr"`
	AliquotaInss     float64 `xml:"AliquotaInss" json:"AliquotaInss" yaml:"AliquotaInss"`
	AliquotaPis      float64 `xml:"AliquotaPis" json:"AliquotaPis" yaml:"AliquotaPis"`
	AliquotaCofins   float64 `xml:"AliquotaCofins" json:"AliquotaCofins" yaml:"AliquotaCofins"`
	AliquotaCsll     float64 `xml:"AliquotaCsll" json:"AliquotaCsll" yaml:"AliquotaCsll"`
	ValorIr          float64 `xml:"ValorIr" json:"ValorIr" yaml:"ValorIr"`
	ValorInss        float64 `xml:"ValorInss" json:"ValorInss" yaml:"ValorInss"`
	ValorPis         float64 `xml:"ValorPis" json:"ValorPis" yaml:"ValorPis"`
	ValorCofins      float64 `xml:"ValorCofins" json:"ValorCofins" yaml:"ValorCofins"`
	ValorCsll        float64 `xml:"ValorCsll" json:"ValorCsll" yaml:"ValorCsll"`
}

// Operation wrapper for CancelarNfse.
// OperationCancelarNfseEnvioSoapIn was auto-generated from WSDL.
type OperationCancelarNfseEnvioSoapIn struct {
	Parameters *CancelarNfseEnvio `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for CancelarNfse.
// OperationCancelarNfseEnvioSoapOut was auto-generated from WSDL.
type OperationCancelarNfseEnvioSoapOut struct {
	Parameters *CancelarNfseEnvioResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for ConsultarLoteRps.
// OperationConsultarLoteRpsEnvioSoapIn was auto-generated from
// WSDL.
type OperationConsultarLoteRpsEnvioSoapIn struct {
	Parameters *ConsultarLoteRpsEnvio `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for ConsultarLoteRps.
// OperationConsultarLoteRpsEnvioSoapOut was auto-generated from
// WSDL.
type OperationConsultarLoteRpsEnvioSoapOut struct {
	Parameters *ConsultarLoteRpsEnvioResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for ConsultarNfsePorRps.
// OperationConsultarNfseRpsEnvioSoapIn was auto-generated from
// WSDL.
type OperationConsultarNfseRpsEnvioSoapIn struct {
	Parameters *ConsultarNfseRpsEnvio `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for ConsultarNfsePorRps.
// OperationConsultarNfseRpsEnvioSoapOut was auto-generated from
// WSDL.
type OperationConsultarNfseRpsEnvioSoapOut struct {
	Parameters *ConsultarNfseRpsEnvioResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for RecepcionarLoteRps.
// OperationEnviarLoteRpsEnvioSoapIn was auto-generated from WSDL.
type OperationEnviarLoteRpsEnvioSoapIn struct {
	Parameters *EnviarLoteRpsEnvio `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for RecepcionarLoteRps.
// OperationEnviarLoteRpsEnvioSoapOut was auto-generated from WSDL.
type OperationEnviarLoteRpsEnvioSoapOut struct {
	Parameters *EnviarLoteRpsEnvioResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// servicosSoap implements the ServicosSoap interface.
type servicosSoap struct {
	cli *soap.Client
}

// 9.2.4 Cancelamento NFS-e Esse serviço será executado através
// da chamada ao método CancelarNfse, passando a mensagem XML
// como parâmetro com a estrutura definida na tabela que segue.
func (p *servicosSoap) CancelarNfse(parameters *CancelarNfseEnvio) (*CancelarNfseEnvioResponse, error) {
	α := struct {
		OperationCancelarNfseEnvioSoapIn `xml:"tns:CancelarNfse"`
	}{
		OperationCancelarNfseEnvioSoapIn{
			parameters,
		},
	}

	γ := struct {
		OperationCancelarNfseEnvioSoapOut `xml:"CancelarNfseResponse"`
	}{}
	if err := p.cli.RoundTripSoap12("https://nfemwshomologacao.joinville.sc.gov.br/CancelarNfseEnvio", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// 9.2.6 Consulta de Lote de RPS Esse serviço será executado
// pelo método ConsultarLoteRps, passando a mensagem XML como
// parâmetro com a estrutura definida na tabela que segue.
func (p *servicosSoap) ConsultarLoteRps(parameters *ConsultarLoteRpsEnvio) (*ConsultarLoteRpsEnvioResponse, error) {
	α := struct {
		OperationConsultarLoteRpsEnvioSoapIn `xml:"tns:ConsultarLoteRps"`
	}{
		OperationConsultarLoteRpsEnvioSoapIn{
			parameters,
		},
	}

	γ := struct {
		OperationConsultarLoteRpsEnvioSoapOut `xml:"ConsultarLoteRpsResponse"`
	}{}
	if err := p.cli.RoundTripSoap12("https://nfemwshomologacao.joinville.sc.gov.br/ConsultarLoteRpsEnvio", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// 9.2.7 Consulta de NFS-e por RPS Esse serviço será executado
// pelo método ConsultarNfsePorRps, passando a mensagem XML como
// parâmetro com a estrutura definida na tabela que segue.
func (p *servicosSoap) ConsultarNfsePorRps(parameters *ConsultarNfseRpsEnvio) (*ConsultarNfseRpsEnvioResponse, error) {
	α := struct {
		OperationConsultarNfseRpsEnvioSoapIn `xml:"tns:ConsultarNfsePorRps"`
	}{
		OperationConsultarNfseRpsEnvioSoapIn{
			parameters,
		},
	}

	γ := struct {
		OperationConsultarNfseRpsEnvioSoapOut `xml:"ConsultarNfsePorRpsResponse"`
	}{}
	if err := p.cli.RoundTripSoap12("https://nfemwshomologacao.joinville.sc.gov.br/ConsultarNfseRpsEnvio", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// 9.2.1 Recepção de Lote de RPS Esse serviço será executado,
// pelo o método RecepcionarLoteRps, passando a mensagem XML como
// parâmetro com a estrutura definida na tabela que segue.
func (p *servicosSoap) RecepcionarLoteRps(parameters *EnviarLoteRpsEnvio) (*EnviarLoteRpsEnvioResponse, error) {
	α := struct {
		OperationEnviarLoteRpsEnvioSoapIn `xml:"tns:RecepcionarLoteRps"`
	}{
		OperationEnviarLoteRpsEnvioSoapIn{
			parameters,
		},
	}

	γ := struct {
		OperationEnviarLoteRpsEnvioSoapOut `xml:"RecepcionarLoteRpsResponse"`
	}{}
	if err := p.cli.RoundTripSoap12("https://nfemwshomologacao.joinville.sc.gov.br/EnviarLoteRpsEnvio", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}
