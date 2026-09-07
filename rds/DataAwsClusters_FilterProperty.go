package rds


// Experimental.
type DataAwsClusters_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_clusters#name DataAwsClusters#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_clusters#values DataAwsClusters#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

