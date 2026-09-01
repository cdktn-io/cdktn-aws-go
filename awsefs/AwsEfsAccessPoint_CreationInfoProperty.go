package awsefs


// Experimental.
type AwsEfsAccessPoint_CreationInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#owner_gid AwsEfsAccessPoint#owner_gid}.
	// Experimental.
	OwnerGid *float64 `field:"required" json:"ownerGid" yaml:"ownerGid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#owner_uid AwsEfsAccessPoint#owner_uid}.
	// Experimental.
	OwnerUid *float64 `field:"required" json:"ownerUid" yaml:"ownerUid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#permissions AwsEfsAccessPoint#permissions}.
	// Experimental.
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
}

