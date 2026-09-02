package awskendra


// Experimental.
type TfDataSource_ProxyConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#host TfDataSource#host}.
	// Experimental.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#port TfDataSource#port}.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#credentials TfDataSource#credentials}.
	// Experimental.
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
}

