package awsworkspacesweb


// Experimental.
type TfUserSettings_BlocklistProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#domain TfUserSettings#domain}.
	// Experimental.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#name TfUserSettings#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_settings#path TfUserSettings#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

