package cloudwatchevidently


// Experimental.
type AwsLaunch_MetricDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#entity_id_key AwsLaunch#entity_id_key}.
	// Experimental.
	EntityIdKey *string `field:"required" json:"entityIdKey" yaml:"entityIdKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#name AwsLaunch#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#value_key AwsLaunch#value_key}.
	// Experimental.
	ValueKey *string `field:"required" json:"valueKey" yaml:"valueKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#event_pattern AwsLaunch#event_pattern}.
	// Experimental.
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#unit_label AwsLaunch#unit_label}.
	// Experimental.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
}

