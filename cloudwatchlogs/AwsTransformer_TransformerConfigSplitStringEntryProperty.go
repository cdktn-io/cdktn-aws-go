package cloudwatchlogs


// Experimental.
type AwsTransformer_TransformerConfigSplitStringEntryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#delimiter AwsTransformer#delimiter}.
	// Experimental.
	Delimiter *string `field:"required" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source AwsTransformer#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
}

