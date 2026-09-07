package sagemakerai


// Experimental.
type AwsFeatureGroup_DataCatalogConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#catalog AwsFeatureGroup#catalog}.
	// Experimental.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#database AwsFeatureGroup#database}.
	// Experimental.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#table_name AwsFeatureGroup#table_name}.
	// Experimental.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

