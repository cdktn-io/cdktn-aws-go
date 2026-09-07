package globalaccelerator


// Experimental.
type AwsCustomRoutingEndpointGroup_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_endpoint_group#create AwsCustomRoutingEndpointGroup#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_endpoint_group#delete AwsCustomRoutingEndpointGroup#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

