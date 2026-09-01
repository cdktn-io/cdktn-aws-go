package awsglobalaccelerator


// Experimental.
type AwsGlobalacceleratorEndpointGroup_PortOverrideProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_endpoint_group#endpoint_port AwsGlobalacceleratorEndpointGroup#endpoint_port}.
	// Experimental.
	EndpointPort *float64 `field:"required" json:"endpointPort" yaml:"endpointPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_endpoint_group#listener_port AwsGlobalacceleratorEndpointGroup#listener_port}.
	// Experimental.
	ListenerPort *float64 `field:"required" json:"listenerPort" yaml:"listenerPort"`
}

