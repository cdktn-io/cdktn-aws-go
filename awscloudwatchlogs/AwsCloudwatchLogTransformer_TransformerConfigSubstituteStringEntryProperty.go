package awscloudwatchlogs


// Experimental.
type AwsCloudwatchLogTransformer_TransformerConfigSubstituteStringEntryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#from AwsCloudwatchLogTransformer#from}.
	// Experimental.
	From *string `field:"required" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source AwsCloudwatchLogTransformer#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#to AwsCloudwatchLogTransformer#to}.
	// Experimental.
	To *string `field:"required" json:"to" yaml:"to"`
}

