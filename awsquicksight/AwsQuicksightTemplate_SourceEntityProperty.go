package awsquicksight


// Experimental.
type AwsQuicksightTemplate_SourceEntityProperty struct {
	// source_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_template#source_analysis AwsQuicksightTemplate#source_analysis}
	// Experimental.
	SourceAnalysis *AwsQuicksightTemplate_SourceAnalysisProperty `field:"optional" json:"sourceAnalysis" yaml:"sourceAnalysis"`
	// source_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_template#source_template AwsQuicksightTemplate#source_template}
	// Experimental.
	SourceTemplate *AwsQuicksightTemplate_SourceTemplateProperty `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

