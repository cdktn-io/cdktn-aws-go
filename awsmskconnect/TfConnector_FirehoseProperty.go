package awsmskconnect


// Experimental.
type TfConnector_FirehoseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#enabled TfConnector#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#delivery_stream TfConnector#delivery_stream}.
	// Experimental.
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

