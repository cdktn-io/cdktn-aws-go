package cloudwatchlogs


// Experimental.
type AwsTransformer_ListToMapProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#key AwsTransformer#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#source AwsTransformer#source}.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#flatten AwsTransformer#flatten}.
	// Experimental.
	Flatten interface{} `field:"optional" json:"flatten" yaml:"flatten"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#flattened_element AwsTransformer#flattened_element}.
	// Experimental.
	FlattenedElement *string `field:"optional" json:"flattenedElement" yaml:"flattenedElement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#target AwsTransformer#target}.
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_transformer#value_key AwsTransformer#value_key}.
	// Experimental.
	ValueKey *string `field:"optional" json:"valueKey" yaml:"valueKey"`
}

