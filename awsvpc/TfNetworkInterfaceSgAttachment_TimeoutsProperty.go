package awsvpc


// Experimental.
type TfNetworkInterfaceSgAttachment_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface_sg_attachment#create TfNetworkInterfaceSgAttachment#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface_sg_attachment#delete TfNetworkInterfaceSgAttachment#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface_sg_attachment#read TfNetworkInterfaceSgAttachment#read}.
	// Experimental.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

