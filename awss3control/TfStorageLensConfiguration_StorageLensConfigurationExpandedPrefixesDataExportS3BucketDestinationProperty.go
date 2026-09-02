package awss3control


// Experimental.
type TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#account_id TfStorageLensConfiguration#account_id}.
	// Experimental.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#arn TfStorageLensConfiguration#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#format TfStorageLensConfiguration#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#output_schema_version TfStorageLensConfiguration#output_schema_version}.
	// Experimental.
	OutputSchemaVersion *string `field:"required" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#encryption TfStorageLensConfiguration#encryption}
	// Experimental.
	Encryption *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportS3BucketDestinationEncryptionProperty `field:"optional" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#prefix TfStorageLensConfiguration#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

