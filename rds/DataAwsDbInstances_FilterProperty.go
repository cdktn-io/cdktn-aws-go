package rds


// Experimental.
type DataAwsDbInstances_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/db_instances#name DataAwsDbInstances#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/db_instances#values DataAwsDbInstances#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

