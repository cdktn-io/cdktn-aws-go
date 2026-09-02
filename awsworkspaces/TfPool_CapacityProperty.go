package awsworkspaces


// Experimental.
type TfPool_CapacityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_pool#desired_user_sessions TfPool#desired_user_sessions}.
	// Experimental.
	DesiredUserSessions *float64 `field:"required" json:"desiredUserSessions" yaml:"desiredUserSessions"`
}

