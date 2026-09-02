package awsappmesh


// Experimental.
type TfVirtualGateway_HealthCheckProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#healthy_threshold TfVirtualGateway#healthy_threshold}.
	// Experimental.
	HealthyThreshold *float64 `field:"required" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#interval_millis TfVirtualGateway#interval_millis}.
	// Experimental.
	IntervalMillis *float64 `field:"required" json:"intervalMillis" yaml:"intervalMillis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#protocol TfVirtualGateway#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#timeout_millis TfVirtualGateway#timeout_millis}.
	// Experimental.
	TimeoutMillis *float64 `field:"required" json:"timeoutMillis" yaml:"timeoutMillis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#unhealthy_threshold TfVirtualGateway#unhealthy_threshold}.
	// Experimental.
	UnhealthyThreshold *float64 `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#path TfVirtualGateway#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#port TfVirtualGateway#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

