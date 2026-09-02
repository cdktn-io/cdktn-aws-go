package awsrds


// Experimental.
type DataTfSnapshots_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_snapshots#name DataTfSnapshots#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_snapshots#values DataTfSnapshots#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

