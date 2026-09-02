package awstransferfamily


// Experimental.
type TfConnector_As2ConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#compression TfConnector#compression}.
	// Experimental.
	Compression *string `field:"required" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#encryption_algorithm TfConnector#encryption_algorithm}.
	// Experimental.
	EncryptionAlgorithm *string `field:"required" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#local_profile_id TfConnector#local_profile_id}.
	// Experimental.
	LocalProfileId *string `field:"required" json:"localProfileId" yaml:"localProfileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#mdn_response TfConnector#mdn_response}.
	// Experimental.
	MdnResponse *string `field:"required" json:"mdnResponse" yaml:"mdnResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#partner_profile_id TfConnector#partner_profile_id}.
	// Experimental.
	PartnerProfileId *string `field:"required" json:"partnerProfileId" yaml:"partnerProfileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#signing_algorithm TfConnector#signing_algorithm}.
	// Experimental.
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#mdn_signing_algorithm TfConnector#mdn_signing_algorithm}.
	// Experimental.
	MdnSigningAlgorithm *string `field:"optional" json:"mdnSigningAlgorithm" yaml:"mdnSigningAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#message_subject TfConnector#message_subject}.
	// Experimental.
	MessageSubject *string `field:"optional" json:"messageSubject" yaml:"messageSubject"`
}

