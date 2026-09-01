package awseks


// Experimental.
type AwsEksCluster_KubeApiServerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#event_ttl AwsEksCluster#event_ttl}.
	// Experimental.
	EventTtl *string `field:"optional" json:"eventTtl" yaml:"eventTtl"`
	// service_node_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#service_node_port_range AwsEksCluster#service_node_port_range}
	// Experimental.
	ServiceNodePortRange *AwsEksCluster_ServiceNodePortRangeProperty `field:"optional" json:"serviceNodePortRange" yaml:"serviceNodePortRange"`
}

