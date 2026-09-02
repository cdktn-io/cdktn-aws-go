package awstimestreamquery


// Experimental.
type TfScheduledQuery_MixedMeasureMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#measure_value_type TfScheduledQuery#measure_value_type}.
	// Experimental.
	MeasureValueType *string `field:"required" json:"measureValueType" yaml:"measureValueType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#measure_name TfScheduledQuery#measure_name}.
	// Experimental.
	MeasureName *string `field:"optional" json:"measureName" yaml:"measureName"`
	// multi_measure_attribute_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#multi_measure_attribute_mapping TfScheduledQuery#multi_measure_attribute_mapping}
	// Experimental.
	MultiMeasureAttributeMapping interface{} `field:"optional" json:"multiMeasureAttributeMapping" yaml:"multiMeasureAttributeMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#source_column TfScheduledQuery#source_column}.
	// Experimental.
	SourceColumn *string `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#target_measure_name TfScheduledQuery#target_measure_name}.
	// Experimental.
	TargetMeasureName *string `field:"optional" json:"targetMeasureName" yaml:"targetMeasureName"`
}

