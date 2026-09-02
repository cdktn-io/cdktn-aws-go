package awsvpc


// Experimental.
type TfSubnet_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#create TfSubnet#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/subnet#delete TfSubnet#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

