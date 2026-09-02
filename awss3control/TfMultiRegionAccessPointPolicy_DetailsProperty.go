package awss3control


// Experimental.
type TfMultiRegionAccessPointPolicy_DetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point_policy#name TfMultiRegionAccessPointPolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point_policy#policy TfMultiRegionAccessPointPolicy#policy}.
	// Experimental.
	Policy *string `field:"required" json:"policy" yaml:"policy"`
}

