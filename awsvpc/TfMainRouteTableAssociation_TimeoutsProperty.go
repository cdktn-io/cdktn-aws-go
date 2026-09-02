package awsvpc


// Experimental.
type TfMainRouteTableAssociation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/main_route_table_association#create TfMainRouteTableAssociation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/main_route_table_association#delete TfMainRouteTableAssociation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/main_route_table_association#update TfMainRouteTableAssociation#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

