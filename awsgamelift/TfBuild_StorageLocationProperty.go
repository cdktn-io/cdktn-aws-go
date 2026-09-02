package awsgamelift


// Experimental.
type TfBuild_StorageLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_build#bucket TfBuild#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_build#key TfBuild#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_build#role_arn TfBuild#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_build#object_version TfBuild#object_version}.
	// Experimental.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

