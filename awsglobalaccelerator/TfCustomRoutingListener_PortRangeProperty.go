package awsglobalaccelerator


// Experimental.
type TfCustomRoutingListener_PortRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_listener#from_port TfCustomRoutingListener#from_port}.
	// Experimental.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_custom_routing_listener#to_port TfCustomRoutingListener#to_port}.
	// Experimental.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

