package awssagemakerai


// Experimental.
type TfFeatureGroup_OnlineStoreConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#enable_online_store TfFeatureGroup#enable_online_store}.
	// Experimental.
	EnableOnlineStore interface{} `field:"optional" json:"enableOnlineStore" yaml:"enableOnlineStore"`
	// security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#security_config TfFeatureGroup#security_config}
	// Experimental.
	SecurityConfig *TfFeatureGroup_SecurityConfigProperty `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#storage_type TfFeatureGroup#storage_type}.
	// Experimental.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// ttl_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#ttl_duration TfFeatureGroup#ttl_duration}
	// Experimental.
	TtlDuration *TfFeatureGroup_TtlDurationProperty `field:"optional" json:"ttlDuration" yaml:"ttlDuration"`
}

