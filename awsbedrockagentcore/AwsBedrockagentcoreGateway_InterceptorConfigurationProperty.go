package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGateway_InterceptorConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#interception_points AwsBedrockagentcoreGateway#interception_points}.
	// Experimental.
	InterceptionPoints *[]*string `field:"required" json:"interceptionPoints" yaml:"interceptionPoints"`
	// input_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#input_configuration AwsBedrockagentcoreGateway#input_configuration}
	// Experimental.
	InputConfiguration interface{} `field:"optional" json:"inputConfiguration" yaml:"inputConfiguration"`
	// interceptor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#interceptor AwsBedrockagentcoreGateway#interceptor}
	// Experimental.
	Interceptor interface{} `field:"optional" json:"interceptor" yaml:"interceptor"`
}

