package awsappflow


// Experimental.
type TfFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#aggregation_type TfFlow#aggregation_type}.
	// Experimental.
	AggregationType *string `field:"optional" json:"aggregationType" yaml:"aggregationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#target_file_size TfFlow#target_file_size}.
	// Experimental.
	TargetFileSize *float64 `field:"optional" json:"targetFileSize" yaml:"targetFileSize"`
}

