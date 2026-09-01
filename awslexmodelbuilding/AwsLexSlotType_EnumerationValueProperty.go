package awslexmodelbuilding


// Experimental.
type AwsLexSlotType_EnumerationValueProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_slot_type#value AwsLexSlotType#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_slot_type#synonyms AwsLexSlotType#synonyms}.
	// Experimental.
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
}

