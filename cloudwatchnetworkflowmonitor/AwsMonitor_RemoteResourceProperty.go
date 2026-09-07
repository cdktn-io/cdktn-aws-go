package cloudwatchnetworkflowmonitor


// Experimental.
type AwsMonitor_RemoteResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkflowmonitor_monitor#identifier AwsMonitor#identifier}.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkflowmonitor_monitor#type AwsMonitor#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

