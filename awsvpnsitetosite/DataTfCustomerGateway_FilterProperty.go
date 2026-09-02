package awsvpnsitetosite


// Experimental.
type DataTfCustomerGateway_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/customer_gateway#name DataTfCustomerGateway#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/customer_gateway#values DataTfCustomerGateway#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

