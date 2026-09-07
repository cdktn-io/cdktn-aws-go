package elasticache


// Experimental.
type AwsGlobalReplicationGroup_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#create AwsGlobalReplicationGroup#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#delete AwsGlobalReplicationGroup#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#update AwsGlobalReplicationGroup#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

