package awsec2


// Experimental.
type TfLaunchTemplate_BaselineEbsBandwidthMbpsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#max TfLaunchTemplate#max}.
	// Experimental.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#min TfLaunchTemplate#min}.
	// Experimental.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

