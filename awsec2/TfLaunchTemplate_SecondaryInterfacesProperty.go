package awsec2


// Experimental.
type TfLaunchTemplate_SecondaryInterfacesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#delete_on_termination TfLaunchTemplate#delete_on_termination}.
	// Experimental.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#device_index TfLaunchTemplate#device_index}.
	// Experimental.
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#interface_type TfLaunchTemplate#interface_type}.
	// Experimental.
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#network_card_index TfLaunchTemplate#network_card_index}.
	// Experimental.
	NetworkCardIndex *float64 `field:"optional" json:"networkCardIndex" yaml:"networkCardIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#private_ip_address_count TfLaunchTemplate#private_ip_address_count}.
	// Experimental.
	PrivateIpAddressCount *float64 `field:"optional" json:"privateIpAddressCount" yaml:"privateIpAddressCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#private_ip_addresses TfLaunchTemplate#private_ip_addresses}.
	// Experimental.
	PrivateIpAddresses *[]*string `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#secondary_subnet_id TfLaunchTemplate#secondary_subnet_id}.
	// Experimental.
	SecondarySubnetId *string `field:"optional" json:"secondarySubnetId" yaml:"secondarySubnetId"`
}

