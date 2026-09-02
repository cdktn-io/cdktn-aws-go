package awsvpc


// Experimental.
type TfNetworkInterface_AttachmentProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#device_index TfNetworkInterface#device_index}.
	// Experimental.
	DeviceIndex *float64 `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#instance TfNetworkInterface#instance}.
	// Experimental.
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#network_card_index TfNetworkInterface#network_card_index}.
	// Experimental.
	NetworkCardIndex *float64 `field:"optional" json:"networkCardIndex" yaml:"networkCardIndex"`
}

