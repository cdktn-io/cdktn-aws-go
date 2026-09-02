package awsapprunner


// Experimental.
type TfService_AuthenticationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#access_role_arn TfService#access_role_arn}.
	// Experimental.
	AccessRoleArn *string `field:"optional" json:"accessRoleArn" yaml:"accessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#connection_arn TfService#connection_arn}.
	// Experimental.
	ConnectionArn *string `field:"optional" json:"connectionArn" yaml:"connectionArn"`
}

