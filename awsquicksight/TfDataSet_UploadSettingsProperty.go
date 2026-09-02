package awsquicksight


// Experimental.
type TfDataSet_UploadSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#contains_header TfDataSet#contains_header}.
	// Experimental.
	ContainsHeader interface{} `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#delimiter TfDataSet#delimiter}.
	// Experimental.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#format TfDataSet#format}.
	// Experimental.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#start_from_row TfDataSet#start_from_row}.
	// Experimental.
	StartFromRow *float64 `field:"optional" json:"startFromRow" yaml:"startFromRow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#text_qualifier TfDataSet#text_qualifier}.
	// Experimental.
	TextQualifier *string `field:"optional" json:"textQualifier" yaml:"textQualifier"`
}

