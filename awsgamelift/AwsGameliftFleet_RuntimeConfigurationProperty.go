package awsgamelift


// Experimental.
type AwsGameliftFleet_RuntimeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#game_session_activation_timeout_seconds AwsGameliftFleet#game_session_activation_timeout_seconds}.
	// Experimental.
	GameSessionActivationTimeoutSeconds *float64 `field:"optional" json:"gameSessionActivationTimeoutSeconds" yaml:"gameSessionActivationTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#max_concurrent_game_session_activations AwsGameliftFleet#max_concurrent_game_session_activations}.
	// Experimental.
	MaxConcurrentGameSessionActivations *float64 `field:"optional" json:"maxConcurrentGameSessionActivations" yaml:"maxConcurrentGameSessionActivations"`
	// server_process block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#server_process AwsGameliftFleet#server_process}
	// Experimental.
	ServerProcess interface{} `field:"optional" json:"serverProcess" yaml:"serverProcess"`
}

