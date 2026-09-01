package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_ListenerProperty struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#port_mapping AwsAppmeshVirtualNode#port_mapping}
	// Experimental.
	PortMapping *AwsAppmeshVirtualNode_PortMappingProperty `field:"required" json:"portMapping" yaml:"portMapping"`
	// connection_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#connection_pool AwsAppmeshVirtualNode#connection_pool}
	// Experimental.
	ConnectionPool *AwsAppmeshVirtualNode_ConnectionPoolProperty `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#health_check AwsAppmeshVirtualNode#health_check}
	// Experimental.
	HealthCheck *AwsAppmeshVirtualNode_HealthCheckProperty `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// outlier_detection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#outlier_detection AwsAppmeshVirtualNode#outlier_detection}
	// Experimental.
	OutlierDetection *AwsAppmeshVirtualNode_OutlierDetectionProperty `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#timeout AwsAppmeshVirtualNode#timeout}
	// Experimental.
	Timeout *AwsAppmeshVirtualNode_TimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#tls AwsAppmeshVirtualNode#tls}
	// Experimental.
	Tls *AwsAppmeshVirtualNode_SpecListenerTlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

