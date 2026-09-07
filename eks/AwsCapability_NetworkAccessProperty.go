package eks


// Experimental.
type AwsCapability_NetworkAccessProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#vpce_ids AwsCapability#vpce_ids}.
	// Experimental.
	VpceIds *[]*string `field:"optional" json:"vpceIds" yaml:"vpceIds"`
}

