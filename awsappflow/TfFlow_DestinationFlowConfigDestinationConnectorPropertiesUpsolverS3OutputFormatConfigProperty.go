package awsappflow


// Experimental.
type TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty struct {
	// prefix_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#prefix_config TfFlow#prefix_config}
	// Experimental.
	PrefixConfig *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty `field:"required" json:"prefixConfig" yaml:"prefixConfig"`
	// aggregation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#aggregation_config TfFlow#aggregation_config}
	// Experimental.
	AggregationConfig *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty `field:"optional" json:"aggregationConfig" yaml:"aggregationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#file_type TfFlow#file_type}.
	// Experimental.
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
}

