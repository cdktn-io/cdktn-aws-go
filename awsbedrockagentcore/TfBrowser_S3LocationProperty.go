package awsbedrockagentcore


// Experimental.
type TfBrowser_S3LocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_browser#bucket TfBrowser#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_browser#prefix TfBrowser#prefix}.
	// Experimental.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

