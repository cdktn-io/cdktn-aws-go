package awsverifiedpermissions


// Experimental.
type TfPolicy_ResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#entity_id TfPolicy#entity_id}.
	// Experimental.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedpermissions_policy#entity_type TfPolicy#entity_type}.
	// Experimental.
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
}

