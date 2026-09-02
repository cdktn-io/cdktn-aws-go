package awsec2


// Experimental.
type TfLaunchTemplate_NetworkInterfacesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#associate_carrier_ip_address TfLaunchTemplate#associate_carrier_ip_address}.
	// Experimental.
	AssociateCarrierIpAddress *string `field:"optional" json:"associateCarrierIpAddress" yaml:"associateCarrierIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#associate_public_ip_address TfLaunchTemplate#associate_public_ip_address}.
	// Experimental.
	AssociatePublicIpAddress *string `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// connection_tracking_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#connection_tracking_specification TfLaunchTemplate#connection_tracking_specification}
	// Experimental.
	ConnectionTrackingSpecification *TfLaunchTemplate_ConnectionTrackingSpecificationProperty `field:"optional" json:"connectionTrackingSpecification" yaml:"connectionTrackingSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#delete_on_termination TfLaunchTemplate#delete_on_termination}.
	// Experimental.
	DeleteOnTermination *string `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#description TfLaunchTemplate#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#device_index TfLaunchTemplate#device_index}.
	// Experimental.
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ena_queue_count TfLaunchTemplate#ena_queue_count}.
	// Experimental.
	EnaQueueCount *float64 `field:"optional" json:"enaQueueCount" yaml:"enaQueueCount"`
	// ena_srd_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ena_srd_specification TfLaunchTemplate#ena_srd_specification}
	// Experimental.
	EnaSrdSpecification *TfLaunchTemplate_EnaSrdSpecificationProperty `field:"optional" json:"enaSrdSpecification" yaml:"enaSrdSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#interface_type TfLaunchTemplate#interface_type}.
	// Experimental.
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ipv4_address_count TfLaunchTemplate#ipv4_address_count}.
	// Experimental.
	Ipv4AddressCount *float64 `field:"optional" json:"ipv4AddressCount" yaml:"ipv4AddressCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ipv4_addresses TfLaunchTemplate#ipv4_addresses}.
	// Experimental.
	Ipv4Addresses *[]*string `field:"optional" json:"ipv4Addresses" yaml:"ipv4Addresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ipv4_prefix_count TfLaunchTemplate#ipv4_prefix_count}.
	// Experimental.
	Ipv4PrefixCount *float64 `field:"optional" json:"ipv4PrefixCount" yaml:"ipv4PrefixCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ipv4_prefixes TfLaunchTemplate#ipv4_prefixes}.
	// Experimental.
	Ipv4Prefixes *[]*string `field:"optional" json:"ipv4Prefixes" yaml:"ipv4Prefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ipv6_address_count TfLaunchTemplate#ipv6_address_count}.
	// Experimental.
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ipv6_addresses TfLaunchTemplate#ipv6_addresses}.
	// Experimental.
	Ipv6Addresses *[]*string `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ipv6_prefix_count TfLaunchTemplate#ipv6_prefix_count}.
	// Experimental.
	Ipv6PrefixCount *float64 `field:"optional" json:"ipv6PrefixCount" yaml:"ipv6PrefixCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ipv6_prefixes TfLaunchTemplate#ipv6_prefixes}.
	// Experimental.
	Ipv6Prefixes *[]*string `field:"optional" json:"ipv6Prefixes" yaml:"ipv6Prefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#network_card_index TfLaunchTemplate#network_card_index}.
	// Experimental.
	NetworkCardIndex *float64 `field:"optional" json:"networkCardIndex" yaml:"networkCardIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#network_interface_id TfLaunchTemplate#network_interface_id}.
	// Experimental.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#primary_ipv6 TfLaunchTemplate#primary_ipv6}.
	// Experimental.
	PrimaryIpv6 *string `field:"optional" json:"primaryIpv6" yaml:"primaryIpv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#private_ip_address TfLaunchTemplate#private_ip_address}.
	// Experimental.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#security_groups TfLaunchTemplate#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#subnet_id TfLaunchTemplate#subnet_id}.
	// Experimental.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

