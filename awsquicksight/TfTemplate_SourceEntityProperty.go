package awsquicksight


// Experimental.
type TfTemplate_SourceEntityProperty struct {
	// source_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_template#source_analysis TfTemplate#source_analysis}
	// Experimental.
	SourceAnalysis *TfTemplate_SourceAnalysisProperty `field:"optional" json:"sourceAnalysis" yaml:"sourceAnalysis"`
	// source_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_template#source_template TfTemplate#source_template}
	// Experimental.
	SourceTemplate *TfTemplate_SourceTemplateProperty `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}

