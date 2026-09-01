package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecProperty struct {
	// backend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#backend AwsAppmeshVirtualNode#backend}
	// Experimental.
	Backend interface{} `field:"optional" json:"backend" yaml:"backend"`
	// backend_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#backend_defaults AwsAppmeshVirtualNode#backend_defaults}
	// Experimental.
	BackendDefaults *AwsAppmeshVirtualNode_BackendDefaultsProperty `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#listener AwsAppmeshVirtualNode#listener}
	// Experimental.
	Listener interface{} `field:"optional" json:"listener" yaml:"listener"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#logging AwsAppmeshVirtualNode#logging}
	// Experimental.
	Logging *AwsAppmeshVirtualNode_LoggingProperty `field:"optional" json:"logging" yaml:"logging"`
	// service_discovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#service_discovery AwsAppmeshVirtualNode#service_discovery}
	// Experimental.
	ServiceDiscovery *AwsAppmeshVirtualNode_ServiceDiscoveryProperty `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}

