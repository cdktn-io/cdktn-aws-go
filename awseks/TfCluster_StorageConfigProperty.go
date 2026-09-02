package awseks


// Experimental.
type TfCluster_StorageConfigProperty struct {
	// block_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#block_storage TfCluster#block_storage}
	// Experimental.
	BlockStorage *TfCluster_BlockStorageProperty `field:"optional" json:"blockStorage" yaml:"blockStorage"`
}

