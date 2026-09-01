package awssagemakerai


// Experimental.
type AwsSagemakerFeatureGroup_FeatureDefinitionProperty struct {
	// collection_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#collection_config AwsSagemakerFeatureGroup#collection_config}
	// Experimental.
	CollectionConfig *AwsSagemakerFeatureGroup_CollectionConfigProperty `field:"optional" json:"collectionConfig" yaml:"collectionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#collection_type AwsSagemakerFeatureGroup#collection_type}.
	// Experimental.
	CollectionType *string `field:"optional" json:"collectionType" yaml:"collectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#feature_name AwsSagemakerFeatureGroup#feature_name}.
	// Experimental.
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#feature_type AwsSagemakerFeatureGroup#feature_type}.
	// Experimental.
	FeatureType *string `field:"optional" json:"featureType" yaml:"featureType"`
}

