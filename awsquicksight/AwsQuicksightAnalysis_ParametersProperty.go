package awsquicksight


// Experimental.
type AwsQuicksightAnalysis_ParametersProperty struct {
	// date_time_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#date_time_parameters AwsQuicksightAnalysis#date_time_parameters}
	// Experimental.
	DateTimeParameters interface{} `field:"optional" json:"dateTimeParameters" yaml:"dateTimeParameters"`
	// decimal_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#decimal_parameters AwsQuicksightAnalysis#decimal_parameters}
	// Experimental.
	DecimalParameters interface{} `field:"optional" json:"decimalParameters" yaml:"decimalParameters"`
	// integer_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#integer_parameters AwsQuicksightAnalysis#integer_parameters}
	// Experimental.
	IntegerParameters interface{} `field:"optional" json:"integerParameters" yaml:"integerParameters"`
	// string_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_analysis#string_parameters AwsQuicksightAnalysis#string_parameters}
	// Experimental.
	StringParameters interface{} `field:"optional" json:"stringParameters" yaml:"stringParameters"`
}

