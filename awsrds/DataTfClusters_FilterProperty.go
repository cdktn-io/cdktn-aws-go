package awsrds


// Experimental.
type DataTfClusters_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_clusters#name DataTfClusters#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_clusters#values DataTfClusters#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

