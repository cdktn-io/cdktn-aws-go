package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/securityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/securityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAutomationRule_CriteriaPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsAccountId() AwsAutomationRule_AwsAccountIdPropertyList
	// Experimental.
	AwsAccountIdInput() interface{}
	// Experimental.
	AwsAccountName() AwsAutomationRule_AwsAccountNamePropertyList
	// Experimental.
	AwsAccountNameInput() interface{}
	// Experimental.
	CompanyName() AwsAutomationRule_CompanyNamePropertyList
	// Experimental.
	CompanyNameInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// Experimental.
	ComplianceAssociatedStandardsId() AwsAutomationRule_ComplianceAssociatedStandardsIdPropertyList
	// Experimental.
	ComplianceAssociatedStandardsIdInput() interface{}
	// Experimental.
	ComplianceSecurityControlId() AwsAutomationRule_ComplianceSecurityControlIdPropertyList
	// Experimental.
	ComplianceSecurityControlIdInput() interface{}
	// Experimental.
	ComplianceStatus() AwsAutomationRule_ComplianceStatusPropertyList
	// Experimental.
	ComplianceStatusInput() interface{}
	// Experimental.
	Confidence() AwsAutomationRule_ConfidencePropertyList
	// Experimental.
	ConfidenceInput() interface{}
	// Experimental.
	CreatedAt() AwsAutomationRule_CreatedAtPropertyList
	// Experimental.
	CreatedAtInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Criticality() AwsAutomationRule_CriticalityPropertyList
	// Experimental.
	CriticalityInput() interface{}
	// Experimental.
	Description() AwsAutomationRule_DescriptionPropertyList
	// Experimental.
	DescriptionInput() interface{}
	// Experimental.
	FirstObservedAt() AwsAutomationRule_FirstObservedAtPropertyList
	// Experimental.
	FirstObservedAtInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	GeneratorId() AwsAutomationRule_GeneratorIdPropertyList
	// Experimental.
	GeneratorIdInput() interface{}
	// Experimental.
	Id() AwsAutomationRule_IdPropertyList
	// Experimental.
	IdInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LastObservedAt() AwsAutomationRule_LastObservedAtPropertyList
	// Experimental.
	LastObservedAtInput() interface{}
	// Experimental.
	NoteText() AwsAutomationRule_NoteTextPropertyList
	// Experimental.
	NoteTextInput() interface{}
	// Experimental.
	NoteUpdatedAt() AwsAutomationRule_NoteUpdatedAtPropertyList
	// Experimental.
	NoteUpdatedAtInput() interface{}
	// Experimental.
	NoteUpdatedBy() AwsAutomationRule_NoteUpdatedByPropertyList
	// Experimental.
	NoteUpdatedByInput() interface{}
	// Experimental.
	ProductArn() AwsAutomationRule_ProductArnPropertyList
	// Experimental.
	ProductArnInput() interface{}
	// Experimental.
	ProductName() AwsAutomationRule_ProductNamePropertyList
	// Experimental.
	ProductNameInput() interface{}
	// Experimental.
	RecordState() AwsAutomationRule_RecordStatePropertyList
	// Experimental.
	RecordStateInput() interface{}
	// Experimental.
	RelatedFindingsId() AwsAutomationRule_RelatedFindingsIdPropertyList
	// Experimental.
	RelatedFindingsIdInput() interface{}
	// Experimental.
	RelatedFindingsProductArn() AwsAutomationRule_RelatedFindingsProductArnPropertyList
	// Experimental.
	RelatedFindingsProductArnInput() interface{}
	// Experimental.
	ResourceApplicationArn() AwsAutomationRule_ResourceApplicationArnPropertyList
	// Experimental.
	ResourceApplicationArnInput() interface{}
	// Experimental.
	ResourceApplicationName() AwsAutomationRule_ResourceApplicationNamePropertyList
	// Experimental.
	ResourceApplicationNameInput() interface{}
	// Experimental.
	ResourceDetailsOther() AwsAutomationRule_ResourceDetailsOtherPropertyList
	// Experimental.
	ResourceDetailsOtherInput() interface{}
	// Experimental.
	ResourceId() AwsAutomationRule_ResourceIdPropertyList
	// Experimental.
	ResourceIdInput() interface{}
	// Experimental.
	ResourcePartition() AwsAutomationRule_ResourcePartitionPropertyList
	// Experimental.
	ResourcePartitionInput() interface{}
	// Experimental.
	ResourceRegion() AwsAutomationRule_ResourceRegionPropertyList
	// Experimental.
	ResourceRegionInput() interface{}
	// Experimental.
	ResourceTags() AwsAutomationRule_ResourceTagsPropertyList
	// Experimental.
	ResourceTagsInput() interface{}
	// Experimental.
	ResourceType() AwsAutomationRule_ResourceTypePropertyList
	// Experimental.
	ResourceTypeInput() interface{}
	// Experimental.
	SeverityLabel() AwsAutomationRule_SeverityLabelPropertyList
	// Experimental.
	SeverityLabelInput() interface{}
	// Experimental.
	SourceUrl() AwsAutomationRule_SourceUrlPropertyList
	// Experimental.
	SourceUrlInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Title() AwsAutomationRule_TitlePropertyList
	// Experimental.
	TitleInput() interface{}
	// Experimental.
	Type() AwsAutomationRule_TypePropertyList
	// Experimental.
	TypeInput() interface{}
	// Experimental.
	UpdatedAt() AwsAutomationRule_UpdatedAtPropertyList
	// Experimental.
	UpdatedAtInput() interface{}
	// Experimental.
	UserDefinedFields() AwsAutomationRule_UserDefinedFieldsPropertyList
	// Experimental.
	UserDefinedFieldsInput() interface{}
	// Experimental.
	VerificationState() AwsAutomationRule_VerificationStatePropertyList
	// Experimental.
	VerificationStateInput() interface{}
	// Experimental.
	WorkflowStatus() AwsAutomationRule_WorkflowStatusPropertyList
	// Experimental.
	WorkflowStatusInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutAwsAccountId(value interface{})
	// Experimental.
	PutAwsAccountName(value interface{})
	// Experimental.
	PutCompanyName(value interface{})
	// Experimental.
	PutComplianceAssociatedStandardsId(value interface{})
	// Experimental.
	PutComplianceSecurityControlId(value interface{})
	// Experimental.
	PutComplianceStatus(value interface{})
	// Experimental.
	PutConfidence(value interface{})
	// Experimental.
	PutCreatedAt(value interface{})
	// Experimental.
	PutCriticality(value interface{})
	// Experimental.
	PutDescription(value interface{})
	// Experimental.
	PutFirstObservedAt(value interface{})
	// Experimental.
	PutGeneratorId(value interface{})
	// Experimental.
	PutId(value interface{})
	// Experimental.
	PutLastObservedAt(value interface{})
	// Experimental.
	PutNoteText(value interface{})
	// Experimental.
	PutNoteUpdatedAt(value interface{})
	// Experimental.
	PutNoteUpdatedBy(value interface{})
	// Experimental.
	PutProductArn(value interface{})
	// Experimental.
	PutProductName(value interface{})
	// Experimental.
	PutRecordState(value interface{})
	// Experimental.
	PutRelatedFindingsId(value interface{})
	// Experimental.
	PutRelatedFindingsProductArn(value interface{})
	// Experimental.
	PutResourceApplicationArn(value interface{})
	// Experimental.
	PutResourceApplicationName(value interface{})
	// Experimental.
	PutResourceDetailsOther(value interface{})
	// Experimental.
	PutResourceId(value interface{})
	// Experimental.
	PutResourcePartition(value interface{})
	// Experimental.
	PutResourceRegion(value interface{})
	// Experimental.
	PutResourceTags(value interface{})
	// Experimental.
	PutResourceType(value interface{})
	// Experimental.
	PutSeverityLabel(value interface{})
	// Experimental.
	PutSourceUrl(value interface{})
	// Experimental.
	PutTitle(value interface{})
	// Experimental.
	PutType(value interface{})
	// Experimental.
	PutUpdatedAt(value interface{})
	// Experimental.
	PutUserDefinedFields(value interface{})
	// Experimental.
	PutVerificationState(value interface{})
	// Experimental.
	PutWorkflowStatus(value interface{})
	// Experimental.
	ResetAwsAccountId()
	// Experimental.
	ResetAwsAccountName()
	// Experimental.
	ResetCompanyName()
	// Experimental.
	ResetComplianceAssociatedStandardsId()
	// Experimental.
	ResetComplianceSecurityControlId()
	// Experimental.
	ResetComplianceStatus()
	// Experimental.
	ResetConfidence()
	// Experimental.
	ResetCreatedAt()
	// Experimental.
	ResetCriticality()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetFirstObservedAt()
	// Experimental.
	ResetGeneratorId()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLastObservedAt()
	// Experimental.
	ResetNoteText()
	// Experimental.
	ResetNoteUpdatedAt()
	// Experimental.
	ResetNoteUpdatedBy()
	// Experimental.
	ResetProductArn()
	// Experimental.
	ResetProductName()
	// Experimental.
	ResetRecordState()
	// Experimental.
	ResetRelatedFindingsId()
	// Experimental.
	ResetRelatedFindingsProductArn()
	// Experimental.
	ResetResourceApplicationArn()
	// Experimental.
	ResetResourceApplicationName()
	// Experimental.
	ResetResourceDetailsOther()
	// Experimental.
	ResetResourceId()
	// Experimental.
	ResetResourcePartition()
	// Experimental.
	ResetResourceRegion()
	// Experimental.
	ResetResourceTags()
	// Experimental.
	ResetResourceType()
	// Experimental.
	ResetSeverityLabel()
	// Experimental.
	ResetSourceUrl()
	// Experimental.
	ResetTitle()
	// Experimental.
	ResetType()
	// Experimental.
	ResetUpdatedAt()
	// Experimental.
	ResetUserDefinedFields()
	// Experimental.
	ResetVerificationState()
	// Experimental.
	ResetWorkflowStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAutomationRule_CriteriaPropertyOutputReference
type jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) AwsAccountId() AwsAutomationRule_AwsAccountIdPropertyList {
	var returns AwsAutomationRule_AwsAccountIdPropertyList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) AwsAccountName() AwsAutomationRule_AwsAccountNamePropertyList {
	var returns AwsAutomationRule_AwsAccountNamePropertyList
	_jsii_.Get(
		j,
		"awsAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) AwsAccountNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) CompanyName() AwsAutomationRule_CompanyNamePropertyList {
	var returns AwsAutomationRule_CompanyNamePropertyList
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) CompanyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ComplianceAssociatedStandardsId() AwsAutomationRule_ComplianceAssociatedStandardsIdPropertyList {
	var returns AwsAutomationRule_ComplianceAssociatedStandardsIdPropertyList
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ComplianceAssociatedStandardsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ComplianceSecurityControlId() AwsAutomationRule_ComplianceSecurityControlIdPropertyList {
	var returns AwsAutomationRule_ComplianceSecurityControlIdPropertyList
	_jsii_.Get(
		j,
		"complianceSecurityControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ComplianceSecurityControlIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityControlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ComplianceStatus() AwsAutomationRule_ComplianceStatusPropertyList {
	var returns AwsAutomationRule_ComplianceStatusPropertyList
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ComplianceStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) Confidence() AwsAutomationRule_ConfidencePropertyList {
	var returns AwsAutomationRule_ConfidencePropertyList
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) CreatedAt() AwsAutomationRule_CreatedAtPropertyList {
	var returns AwsAutomationRule_CreatedAtPropertyList
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) CreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) Criticality() AwsAutomationRule_CriticalityPropertyList {
	var returns AwsAutomationRule_CriticalityPropertyList
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) CriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) Description() AwsAutomationRule_DescriptionPropertyList {
	var returns AwsAutomationRule_DescriptionPropertyList
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) DescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) FirstObservedAt() AwsAutomationRule_FirstObservedAtPropertyList {
	var returns AwsAutomationRule_FirstObservedAtPropertyList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) GeneratorId() AwsAutomationRule_GeneratorIdPropertyList {
	var returns AwsAutomationRule_GeneratorIdPropertyList
	_jsii_.Get(
		j,
		"generatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) GeneratorIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) Id() AwsAutomationRule_IdPropertyList {
	var returns AwsAutomationRule_IdPropertyList
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) IdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) LastObservedAt() AwsAutomationRule_LastObservedAtPropertyList {
	var returns AwsAutomationRule_LastObservedAtPropertyList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) NoteText() AwsAutomationRule_NoteTextPropertyList {
	var returns AwsAutomationRule_NoteTextPropertyList
	_jsii_.Get(
		j,
		"noteText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) NoteTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedAt() AwsAutomationRule_NoteUpdatedAtPropertyList {
	var returns AwsAutomationRule_NoteUpdatedAtPropertyList
	_jsii_.Get(
		j,
		"noteUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedBy() AwsAutomationRule_NoteUpdatedByPropertyList {
	var returns AwsAutomationRule_NoteUpdatedByPropertyList
	_jsii_.Get(
		j,
		"noteUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ProductArn() AwsAutomationRule_ProductArnPropertyList {
	var returns AwsAutomationRule_ProductArnPropertyList
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ProductName() AwsAutomationRule_ProductNamePropertyList {
	var returns AwsAutomationRule_ProductNamePropertyList
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ProductNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) RecordState() AwsAutomationRule_RecordStatePropertyList {
	var returns AwsAutomationRule_RecordStatePropertyList
	_jsii_.Get(
		j,
		"recordState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) RecordStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsId() AwsAutomationRule_RelatedFindingsIdPropertyList {
	var returns AwsAutomationRule_RelatedFindingsIdPropertyList
	_jsii_.Get(
		j,
		"relatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsProductArn() AwsAutomationRule_RelatedFindingsProductArnPropertyList {
	var returns AwsAutomationRule_RelatedFindingsProductArnPropertyList
	_jsii_.Get(
		j,
		"relatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationArn() AwsAutomationRule_ResourceApplicationArnPropertyList {
	var returns AwsAutomationRule_ResourceApplicationArnPropertyList
	_jsii_.Get(
		j,
		"resourceApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceApplicationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationName() AwsAutomationRule_ResourceApplicationNamePropertyList {
	var returns AwsAutomationRule_ResourceApplicationNamePropertyList
	_jsii_.Get(
		j,
		"resourceApplicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceApplicationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceDetailsOther() AwsAutomationRule_ResourceDetailsOtherPropertyList {
	var returns AwsAutomationRule_ResourceDetailsOtherPropertyList
	_jsii_.Get(
		j,
		"resourceDetailsOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceDetailsOtherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceDetailsOtherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceId() AwsAutomationRule_ResourceIdPropertyList {
	var returns AwsAutomationRule_ResourceIdPropertyList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourcePartition() AwsAutomationRule_ResourcePartitionPropertyList {
	var returns AwsAutomationRule_ResourcePartitionPropertyList
	_jsii_.Get(
		j,
		"resourcePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourcePartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceRegion() AwsAutomationRule_ResourceRegionPropertyList {
	var returns AwsAutomationRule_ResourceRegionPropertyList
	_jsii_.Get(
		j,
		"resourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceTags() AwsAutomationRule_ResourceTagsPropertyList {
	var returns AwsAutomationRule_ResourceTagsPropertyList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceType() AwsAutomationRule_ResourceTypePropertyList {
	var returns AwsAutomationRule_ResourceTypePropertyList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) SeverityLabel() AwsAutomationRule_SeverityLabelPropertyList {
	var returns AwsAutomationRule_SeverityLabelPropertyList
	_jsii_.Get(
		j,
		"severityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) SeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) SourceUrl() AwsAutomationRule_SourceUrlPropertyList {
	var returns AwsAutomationRule_SourceUrlPropertyList
	_jsii_.Get(
		j,
		"sourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) SourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) Title() AwsAutomationRule_TitlePropertyList {
	var returns AwsAutomationRule_TitlePropertyList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) Type() AwsAutomationRule_TypePropertyList {
	var returns AwsAutomationRule_TypePropertyList
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) TypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) UpdatedAt() AwsAutomationRule_UpdatedAtPropertyList {
	var returns AwsAutomationRule_UpdatedAtPropertyList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) UserDefinedFields() AwsAutomationRule_UserDefinedFieldsPropertyList {
	var returns AwsAutomationRule_UserDefinedFieldsPropertyList
	_jsii_.Get(
		j,
		"userDefinedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) UserDefinedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) VerificationState() AwsAutomationRule_VerificationStatePropertyList {
	var returns AwsAutomationRule_VerificationStatePropertyList
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) VerificationStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) WorkflowStatus() AwsAutomationRule_WorkflowStatusPropertyList {
	var returns AwsAutomationRule_WorkflowStatusPropertyList
	_jsii_.Get(
		j,
		"workflowStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) WorkflowStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowStatusInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAutomationRule_CriteriaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsAutomationRule_CriteriaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAutomationRule_CriteriaPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsAutomationRule.CriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAutomationRule_CriteriaPropertyOutputReference_Override(a AwsAutomationRule_CriteriaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsAutomationRule.CriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutAwsAccountId(value interface{}) {
	if err := a.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutAwsAccountName(value interface{}) {
	if err := a.validatePutAwsAccountNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsAccountName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutCompanyName(value interface{}) {
	if err := a.validatePutCompanyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCompanyName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutComplianceAssociatedStandardsId(value interface{}) {
	if err := a.validatePutComplianceAssociatedStandardsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComplianceAssociatedStandardsId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutComplianceSecurityControlId(value interface{}) {
	if err := a.validatePutComplianceSecurityControlIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComplianceSecurityControlId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutComplianceStatus(value interface{}) {
	if err := a.validatePutComplianceStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComplianceStatus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutConfidence(value interface{}) {
	if err := a.validatePutConfidenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfidence",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutCreatedAt(value interface{}) {
	if err := a.validatePutCreatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCreatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutCriticality(value interface{}) {
	if err := a.validatePutCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCriticality",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutDescription(value interface{}) {
	if err := a.validatePutDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDescription",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutFirstObservedAt(value interface{}) {
	if err := a.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutGeneratorId(value interface{}) {
	if err := a.validatePutGeneratorIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGeneratorId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutId(value interface{}) {
	if err := a.validatePutIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutLastObservedAt(value interface{}) {
	if err := a.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutNoteText(value interface{}) {
	if err := a.validatePutNoteTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoteText",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutNoteUpdatedAt(value interface{}) {
	if err := a.validatePutNoteUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoteUpdatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutNoteUpdatedBy(value interface{}) {
	if err := a.validatePutNoteUpdatedByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoteUpdatedBy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutProductArn(value interface{}) {
	if err := a.validatePutProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProductArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutProductName(value interface{}) {
	if err := a.validatePutProductNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProductName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutRecordState(value interface{}) {
	if err := a.validatePutRecordStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRecordState",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutRelatedFindingsId(value interface{}) {
	if err := a.validatePutRelatedFindingsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelatedFindingsId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutRelatedFindingsProductArn(value interface{}) {
	if err := a.validatePutRelatedFindingsProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelatedFindingsProductArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutResourceApplicationArn(value interface{}) {
	if err := a.validatePutResourceApplicationArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceApplicationArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutResourceApplicationName(value interface{}) {
	if err := a.validatePutResourceApplicationNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceApplicationName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutResourceDetailsOther(value interface{}) {
	if err := a.validatePutResourceDetailsOtherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceDetailsOther",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutResourceId(value interface{}) {
	if err := a.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutResourcePartition(value interface{}) {
	if err := a.validatePutResourcePartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourcePartition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutResourceRegion(value interface{}) {
	if err := a.validatePutResourceRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceRegion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutResourceTags(value interface{}) {
	if err := a.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutResourceType(value interface{}) {
	if err := a.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutSeverityLabel(value interface{}) {
	if err := a.validatePutSeverityLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSeverityLabel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutSourceUrl(value interface{}) {
	if err := a.validatePutSourceUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceUrl",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutTitle(value interface{}) {
	if err := a.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTitle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutType(value interface{}) {
	if err := a.validatePutTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutUpdatedAt(value interface{}) {
	if err := a.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutUserDefinedFields(value interface{}) {
	if err := a.validatePutUserDefinedFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserDefinedFields",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutVerificationState(value interface{}) {
	if err := a.validatePutVerificationStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVerificationState",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) PutWorkflowStatus(value interface{}) {
	if err := a.validatePutWorkflowStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkflowStatus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetAwsAccountName() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsAccountName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetCompanyName() {
	_jsii_.InvokeVoid(
		a,
		"resetCompanyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetComplianceAssociatedStandardsId() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceAssociatedStandardsId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetComplianceSecurityControlId() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceSecurityControlId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetComplianceStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetConfidence() {
	_jsii_.InvokeVoid(
		a,
		"resetConfidence",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		a,
		"resetCriticality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetGeneratorId() {
	_jsii_.InvokeVoid(
		a,
		"resetGeneratorId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetNoteText() {
	_jsii_.InvokeVoid(
		a,
		"resetNoteText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetNoteUpdatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetNoteUpdatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetNoteUpdatedBy() {
	_jsii_.InvokeVoid(
		a,
		"resetNoteUpdatedBy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetProductArn() {
	_jsii_.InvokeVoid(
		a,
		"resetProductArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetProductName() {
	_jsii_.InvokeVoid(
		a,
		"resetProductName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetRecordState() {
	_jsii_.InvokeVoid(
		a,
		"resetRecordState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetRelatedFindingsId() {
	_jsii_.InvokeVoid(
		a,
		"resetRelatedFindingsId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetResourceApplicationArn() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceApplicationArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetResourceApplicationName() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceApplicationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetResourceDetailsOther() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceDetailsOther",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetResourcePartition() {
	_jsii_.InvokeVoid(
		a,
		"resetResourcePartition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetResourceRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetSeverityLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetSeverityLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetSourceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		a,
		"resetTitle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetUserDefinedFields() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDefinedFields",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetVerificationState() {
	_jsii_.InvokeVoid(
		a,
		"resetVerificationState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ResetWorkflowStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkflowStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutomationRule_CriteriaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

