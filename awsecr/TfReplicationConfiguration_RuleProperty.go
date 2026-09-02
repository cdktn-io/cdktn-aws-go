package awsecr


// Experimental.
type TfReplicationConfiguration_RuleProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_replication_configuration#destination TfReplicationConfiguration#destination}
	// Experimental.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// repository_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_replication_configuration#repository_filter TfReplicationConfiguration#repository_filter}
	// Experimental.
	RepositoryFilter interface{} `field:"optional" json:"repositoryFilter" yaml:"repositoryFilter"`
}

