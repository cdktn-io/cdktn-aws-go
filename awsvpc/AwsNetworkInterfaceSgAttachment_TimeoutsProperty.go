package awsvpc


// Experimental.
type AwsNetworkInterfaceSgAttachment_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface_sg_attachment#create AwsNetworkInterfaceSgAttachment#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface_sg_attachment#delete AwsNetworkInterfaceSgAttachment#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface_sg_attachment#read AwsNetworkInterfaceSgAttachment#read}.
	// Experimental.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

