package awsglue


// Experimental.
type TfMlTransform_FindMatchesParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_ml_transform#accuracy_cost_trade_off TfMlTransform#accuracy_cost_trade_off}.
	// Experimental.
	AccuracyCostTradeOff *float64 `field:"optional" json:"accuracyCostTradeOff" yaml:"accuracyCostTradeOff"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_ml_transform#enforce_provided_labels TfMlTransform#enforce_provided_labels}.
	// Experimental.
	EnforceProvidedLabels interface{} `field:"optional" json:"enforceProvidedLabels" yaml:"enforceProvidedLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_ml_transform#precision_recall_trade_off TfMlTransform#precision_recall_trade_off}.
	// Experimental.
	PrecisionRecallTradeOff *float64 `field:"optional" json:"precisionRecallTradeOff" yaml:"precisionRecallTradeOff"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_ml_transform#primary_key_column_name TfMlTransform#primary_key_column_name}.
	// Experimental.
	PrimaryKeyColumnName *string `field:"optional" json:"primaryKeyColumnName" yaml:"primaryKeyColumnName"`
}

