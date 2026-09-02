package awsmskconnect


// Experimental.
type TfConnector_AutoscalingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#max_worker_count TfConnector#max_worker_count}.
	// Experimental.
	MaxWorkerCount *float64 `field:"required" json:"maxWorkerCount" yaml:"maxWorkerCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#min_worker_count TfConnector#min_worker_count}.
	// Experimental.
	MinWorkerCount *float64 `field:"required" json:"minWorkerCount" yaml:"minWorkerCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#mcu_count TfConnector#mcu_count}.
	// Experimental.
	McuCount *float64 `field:"optional" json:"mcuCount" yaml:"mcuCount"`
	// scale_in_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#scale_in_policy TfConnector#scale_in_policy}
	// Experimental.
	ScaleInPolicy *TfConnector_ScaleInPolicyProperty `field:"optional" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// scale_out_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#scale_out_policy TfConnector#scale_out_policy}
	// Experimental.
	ScaleOutPolicy *TfConnector_ScaleOutPolicyProperty `field:"optional" json:"scaleOutPolicy" yaml:"scaleOutPolicy"`
}

