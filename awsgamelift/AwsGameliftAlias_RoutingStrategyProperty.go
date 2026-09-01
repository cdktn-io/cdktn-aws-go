package awsgamelift


// Experimental.
type AwsGameliftAlias_RoutingStrategyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_alias#type AwsGameliftAlias#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_alias#fleet_id AwsGameliftAlias#fleet_id}.
	// Experimental.
	FleetId *string `field:"optional" json:"fleetId" yaml:"fleetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_alias#message AwsGameliftAlias#message}.
	// Experimental.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

