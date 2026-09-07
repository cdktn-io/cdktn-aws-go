package cleanrooms


// Experimental.
type AwsMembership_DefaultResultConfigurationProperty struct {
	// output_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_membership#output_configuration AwsMembership#output_configuration}
	// Experimental.
	OutputConfiguration interface{} `field:"optional" json:"outputConfiguration" yaml:"outputConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cleanrooms_membership#role_arn AwsMembership#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

