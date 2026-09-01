package awsmsk


// Experimental.
type AwsMskCluster_ConnectivityInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#network_type AwsMskCluster#network_type}.
	// Experimental.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// public_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#public_access AwsMskCluster#public_access}
	// Experimental.
	PublicAccess *AwsMskCluster_PublicAccessProperty `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// vpc_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#vpc_connectivity AwsMskCluster#vpc_connectivity}
	// Experimental.
	VpcConnectivity *AwsMskCluster_VpcConnectivityProperty `field:"optional" json:"vpcConnectivity" yaml:"vpcConnectivity"`
}

