package awsmskconnect


// Experimental.
type TfConnector_CustomPluginProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#arn TfConnector#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#revision TfConnector#revision}.
	// Experimental.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

