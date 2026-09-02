package awsbedrockagentcore


// Experimental.
type TfBrowser_EnterprisePolicyProperty struct {
	// location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_browser#location TfBrowser#location}
	// Experimental.
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_browser#type TfBrowser#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

