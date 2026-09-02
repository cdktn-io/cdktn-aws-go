package awsvpc


// Experimental.
type DataTfPeeringConnections_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_peering_connections#name DataTfPeeringConnections#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/vpc_peering_connections#values DataTfPeeringConnections#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

