package awsappstream20


// Experimental.
type AwsAppstreamStack_UserSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#action AwsAppstreamStack#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#permission AwsAppstreamStack#permission}.
	// Experimental.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

