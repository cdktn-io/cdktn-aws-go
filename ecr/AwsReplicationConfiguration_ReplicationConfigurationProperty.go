package ecr


// Experimental.
type AwsReplicationConfiguration_ReplicationConfigurationProperty struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_replication_configuration#rule AwsReplicationConfiguration#rule}
	// Experimental.
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

