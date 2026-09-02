package awsefs


// Experimental.
type TfAccessPoint_PosixUserProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#gid TfAccessPoint#gid}.
	// Experimental.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#uid TfAccessPoint#uid}.
	// Experimental.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#secondary_gids TfAccessPoint#secondary_gids}.
	// Experimental.
	SecondaryGids *[]*float64 `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

