package appflow


// Experimental.
type AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty struct {
	// prefix_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#prefix_config AwsFlow#prefix_config}
	// Experimental.
	PrefixConfig *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty `field:"required" json:"prefixConfig" yaml:"prefixConfig"`
	// aggregation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#aggregation_config AwsFlow#aggregation_config}
	// Experimental.
	AggregationConfig *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty `field:"optional" json:"aggregationConfig" yaml:"aggregationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#file_type AwsFlow#file_type}.
	// Experimental.
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
}

