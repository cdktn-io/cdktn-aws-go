package awsecr


// Experimental.
type TfReplicationConfiguration_ReplicationConfigurationProperty struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_replication_configuration#rule TfReplicationConfiguration#rule}
	// Experimental.
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

