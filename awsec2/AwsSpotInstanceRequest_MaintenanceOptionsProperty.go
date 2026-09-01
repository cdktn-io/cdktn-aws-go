package awsec2


// Experimental.
type AwsSpotInstanceRequest_MaintenanceOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#auto_recovery AwsSpotInstanceRequest#auto_recovery}.
	// Experimental.
	AutoRecovery *string `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
}

