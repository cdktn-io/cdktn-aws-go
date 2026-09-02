package awsvpc


// Experimental.
type TfNatGateway_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#create TfNatGateway#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#delete TfNatGateway#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/nat_gateway#update TfNatGateway#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

