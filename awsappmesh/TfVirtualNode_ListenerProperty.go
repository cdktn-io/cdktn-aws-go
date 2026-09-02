package awsappmesh


// Experimental.
type TfVirtualNode_ListenerProperty struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#port_mapping TfVirtualNode#port_mapping}
	// Experimental.
	PortMapping *TfVirtualNode_PortMappingProperty `field:"required" json:"portMapping" yaml:"portMapping"`
	// connection_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#connection_pool TfVirtualNode#connection_pool}
	// Experimental.
	ConnectionPool *TfVirtualNode_ConnectionPoolProperty `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#health_check TfVirtualNode#health_check}
	// Experimental.
	HealthCheck *TfVirtualNode_HealthCheckProperty `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// outlier_detection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#outlier_detection TfVirtualNode#outlier_detection}
	// Experimental.
	OutlierDetection *TfVirtualNode_OutlierDetectionProperty `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#timeout TfVirtualNode#timeout}
	// Experimental.
	Timeout *TfVirtualNode_TimeoutProperty `field:"optional" json:"timeout" yaml:"timeout"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#tls TfVirtualNode#tls}
	// Experimental.
	Tls *TfVirtualNode_SpecListenerTlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

