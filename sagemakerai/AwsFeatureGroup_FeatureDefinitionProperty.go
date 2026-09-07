package sagemakerai


// Experimental.
type AwsFeatureGroup_FeatureDefinitionProperty struct {
	// collection_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#collection_config AwsFeatureGroup#collection_config}
	// Experimental.
	CollectionConfig *AwsFeatureGroup_CollectionConfigProperty `field:"optional" json:"collectionConfig" yaml:"collectionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#collection_type AwsFeatureGroup#collection_type}.
	// Experimental.
	CollectionType *string `field:"optional" json:"collectionType" yaml:"collectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#feature_name AwsFeatureGroup#feature_name}.
	// Experimental.
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#feature_type AwsFeatureGroup#feature_type}.
	// Experimental.
	FeatureType *string `field:"optional" json:"featureType" yaml:"featureType"`
}

