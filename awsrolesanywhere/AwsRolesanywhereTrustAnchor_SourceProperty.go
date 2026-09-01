package awsrolesanywhere


// Experimental.
type AwsRolesanywhereTrustAnchor_SourceProperty struct {
	// source_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#source_data AwsRolesanywhereTrustAnchor#source_data}
	// Experimental.
	SourceData *AwsRolesanywhereTrustAnchor_SourceDataProperty `field:"required" json:"sourceData" yaml:"sourceData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#source_type AwsRolesanywhereTrustAnchor#source_type}.
	// Experimental.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
}

