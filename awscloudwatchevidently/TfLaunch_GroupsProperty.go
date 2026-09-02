package awscloudwatchevidently


// Experimental.
type TfLaunch_GroupsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#feature TfLaunch#feature}.
	// Experimental.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#name TfLaunch#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#variation TfLaunch#variation}.
	// Experimental.
	Variation *string `field:"required" json:"variation" yaml:"variation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#description TfLaunch#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

