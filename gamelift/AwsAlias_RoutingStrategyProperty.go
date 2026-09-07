package gamelift


// Experimental.
type AwsAlias_RoutingStrategyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_alias#type AwsAlias#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_alias#fleet_id AwsAlias#fleet_id}.
	// Experimental.
	FleetId *string `field:"optional" json:"fleetId" yaml:"fleetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_alias#message AwsAlias#message}.
	// Experimental.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

