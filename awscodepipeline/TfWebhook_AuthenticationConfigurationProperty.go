package awscodepipeline


// Experimental.
type TfWebhook_AuthenticationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_webhook#allowed_ip_range TfWebhook#allowed_ip_range}.
	// Experimental.
	AllowedIpRange *string `field:"optional" json:"allowedIpRange" yaml:"allowedIpRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_webhook#secret_token TfWebhook#secret_token}.
	// Experimental.
	SecretToken *string `field:"optional" json:"secretToken" yaml:"secretToken"`
}

