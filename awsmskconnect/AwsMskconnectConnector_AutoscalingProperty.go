package awsmskconnect


// Experimental.
type AwsMskconnectConnector_AutoscalingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#max_worker_count AwsMskconnectConnector#max_worker_count}.
	// Experimental.
	MaxWorkerCount *float64 `field:"required" json:"maxWorkerCount" yaml:"maxWorkerCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#min_worker_count AwsMskconnectConnector#min_worker_count}.
	// Experimental.
	MinWorkerCount *float64 `field:"required" json:"minWorkerCount" yaml:"minWorkerCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#mcu_count AwsMskconnectConnector#mcu_count}.
	// Experimental.
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
	// scale_in_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#scale_in_policy AwsMskconnectConnector#scale_in_policy}
	// Experimental.
	ScaleInPolicy *AwsMskconnectConnector_ScaleInPolicyProperty `field:"optional" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// scale_out_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#scale_out_policy AwsMskconnectConnector#scale_out_policy}
	// Experimental.
	ScaleOutPolicy *AwsMskconnectConnector_ScaleOutPolicyProperty `field:"optional" json:"scaleOutPolicy" yaml:"scaleOutPolicy"`
}

