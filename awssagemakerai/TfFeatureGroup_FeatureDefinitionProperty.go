package awssagemakerai


// Experimental.
type TfFeatureGroup_FeatureDefinitionProperty struct {
	// collection_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#collection_config TfFeatureGroup#collection_config}
	// Experimental.
	CollectionConfig *TfFeatureGroup_CollectionConfigProperty `field:"optional" json:"collectionConfig" yaml:"collectionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#collection_type TfFeatureGroup#collection_type}.
	// Experimental.
	CollectionType *string `field:"optional" json:"collectionType" yaml:"collectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#feature_name TfFeatureGroup#feature_name}.
	// Experimental.
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#feature_type TfFeatureGroup#feature_type}.
	// Experimental.
	FeatureType *string `field:"optional" json:"featureType" yaml:"featureType"`
}

