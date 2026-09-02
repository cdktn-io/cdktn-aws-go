package awsmsk


// Experimental.
type TfReplicator_FirehoseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#enabled TfReplicator#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_replicator#delivery_stream TfReplicator#delivery_stream}.
	// Experimental.
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

