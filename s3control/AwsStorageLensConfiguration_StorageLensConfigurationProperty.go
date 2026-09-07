package s3control


// Experimental.
type AwsStorageLensConfiguration_StorageLensConfigurationProperty struct {
	// account_level block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#account_level AwsStorageLensConfiguration#account_level}
	// Experimental.
	AccountLevel *AwsStorageLensConfiguration_AccountLevelProperty `field:"required" json:"accountLevel" yaml:"accountLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#enabled AwsStorageLensConfiguration#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// aws_org block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#aws_org AwsStorageLensConfiguration#aws_org}
	// Experimental.
	AwsOrg *AwsStorageLensConfiguration_AwsOrgProperty `field:"optional" json:"awsOrg" yaml:"awsOrg"`
	// data_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#data_export AwsStorageLensConfiguration#data_export}
	// Experimental.
	DataExport *AwsStorageLensConfiguration_DataExportProperty `field:"optional" json:"dataExport" yaml:"dataExport"`
	// exclude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#exclude AwsStorageLensConfiguration#exclude}
	// Experimental.
	Exclude *AwsStorageLensConfiguration_ExcludeProperty `field:"optional" json:"exclude" yaml:"exclude"`
	// expanded_prefixes_data_export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#expanded_prefixes_data_export AwsStorageLensConfiguration#expanded_prefixes_data_export}
	// Experimental.
	ExpandedPrefixesDataExport *AwsStorageLensConfiguration_ExpandedPrefixesDataExportProperty `field:"optional" json:"expandedPrefixesDataExport" yaml:"expandedPrefixesDataExport"`
	// include block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#include AwsStorageLensConfiguration#include}
	// Experimental.
	Include *AwsStorageLensConfiguration_IncludeProperty `field:"optional" json:"include" yaml:"include"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#prefix_delimiter AwsStorageLensConfiguration#prefix_delimiter}.
	// Experimental.
	PrefixDelimiter *string `field:"optional" json:"prefixDelimiter" yaml:"prefixDelimiter"`
}

