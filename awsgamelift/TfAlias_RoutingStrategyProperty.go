package awsgamelift


// Experimental.
type TfAlias_RoutingStrategyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_alias#type TfAlias#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_alias#fleet_id TfAlias#fleet_id}.
	// Experimental.
	FleetId *string `field:"optional" json:"fleetId" yaml:"fleetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_alias#message TfAlias#message}.
	// Experimental.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

