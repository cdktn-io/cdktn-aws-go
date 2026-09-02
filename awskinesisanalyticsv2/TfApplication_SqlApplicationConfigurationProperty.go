package awskinesisanalyticsv2


// Experimental.
type TfApplication_SqlApplicationConfigurationProperty struct {
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input TfApplication#input}
	// Experimental.
	Input *TfApplication_InputProperty `field:"optional" json:"input" yaml:"input"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#output TfApplication#output}
	// Experimental.
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// reference_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#reference_data_source TfApplication#reference_data_source}
	// Experimental.
	ReferenceDataSource *TfApplication_ReferenceDataSourceProperty `field:"optional" json:"referenceDataSource" yaml:"referenceDataSource"`
}

