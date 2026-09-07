package bedrockagentcore


// Experimental.
type AwsGateway_InterceptorConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#interception_points AwsGateway#interception_points}.
	// Experimental.
	InterceptionPoints *[]*string `field:"required" json:"interceptionPoints" yaml:"interceptionPoints"`
	// input_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#input_configuration AwsGateway#input_configuration}
	// Experimental.
	InputConfiguration interface{} `field:"optional" json:"inputConfiguration" yaml:"inputConfiguration"`
	// interceptor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#interceptor AwsGateway#interceptor}
	// Experimental.
	Interceptor interface{} `field:"optional" json:"interceptor" yaml:"interceptor"`
}

