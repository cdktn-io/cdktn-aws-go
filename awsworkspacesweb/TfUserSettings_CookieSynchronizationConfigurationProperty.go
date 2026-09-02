package awsworkspacesweb


// Experimental.
type TfUserSettings_CookieSynchronizationConfigurationProperty struct {
	// allowlist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#allowlist TfUserSettings#allowlist}
	// Experimental.
	Allowlist interface{} `field:"optional" json:"allowlist" yaml:"allowlist"`
	// blocklist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#blocklist TfUserSettings#blocklist}
	// Experimental.
	Blocklist interface{} `field:"optional" json:"blocklist" yaml:"blocklist"`
}

