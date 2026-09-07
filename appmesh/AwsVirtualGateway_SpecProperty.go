package appmesh


// Experimental.
type AwsVirtualGateway_SpecProperty struct {
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#listener AwsVirtualGateway#listener}
	// Experimental.
	Listener interface{} `field:"required" json:"listener" yaml:"listener"`
	// backend_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#backend_defaults AwsVirtualGateway#backend_defaults}
	// Experimental.
	BackendDefaults *AwsVirtualGateway_BackendDefaultsProperty `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#logging AwsVirtualGateway#logging}
	// Experimental.
	Logging *AwsVirtualGateway_LoggingProperty `field:"optional" json:"logging" yaml:"logging"`
}

