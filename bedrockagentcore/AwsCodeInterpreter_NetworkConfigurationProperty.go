package bedrockagentcore


// Experimental.
type AwsCodeInterpreter_NetworkConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_code_interpreter#network_mode AwsCodeInterpreter#network_mode}.
	// Experimental.
	NetworkMode *string `field:"required" json:"networkMode" yaml:"networkMode"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_code_interpreter#vpc_config AwsCodeInterpreter#vpc_config}
	// Experimental.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

