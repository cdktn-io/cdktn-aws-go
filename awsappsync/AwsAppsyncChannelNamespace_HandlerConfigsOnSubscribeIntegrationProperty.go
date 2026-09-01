package awsappsync


// Experimental.
type AwsAppsyncChannelNamespace_HandlerConfigsOnSubscribeIntegrationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_channel_namespace#data_source_name AwsAppsyncChannelNamespace#data_source_name}.
	// Experimental.
	DataSourceName *string `field:"required" json:"dataSourceName" yaml:"dataSourceName"`
	// lambda_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_channel_namespace#lambda_config AwsAppsyncChannelNamespace#lambda_config}
	// Experimental.
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
}

