package awsquicksight


// Experimental.
type AwsQuicksightDataSet_LogicalTableMapProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#alias AwsQuicksightDataSet#alias}.
	// Experimental.
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#logical_table_map_id AwsQuicksightDataSet#logical_table_map_id}.
	// Experimental.
	LogicalTableMapId *string `field:"required" json:"logicalTableMapId" yaml:"logicalTableMapId"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#source AwsQuicksightDataSet#source}
	// Experimental.
	Source *AwsQuicksightDataSet_SourceProperty `field:"required" json:"source" yaml:"source"`
	// data_transforms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#data_transforms AwsQuicksightDataSet#data_transforms}
	// Experimental.
	DataTransforms interface{} `field:"optional" json:"dataTransforms" yaml:"dataTransforms"`
}

