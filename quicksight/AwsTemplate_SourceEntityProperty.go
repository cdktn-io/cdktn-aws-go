package quicksight


// Experimental.
type AwsTemplate_SourceEntityProperty struct {
	// source_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_template#source_analysis AwsTemplate#source_analysis}
	// Experimental.
	SourceAnalysis *AwsTemplate_SourceAnalysisProperty `field:"optional" json:"sourceAnalysis" yaml:"sourceAnalysis"`
	// source_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_template#source_template AwsTemplate#source_template}
	// Experimental.
	SourceTemplate *AwsTemplate_SourceTemplateProperty `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

