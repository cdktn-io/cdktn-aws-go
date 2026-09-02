package awsbackup


// Experimental.
type TfRestoreTestingSelection_ProtectedResourceConditionsProperty struct {
	// string_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#string_equals TfRestoreTestingSelection#string_equals}
	// Experimental.
	StringEquals interface{} `field:"optional" json:"stringEquals" yaml:"stringEquals"`
	// string_not_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_restore_testing_selection#string_not_equals TfRestoreTestingSelection#string_not_equals}
	// Experimental.
	StringNotEquals interface{} `field:"optional" json:"stringNotEquals" yaml:"stringNotEquals"`
}

