package awscloudwatchlogs


// Experimental.
type TfTransformer_CsvProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#columns TfTransformer#columns}.
	// Experimental.
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#delimiter TfTransformer#delimiter}.
	// Experimental.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#quote_character TfTransformer#quote_character}.
	// Experimental.
	QuoteCharacter *string `field:"optional" json:"quoteCharacter" yaml:"quoteCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source TfTransformer#source}.
	// Experimental.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

