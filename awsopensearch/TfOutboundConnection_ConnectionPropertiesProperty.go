package awsopensearch


// Experimental.
type TfOutboundConnection_ConnectionPropertiesProperty struct {
	// cross_cluster_search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_outbound_connection#cross_cluster_search TfOutboundConnection#cross_cluster_search}
	// Experimental.
	CrossClusterSearch *TfOutboundConnection_CrossClusterSearchProperty `field:"optional" json:"crossClusterSearch" yaml:"crossClusterSearch"`
}

