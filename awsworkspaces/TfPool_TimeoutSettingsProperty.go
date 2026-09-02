package awsworkspaces


// Experimental.
type TfPool_TimeoutSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_pool#disconnect_timeout_in_seconds TfPool#disconnect_timeout_in_seconds}.
	// Experimental.
	DisconnectTimeoutInSeconds *float64 `field:"optional" json:"disconnectTimeoutInSeconds" yaml:"disconnectTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_pool#idle_disconnect_timeout_in_seconds TfPool#idle_disconnect_timeout_in_seconds}.
	// Experimental.
	IdleDisconnectTimeoutInSeconds *float64 `field:"optional" json:"idleDisconnectTimeoutInSeconds" yaml:"idleDisconnectTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_pool#max_user_duration_in_seconds TfPool#max_user_duration_in_seconds}.
	// Experimental.
	MaxUserDurationInSeconds *float64 `field:"optional" json:"maxUserDurationInSeconds" yaml:"maxUserDurationInSeconds"`
}

