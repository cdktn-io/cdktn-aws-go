package awsvpc


// Experimental.
type TfPeeringConnection_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#create TfPeeringConnection#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#delete TfPeeringConnection#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_peering_connection#update TfPeeringConnection#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

