package awsappmesh


// Experimental.
type AwsAppmeshVirtualRouter_ListenerProperty struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_router#port_mapping AwsAppmeshVirtualRouter#port_mapping}
	// Experimental.
	PortMapping *AwsAppmeshVirtualRouter_PortMappingProperty `field:"required" json:"portMapping" yaml:"portMapping"`
}

