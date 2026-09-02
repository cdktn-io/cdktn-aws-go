package awsrolesanywhere


// Experimental.
type TfTrustAnchor_SourceProperty struct {
	// source_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#source_data TfTrustAnchor#source_data}
	// Experimental.
	SourceData *TfTrustAnchor_SourceDataProperty `field:"required" json:"sourceData" yaml:"sourceData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#source_type TfTrustAnchor#source_type}.
	// Experimental.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
}

