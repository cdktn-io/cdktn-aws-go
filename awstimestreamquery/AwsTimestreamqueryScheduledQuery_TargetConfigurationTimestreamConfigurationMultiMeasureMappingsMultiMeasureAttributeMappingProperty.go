package awstimestreamquery


// Experimental.
type AwsTimestreamqueryScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#measure_value_type AwsTimestreamqueryScheduledQuery#measure_value_type}.
	// Experimental.
	MeasureValueType *string `field:"required" json:"measureValueType" yaml:"measureValueType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#source_column AwsTimestreamqueryScheduledQuery#source_column}.
	// Experimental.
	SourceColumn *string `field:"required" json:"sourceColumn" yaml:"sourceColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#target_multi_measure_attribute_name AwsTimestreamqueryScheduledQuery#target_multi_measure_attribute_name}.
	// Experimental.
	TargetMultiMeasureAttributeName *string `field:"optional" json:"targetMultiMeasureAttributeName" yaml:"targetMultiMeasureAttributeName"`
}

