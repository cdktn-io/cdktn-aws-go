package awsgamelift


// Experimental.
type TfFleet_ResourceCreationLimitPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#new_game_sessions_per_creator TfFleet#new_game_sessions_per_creator}.
	// Experimental.
	NewGameSessionsPerCreator *float64 `field:"optional" json:"newGameSessionsPerCreator" yaml:"newGameSessionsPerCreator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_fleet#policy_period_in_minutes TfFleet#policy_period_in_minutes}.
	// Experimental.
	PolicyPeriodInMinutes *float64 `field:"optional" json:"policyPeriodInMinutes" yaml:"policyPeriodInMinutes"`
}

