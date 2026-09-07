package rds


// Experimental.
type DataAwsSnapshots_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_snapshots#name DataAwsSnapshots#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_snapshots#values DataAwsSnapshots#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

