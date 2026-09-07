package appmesh


// Experimental.
type AwsVirtualGateway_ListenerProperty struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#port_mapping AwsVirtualGateway#port_mapping}
	// Experimental.
	PortMapping *AwsVirtualGateway_PortMappingProperty `field:"required" json:"portMapping" yaml:"portMapping"`
	// connection_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#connection_pool AwsVirtualGateway#connection_pool}
	// Experimental.
	ConnectionPool *AwsVirtualGateway_ConnectionPoolProperty `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#health_check AwsVirtualGateway#health_check}
	// Experimental.
	HealthCheck *AwsVirtualGateway_HealthCheckProperty `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#tls AwsVirtualGateway#tls}
	// Experimental.
	Tls *AwsVirtualGateway_SpecListenerTlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

