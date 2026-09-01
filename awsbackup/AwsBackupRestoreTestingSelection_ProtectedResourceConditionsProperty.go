package awsbackup


// Experimental.
type AwsBackupRestoreTestingSelection_ProtectedResourceConditionsProperty struct {
	// string_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#string_equals AwsBackupRestoreTestingSelection#string_equals}
	// Experimental.
	StringEquals interface{} `field:"optional" json:"stringEquals" yaml:"stringEquals"`
	// string_not_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#string_not_equals AwsBackupRestoreTestingSelection#string_not_equals}
	// Experimental.
	StringNotEquals interface{} `field:"optional" json:"stringNotEquals" yaml:"stringNotEquals"`
}

