package cloudwatchevidently


// Experimental.
type AwsLaunch_GroupsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#feature AwsLaunch#feature}.
	// Experimental.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#name AwsLaunch#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#variation AwsLaunch#variation}.
	// Experimental.
	Variation *string `field:"required" json:"variation" yaml:"variation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#description AwsLaunch#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

