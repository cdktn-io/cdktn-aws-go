package awslexmodelbuilding


// Experimental.
type TfSlotType_EnumerationValueProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_slot_type#value TfSlotType#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_slot_type#synonyms TfSlotType#synonyms}.
	// Experimental.
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
}

