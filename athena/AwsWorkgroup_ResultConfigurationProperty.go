package athena


// Experimental.
type AwsWorkgroup_ResultConfigurationProperty struct {
	// acl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#acl_configuration AwsWorkgroup#acl_configuration}
	// Experimental.
	AclConfiguration *AwsWorkgroup_AclConfigurationProperty `field:"optional" json:"aclConfiguration" yaml:"aclConfiguration"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#encryption_configuration AwsWorkgroup#encryption_configuration}
	// Experimental.
	EncryptionConfiguration *AwsWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationProperty `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#expected_bucket_owner AwsWorkgroup#expected_bucket_owner}.
	// Experimental.
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#output_location AwsWorkgroup#output_location}.
	// Experimental.
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
}

