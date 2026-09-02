package awslexv2models


// Experimental.
type TfIntent_AudioAndDtmfInputSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#start_timeout_ms TfIntent#start_timeout_ms}.
	// Experimental.
	StartTimeoutMs *float64 `field:"required" json:"startTimeoutMs" yaml:"startTimeoutMs"`
	// audio_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#audio_specification TfIntent#audio_specification}
	// Experimental.
	AudioSpecification interface{} `field:"optional" json:"audioSpecification" yaml:"audioSpecification"`
	// dtmf_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#dtmf_specification TfIntent#dtmf_specification}
	// Experimental.
	DtmfSpecification interface{} `field:"optional" json:"dtmfSpecification" yaml:"dtmfSpecification"`
}

