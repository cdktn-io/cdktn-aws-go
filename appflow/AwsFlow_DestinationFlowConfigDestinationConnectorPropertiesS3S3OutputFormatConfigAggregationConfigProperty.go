package appflow


// Experimental.
type AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#aggregation_type AwsFlow#aggregation_type}.
	// Experimental.
	AggregationType *string `field:"optional" json:"aggregationType" yaml:"aggregationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#target_file_size AwsFlow#target_file_size}.
	// Experimental.
	TargetFileSize *float64 `field:"optional" json:"targetFileSize" yaml:"targetFileSize"`
}

