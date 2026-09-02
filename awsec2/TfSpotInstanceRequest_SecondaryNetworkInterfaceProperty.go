package awsec2


// Experimental.
type TfSpotInstanceRequest_SecondaryNetworkInterfaceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#network_card_index TfSpotInstanceRequest#network_card_index}.
	// Experimental.
	NetworkCardIndex *float64 `field:"required" json:"networkCardIndex" yaml:"networkCardIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#secondary_subnet_id TfSpotInstanceRequest#secondary_subnet_id}.
	// Experimental.
	SecondarySubnetId *string `field:"required" json:"secondarySubnetId" yaml:"secondarySubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#delete_on_termination TfSpotInstanceRequest#delete_on_termination}.
	// Experimental.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#device_index TfSpotInstanceRequest#device_index}.
	// Experimental.
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#interface_type TfSpotInstanceRequest#interface_type}.
	// Experimental.
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#private_ip_address_count TfSpotInstanceRequest#private_ip_address_count}.
	// Experimental.
	PrivateIpAddressCount *float64 `field:"optional" json:"privateIpAddressCount" yaml:"privateIpAddressCount"`
}

