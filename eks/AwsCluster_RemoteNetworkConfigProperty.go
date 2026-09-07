package eks


// Experimental.
type AwsCluster_RemoteNetworkConfigProperty struct {
	// remote_node_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#remote_node_networks AwsCluster#remote_node_networks}
	// Experimental.
	RemoteNodeNetworks *AwsCluster_RemoteNodeNetworksProperty `field:"optional" json:"remoteNodeNetworks" yaml:"remoteNodeNetworks"`
	// remote_pod_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#remote_pod_networks AwsCluster#remote_pod_networks}
	// Experimental.
	RemotePodNetworks *AwsCluster_RemotePodNetworksProperty `field:"optional" json:"remotePodNetworks" yaml:"remotePodNetworks"`
}

