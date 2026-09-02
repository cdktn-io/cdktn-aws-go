package awsathena


// Experimental.
type TfWorkgroup_ManagedQueryResultsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#enabled TfWorkgroup#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#encryption_configuration TfWorkgroup#encryption_configuration}
	// Experimental.
	EncryptionConfiguration *TfWorkgroup_ConfigurationManagedQueryResultsConfigurationEncryptionConfigurationProperty `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

