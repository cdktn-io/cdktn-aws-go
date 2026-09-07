package vpc


// Experimental.
type DataAwsEndpointService_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_endpoint_service#name DataAwsEndpointService#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_endpoint_service#values DataAwsEndpointService#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

