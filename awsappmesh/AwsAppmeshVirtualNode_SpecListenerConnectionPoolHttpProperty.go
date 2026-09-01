package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecListenerConnectionPoolHttpProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#max_connections AwsAppmeshVirtualNode#max_connections}.
	// Experimental.
	MaxConnections *float64 `field:"required" json:"maxConnections" yaml:"maxConnections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#max_pending_requests AwsAppmeshVirtualNode#max_pending_requests}.
	// Experimental.
	MaxPendingRequests *float64 `field:"optional" json:"maxPendingRequests" yaml:"maxPendingRequests"`
}

