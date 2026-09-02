package awsbedrockagentcore


// Experimental.
type TfBrowser_RecordingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_browser#enabled TfBrowser#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// s3_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_browser#s3_location TfBrowser#s3_location}
	// Experimental.
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

