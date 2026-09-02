package awseks


// Experimental.
type TfCluster_RemoteNetworkConfigProperty struct {
	// remote_node_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#remote_node_networks TfCluster#remote_node_networks}
	// Experimental.
	RemoteNodeNetworks *TfCluster_RemoteNodeNetworksProperty `field:"optional" json:"remoteNodeNetworks" yaml:"remoteNodeNetworks"`
	// remote_pod_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#remote_pod_networks TfCluster#remote_pod_networks}
	// Experimental.
	RemotePodNetworks *TfCluster_RemotePodNetworksProperty `field:"optional" json:"remotePodNetworks" yaml:"remotePodNetworks"`
}

