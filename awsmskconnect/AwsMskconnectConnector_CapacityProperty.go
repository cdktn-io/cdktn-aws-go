package awsmskconnect


// Experimental.
type AwsMskconnectConnector_CapacityProperty struct {
	// autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#autoscaling AwsMskconnectConnector#autoscaling}
	// Experimental.
	Autoscaling *AwsMskconnectConnector_AutoscalingProperty `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// provisioned_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#provisioned_capacity AwsMskconnectConnector#provisioned_capacity}
	// Experimental.
	ProvisionedCapacity *AwsMskconnectConnector_ProvisionedCapacityProperty `field:"optional" json:"provisionedCapacity" yaml:"provisionedCapacity"`
}

