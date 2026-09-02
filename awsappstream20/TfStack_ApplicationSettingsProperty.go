package awsappstream20


// Experimental.
type TfStack_ApplicationSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#enabled TfStack#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#settings_group TfStack#settings_group}.
	// Experimental.
	SettingsGroup *string `field:"optional" json:"settingsGroup" yaml:"settingsGroup"`
}

