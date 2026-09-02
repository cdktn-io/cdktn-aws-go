package awscodepipeline


// Experimental.
type TfCodepipeline_ArtifactStoreProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#location TfCodepipeline#location}.
	// Experimental.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#type TfCodepipeline#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#encryption_key TfCodepipeline#encryption_key}
	// Experimental.
	EncryptionKey *TfCodepipeline_EncryptionKeyProperty `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#region TfCodepipeline#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

