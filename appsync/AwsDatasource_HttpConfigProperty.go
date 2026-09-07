package appsync


// Experimental.
type AwsDatasource_HttpConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#endpoint AwsDatasource#endpoint}.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// authorization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#authorization_config AwsDatasource#authorization_config}
	// Experimental.
	AuthorizationConfig *AwsDatasource_AuthorizationConfigProperty `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

