package awsbackup


// Experimental.
type AwsBackupPlan_RuleLifecycleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#cold_storage_after AwsBackupPlan#cold_storage_after}.
	// Experimental.
	ColdStorageAfter *float64 `field:"optional" json:"coldStorageAfter" yaml:"coldStorageAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#delete_after AwsBackupPlan#delete_after}.
	// Experimental.
	DeleteAfter *float64 `field:"optional" json:"deleteAfter" yaml:"deleteAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_plan#opt_in_to_archive_for_supported_resources AwsBackupPlan#opt_in_to_archive_for_supported_resources}.
	// Experimental.
	OptInToArchiveForSupportedResources interface{} `field:"optional" json:"optInToArchiveForSupportedResources" yaml:"optInToArchiveForSupportedResources"`
}

