package awseks


// Experimental.
type TfCluster_KubeApiServerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#event_ttl TfCluster#event_ttl}.
	// Experimental.
	EventTtl *string `field:"optional" json:"eventTtl" yaml:"eventTtl"`
	// service_node_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#service_node_port_range TfCluster#service_node_port_range}
	// Experimental.
	ServiceNodePortRange *TfCluster_ServiceNodePortRangeProperty `field:"optional" json:"serviceNodePortRange" yaml:"serviceNodePortRange"`
}

