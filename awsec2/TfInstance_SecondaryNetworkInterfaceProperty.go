package awsec2


// Experimental.
type TfInstance_SecondaryNetworkInterfaceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#network_card_index TfInstance#network_card_index}.
	// Experimental.
	NetworkCardIndex *float64 `field:"required" json:"networkCardIndex" yaml:"networkCardIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#secondary_subnet_id TfInstance#secondary_subnet_id}.
	// Experimental.
	SecondarySubnetId *string `field:"required" json:"secondarySubnetId" yaml:"secondarySubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#delete_on_termination TfInstance#delete_on_termination}.
	// Experimental.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#device_index TfInstance#device_index}.
	// Experimental.
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#interface_type TfInstance#interface_type}.
	// Experimental.
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#private_ip_address_count TfInstance#private_ip_address_count}.
	// Experimental.
	PrivateIpAddressCount *float64 `field:"optional" json:"privateIpAddressCount" yaml:"privateIpAddressCount"`
}

