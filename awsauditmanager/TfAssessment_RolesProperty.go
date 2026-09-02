package awsauditmanager


// Experimental.
type TfAssessment_RolesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_assessment#role_arn TfAssessment#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_assessment#role_type TfAssessment#role_type}.
	// Experimental.
	RoleType *string `field:"required" json:"roleType" yaml:"roleType"`
}

