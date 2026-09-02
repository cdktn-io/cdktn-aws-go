package awsmsk


// Experimental.
type TfCluster_BrokerNodeGroupInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#client_subnets TfCluster#client_subnets}.
	// Experimental.
	ClientSubnets *[]*string `field:"required" json:"clientSubnets" yaml:"clientSubnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#instance_type TfCluster#instance_type}.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#security_groups TfCluster#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#az_distribution TfCluster#az_distribution}.
	// Experimental.
	AzDistribution *string `field:"optional" json:"azDistribution" yaml:"azDistribution"`
	// connectivity_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#connectivity_info TfCluster#connectivity_info}
	// Experimental.
	ConnectivityInfo *TfCluster_ConnectivityInfoProperty `field:"optional" json:"connectivityInfo" yaml:"connectivityInfo"`
	// storage_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#storage_info TfCluster#storage_info}
	// Experimental.
	StorageInfo *TfCluster_StorageInfoProperty `field:"optional" json:"storageInfo" yaml:"storageInfo"`
}

