package awsquicksight


// Experimental.
type TfDataSet_JoinInstructionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#left_operand TfDataSet#left_operand}.
	// Experimental.
	LeftOperand *string `field:"required" json:"leftOperand" yaml:"leftOperand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#on_clause TfDataSet#on_clause}.
	// Experimental.
	OnClause *string `field:"required" json:"onClause" yaml:"onClause"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#right_operand TfDataSet#right_operand}.
	// Experimental.
	RightOperand *string `field:"required" json:"rightOperand" yaml:"rightOperand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#type TfDataSet#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// left_join_key_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#left_join_key_properties TfDataSet#left_join_key_properties}
	// Experimental.
	LeftJoinKeyProperties *TfDataSet_LeftJoinKeyPropertiesProperty `field:"optional" json:"leftJoinKeyProperties" yaml:"leftJoinKeyProperties"`
	// right_join_key_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#right_join_key_properties TfDataSet#right_join_key_properties}
	// Experimental.
	RightJoinKeyProperties *TfDataSet_RightJoinKeyPropertiesProperty `field:"optional" json:"rightJoinKeyProperties" yaml:"rightJoinKeyProperties"`
}

