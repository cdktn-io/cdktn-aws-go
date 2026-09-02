package awscodebuild


// Experimental.
type TfProject_SecondaryArtifactsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#artifact_identifier TfProject#artifact_identifier}.
	// Experimental.
	ArtifactIdentifier *string `field:"required" json:"artifactIdentifier" yaml:"artifactIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#type TfProject#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#bucket_owner_access TfProject#bucket_owner_access}.
	// Experimental.
	BucketOwnerAccess *string `field:"optional" json:"bucketOwnerAccess" yaml:"bucketOwnerAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#encryption_disabled TfProject#encryption_disabled}.
	// Experimental.
	EncryptionDisabled interface{} `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#location TfProject#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#name TfProject#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#namespace_type TfProject#namespace_type}.
	// Experimental.
	NamespaceType *string `field:"optional" json:"namespaceType" yaml:"namespaceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#override_artifact_name TfProject#override_artifact_name}.
	// Experimental.
	OverrideArtifactName interface{} `field:"optional" json:"overrideArtifactName" yaml:"overrideArtifactName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#packaging TfProject#packaging}.
	// Experimental.
	Packaging *string `field:"optional" json:"packaging" yaml:"packaging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#path TfProject#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

