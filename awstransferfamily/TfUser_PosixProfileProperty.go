package awstransferfamily


// Experimental.
type TfUser_PosixProfileProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_user#gid TfUser#gid}.
	// Experimental.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_user#uid TfUser#uid}.
	// Experimental.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_user#secondary_gids TfUser#secondary_gids}.
	// Experimental.
	SecondaryGids *[]*float64 `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

