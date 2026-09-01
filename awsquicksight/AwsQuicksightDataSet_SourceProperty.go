package awsquicksight


// Experimental.
type AwsQuicksightDataSet_SourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#data_set_arn AwsQuicksightDataSet#data_set_arn}.
	// Experimental.
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// join_instruction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#join_instruction AwsQuicksightDataSet#join_instruction}
	// Experimental.
	JoinInstruction *AwsQuicksightDataSet_JoinInstructionProperty `field:"optional" json:"joinInstruction" yaml:"joinInstruction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#physical_table_id AwsQuicksightDataSet#physical_table_id}.
	// Experimental.
	PhysicalTableId *string `field:"optional" json:"physicalTableId" yaml:"physicalTableId"`
}

