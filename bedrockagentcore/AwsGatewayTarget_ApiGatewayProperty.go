package bedrockagentcore


// Experimental.
type AwsGatewayTarget_ApiGatewayProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#rest_api_id AwsGatewayTarget#rest_api_id}.
	// Experimental.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#stage AwsGatewayTarget#stage}.
	// Experimental.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// api_gateway_tool_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#api_gateway_tool_configuration AwsGatewayTarget#api_gateway_tool_configuration}
	// Experimental.
	ApiGatewayToolConfiguration interface{} `field:"optional" json:"apiGatewayToolConfiguration" yaml:"apiGatewayToolConfiguration"`
}

