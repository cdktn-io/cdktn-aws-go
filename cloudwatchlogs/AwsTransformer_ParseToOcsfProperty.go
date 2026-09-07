package cloudwatchlogs


// Experimental.
type AwsTransformer_ParseToOcsfProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#event_source AwsTransformer#event_source}.
	// Experimental.
	EventSource *string `field:"required" json:"eventSource" yaml:"eventSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#ocsf_version AwsTransformer#ocsf_version}.
	// Experimental.
	OcsfVersion *string `field:"required" json:"ocsfVersion" yaml:"ocsfVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source AwsTransformer#source}.
	// Experimental.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

