package awsnetworkmanager


// Experimental.
type AwsNetworkmanagerVpcAttachment_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_vpc_attachment#create AwsNetworkmanagerVpcAttachment#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_vpc_attachment#delete AwsNetworkmanagerVpcAttachment#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkmanager_vpc_attachment#update AwsNetworkmanagerVpcAttachment#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

