package timestreamquery


// Experimental.
type AwsScheduledQuery_DimensionMappingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#dimension_value_type AwsScheduledQuery#dimension_value_type}.
	// Experimental.
	DimensionValueType *string `field:"required" json:"dimensionValueType" yaml:"dimensionValueType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#name AwsScheduledQuery#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

