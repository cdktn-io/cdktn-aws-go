package awssagemakerai


// Experimental.
type TfFeatureGroup_DataCatalogConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#catalog TfFeatureGroup#catalog}.
	// Experimental.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#database TfFeatureGroup#database}.
	// Experimental.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#table_name TfFeatureGroup#table_name}.
	// Experimental.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

