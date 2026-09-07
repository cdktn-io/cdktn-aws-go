package msk


// Experimental.
type AwsCluster_VpcConnectivityProperty struct {
	// client_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#client_authentication AwsCluster#client_authentication}
	// Experimental.
	ClientAuthentication *AwsCluster_BrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationProperty `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
}

