package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_OutputContextProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#name AwsLexv2ModelsIntent#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#time_to_live_in_seconds AwsLexv2ModelsIntent#time_to_live_in_seconds}.
	// Experimental.
	TimeToLiveInSeconds *float64 `field:"required" json:"timeToLiveInSeconds" yaml:"timeToLiveInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#turns_to_live AwsLexv2ModelsIntent#turns_to_live}.
	// Experimental.
	TurnsToLive *float64 `field:"required" json:"turnsToLive" yaml:"turnsToLive"`
}

