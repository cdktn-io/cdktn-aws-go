package awselasticache


// Experimental.
type TfReplicationGroup_NodeGroupConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#node_group_id TfReplicationGroup#node_group_id}.
	// Experimental.
	NodeGroupId *string `field:"optional" json:"nodeGroupId" yaml:"nodeGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#primary_availability_zone TfReplicationGroup#primary_availability_zone}.
	// Experimental.
	PrimaryAvailabilityZone *string `field:"optional" json:"primaryAvailabilityZone" yaml:"primaryAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#primary_outpost_arn TfReplicationGroup#primary_outpost_arn}.
	// Experimental.
	PrimaryOutpostArn *string `field:"optional" json:"primaryOutpostArn" yaml:"primaryOutpostArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#replica_availability_zones TfReplicationGroup#replica_availability_zones}.
	// Experimental.
	ReplicaAvailabilityZones *[]*string `field:"optional" json:"replicaAvailabilityZones" yaml:"replicaAvailabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#replica_count TfReplicationGroup#replica_count}.
	// Experimental.
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#replica_outpost_arns TfReplicationGroup#replica_outpost_arns}.
	// Experimental.
	ReplicaOutpostArns *[]*string `field:"optional" json:"replicaOutpostArns" yaml:"replicaOutpostArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#slots TfReplicationGroup#slots}.
	// Experimental.
	Slots *string `field:"optional" json:"slots" yaml:"slots"`
}

