package awsbedrockagentcore


// Experimental.
type TfBrowser_VpcConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_browser#security_groups TfBrowser#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_browser#subnets TfBrowser#subnets}.
	// Experimental.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

