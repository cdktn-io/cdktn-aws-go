package awscloudwatchevidently


// Experimental.
type TfLaunch_MetricDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#entity_id_key TfLaunch#entity_id_key}.
	// Experimental.
	EntityIdKey *string `field:"required" json:"entityIdKey" yaml:"entityIdKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#name TfLaunch#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#value_key TfLaunch#value_key}.
	// Experimental.
	ValueKey *string `field:"required" json:"valueKey" yaml:"valueKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#event_pattern TfLaunch#event_pattern}.
	// Experimental.
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/evidently_launch#unit_label TfLaunch#unit_label}.
	// Experimental.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
}

