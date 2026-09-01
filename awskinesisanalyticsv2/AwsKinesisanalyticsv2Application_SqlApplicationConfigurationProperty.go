package awskinesisanalyticsv2


// Experimental.
type AwsKinesisanalyticsv2Application_SqlApplicationConfigurationProperty struct {
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#input AwsKinesisanalyticsv2Application#input}
	// Experimental.
	Input *AwsKinesisanalyticsv2Application_InputProperty `field:"optional" json:"input" yaml:"input"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#output AwsKinesisanalyticsv2Application#output}
	// Experimental.
	Output interface{} `field:"optional" json:"output" yaml:"output"`
	// reference_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#reference_data_source AwsKinesisanalyticsv2Application#reference_data_source}
	// Experimental.
	ReferenceDataSource *AwsKinesisanalyticsv2Application_ReferenceDataSourceProperty `field:"optional" json:"referenceDataSource" yaml:"referenceDataSource"`
}

