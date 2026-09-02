package awscloudwatchlogs


// Experimental.
type TfTransformer_TransformerConfigSplitStringEntryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#delimiter TfTransformer#delimiter}.
	// Experimental.
	Delimiter *string `field:"required" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source TfTransformer#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
}

