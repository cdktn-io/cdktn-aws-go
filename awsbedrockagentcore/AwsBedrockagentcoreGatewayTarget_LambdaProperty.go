package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGatewayTarget_LambdaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#lambda_arn AwsBedrockagentcoreGatewayTarget#lambda_arn}.
	// Experimental.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// tool_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#tool_schema AwsBedrockagentcoreGatewayTarget#tool_schema}
	// Experimental.
	ToolSchema interface{} `field:"optional" json:"toolSchema" yaml:"toolSchema"`
}

