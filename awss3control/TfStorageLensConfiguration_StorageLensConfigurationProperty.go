package awss3control


// Experimental.
type TfStorageLensConfiguration_StorageLensConfigurationProperty struct {
	// account_level block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#account_level TfStorageLensConfiguration#account_level}
	// Experimental.
	AccountLevel *TfStorageLensConfiguration_AccountLevelProperty `field:"required" json:"accountLevel" yaml:"accountLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#enabled TfStorageLensConfiguration#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// aws_org block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#aws_org TfStorageLensConfiguration#aws_org}
	// Experimental.
	AwsOrg *TfStorageLensConfiguration_AwsOrgProperty `field:"optional" json:"awsOrg" yaml:"awsOrg"`
	// data_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#data_export TfStorageLensConfiguration#data_export}
	// Experimental.
	DataExport *TfStorageLensConfiguration_DataExportProperty `field:"optional" json:"dataExport" yaml:"dataExport"`
	// exclude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#exclude TfStorageLensConfiguration#exclude}
	// Experimental.
	Exclude *TfStorageLensConfiguration_ExcludeProperty `field:"optional" json:"exclude" yaml:"exclude"`
	// expanded_prefixes_data_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#expanded_prefixes_data_export TfStorageLensConfiguration#expanded_prefixes_data_export}
	// Experimental.
	ExpandedPrefixesDataExport *TfStorageLensConfiguration_ExpandedPrefixesDataExportProperty `field:"optional" json:"expandedPrefixesDataExport" yaml:"expandedPrefixesDataExport"`
	// include block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#include TfStorageLensConfiguration#include}
	// Experimental.
	Include *TfStorageLensConfiguration_IncludeProperty `field:"optional" json:"include" yaml:"include"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#prefix_delimiter TfStorageLensConfiguration#prefix_delimiter}.
	// Experimental.
	PrefixDelimiter *string `field:"optional" json:"prefixDelimiter" yaml:"prefixDelimiter"`
}

