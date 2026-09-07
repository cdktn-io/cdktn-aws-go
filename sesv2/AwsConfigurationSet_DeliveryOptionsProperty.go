package sesv2


// Experimental.
type AwsConfigurationSet_DeliveryOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#max_delivery_seconds AwsConfigurationSet#max_delivery_seconds}.
	// Experimental.
	MaxDeliverySeconds *float64 `field:"optional" json:"maxDeliverySeconds" yaml:"maxDeliverySeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#sending_pool_name AwsConfigurationSet#sending_pool_name}.
	// Experimental.
	SendingPoolName *string `field:"optional" json:"sendingPoolName" yaml:"sendingPoolName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#tls_policy AwsConfigurationSet#tls_policy}.
	// Experimental.
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

