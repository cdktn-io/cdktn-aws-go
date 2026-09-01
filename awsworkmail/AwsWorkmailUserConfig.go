package awsworkmail

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWorkmailUserConfig struct {
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
	// Display name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#display_name AwsWorkmailUser#display_name}
	// Experimental.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Primary email address used to register the user with WorkMail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#email AwsWorkmailUser#email}
	// Experimental.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Username of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#name AwsWorkmailUser#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Identifier of the WorkMail organization where the user is managed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#organization_id AwsWorkmailUser#organization_id}
	// Experimental.
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// City where the user is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#city AwsWorkmailUser#city}
	// Experimental.
	City *string `field:"optional" json:"city" yaml:"city"`
	// Company associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#company AwsWorkmailUser#company}
	// Experimental.
	Company *string `field:"optional" json:"company" yaml:"company"`
	// Country where the user is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#country AwsWorkmailUser#country}
	// Experimental.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Department associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#department AwsWorkmailUser#department}
	// Experimental.
	Department *string `field:"optional" json:"department" yaml:"department"`
	// First name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#first_name AwsWorkmailUser#first_name}
	// Experimental.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Whether to hide the user from the global address list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#hidden_from_global_address_list AwsWorkmailUser#hidden_from_global_address_list}
	// Experimental.
	HiddenFromGlobalAddressList interface{} `field:"optional" json:"hiddenFromGlobalAddressList" yaml:"hiddenFromGlobalAddressList"`
	// User ID from IAM Identity Center associated with the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#identity_provider_user_id AwsWorkmailUser#identity_provider_user_id}
	// Experimental.
	IdentityProviderUserId *string `field:"optional" json:"identityProviderUserId" yaml:"identityProviderUserId"`
	// Initials of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#initials AwsWorkmailUser#initials}
	// Experimental.
	Initials *string `field:"optional" json:"initials" yaml:"initials"`
	// Job title of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#job_title AwsWorkmailUser#job_title}
	// Experimental.
	JobTitle *string `field:"optional" json:"jobTitle" yaml:"jobTitle"`
	// Last name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#last_name AwsWorkmailUser#last_name}
	// Experimental.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// Office where the user is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#office AwsWorkmailUser#office}
	// Experimental.
	Office *string `field:"optional" json:"office" yaml:"office"`
	// Password to set for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#password AwsWorkmailUser#password}
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#region AwsWorkmailUser#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Street address of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#street AwsWorkmailUser#street}
	// Experimental.
	Street *string `field:"optional" json:"street" yaml:"street"`
	// Telephone number of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#telephone AwsWorkmailUser#telephone}
	// Experimental.
	Telephone *string `field:"optional" json:"telephone" yaml:"telephone"`
	// Role assigned to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#user_role AwsWorkmailUser#user_role}
	// Experimental.
	UserRole *string `field:"optional" json:"userRole" yaml:"userRole"`
	// ZIP or postal code of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workmail_user#zip_code AwsWorkmailUser#zip_code}
	// Experimental.
	ZipCode *string `field:"optional" json:"zipCode" yaml:"zipCode"`
}

