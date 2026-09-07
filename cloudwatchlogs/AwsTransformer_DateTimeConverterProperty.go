package cloudwatchlogs


// Experimental.
type AwsTransformer_DateTimeConverterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#match_patterns AwsTransformer#match_patterns}.
	// Experimental.
	MatchPatterns *[]*string `field:"required" json:"matchPatterns" yaml:"matchPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source AwsTransformer#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#target AwsTransformer#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#locale AwsTransformer#locale}.
	// Experimental.
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source_timezone AwsTransformer#source_timezone}.
	// Experimental.
	SourceTimezone *string `field:"optional" json:"sourceTimezone" yaml:"sourceTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#target_format AwsTransformer#target_format}.
	// Experimental.
	TargetFormat *string `field:"optional" json:"targetFormat" yaml:"targetFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#target_timezone AwsTransformer#target_timezone}.
	// Experimental.
	TargetTimezone *string `field:"optional" json:"targetTimezone" yaml:"targetTimezone"`
}

