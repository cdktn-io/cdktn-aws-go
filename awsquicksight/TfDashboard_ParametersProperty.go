package awsquicksight


// Experimental.
type TfDashboard_ParametersProperty struct {
	// date_time_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#date_time_parameters TfDashboard#date_time_parameters}
	// Experimental.
	DateTimeParameters interface{} `field:"optional" json:"dateTimeParameters" yaml:"dateTimeParameters"`
	// decimal_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#decimal_parameters TfDashboard#decimal_parameters}
	// Experimental.
	DecimalParameters interface{} `field:"optional" json:"decimalParameters" yaml:"decimalParameters"`
	// integer_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#integer_parameters TfDashboard#integer_parameters}
	// Experimental.
	IntegerParameters interface{} `field:"optional" json:"integerParameters" yaml:"integerParameters"`
	// string_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#string_parameters TfDashboard#string_parameters}
	// Experimental.
	StringParameters interface{} `field:"optional" json:"stringParameters" yaml:"stringParameters"`
}

