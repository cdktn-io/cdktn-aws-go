package awsappmesh


// Experimental.
type TfVirtualNode_SpecProperty struct {
	// backend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#backend TfVirtualNode#backend}
	// Experimental.
	Backend interface{} `field:"optional" json:"backend" yaml:"backend"`
	// backend_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#backend_defaults TfVirtualNode#backend_defaults}
	// Experimental.
	BackendDefaults *TfVirtualNode_BackendDefaultsProperty `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#listener TfVirtualNode#listener}
	// Experimental.
	Listener interface{} `field:"optional" json:"listener" yaml:"listener"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#logging TfVirtualNode#logging}
	// Experimental.
	Logging *TfVirtualNode_LoggingProperty `field:"optional" json:"logging" yaml:"logging"`
	// service_discovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#service_discovery TfVirtualNode#service_discovery}
	// Experimental.
	ServiceDiscovery *TfVirtualNode_ServiceDiscoveryProperty `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}

