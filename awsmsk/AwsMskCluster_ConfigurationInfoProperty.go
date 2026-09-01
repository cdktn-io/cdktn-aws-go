package awsmsk


// Experimental.
type AwsMskCluster_ConfigurationInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#arn AwsMskCluster#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#revision AwsMskCluster#revision}.
	// Experimental.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

