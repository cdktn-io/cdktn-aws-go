package awsappmesh


// Experimental.
type TfVirtualGateway_SpecProperty struct {
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#listener TfVirtualGateway#listener}
	// Experimental.
	Listener interface{} `field:"required" json:"listener" yaml:"listener"`
	// backend_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#backend_defaults TfVirtualGateway#backend_defaults}
	// Experimental.
	BackendDefaults *TfVirtualGateway_BackendDefaultsProperty `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#logging TfVirtualGateway#logging}
	// Experimental.
	Logging *TfVirtualGateway_LoggingProperty `field:"optional" json:"logging" yaml:"logging"`
}

