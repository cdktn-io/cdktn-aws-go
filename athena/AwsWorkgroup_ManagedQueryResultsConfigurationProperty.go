package athena


// Experimental.
type AwsWorkgroup_ManagedQueryResultsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#enabled AwsWorkgroup#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#encryption_configuration AwsWorkgroup#encryption_configuration}
	// Experimental.
	EncryptionConfiguration *AwsWorkgroup_ConfigurationManagedQueryResultsConfigurationEncryptionConfigurationProperty `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

