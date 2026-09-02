package awsauditmanager


// Experimental.
type TfAssessment_AssessmentReportsDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_assessment#destination TfAssessment#destination}.
	// Experimental.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_assessment#destination_type TfAssessment#destination_type}.
	// Experimental.
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
}

