package awssagemakerai


// Experimental.
type AwsSagemakerFeatureGroup_OnlineStoreConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#enable_online_store AwsSagemakerFeatureGroup#enable_online_store}.
	// Experimental.
	EnableOnlineStore interface{} `field:"optional" json:"enableOnlineStore" yaml:"enableOnlineStore"`
	// security_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#security_config AwsSagemakerFeatureGroup#security_config}
	// Experimental.
	SecurityConfig *AwsSagemakerFeatureGroup_SecurityConfigProperty `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#storage_type AwsSagemakerFeatureGroup#storage_type}.
	// Experimental.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// ttl_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_feature_group#ttl_duration AwsSagemakerFeatureGroup#ttl_duration}
	// Experimental.
	TtlDuration *AwsSagemakerFeatureGroup_TtlDurationProperty `field:"optional" json:"ttlDuration" yaml:"ttlDuration"`
}

