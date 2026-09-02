package awsathena


// Experimental.
type TfWorkgroup_ResultConfigurationProperty struct {
	// acl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#acl_configuration TfWorkgroup#acl_configuration}
	// Experimental.
	AclConfiguration *TfWorkgroup_AclConfigurationProperty `field:"optional" json:"aclConfiguration" yaml:"aclConfiguration"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#encryption_configuration TfWorkgroup#encryption_configuration}
	// Experimental.
	EncryptionConfiguration *TfWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationProperty `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#expected_bucket_owner TfWorkgroup#expected_bucket_owner}.
	// Experimental.
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#output_location TfWorkgroup#output_location}.
	// Experimental.
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
}

