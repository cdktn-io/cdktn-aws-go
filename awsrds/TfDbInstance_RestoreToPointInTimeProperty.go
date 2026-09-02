package awsrds


// Experimental.
type TfDbInstance_RestoreToPointInTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#restore_time TfDbInstance#restore_time}.
	// Experimental.
	RestoreTime *string `field:"optional" json:"restoreTime" yaml:"restoreTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#source_db_instance_automated_backups_arn TfDbInstance#source_db_instance_automated_backups_arn}.
	// Experimental.
	SourceDbInstanceAutomatedBackupsArn *string `field:"optional" json:"sourceDbInstanceAutomatedBackupsArn" yaml:"sourceDbInstanceAutomatedBackupsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#source_db_instance_identifier TfDbInstance#source_db_instance_identifier}.
	// Experimental.
	SourceDbInstanceIdentifier *string `field:"optional" json:"sourceDbInstanceIdentifier" yaml:"sourceDbInstanceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#source_dbi_resource_id TfDbInstance#source_dbi_resource_id}.
	// Experimental.
	SourceDbiResourceId *string `field:"optional" json:"sourceDbiResourceId" yaml:"sourceDbiResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#use_latest_restorable_time TfDbInstance#use_latest_restorable_time}.
	// Experimental.
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
}

