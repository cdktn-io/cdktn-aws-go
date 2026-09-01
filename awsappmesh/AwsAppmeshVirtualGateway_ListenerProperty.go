package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_ListenerProperty struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#port_mapping AwsAppmeshVirtualGateway#port_mapping}
	// Experimental.
	PortMapping *AwsAppmeshVirtualGateway_PortMappingProperty `field:"required" json:"portMapping" yaml:"portMapping"`
	// connection_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#connection_pool AwsAppmeshVirtualGateway#connection_pool}
	// Experimental.
	ConnectionPool *AwsAppmeshVirtualGateway_ConnectionPoolProperty `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#health_check AwsAppmeshVirtualGateway#health_check}
	// Experimental.
	HealthCheck *AwsAppmeshVirtualGateway_HealthCheckProperty `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#tls AwsAppmeshVirtualGateway#tls}
	// Experimental.
	Tls *AwsAppmeshVirtualGateway_SpecListenerTlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

