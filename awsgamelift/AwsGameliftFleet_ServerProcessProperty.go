package awsgamelift


// Experimental.
type AwsGameliftFleet_ServerProcessProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#concurrent_executions AwsGameliftFleet#concurrent_executions}.
	// Experimental.
	ConcurrentExecutions *float64 `field:"required" json:"concurrentExecutions" yaml:"concurrentExecutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#launch_path AwsGameliftFleet#launch_path}.
	// Experimental.
	LaunchPath *string `field:"required" json:"launchPath" yaml:"launchPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#parameters AwsGameliftFleet#parameters}.
	// Experimental.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

