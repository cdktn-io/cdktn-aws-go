package kendra


// Experimental.
type AwsDataSource_BasicAuthenticationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#credentials AwsDataSource#credentials}.
	// Experimental.
	Credentials *string `field:"required" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#host AwsDataSource#host}.
	// Experimental.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#port AwsDataSource#port}.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

