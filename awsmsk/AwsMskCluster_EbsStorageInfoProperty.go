package awsmsk


// Experimental.
type AwsMskCluster_EbsStorageInfoProperty struct {
	// provisioned_throughput block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#provisioned_throughput AwsMskCluster#provisioned_throughput}
	// Experimental.
	ProvisionedThroughput *AwsMskCluster_ProvisionedThroughputProperty `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#volume_size AwsMskCluster#volume_size}.
	// Experimental.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

