package codebuild


// Experimental.
type AwsWebhook_ScopeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_webhook#name AwsWebhook#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_webhook#scope AwsWebhook#scope}.
	// Experimental.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_webhook#domain AwsWebhook#domain}.
	// Experimental.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

