package codebuild


// Experimental.
type AwsProject_CacheProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#cache_namespace AwsProject#cache_namespace}.
	// Experimental.
	CacheNamespace *string `field:"optional" json:"cacheNamespace" yaml:"cacheNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#location AwsProject#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#modes AwsProject#modes}.
	// Experimental.
	Modes *[]*string `field:"optional" json:"modes" yaml:"modes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#type AwsProject#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

