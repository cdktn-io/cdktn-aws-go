package lexv2models


// Experimental.
type AwsIntent_DtmfSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#deletion_character AwsIntent#deletion_character}.
	// Experimental.
	DeletionCharacter *string `field:"required" json:"deletionCharacter" yaml:"deletionCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#end_character AwsIntent#end_character}.
	// Experimental.
	EndCharacter *string `field:"required" json:"endCharacter" yaml:"endCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#end_timeout_ms AwsIntent#end_timeout_ms}.
	// Experimental.
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#max_length AwsIntent#max_length}.
	// Experimental.
	MaxLength *float64 `field:"required" json:"maxLength" yaml:"maxLength"`
}

