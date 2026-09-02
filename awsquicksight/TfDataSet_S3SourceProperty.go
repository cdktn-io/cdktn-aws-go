package awsquicksight


// Experimental.
type TfDataSet_S3SourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#data_source_arn TfDataSet#data_source_arn}.
	// Experimental.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// input_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#input_columns TfDataSet#input_columns}
	// Experimental.
	InputColumns interface{} `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// upload_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#upload_settings TfDataSet#upload_settings}
	// Experimental.
	UploadSettings *TfDataSet_UploadSettingsProperty `field:"required" json:"uploadSettings" yaml:"uploadSettings"`
}

