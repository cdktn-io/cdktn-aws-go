package appmesh


// Experimental.
type AwsVirtualNode_ListenerProperty struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#port_mapping AwsVirtualNode#port_mapping}
	// Experimental.
	PortMapping *AwsVirtualNode_PortMappingProperty `field:"required" json:"portMapping" yaml:"portMapping"`
	// connection_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#connection_pool AwsVirtualNode#connection_pool}
	// Experimental.
	ConnectionPool *AwsVirtualNode_ConnectionPoolProperty `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#health_check AwsVirtualNode#health_check}
	// Experimental.
	HealthCheck *AwsVirtualNode_HealthCheckProperty `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// outlier_detection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#outlier_detection AwsVirtualNode#outlier_detection}
	// Experimental.
	OutlierDetection *AwsVirtualNode_OutlierDetectionProperty `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#timeout AwsVirtualNode#timeout}
	// Experimental.
	Timeout *AwsVirtualNode_TimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#tls AwsVirtualNode#tls}
	// Experimental.
	Tls *AwsVirtualNode_SpecListenerTlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

