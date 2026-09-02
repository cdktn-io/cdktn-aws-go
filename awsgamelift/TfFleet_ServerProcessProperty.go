package awsgamelift


// Experimental.
type TfFleet_ServerProcessProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#concurrent_executions TfFleet#concurrent_executions}.
	// Experimental.
	ConcurrentExecutions *float64 `field:"required" json:"concurrentExecutions" yaml:"concurrentExecutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#launch_path TfFleet#launch_path}.
	// Experimental.
	LaunchPath *string `field:"required" json:"launchPath" yaml:"launchPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#parameters TfFleet#parameters}.
	// Experimental.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

