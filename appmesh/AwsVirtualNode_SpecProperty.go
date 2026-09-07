package appmesh


// Experimental.
type AwsVirtualNode_SpecProperty struct {
	// backend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#backend AwsVirtualNode#backend}
	// Experimental.
	Backend interface{} `field:"optional" json:"backend" yaml:"backend"`
	// backend_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#backend_defaults AwsVirtualNode#backend_defaults}
	// Experimental.
	BackendDefaults *AwsVirtualNode_BackendDefaultsProperty `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#listener AwsVirtualNode#listener}
	// Experimental.
	Listener interface{} `field:"optional" json:"listener" yaml:"listener"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#logging AwsVirtualNode#logging}
	// Experimental.
	Logging *AwsVirtualNode_LoggingProperty `field:"optional" json:"logging" yaml:"logging"`
	// service_discovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#service_discovery AwsVirtualNode#service_discovery}
	// Experimental.
	ServiceDiscovery *AwsVirtualNode_ServiceDiscoveryProperty `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}

