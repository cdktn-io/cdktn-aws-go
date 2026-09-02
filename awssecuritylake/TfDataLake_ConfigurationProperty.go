package awssecuritylake


// Experimental.
type TfDataLake_ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#region TfDataLake#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#encryption_configuration TfDataLake#encryption_configuration}.
	// Experimental.
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// lifecycle_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#lifecycle_configuration TfDataLake#lifecycle_configuration}
	// Experimental.
	LifecycleConfiguration interface{} `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// replication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_data_lake#replication_configuration TfDataLake#replication_configuration}
	// Experimental.
	ReplicationConfiguration interface{} `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
}

