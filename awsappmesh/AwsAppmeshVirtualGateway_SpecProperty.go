package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_SpecProperty struct {
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#listener AwsAppmeshVirtualGateway#listener}
	// Experimental.
	Listener interface{} `field:"required" json:"listener" yaml:"listener"`
	// backend_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#backend_defaults AwsAppmeshVirtualGateway#backend_defaults}
	// Experimental.
	BackendDefaults *AwsAppmeshVirtualGateway_BackendDefaultsProperty `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#logging AwsAppmeshVirtualGateway#logging}
	// Experimental.
	Logging *AwsAppmeshVirtualGateway_LoggingProperty `field:"optional" json:"logging" yaml:"logging"`
}

