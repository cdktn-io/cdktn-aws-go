package awsappflow


// Experimental.
type TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigProperty struct {
	// aggregation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#aggregation_config TfFlow#aggregation_config}
	// Experimental.
	AggregationConfig *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty `field:"optional" json:"aggregationConfig" yaml:"aggregationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#file_type TfFlow#file_type}.
	// Experimental.
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
	// prefix_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#prefix_config TfFlow#prefix_config}
	// Experimental.
	PrefixConfig *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigProperty `field:"optional" json:"prefixConfig" yaml:"prefixConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#preserve_source_data_typing TfFlow#preserve_source_data_typing}.
	// Experimental.
	PreserveSourceDataTyping interface{} `field:"optional" json:"preserveSourceDataTyping" yaml:"preserveSourceDataTyping"`
}

