package awsappmesh


// Experimental.
type TfVirtualNode_SpecListenerConnectionPoolHttpProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#max_connections TfVirtualNode#max_connections}.
	// Experimental.
	MaxConnections *float64 `field:"required" json:"maxConnections" yaml:"maxConnections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#max_pending_requests TfVirtualNode#max_pending_requests}.
	// Experimental.
	MaxPendingRequests *float64 `field:"optional" json:"maxPendingRequests" yaml:"maxPendingRequests"`
}

