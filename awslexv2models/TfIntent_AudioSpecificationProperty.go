package awslexv2models


// Experimental.
type TfIntent_AudioSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#end_timeout_ms TfIntent#end_timeout_ms}.
	// Experimental.
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#max_length_ms TfIntent#max_length_ms}.
	// Experimental.
	MaxLengthMs *float64 `field:"required" json:"maxLengthMs" yaml:"maxLengthMs"`
}

