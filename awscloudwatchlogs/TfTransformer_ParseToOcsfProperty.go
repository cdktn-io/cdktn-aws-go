package awscloudwatchlogs


// Experimental.
type TfTransformer_ParseToOcsfProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#event_source TfTransformer#event_source}.
	// Experimental.
	EventSource *string `field:"required" json:"eventSource" yaml:"eventSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#ocsf_version TfTransformer#ocsf_version}.
	// Experimental.
	OcsfVersion *string `field:"required" json:"ocsfVersion" yaml:"ocsfVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source TfTransformer#source}.
	// Experimental.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

