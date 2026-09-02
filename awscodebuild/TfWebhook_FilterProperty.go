package awscodebuild


// Experimental.
type TfWebhook_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_webhook#pattern TfWebhook#pattern}.
	// Experimental.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_webhook#type TfWebhook#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_webhook#exclude_matched_pattern TfWebhook#exclude_matched_pattern}.
	// Experimental.
	ExcludeMatchedPattern interface{} `field:"optional" json:"excludeMatchedPattern" yaml:"excludeMatchedPattern"`
}

