package msk


// Experimental.
type AwsCluster_BrokerNodeGroupInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#client_subnets AwsCluster#client_subnets}.
	// Experimental.
	ClientSubnets *[]*string `field:"required" json:"clientSubnets" yaml:"clientSubnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#instance_type AwsCluster#instance_type}.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#security_groups AwsCluster#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#az_distribution AwsCluster#az_distribution}.
	// Experimental.
	AzDistribution *string `field:"optional" json:"azDistribution" yaml:"azDistribution"`
	// connectivity_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#connectivity_info AwsCluster#connectivity_info}
	// Experimental.
	ConnectivityInfo *AwsCluster_ConnectivityInfoProperty `field:"optional" json:"connectivityInfo" yaml:"connectivityInfo"`
	// storage_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#storage_info AwsCluster#storage_info}
	// Experimental.
	StorageInfo *AwsCluster_StorageInfoProperty `field:"optional" json:"storageInfo" yaml:"storageInfo"`
}

