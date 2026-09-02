package awscloudwatchnetworkflowmonitor


// Experimental.
type TfMonitor_LocalResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkflowmonitor_monitor#identifier TfMonitor#identifier}.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkflowmonitor_monitor#type TfMonitor#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

