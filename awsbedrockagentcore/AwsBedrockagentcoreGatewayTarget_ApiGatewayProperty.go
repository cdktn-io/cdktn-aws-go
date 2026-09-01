package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGatewayTarget_ApiGatewayProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#rest_api_id AwsBedrockagentcoreGatewayTarget#rest_api_id}.
	// Experimental.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#stage AwsBedrockagentcoreGatewayTarget#stage}.
	// Experimental.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// api_gateway_tool_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#api_gateway_tool_configuration AwsBedrockagentcoreGatewayTarget#api_gateway_tool_configuration}
	// Experimental.
	ApiGatewayToolConfiguration interface{} `field:"optional" json:"apiGatewayToolConfiguration" yaml:"apiGatewayToolConfiguration"`
}

