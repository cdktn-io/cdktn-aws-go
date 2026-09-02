package awsappmesh


// Experimental.
type TfVirtualGateway_ListenerProperty struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#port_mapping TfVirtualGateway#port_mapping}
	// Experimental.
	PortMapping *TfVirtualGateway_PortMappingProperty `field:"required" json:"portMapping" yaml:"portMapping"`
	// connection_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#connection_pool TfVirtualGateway#connection_pool}
	// Experimental.
	ConnectionPool *TfVirtualGateway_ConnectionPoolProperty `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#health_check TfVirtualGateway#health_check}
	// Experimental.
	HealthCheck *TfVirtualGateway_HealthCheckProperty `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#tls TfVirtualGateway#tls}
	// Experimental.
	Tls *TfVirtualGateway_SpecListenerTlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

