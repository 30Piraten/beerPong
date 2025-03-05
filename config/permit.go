package config

import (
	"github.com/permitio/permit-golang/pkg/config"
	"github.com/permitio/permit-golang/pkg/permit"
)

var PermitClient *permit.Client

func PermitInit() {

	apiKey := CheckEnv("API_KEY")

	permitConfig := config.NewConfigBuilder(apiKey).WithPdpUrl("https://cloudpdp.api.permit.io").Build()
	PermitClient = permit.New(permitConfig)
}
