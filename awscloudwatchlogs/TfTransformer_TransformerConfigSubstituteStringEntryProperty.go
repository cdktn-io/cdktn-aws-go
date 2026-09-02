package awscloudwatchlogs


// Experimental.
type TfTransformer_TransformerConfigSubstituteStringEntryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#from TfTransformer#from}.
	// Experimental.
	From *string `field:"required" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source TfTransformer#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#to TfTransformer#to}.
	// Experimental.
	To *string `field:"required" json:"to" yaml:"to"`
}

