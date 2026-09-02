package awsmetadatasources


// Experimental.
type DataTfRegions_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/regions#name DataTfRegions#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/regions#values DataTfRegions#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

