package awsec2


// Experimental.
type AwsLaunchTemplate_AcceleratorCountProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#max AwsLaunchTemplate#max}.
	// Experimental.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#min AwsLaunchTemplate#min}.
	// Experimental.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

