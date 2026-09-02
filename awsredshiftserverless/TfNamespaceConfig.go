package awsredshiftserverless

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNamespaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#namespace_name TfNamespace#namespace_name}.
	// Experimental.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#admin_password_secret_kms_key_id TfNamespace#admin_password_secret_kms_key_id}.
	// Experimental.
	AdminPasswordSecretKmsKeyId *string `field:"optional" json:"adminPasswordSecretKmsKeyId" yaml:"adminPasswordSecretKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#admin_username TfNamespace#admin_username}.
	// Experimental.
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#admin_user_password TfNamespace#admin_user_password}.
	// Experimental.
	AdminUserPassword *string `field:"optional" json:"adminUserPassword" yaml:"adminUserPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#admin_user_password_wo TfNamespace#admin_user_password_wo}.
	// Experimental.
	AdminUserPasswordWo *string `field:"optional" json:"adminUserPasswordWo" yaml:"adminUserPasswordWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#admin_user_password_wo_version TfNamespace#admin_user_password_wo_version}.
	// Experimental.
	AdminUserPasswordWoVersion *float64 `field:"optional" json:"adminUserPasswordWoVersion" yaml:"adminUserPasswordWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#db_name TfNamespace#db_name}.
	// Experimental.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#default_iam_role_arn TfNamespace#default_iam_role_arn}.
	// Experimental.
	DefaultIamRoleArn *string `field:"optional" json:"defaultIamRoleArn" yaml:"defaultIamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#iam_roles TfNamespace#iam_roles}.
	// Experimental.
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#id TfNamespace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#kms_key_id TfNamespace#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#log_exports TfNamespace#log_exports}.
	// Experimental.
	LogExports *[]*string `field:"optional" json:"logExports" yaml:"logExports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#manage_admin_password TfNamespace#manage_admin_password}.
	// Experimental.
	ManageAdminPassword interface{} `field:"optional" json:"manageAdminPassword" yaml:"manageAdminPassword"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#region TfNamespace#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#tags TfNamespace#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_namespace#tags_all TfNamespace#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

