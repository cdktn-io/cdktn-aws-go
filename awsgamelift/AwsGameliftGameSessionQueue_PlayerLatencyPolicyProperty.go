package awsgamelift


// Experimental.
type AwsGameliftGameSessionQueue_PlayerLatencyPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_session_queue#maximum_individual_player_latency_milliseconds AwsGameliftGameSessionQueue#maximum_individual_player_latency_milliseconds}.
	// Experimental.
	MaximumIndividualPlayerLatencyMilliseconds *float64 `field:"required" json:"maximumIndividualPlayerLatencyMilliseconds" yaml:"maximumIndividualPlayerLatencyMilliseconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_session_queue#policy_duration_seconds AwsGameliftGameSessionQueue#policy_duration_seconds}.
	// Experimental.
	PolicyDurationSeconds *float64 `field:"optional" json:"policyDurationSeconds" yaml:"policyDurationSeconds"`
}

