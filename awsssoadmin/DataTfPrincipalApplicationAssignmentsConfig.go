package awsssoadmin

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfPrincipalApplicationAssignmentsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ssoadmin_principal_application_assignments#instance_arn DataTfPrincipalApplicationAssignments#instance_arn}.
	// Experimental.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ssoadmin_principal_application_assignments#principal_id DataTfPrincipalApplicationAssignments#principal_id}.
	// Experimental.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ssoadmin_principal_application_assignments#principal_type DataTfPrincipalApplicationAssignments#principal_type}.
	// Experimental.
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// application_assignments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ssoadmin_principal_application_assignments#application_assignments DataTfPrincipalApplicationAssignments#application_assignments}
	// Experimental.
	ApplicationAssignments interface{} `field:"optional" json:"applicationAssignments" yaml:"applicationAssignments"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ssoadmin_principal_application_assignments#region DataTfPrincipalApplicationAssignments#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

