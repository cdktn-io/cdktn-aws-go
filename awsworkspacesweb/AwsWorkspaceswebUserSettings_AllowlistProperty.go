package awsworkspacesweb


// Experimental.
type AwsWorkspaceswebUserSettings_AllowlistProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#domain AwsWorkspaceswebUserSettings#domain}.
	// Experimental.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#name AwsWorkspaceswebUserSettings#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#path AwsWorkspaceswebUserSettings#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

