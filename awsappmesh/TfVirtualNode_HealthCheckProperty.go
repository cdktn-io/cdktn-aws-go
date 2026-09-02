package awsappmesh


// Experimental.
type TfVirtualNode_HealthCheckProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#healthy_threshold TfVirtualNode#healthy_threshold}.
	// Experimental.
	HealthyThreshold *float64 `field:"required" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#interval_millis TfVirtualNode#interval_millis}.
	// Experimental.
	IntervalMillis *float64 `field:"required" json:"intervalMillis" yaml:"intervalMillis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#protocol TfVirtualNode#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#timeout_millis TfVirtualNode#timeout_millis}.
	// Experimental.
	TimeoutMillis *float64 `field:"required" json:"timeoutMillis" yaml:"timeoutMillis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#unhealthy_threshold TfVirtualNode#unhealthy_threshold}.
	// Experimental.
	UnhealthyThreshold *float64 `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#path TfVirtualNode#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#port TfVirtualNode#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

