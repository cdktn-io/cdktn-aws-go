package awsglobalaccelerator


// Experimental.
type AwsGlobalacceleratorCustomRoutingEndpointGroup_DestinationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_endpoint_group#from_port AwsGlobalacceleratorCustomRoutingEndpointGroup#from_port}.
	// Experimental.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_endpoint_group#protocols AwsGlobalacceleratorCustomRoutingEndpointGroup#protocols}.
	// Experimental.
	Protocols *[]*string `field:"required" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_endpoint_group#to_port AwsGlobalacceleratorCustomRoutingEndpointGroup#to_port}.
	// Experimental.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

