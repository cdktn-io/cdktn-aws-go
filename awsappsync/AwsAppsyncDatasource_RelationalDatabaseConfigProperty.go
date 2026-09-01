package awsappsync


// Experimental.
type AwsAppsyncDatasource_RelationalDatabaseConfigProperty struct {
	// http_endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#http_endpoint_config AwsAppsyncDatasource#http_endpoint_config}
	// Experimental.
	HttpEndpointConfig *AwsAppsyncDatasource_HttpEndpointConfigProperty `field:"optional" json:"httpEndpointConfig" yaml:"httpEndpointConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#source_type AwsAppsyncDatasource#source_type}.
	// Experimental.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

