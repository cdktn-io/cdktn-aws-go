package awsdirectoryservice


// Experimental.
type TfRegion_VpcSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_region#subnet_ids TfRegion#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_region#vpc_id TfRegion#vpc_id}.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

