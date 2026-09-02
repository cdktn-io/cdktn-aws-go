package awsgamelift


// Experimental.
type TfFleet_Ec2InboundPermissionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#from_port TfFleet#from_port}.
	// Experimental.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#ip_range TfFleet#ip_range}.
	// Experimental.
	IpRange *string `field:"required" json:"ipRange" yaml:"ipRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#protocol TfFleet#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#to_port TfFleet#to_port}.
	// Experimental.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

