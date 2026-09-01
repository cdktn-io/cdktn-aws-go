package awsworkspacesweb


// Experimental.
type AwsWorkspaceswebUserSettings_ToolbarConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#hidden_toolbar_items AwsWorkspaceswebUserSettings#hidden_toolbar_items}.
	// Experimental.
	HiddenToolbarItems *[]*string `field:"optional" json:"hiddenToolbarItems" yaml:"hiddenToolbarItems"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#max_display_resolution AwsWorkspaceswebUserSettings#max_display_resolution}.
	// Experimental.
	MaxDisplayResolution *string `field:"optional" json:"maxDisplayResolution" yaml:"maxDisplayResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#toolbar_type AwsWorkspaceswebUserSettings#toolbar_type}.
	// Experimental.
	ToolbarType *string `field:"optional" json:"toolbarType" yaml:"toolbarType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#visual_mode AwsWorkspaceswebUserSettings#visual_mode}.
	// Experimental.
	VisualMode *string `field:"optional" json:"visualMode" yaml:"visualMode"`
}

