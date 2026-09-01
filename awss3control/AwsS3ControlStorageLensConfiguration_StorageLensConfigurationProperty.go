package awss3control


// Experimental.
type AwsS3ControlStorageLensConfiguration_StorageLensConfigurationProperty struct {
	// account_level block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#account_level AwsS3ControlStorageLensConfiguration#account_level}
	// Experimental.
	AccountLevel *AwsS3ControlStorageLensConfiguration_AccountLevelProperty `field:"required" json:"accountLevel" yaml:"accountLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#enabled AwsS3ControlStorageLensConfiguration#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// aws_org block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#aws_org AwsS3ControlStorageLensConfiguration#aws_org}
	// Experimental.
	AwsOrg *AwsS3ControlStorageLensConfiguration_AwsOrgProperty `field:"optional" json:"awsOrg" yaml:"awsOrg"`
	// data_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#data_export AwsS3ControlStorageLensConfiguration#data_export}
	// Experimental.
	DataExport *AwsS3ControlStorageLensConfiguration_DataExportProperty `field:"optional" json:"dataExport" yaml:"dataExport"`
	// exclude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#exclude AwsS3ControlStorageLensConfiguration#exclude}
	// Experimental.
	Exclude *AwsS3ControlStorageLensConfiguration_ExcludeProperty `field:"optional" json:"exclude" yaml:"exclude"`
	// expanded_prefixes_data_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#expanded_prefixes_data_export AwsS3ControlStorageLensConfiguration#expanded_prefixes_data_export}
	// Experimental.
	ExpandedPrefixesDataExport *AwsS3ControlStorageLensConfiguration_ExpandedPrefixesDataExportProperty `field:"optional" json:"expandedPrefixesDataExport" yaml:"expandedPrefixesDataExport"`
	// include block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#include AwsS3ControlStorageLensConfiguration#include}
	// Experimental.
	Include *AwsS3ControlStorageLensConfiguration_IncludeProperty `field:"optional" json:"include" yaml:"include"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#prefix_delimiter AwsS3ControlStorageLensConfiguration#prefix_delimiter}.
	// Experimental.
	PrefixDelimiter *string `field:"optional" json:"prefixDelimiter" yaml:"prefixDelimiter"`
}

