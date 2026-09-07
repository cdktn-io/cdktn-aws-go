package sagemakerai


// Experimental.
type AwsFeatureGroup_OnlineStoreConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#enable_online_store AwsFeatureGroup#enable_online_store}.
	// Experimental.
	EnableOnlineStore interface{} `field:"optional" json:"enableOnlineStore" yaml:"enableOnlineStore"`
	// security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#security_config AwsFeatureGroup#security_config}
	// Experimental.
	SecurityConfig *AwsFeatureGroup_SecurityConfigProperty `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#storage_type AwsFeatureGroup#storage_type}.
	// Experimental.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// ttl_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#ttl_duration AwsFeatureGroup#ttl_duration}
	// Experimental.
	TtlDuration *AwsFeatureGroup_TtlDurationProperty `field:"optional" json:"ttlDuration" yaml:"ttlDuration"`
}

