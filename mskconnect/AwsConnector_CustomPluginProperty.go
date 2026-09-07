package mskconnect


// Experimental.
type AwsConnector_CustomPluginProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#arn AwsConnector#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#revision AwsConnector#revision}.
	// Experimental.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

