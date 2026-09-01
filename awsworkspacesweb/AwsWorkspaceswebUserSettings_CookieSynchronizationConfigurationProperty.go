package awsworkspacesweb


// Experimental.
type AwsWorkspaceswebUserSettings_CookieSynchronizationConfigurationProperty struct {
	// allowlist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#allowlist AwsWorkspaceswebUserSettings#allowlist}
	// Experimental.
	Allowlist interface{} `field:"optional" json:"allowlist" yaml:"allowlist"`
	// blocklist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#blocklist AwsWorkspaceswebUserSettings#blocklist}
	// Experimental.
	Blocklist interface{} `field:"optional" json:"blocklist" yaml:"blocklist"`
}

