package awscloudwatchlogs


// Experimental.
type AwsCloudwatchLogTransformer_ParseToOcsfProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#event_source AwsCloudwatchLogTransformer#event_source}.
	// Experimental.
	EventSource *string `field:"required" json:"eventSource" yaml:"eventSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#ocsf_version AwsCloudwatchLogTransformer#ocsf_version}.
	// Experimental.
	OcsfVersion *string `field:"required" json:"ocsfVersion" yaml:"ocsfVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source AwsCloudwatchLogTransformer#source}.
	// Experimental.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

