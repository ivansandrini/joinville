package main

import (
	"fmt"
	"github.com/fiorix/wsdl2go/soap"
	"github.com/ivansandrini/joinville/joinville"
	"github.com/prometheus/common/log"
)

func main() {
	cli := soap.Client{
		URL:       "http://e-gov.betha.com.br/e-nota-contribuinte-ws/consultarNfsePorRps",
		Namespace: joinville.Namespace,
	}

	service := joinville.NewServicosSoap(&cli)
	resp, err := service.ConsultarLoteRps(&joinville.ConsultarLoteRpsEnvio{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %v", resp)
}
