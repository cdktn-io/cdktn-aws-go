package awsbedrockagentcore


// Experimental.
type TfGateway_InterceptorConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#interception_points TfGateway#interception_points}.
	// Experimental.
	InterceptionPoints *[]*string `field:"required" json:"interceptionPoints" yaml:"interceptionPoints"`
	// input_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#input_configuration TfGateway#input_configuration}
	// Experimental.
	InputConfiguration interface{} `field:"optional" json:"inputConfiguration" yaml:"inputConfiguration"`
	// interceptor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway#interceptor TfGateway#interceptor}
	// Experimental.
	Interceptor interface{} `field:"optional" json:"interceptor" yaml:"interceptor"`
}

