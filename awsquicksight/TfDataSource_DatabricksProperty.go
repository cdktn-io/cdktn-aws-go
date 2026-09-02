package awsquicksight


// Experimental.
type TfDataSource_DatabricksProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#host TfDataSource#host}.
	// Experimental.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#port TfDataSource#port}.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#sql_endpoint_path TfDataSource#sql_endpoint_path}.
	// Experimental.
	SqlEndpointPath *string `field:"required" json:"sqlEndpointPath" yaml:"sqlEndpointPath"`
}

