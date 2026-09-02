package awscloudwatchlogs


// Experimental.
type TfDelivery_S3DeliveryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_delivery#enable_hive_compatible_path TfDelivery#enable_hive_compatible_path}.
	// Experimental.
	EnableHiveCompatiblePath interface{} `field:"optional" json:"enableHiveCompatiblePath" yaml:"enableHiveCompatiblePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_delivery#suffix_path TfDelivery#suffix_path}.
	// Experimental.
	SuffixPath *string `field:"optional" json:"suffixPath" yaml:"suffixPath"`
}

