package workspaces


// Experimental.
type AwsPool_ApplicationSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_pool#settings_group AwsPool#settings_group}.
	// Experimental.
	SettingsGroup *string `field:"optional" json:"settingsGroup" yaml:"settingsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_pool#status AwsPool#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

