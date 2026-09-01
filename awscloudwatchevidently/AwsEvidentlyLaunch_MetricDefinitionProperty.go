package awscloudwatchevidently


// Experimental.
type AwsEvidentlyLaunch_MetricDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#entity_id_key AwsEvidentlyLaunch#entity_id_key}.
	// Experimental.
	EntityIdKey *string `field:"required" json:"entityIdKey" yaml:"entityIdKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#name AwsEvidentlyLaunch#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#value_key AwsEvidentlyLaunch#value_key}.
	// Experimental.
	ValueKey *string `field:"required" json:"valueKey" yaml:"valueKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#event_pattern AwsEvidentlyLaunch#event_pattern}.
	// Experimental.
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#unit_label AwsEvidentlyLaunch#unit_label}.
	// Experimental.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
}

