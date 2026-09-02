package awsecs


// Experimental.
type TfExpressGatewayService_AwsLogsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#log_group TfExpressGatewayService#log_group}.
	// Experimental.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#log_stream_prefix TfExpressGatewayService#log_stream_prefix}.
	// Experimental.
	LogStreamPrefix *string `field:"optional" json:"logStreamPrefix" yaml:"logStreamPrefix"`
}

