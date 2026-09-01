package awscloudwatchlogs


// Experimental.
type AwsCloudwatchLogTransformer_TransformerConfigSplitStringEntryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#delimiter AwsCloudwatchLogTransformer#delimiter}.
	// Experimental.
	Delimiter *string `field:"required" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source AwsCloudwatchLogTransformer#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
}

