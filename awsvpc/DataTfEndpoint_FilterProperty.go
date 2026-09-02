package awsvpc


// Experimental.
type DataTfEndpoint_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_endpoint#name DataTfEndpoint#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_endpoint#values DataTfEndpoint#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

