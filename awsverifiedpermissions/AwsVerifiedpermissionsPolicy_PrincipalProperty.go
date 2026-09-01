package awsverifiedpermissions


// Experimental.
type AwsVerifiedpermissionsPolicy_PrincipalProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#entity_id AwsVerifiedpermissionsPolicy#entity_id}.
	// Experimental.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#entity_type AwsVerifiedpermissionsPolicy#entity_type}.
	// Experimental.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
}

