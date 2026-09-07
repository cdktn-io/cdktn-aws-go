package verifiedpermissions


// Experimental.
type AwsPolicy_ResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#entity_id AwsPolicy#entity_id}.
	// Experimental.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#entity_type AwsPolicy#entity_type}.
	// Experimental.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
}

