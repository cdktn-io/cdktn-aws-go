package awsrds


// Experimental.
type DataTfEngineVersion_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#name DataTfEngineVersion#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/rds_engine_version#values DataTfEngineVersion#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

