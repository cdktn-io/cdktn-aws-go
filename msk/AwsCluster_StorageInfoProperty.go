package msk


// Experimental.
type AwsCluster_StorageInfoProperty struct {
	// ebs_storage_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#ebs_storage_info AwsCluster#ebs_storage_info}
	// Experimental.
	EbsStorageInfo *AwsCluster_EbsStorageInfoProperty `field:"optional" json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
}

