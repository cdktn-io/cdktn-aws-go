package autoscaling


// Experimental.
type DataAwsGroups_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/autoscaling_groups#name DataAwsGroups#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/autoscaling_groups#values DataAwsGroups#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

