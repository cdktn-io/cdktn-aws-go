package awsmskconnect


// Experimental.
type TfConnector_CapacityProperty struct {
	// autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#autoscaling TfConnector#autoscaling}
	// Experimental.
	Autoscaling *TfConnector_AutoscalingProperty `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// provisioned_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#provisioned_capacity TfConnector#provisioned_capacity}
	// Experimental.
	ProvisionedCapacity *TfConnector_ProvisionedCapacityProperty `field:"optional" json:"provisionedCapacity" yaml:"provisionedCapacity"`
}

