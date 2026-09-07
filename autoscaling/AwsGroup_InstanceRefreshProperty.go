package autoscaling


// Experimental.
type AwsGroup_InstanceRefreshProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#strategy AwsGroup#strategy}.
	// Experimental.
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#preferences AwsGroup#preferences}
	// Experimental.
	Preferences *AwsGroup_PreferencesProperty `field:"optional" json:"preferences" yaml:"preferences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#triggers AwsGroup#triggers}.
	// Experimental.
	Triggers *[]*string `field:"optional" json:"triggers" yaml:"triggers"`
}

