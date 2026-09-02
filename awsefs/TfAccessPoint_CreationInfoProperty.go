package awsefs


// Experimental.
type TfAccessPoint_CreationInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#owner_gid TfAccessPoint#owner_gid}.
	// Experimental.
	OwnerGid *float64 `field:"required" json:"ownerGid" yaml:"ownerGid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#owner_uid TfAccessPoint#owner_uid}.
	// Experimental.
	OwnerUid *float64 `field:"required" json:"ownerUid" yaml:"ownerUid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#permissions TfAccessPoint#permissions}.
	// Experimental.
	Permissions *string `field:"required" json:"permissions" yaml:"permissions"`
}

