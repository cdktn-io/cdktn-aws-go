package kinesisanalyticsv2


// Experimental.
type AwsApplication_SqlApplicationConfigurationProperty struct {
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input AwsApplication#input}
	// Experimental.
	Input *AwsApplication_InputProperty `field:"optional" json:"input" yaml:"input"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#output AwsApplication#output}
	// Experimental.
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// reference_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#reference_data_source AwsApplication#reference_data_source}
	// Experimental.
	ReferenceDataSource *AwsApplication_ReferenceDataSourceProperty `field:"optional" json:"referenceDataSource" yaml:"referenceDataSource"`
}

