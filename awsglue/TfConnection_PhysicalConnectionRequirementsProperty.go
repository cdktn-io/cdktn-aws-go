package awsglue


// Experimental.
type TfConnection_PhysicalConnectionRequirementsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#availability_zone TfConnection#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#security_group_id_list TfConnection#security_group_id_list}.
	// Experimental.
	SecurityGroupIdList *[]*string `field:"optional" json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_connection#subnet_id TfConnection#subnet_id}.
	// Experimental.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

