package awsbedrockagentcore


// Experimental.
type TfBrowser_NetworkConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_browser#network_mode TfBrowser#network_mode}.
	// Experimental.
	NetworkMode *string `field:"required" json:"networkMode" yaml:"networkMode"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_browser#vpc_config TfBrowser#vpc_config}
	// Experimental.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

