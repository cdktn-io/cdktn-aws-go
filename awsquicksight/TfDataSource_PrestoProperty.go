package awsquicksight


// Experimental.
type TfDataSource_PrestoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#catalog TfDataSource#catalog}.
	// Experimental.
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#host TfDataSource#host}.
	// Experimental.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#port TfDataSource#port}.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

