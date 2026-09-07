package ecr


// Experimental.
type AwsReplicationConfiguration_RepositoryFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_replication_configuration#filter AwsReplicationConfiguration#filter}.
	// Experimental.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecr_replication_configuration#filter_type AwsReplicationConfiguration#filter_type}.
	// Experimental.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

