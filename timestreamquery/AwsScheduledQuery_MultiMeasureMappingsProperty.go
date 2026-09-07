package timestreamquery


// Experimental.
type AwsScheduledQuery_MultiMeasureMappingsProperty struct {
	// multi_measure_attribute_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#multi_measure_attribute_mapping AwsScheduledQuery#multi_measure_attribute_mapping}
	// Experimental.
	MultiMeasureAttributeMapping interface{} `field:"optional" json:"multiMeasureAttributeMapping" yaml:"multiMeasureAttributeMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#target_multi_measure_name AwsScheduledQuery#target_multi_measure_name}.
	// Experimental.
	TargetMultiMeasureName *string `field:"optional" json:"targetMultiMeasureName" yaml:"targetMultiMeasureName"`
}

