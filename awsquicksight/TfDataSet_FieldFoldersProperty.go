package awsquicksight


// Experimental.
type TfDataSet_FieldFoldersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#field_folders_id TfDataSet#field_folders_id}.
	// Experimental.
	FieldFoldersId *string `field:"required" json:"fieldFoldersId" yaml:"fieldFoldersId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#columns TfDataSet#columns}.
	// Experimental.
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#description TfDataSet#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

