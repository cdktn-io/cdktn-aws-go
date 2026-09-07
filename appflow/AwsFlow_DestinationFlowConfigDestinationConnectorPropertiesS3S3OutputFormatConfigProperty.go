package appflow


// Experimental.
type AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigProperty struct {
	// aggregation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#aggregation_config AwsFlow#aggregation_config}
	// Experimental.
	AggregationConfig *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty `field:"optional" json:"aggregationConfig" yaml:"aggregationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#file_type AwsFlow#file_type}.
	// Experimental.
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
	// prefix_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#prefix_config AwsFlow#prefix_config}
	// Experimental.
	PrefixConfig *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigProperty `field:"optional" json:"prefixConfig" yaml:"prefixConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#preserve_source_data_typing AwsFlow#preserve_source_data_typing}.
	// Experimental.
	PreserveSourceDataTyping interface{} `field:"optional" json:"preserveSourceDataTyping" yaml:"preserveSourceDataTyping"`
}

