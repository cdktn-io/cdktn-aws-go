package awsappmesh


// Experimental.
type TfVirtualNode_OutlierDetectionProperty struct {
	// base_ejection_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#base_ejection_duration TfVirtualNode#base_ejection_duration}
	// Experimental.
	BaseEjectionDuration *TfVirtualNode_BaseEjectionDurationProperty `field:"required" json:"baseEjectionDuration" yaml:"baseEjectionDuration"`
	// interval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#interval TfVirtualNode#interval}
	// Experimental.
	Interval *TfVirtualNode_IntervalProperty `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#max_ejection_percent TfVirtualNode#max_ejection_percent}.
	// Experimental.
	MaxEjectionPercent *float64 `field:"required" json:"maxEjectionPercent" yaml:"maxEjectionPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#max_server_errors TfVirtualNode#max_server_errors}.
	// Experimental.
	MaxServerErrors *float64 `field:"required" json:"maxServerErrors" yaml:"maxServerErrors"`
}

