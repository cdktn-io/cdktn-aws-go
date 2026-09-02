package awsmsk


// Experimental.
type TfCluster_ProvisionedThroughputProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#enabled TfCluster#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#volume_throughput TfCluster#volume_throughput}.
	// Experimental.
	VolumeThroughput *float64 `field:"optional" json:"volumeThroughput" yaml:"volumeThroughput"`
}

