package awsdms


// Experimental.
type TfReplicationConfig_ComputeConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#replication_subnet_group_id TfReplicationConfig#replication_subnet_group_id}.
	// Experimental.
	ReplicationSubnetGroupId *string `field:"required" json:"replicationSubnetGroupId" yaml:"replicationSubnetGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#availability_zone TfReplicationConfig#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#dns_name_servers TfReplicationConfig#dns_name_servers}.
	// Experimental.
	DnsNameServers *string `field:"optional" json:"dnsNameServers" yaml:"dnsNameServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#kms_key_id TfReplicationConfig#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#max_capacity_units TfReplicationConfig#max_capacity_units}.
	// Experimental.
	MaxCapacityUnits *float64 `field:"optional" json:"maxCapacityUnits" yaml:"maxCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#min_capacity_units TfReplicationConfig#min_capacity_units}.
	// Experimental.
	MinCapacityUnits *float64 `field:"optional" json:"minCapacityUnits" yaml:"minCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#multi_az TfReplicationConfig#multi_az}.
	// Experimental.
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#preferred_maintenance_window TfReplicationConfig#preferred_maintenance_window}.
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#vpc_security_group_ids TfReplicationConfig#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

