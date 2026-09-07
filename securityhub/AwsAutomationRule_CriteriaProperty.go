package securityhub


// Experimental.
type AwsAutomationRule_CriteriaProperty struct {
	// aws_account_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#aws_account_id AwsAutomationRule#aws_account_id}
	// Experimental.
	AwsAccountId interface{} `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// aws_account_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#aws_account_name AwsAutomationRule#aws_account_name}
	// Experimental.
	AwsAccountName interface{} `field:"optional" json:"awsAccountName" yaml:"awsAccountName"`
	// company_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#company_name AwsAutomationRule#company_name}
	// Experimental.
	CompanyName interface{} `field:"optional" json:"companyName" yaml:"companyName"`
	// compliance_associated_standards_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#compliance_associated_standards_id AwsAutomationRule#compliance_associated_standards_id}
	// Experimental.
	ComplianceAssociatedStandardsId interface{} `field:"optional" json:"complianceAssociatedStandardsId" yaml:"complianceAssociatedStandardsId"`
	// compliance_security_control_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#compliance_security_control_id AwsAutomationRule#compliance_security_control_id}
	// Experimental.
	ComplianceSecurityControlId interface{} `field:"optional" json:"complianceSecurityControlId" yaml:"complianceSecurityControlId"`
	// compliance_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#compliance_status AwsAutomationRule#compliance_status}
	// Experimental.
	ComplianceStatus interface{} `field:"optional" json:"complianceStatus" yaml:"complianceStatus"`
	// confidence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#confidence AwsAutomationRule#confidence}
	// Experimental.
	Confidence interface{} `field:"optional" json:"confidence" yaml:"confidence"`
	// created_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#created_at AwsAutomationRule#created_at}
	// Experimental.
	CreatedAt interface{} `field:"optional" json:"createdAt" yaml:"createdAt"`
	// criticality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#criticality AwsAutomationRule#criticality}
	// Experimental.
	Criticality interface{} `field:"optional" json:"criticality" yaml:"criticality"`
	// description block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#description AwsAutomationRule#description}
	// Experimental.
	Description interface{} `field:"optional" json:"description" yaml:"description"`
	// first_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#first_observed_at AwsAutomationRule#first_observed_at}
	// Experimental.
	FirstObservedAt interface{} `field:"optional" json:"firstObservedAt" yaml:"firstObservedAt"`
	// generator_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#generator_id AwsAutomationRule#generator_id}
	// Experimental.
	GeneratorId interface{} `field:"optional" json:"generatorId" yaml:"generatorId"`
	// id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#id AwsAutomationRule#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id interface{} `field:"optional" json:"id" yaml:"id"`
	// last_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#last_observed_at AwsAutomationRule#last_observed_at}
	// Experimental.
	LastObservedAt interface{} `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// note_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#note_text AwsAutomationRule#note_text}
	// Experimental.
	NoteText interface{} `field:"optional" json:"noteText" yaml:"noteText"`
	// note_updated_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#note_updated_at AwsAutomationRule#note_updated_at}
	// Experimental.
	NoteUpdatedAt interface{} `field:"optional" json:"noteUpdatedAt" yaml:"noteUpdatedAt"`
	// note_updated_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#note_updated_by AwsAutomationRule#note_updated_by}
	// Experimental.
	NoteUpdatedBy interface{} `field:"optional" json:"noteUpdatedBy" yaml:"noteUpdatedBy"`
	// product_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#product_arn AwsAutomationRule#product_arn}
	// Experimental.
	ProductArn interface{} `field:"optional" json:"productArn" yaml:"productArn"`
	// product_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#product_name AwsAutomationRule#product_name}
	// Experimental.
	ProductName interface{} `field:"optional" json:"productName" yaml:"productName"`
	// record_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#record_state AwsAutomationRule#record_state}
	// Experimental.
	RecordState interface{} `field:"optional" json:"recordState" yaml:"recordState"`
	// related_findings_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#related_findings_id AwsAutomationRule#related_findings_id}
	// Experimental.
	RelatedFindingsId interface{} `field:"optional" json:"relatedFindingsId" yaml:"relatedFindingsId"`
	// related_findings_product_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#related_findings_product_arn AwsAutomationRule#related_findings_product_arn}
	// Experimental.
	RelatedFindingsProductArn interface{} `field:"optional" json:"relatedFindingsProductArn" yaml:"relatedFindingsProductArn"`
	// resource_application_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#resource_application_arn AwsAutomationRule#resource_application_arn}
	// Experimental.
	ResourceApplicationArn interface{} `field:"optional" json:"resourceApplicationArn" yaml:"resourceApplicationArn"`
	// resource_application_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#resource_application_name AwsAutomationRule#resource_application_name}
	// Experimental.
	ResourceApplicationName interface{} `field:"optional" json:"resourceApplicationName" yaml:"resourceApplicationName"`
	// resource_details_other block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#resource_details_other AwsAutomationRule#resource_details_other}
	// Experimental.
	ResourceDetailsOther interface{} `field:"optional" json:"resourceDetailsOther" yaml:"resourceDetailsOther"`
	// resource_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#resource_id AwsAutomationRule#resource_id}
	// Experimental.
	ResourceId interface{} `field:"optional" json:"resourceId" yaml:"resourceId"`
	// resource_partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#resource_partition AwsAutomationRule#resource_partition}
	// Experimental.
	ResourcePartition interface{} `field:"optional" json:"resourcePartition" yaml:"resourcePartition"`
	// resource_region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#resource_region AwsAutomationRule#resource_region}
	// Experimental.
	ResourceRegion interface{} `field:"optional" json:"resourceRegion" yaml:"resourceRegion"`
	// resource_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#resource_tags AwsAutomationRule#resource_tags}
	// Experimental.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#resource_type AwsAutomationRule#resource_type}
	// Experimental.
	ResourceType interface{} `field:"optional" json:"resourceType" yaml:"resourceType"`
	// severity_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#severity_label AwsAutomationRule#severity_label}
	// Experimental.
	SeverityLabel interface{} `field:"optional" json:"severityLabel" yaml:"severityLabel"`
	// source_url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#source_url AwsAutomationRule#source_url}
	// Experimental.
	SourceUrl interface{} `field:"optional" json:"sourceUrl" yaml:"sourceUrl"`
	// title block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#title AwsAutomationRule#title}
	// Experimental.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#type AwsAutomationRule#type}
	// Experimental.
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	// updated_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#updated_at AwsAutomationRule#updated_at}
	// Experimental.
	UpdatedAt interface{} `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// user_defined_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#user_defined_fields AwsAutomationRule#user_defined_fields}
	// Experimental.
	UserDefinedFields interface{} `field:"optional" json:"userDefinedFields" yaml:"userDefinedFields"`
	// verification_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#verification_state AwsAutomationRule#verification_state}
	// Experimental.
	VerificationState interface{} `field:"optional" json:"verificationState" yaml:"verificationState"`
	// workflow_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_automation_rule#workflow_status AwsAutomationRule#workflow_status}
	// Experimental.
	WorkflowStatus interface{} `field:"optional" json:"workflowStatus" yaml:"workflowStatus"`
}

