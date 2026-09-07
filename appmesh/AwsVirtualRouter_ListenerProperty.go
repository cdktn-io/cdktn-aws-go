package appmesh


// Experimental.
type AwsVirtualRouter_ListenerProperty struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_router#port_mapping AwsVirtualRouter#port_mapping}
	// Experimental.
	PortMapping *AwsVirtualRouter_PortMappingProperty `field:"required" json:"portMapping" yaml:"portMapping"`
}

