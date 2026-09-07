package kendra


// Experimental.
type AwsDataSource_S3ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#bucket_name AwsDataSource#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// access_control_list_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#access_control_list_configuration AwsDataSource#access_control_list_configuration}
	// Experimental.
	AccessControlListConfiguration *AwsDataSource_AccessControlListConfigurationProperty `field:"optional" json:"accessControlListConfiguration" yaml:"accessControlListConfiguration"`
	// documents_metadata_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#documents_metadata_configuration AwsDataSource#documents_metadata_configuration}
	// Experimental.
	DocumentsMetadataConfiguration *AwsDataSource_DocumentsMetadataConfigurationProperty `field:"optional" json:"documentsMetadataConfiguration" yaml:"documentsMetadataConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#exclusion_patterns AwsDataSource#exclusion_patterns}.
	// Experimental.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#inclusion_patterns AwsDataSource#inclusion_patterns}.
	// Experimental.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#inclusion_prefixes AwsDataSource#inclusion_prefixes}.
	// Experimental.
	InclusionPrefixes *[]*string `field:"optional" json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
}

