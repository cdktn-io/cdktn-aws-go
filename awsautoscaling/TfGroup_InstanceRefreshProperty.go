package awsautoscaling


// Experimental.
type TfGroup_InstanceRefreshProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#strategy TfGroup#strategy}.
	// Experimental.
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#preferences TfGroup#preferences}
	// Experimental.
	Preferences *TfGroup_PreferencesProperty `field:"optional" json:"preferences" yaml:"preferences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#triggers TfGroup#triggers}.
	// Experimental.
	Triggers *[]*string `field:"optional" json:"triggers" yaml:"triggers"`
}

