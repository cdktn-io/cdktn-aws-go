package codepipeline


// Experimental.
type AwsWebhook_AuthenticationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_webhook#allowed_ip_range AwsWebhook#allowed_ip_range}.
	// Experimental.
	AllowedIpRange *string `field:"optional" json:"allowedIpRange" yaml:"allowedIpRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_webhook#secret_token AwsWebhook#secret_token}.
	// Experimental.
	SecretToken *string `field:"optional" json:"secretToken" yaml:"secretToken"`
}

