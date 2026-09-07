package workspacesweb


// Experimental.
type AwsUserSettings_CookieSynchronizationConfigurationProperty struct {
	// allowlist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#allowlist AwsUserSettings#allowlist}
	// Experimental.
	Allowlist interface{} `field:"optional" json:"allowlist" yaml:"allowlist"`
	// blocklist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#blocklist AwsUserSettings#blocklist}
	// Experimental.
	Blocklist interface{} `field:"optional" json:"blocklist" yaml:"blocklist"`
}

