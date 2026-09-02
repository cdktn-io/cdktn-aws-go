package awsappsync


// Experimental.
type TfDatasource_HttpConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#endpoint TfDatasource#endpoint}.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// authorization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#authorization_config TfDatasource#authorization_config}
	// Experimental.
	AuthorizationConfig *TfDatasource_AuthorizationConfigProperty `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

