package msk


// Experimental.
type AwsCluster_ConnectivityInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#network_type AwsCluster#network_type}.
	// Experimental.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// public_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#public_access AwsCluster#public_access}
	// Experimental.
	PublicAccess *AwsCluster_PublicAccessProperty `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// vpc_connectivity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#vpc_connectivity AwsCluster#vpc_connectivity}
	// Experimental.
	VpcConnectivity *AwsCluster_VpcConnectivityProperty `field:"optional" json:"vpcConnectivity" yaml:"vpcConnectivity"`
}

