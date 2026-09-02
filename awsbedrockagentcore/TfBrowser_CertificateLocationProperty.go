package awsbedrockagentcore


// Experimental.
type TfBrowser_CertificateLocationProperty struct {
	// secrets_manager block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_browser#secrets_manager TfBrowser#secrets_manager}
	// Experimental.
	SecretsManager interface{} `field:"optional" json:"secretsManager" yaml:"secretsManager"`
}

