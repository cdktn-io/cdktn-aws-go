package awsautoscaling


// Experimental.
type DataTfGroups_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/autoscaling_groups#name DataTfGroups#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/autoscaling_groups#values DataTfGroups#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

