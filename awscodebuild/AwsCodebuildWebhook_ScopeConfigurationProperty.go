package awscodebuild


// Experimental.
type AwsCodebuildWebhook_ScopeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_webhook#name AwsCodebuildWebhook#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_webhook#scope AwsCodebuildWebhook#scope}.
	// Experimental.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_webhook#domain AwsCodebuildWebhook#domain}.
	// Experimental.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

