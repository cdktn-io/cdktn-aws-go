package elasticache


// Experimental.
type AwsReplicationGroup_NodeGroupConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#node_group_id AwsReplicationGroup#node_group_id}.
	// Experimental.
	NodeGroupId *string `field:"optional" json:"nodeGroupId" yaml:"nodeGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#primary_availability_zone AwsReplicationGroup#primary_availability_zone}.
	// Experimental.
	PrimaryAvailabilityZone *string `field:"optional" json:"primaryAvailabilityZone" yaml:"primaryAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#primary_outpost_arn AwsReplicationGroup#primary_outpost_arn}.
	// Experimental.
	PrimaryOutpostArn *string `field:"optional" json:"primaryOutpostArn" yaml:"primaryOutpostArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#replica_availability_zones AwsReplicationGroup#replica_availability_zones}.
	// Experimental.
	ReplicaAvailabilityZones *[]*string `field:"optional" json:"replicaAvailabilityZones" yaml:"replicaAvailabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#replica_count AwsReplicationGroup#replica_count}.
	// Experimental.
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#replica_outpost_arns AwsReplicationGroup#replica_outpost_arns}.
	// Experimental.
	ReplicaOutpostArns *[]*string `field:"optional" json:"replicaOutpostArns" yaml:"replicaOutpostArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#slots AwsReplicationGroup#slots}.
	// Experimental.
	Slots *string `field:"optional" json:"slots" yaml:"slots"`
}

