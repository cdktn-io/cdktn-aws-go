package awsvpc


// Experimental.
type TfPeeringConnectionAccepter_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection_accepter#create TfPeeringConnectionAccepter#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection_accepter#update TfPeeringConnectionAccepter#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

