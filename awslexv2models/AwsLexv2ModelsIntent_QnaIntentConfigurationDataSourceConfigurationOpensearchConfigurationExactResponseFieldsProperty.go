package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_QnaIntentConfigurationDataSourceConfigurationOpensearchConfigurationExactResponseFieldsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#answer_field AwsLexv2ModelsIntent#answer_field}.
	// Experimental.
	AnswerField *string `field:"required" json:"answerField" yaml:"answerField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#question_field AwsLexv2ModelsIntent#question_field}.
	// Experimental.
	QuestionField *string `field:"required" json:"questionField" yaml:"questionField"`
}

