package quicksight


// Experimental.
type AwsDataSource_SparkProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#host AwsDataSource#host}.
	// Experimental.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#port AwsDataSource#port}.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

