package quicksight


// Experimental.
type AwsDataSet_JoinInstructionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#left_operand AwsDataSet#left_operand}.
	// Experimental.
	LeftOperand *string `field:"required" json:"leftOperand" yaml:"leftOperand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#on_clause AwsDataSet#on_clause}.
	// Experimental.
	OnClause *string `field:"required" json:"onClause" yaml:"onClause"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#right_operand AwsDataSet#right_operand}.
	// Experimental.
	RightOperand *string `field:"required" json:"rightOperand" yaml:"rightOperand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#type AwsDataSet#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// left_join_key_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#left_join_key_properties AwsDataSet#left_join_key_properties}
	// Experimental.
	LeftJoinKeyProperties *AwsDataSet_LeftJoinKeyPropertiesProperty `field:"optional" json:"leftJoinKeyProperties" yaml:"leftJoinKeyProperties"`
	// right_join_key_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#right_join_key_properties AwsDataSet#right_join_key_properties}
	// Experimental.
	RightJoinKeyProperties *AwsDataSet_RightJoinKeyPropertiesProperty `field:"optional" json:"rightJoinKeyProperties" yaml:"rightJoinKeyProperties"`
}

