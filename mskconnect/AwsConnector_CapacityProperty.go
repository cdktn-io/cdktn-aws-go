package mskconnect


// Experimental.
type AwsConnector_CapacityProperty struct {
	// autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#autoscaling AwsConnector#autoscaling}
	// Experimental.
	Autoscaling *AwsConnector_AutoscalingProperty `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// provisioned_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#provisioned_capacity AwsConnector#provisioned_capacity}
	// Experimental.
	ProvisionedCapacity *AwsConnector_ProvisionedCapacityProperty `field:"optional" json:"provisionedCapacity" yaml:"provisionedCapacity"`
}

