package awscloudwatchlogs


// Experimental.
type TfTransformer_GrokProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#match TfTransformer#match}.
	// Experimental.
	Match *string `field:"required" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source TfTransformer#source}.
	// Experimental.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

