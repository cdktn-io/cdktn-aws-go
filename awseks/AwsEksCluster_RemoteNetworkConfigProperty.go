package awseks


// Experimental.
type AwsEksCluster_RemoteNetworkConfigProperty struct {
	// remote_node_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#remote_node_networks AwsEksCluster#remote_node_networks}
	// Experimental.
	RemoteNodeNetworks *AwsEksCluster_RemoteNodeNetworksProperty `field:"optional" json:"remoteNodeNetworks" yaml:"remoteNodeNetworks"`
	// remote_pod_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#remote_pod_networks AwsEksCluster#remote_pod_networks}
	// Experimental.
	RemotePodNetworks *AwsEksCluster_RemotePodNetworksProperty `field:"optional" json:"remotePodNetworks" yaml:"remotePodNetworks"`
}

