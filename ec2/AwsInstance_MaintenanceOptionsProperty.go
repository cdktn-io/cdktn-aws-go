package ec2


// Experimental.
type AwsInstance_MaintenanceOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#auto_recovery AwsInstance#auto_recovery}.
	// Experimental.
	AutoRecovery *string `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
}

