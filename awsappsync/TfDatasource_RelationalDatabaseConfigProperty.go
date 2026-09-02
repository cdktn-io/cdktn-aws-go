package awsappsync


// Experimental.
type TfDatasource_RelationalDatabaseConfigProperty struct {
	// http_endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#http_endpoint_config TfDatasource#http_endpoint_config}
	// Experimental.
	HttpEndpointConfig *TfDatasource_HttpEndpointConfigProperty `field:"optional" json:"httpEndpointConfig" yaml:"httpEndpointConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#source_type TfDatasource#source_type}.
	// Experimental.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

