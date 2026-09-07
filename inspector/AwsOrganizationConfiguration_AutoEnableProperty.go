package inspector


// Experimental.
type AwsOrganizationConfiguration_AutoEnableProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_organization_configuration#ec2 AwsOrganizationConfiguration#ec2}.
	// Experimental.
	Ec2 interface{} `field:"required" json:"ec2" yaml:"ec2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_organization_configuration#ecr AwsOrganizationConfiguration#ecr}.
	// Experimental.
	Ecr interface{} `field:"required" json:"ecr" yaml:"ecr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_organization_configuration#code_repository AwsOrganizationConfiguration#code_repository}.
	// Experimental.
	CodeRepository interface{} `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_organization_configuration#lambda AwsOrganizationConfiguration#lambda}.
	// Experimental.
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_organization_configuration#lambda_code AwsOrganizationConfiguration#lambda_code}.
	// Experimental.
	LambdaCode interface{} `field:"optional" json:"lambdaCode" yaml:"lambdaCode"`
}

