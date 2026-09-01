package awsathena


// Experimental.
type AwsAthenaWorkgroup_ResultConfigurationProperty struct {
	// acl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#acl_configuration AwsAthenaWorkgroup#acl_configuration}
	// Experimental.
	AclConfiguration *AwsAthenaWorkgroup_AclConfigurationProperty `field:"optional" json:"aclConfiguration" yaml:"aclConfiguration"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#encryption_configuration AwsAthenaWorkgroup#encryption_configuration}
	// Experimental.
	EncryptionConfiguration *AwsAthenaWorkgroup_ConfigurationResultConfigurationEncryptionConfigurationProperty `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#expected_bucket_owner AwsAthenaWorkgroup#expected_bucket_owner}.
	// Experimental.
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#output_location AwsAthenaWorkgroup#output_location}.
	// Experimental.
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
}

