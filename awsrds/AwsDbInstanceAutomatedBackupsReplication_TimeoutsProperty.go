package awsrds


// Experimental.
type AwsDbInstanceAutomatedBackupsReplication_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance_automated_backups_replication#create AwsDbInstanceAutomatedBackupsReplication#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance_automated_backups_replication#delete AwsDbInstanceAutomatedBackupsReplication#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

