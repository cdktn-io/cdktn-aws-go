package appstream20


// Experimental.
type AwsStack_UserSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#action AwsStack#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#permission AwsStack#permission}.
	// Experimental.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

