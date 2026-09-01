package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssecurityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssecurityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsAccountId() AwsSecurityhubAutomationRule_AwsAccountIdPropertyList
	// Experimental.
	AwsAccountIdInput() interface{}
	// Experimental.
	AwsAccountName() AwsSecurityhubAutomationRule_AwsAccountNamePropertyList
	// Experimental.
	AwsAccountNameInput() interface{}
	// Experimental.
	CompanyName() AwsSecurityhubAutomationRule_CompanyNamePropertyList
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
	ComplianceAssociatedStandardsId() AwsSecurityhubAutomationRule_ComplianceAssociatedStandardsIdPropertyList
	// Experimental.
	ComplianceAssociatedStandardsIdInput() interface{}
	// Experimental.
	ComplianceSecurityControlId() AwsSecurityhubAutomationRule_ComplianceSecurityControlIdPropertyList
	// Experimental.
	ComplianceSecurityControlIdInput() interface{}
	// Experimental.
	ComplianceStatus() AwsSecurityhubAutomationRule_ComplianceStatusPropertyList
	// Experimental.
	ComplianceStatusInput() interface{}
	// Experimental.
	Confidence() AwsSecurityhubAutomationRule_ConfidencePropertyList
	// Experimental.
	ConfidenceInput() interface{}
	// Experimental.
	CreatedAt() AwsSecurityhubAutomationRule_CreatedAtPropertyList
	// Experimental.
	CreatedAtInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Criticality() AwsSecurityhubAutomationRule_CriticalityPropertyList
	// Experimental.
	CriticalityInput() interface{}
	// Experimental.
	Description() AwsSecurityhubAutomationRule_DescriptionPropertyList
	// Experimental.
	DescriptionInput() interface{}
	// Experimental.
	FirstObservedAt() AwsSecurityhubAutomationRule_FirstObservedAtPropertyList
	// Experimental.
	FirstObservedAtInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	GeneratorId() AwsSecurityhubAutomationRule_GeneratorIdPropertyList
	// Experimental.
	GeneratorIdInput() interface{}
	// Experimental.
	Id() AwsSecurityhubAutomationRule_IdPropertyList
	// Experimental.
	IdInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LastObservedAt() AwsSecurityhubAutomationRule_LastObservedAtPropertyList
	// Experimental.
	LastObservedAtInput() interface{}
	// Experimental.
	NoteText() AwsSecurityhubAutomationRule_NoteTextPropertyList
	// Experimental.
	NoteTextInput() interface{}
	// Experimental.
	NoteUpdatedAt() AwsSecurityhubAutomationRule_NoteUpdatedAtPropertyList
	// Experimental.
	NoteUpdatedAtInput() interface{}
	// Experimental.
	NoteUpdatedBy() AwsSecurityhubAutomationRule_NoteUpdatedByPropertyList
	// Experimental.
	NoteUpdatedByInput() interface{}
	// Experimental.
	ProductArn() AwsSecurityhubAutomationRule_ProductArnPropertyList
	// Experimental.
	ProductArnInput() interface{}
	// Experimental.
	ProductName() AwsSecurityhubAutomationRule_ProductNamePropertyList
	// Experimental.
	ProductNameInput() interface{}
	// Experimental.
	RecordState() AwsSecurityhubAutomationRule_RecordStatePropertyList
	// Experimental.
	RecordStateInput() interface{}
	// Experimental.
	RelatedFindingsId() AwsSecurityhubAutomationRule_RelatedFindingsIdPropertyList
	// Experimental.
	RelatedFindingsIdInput() interface{}
	// Experimental.
	RelatedFindingsProductArn() AwsSecurityhubAutomationRule_RelatedFindingsProductArnPropertyList
	// Experimental.
	RelatedFindingsProductArnInput() interface{}
	// Experimental.
	ResourceApplicationArn() AwsSecurityhubAutomationRule_ResourceApplicationArnPropertyList
	// Experimental.
	ResourceApplicationArnInput() interface{}
	// Experimental.
	ResourceApplicationName() AwsSecurityhubAutomationRule_ResourceApplicationNamePropertyList
	// Experimental.
	ResourceApplicationNameInput() interface{}
	// Experimental.
	ResourceDetailsOther() AwsSecurityhubAutomationRule_ResourceDetailsOtherPropertyList
	// Experimental.
	ResourceDetailsOtherInput() interface{}
	// Experimental.
	ResourceId() AwsSecurityhubAutomationRule_ResourceIdPropertyList
	// Experimental.
	ResourceIdInput() interface{}
	// Experimental.
	ResourcePartition() AwsSecurityhubAutomationRule_ResourcePartitionPropertyList
	// Experimental.
	ResourcePartitionInput() interface{}
	// Experimental.
	ResourceRegion() AwsSecurityhubAutomationRule_ResourceRegionPropertyList
	// Experimental.
	ResourceRegionInput() interface{}
	// Experimental.
	ResourceTags() AwsSecurityhubAutomationRule_ResourceTagsPropertyList
	// Experimental.
	ResourceTagsInput() interface{}
	// Experimental.
	ResourceType() AwsSecurityhubAutomationRule_ResourceTypePropertyList
	// Experimental.
	ResourceTypeInput() interface{}
	// Experimental.
	SeverityLabel() AwsSecurityhubAutomationRule_SeverityLabelPropertyList
	// Experimental.
	SeverityLabelInput() interface{}
	// Experimental.
	SourceUrl() AwsSecurityhubAutomationRule_SourceUrlPropertyList
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
	Title() AwsSecurityhubAutomationRule_TitlePropertyList
	// Experimental.
	TitleInput() interface{}
	// Experimental.
	Type() AwsSecurityhubAutomationRule_TypePropertyList
	// Experimental.
	TypeInput() interface{}
	// Experimental.
	UpdatedAt() AwsSecurityhubAutomationRule_UpdatedAtPropertyList
	// Experimental.
	UpdatedAtInput() interface{}
	// Experimental.
	UserDefinedFields() AwsSecurityhubAutomationRule_UserDefinedFieldsPropertyList
	// Experimental.
	UserDefinedFieldsInput() interface{}
	// Experimental.
	VerificationState() AwsSecurityhubAutomationRule_VerificationStatePropertyList
	// Experimental.
	VerificationStateInput() interface{}
	// Experimental.
	WorkflowStatus() AwsSecurityhubAutomationRule_WorkflowStatusPropertyList
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

// The jsii proxy struct for AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference
type jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) AwsAccountId() AwsSecurityhubAutomationRule_AwsAccountIdPropertyList {
	var returns AwsSecurityhubAutomationRule_AwsAccountIdPropertyList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) AwsAccountName() AwsSecurityhubAutomationRule_AwsAccountNamePropertyList {
	var returns AwsSecurityhubAutomationRule_AwsAccountNamePropertyList
	_jsii_.Get(
		j,
		"awsAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) AwsAccountNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) CompanyName() AwsSecurityhubAutomationRule_CompanyNamePropertyList {
	var returns AwsSecurityhubAutomationRule_CompanyNamePropertyList
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) CompanyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ComplianceAssociatedStandardsId() AwsSecurityhubAutomationRule_ComplianceAssociatedStandardsIdPropertyList {
	var returns AwsSecurityhubAutomationRule_ComplianceAssociatedStandardsIdPropertyList
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ComplianceAssociatedStandardsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ComplianceSecurityControlId() AwsSecurityhubAutomationRule_ComplianceSecurityControlIdPropertyList {
	var returns AwsSecurityhubAutomationRule_ComplianceSecurityControlIdPropertyList
	_jsii_.Get(
		j,
		"complianceSecurityControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ComplianceSecurityControlIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityControlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ComplianceStatus() AwsSecurityhubAutomationRule_ComplianceStatusPropertyList {
	var returns AwsSecurityhubAutomationRule_ComplianceStatusPropertyList
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ComplianceStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) Confidence() AwsSecurityhubAutomationRule_ConfidencePropertyList {
	var returns AwsSecurityhubAutomationRule_ConfidencePropertyList
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) CreatedAt() AwsSecurityhubAutomationRule_CreatedAtPropertyList {
	var returns AwsSecurityhubAutomationRule_CreatedAtPropertyList
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) CreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) Criticality() AwsSecurityhubAutomationRule_CriticalityPropertyList {
	var returns AwsSecurityhubAutomationRule_CriticalityPropertyList
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) CriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) Description() AwsSecurityhubAutomationRule_DescriptionPropertyList {
	var returns AwsSecurityhubAutomationRule_DescriptionPropertyList
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) DescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) FirstObservedAt() AwsSecurityhubAutomationRule_FirstObservedAtPropertyList {
	var returns AwsSecurityhubAutomationRule_FirstObservedAtPropertyList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) GeneratorId() AwsSecurityhubAutomationRule_GeneratorIdPropertyList {
	var returns AwsSecurityhubAutomationRule_GeneratorIdPropertyList
	_jsii_.Get(
		j,
		"generatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) GeneratorIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) Id() AwsSecurityhubAutomationRule_IdPropertyList {
	var returns AwsSecurityhubAutomationRule_IdPropertyList
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) IdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) LastObservedAt() AwsSecurityhubAutomationRule_LastObservedAtPropertyList {
	var returns AwsSecurityhubAutomationRule_LastObservedAtPropertyList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) NoteText() AwsSecurityhubAutomationRule_NoteTextPropertyList {
	var returns AwsSecurityhubAutomationRule_NoteTextPropertyList
	_jsii_.Get(
		j,
		"noteText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) NoteTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedAt() AwsSecurityhubAutomationRule_NoteUpdatedAtPropertyList {
	var returns AwsSecurityhubAutomationRule_NoteUpdatedAtPropertyList
	_jsii_.Get(
		j,
		"noteUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedBy() AwsSecurityhubAutomationRule_NoteUpdatedByPropertyList {
	var returns AwsSecurityhubAutomationRule_NoteUpdatedByPropertyList
	_jsii_.Get(
		j,
		"noteUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ProductArn() AwsSecurityhubAutomationRule_ProductArnPropertyList {
	var returns AwsSecurityhubAutomationRule_ProductArnPropertyList
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ProductName() AwsSecurityhubAutomationRule_ProductNamePropertyList {
	var returns AwsSecurityhubAutomationRule_ProductNamePropertyList
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ProductNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) RecordState() AwsSecurityhubAutomationRule_RecordStatePropertyList {
	var returns AwsSecurityhubAutomationRule_RecordStatePropertyList
	_jsii_.Get(
		j,
		"recordState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) RecordStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsId() AwsSecurityhubAutomationRule_RelatedFindingsIdPropertyList {
	var returns AwsSecurityhubAutomationRule_RelatedFindingsIdPropertyList
	_jsii_.Get(
		j,
		"relatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsProductArn() AwsSecurityhubAutomationRule_RelatedFindingsProductArnPropertyList {
	var returns AwsSecurityhubAutomationRule_RelatedFindingsProductArnPropertyList
	_jsii_.Get(
		j,
		"relatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationArn() AwsSecurityhubAutomationRule_ResourceApplicationArnPropertyList {
	var returns AwsSecurityhubAutomationRule_ResourceApplicationArnPropertyList
	_jsii_.Get(
		j,
		"resourceApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceApplicationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationName() AwsSecurityhubAutomationRule_ResourceApplicationNamePropertyList {
	var returns AwsSecurityhubAutomationRule_ResourceApplicationNamePropertyList
	_jsii_.Get(
		j,
		"resourceApplicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceApplicationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceDetailsOther() AwsSecurityhubAutomationRule_ResourceDetailsOtherPropertyList {
	var returns AwsSecurityhubAutomationRule_ResourceDetailsOtherPropertyList
	_jsii_.Get(
		j,
		"resourceDetailsOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceDetailsOtherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceDetailsOtherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceId() AwsSecurityhubAutomationRule_ResourceIdPropertyList {
	var returns AwsSecurityhubAutomationRule_ResourceIdPropertyList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourcePartition() AwsSecurityhubAutomationRule_ResourcePartitionPropertyList {
	var returns AwsSecurityhubAutomationRule_ResourcePartitionPropertyList
	_jsii_.Get(
		j,
		"resourcePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourcePartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceRegion() AwsSecurityhubAutomationRule_ResourceRegionPropertyList {
	var returns AwsSecurityhubAutomationRule_ResourceRegionPropertyList
	_jsii_.Get(
		j,
		"resourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceTags() AwsSecurityhubAutomationRule_ResourceTagsPropertyList {
	var returns AwsSecurityhubAutomationRule_ResourceTagsPropertyList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceType() AwsSecurityhubAutomationRule_ResourceTypePropertyList {
	var returns AwsSecurityhubAutomationRule_ResourceTypePropertyList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) SeverityLabel() AwsSecurityhubAutomationRule_SeverityLabelPropertyList {
	var returns AwsSecurityhubAutomationRule_SeverityLabelPropertyList
	_jsii_.Get(
		j,
		"severityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) SeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) SourceUrl() AwsSecurityhubAutomationRule_SourceUrlPropertyList {
	var returns AwsSecurityhubAutomationRule_SourceUrlPropertyList
	_jsii_.Get(
		j,
		"sourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) SourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) Title() AwsSecurityhubAutomationRule_TitlePropertyList {
	var returns AwsSecurityhubAutomationRule_TitlePropertyList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) Type() AwsSecurityhubAutomationRule_TypePropertyList {
	var returns AwsSecurityhubAutomationRule_TypePropertyList
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) TypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) UpdatedAt() AwsSecurityhubAutomationRule_UpdatedAtPropertyList {
	var returns AwsSecurityhubAutomationRule_UpdatedAtPropertyList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) UserDefinedFields() AwsSecurityhubAutomationRule_UserDefinedFieldsPropertyList {
	var returns AwsSecurityhubAutomationRule_UserDefinedFieldsPropertyList
	_jsii_.Get(
		j,
		"userDefinedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) UserDefinedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) VerificationState() AwsSecurityhubAutomationRule_VerificationStatePropertyList {
	var returns AwsSecurityhubAutomationRule_VerificationStatePropertyList
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) VerificationStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) WorkflowStatus() AwsSecurityhubAutomationRule_WorkflowStatusPropertyList {
	var returns AwsSecurityhubAutomationRule_WorkflowStatusPropertyList
	_jsii_.Get(
		j,
		"workflowStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) WorkflowStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowStatusInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSecurityhubAutomationRule_CriteriaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSecurityhubAutomationRule_CriteriaPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsSecurityhubAutomationRule.CriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSecurityhubAutomationRule_CriteriaPropertyOutputReference_Override(a AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsSecurityhubAutomationRule.CriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutAwsAccountId(value interface{}) {
	if err := a.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutAwsAccountName(value interface{}) {
	if err := a.validatePutAwsAccountNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsAccountName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutCompanyName(value interface{}) {
	if err := a.validatePutCompanyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCompanyName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutComplianceAssociatedStandardsId(value interface{}) {
	if err := a.validatePutComplianceAssociatedStandardsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComplianceAssociatedStandardsId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutComplianceSecurityControlId(value interface{}) {
	if err := a.validatePutComplianceSecurityControlIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComplianceSecurityControlId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutComplianceStatus(value interface{}) {
	if err := a.validatePutComplianceStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComplianceStatus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutConfidence(value interface{}) {
	if err := a.validatePutConfidenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfidence",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutCreatedAt(value interface{}) {
	if err := a.validatePutCreatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCreatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutCriticality(value interface{}) {
	if err := a.validatePutCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCriticality",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutDescription(value interface{}) {
	if err := a.validatePutDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDescription",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutFirstObservedAt(value interface{}) {
	if err := a.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutGeneratorId(value interface{}) {
	if err := a.validatePutGeneratorIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGeneratorId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutId(value interface{}) {
	if err := a.validatePutIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutLastObservedAt(value interface{}) {
	if err := a.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutNoteText(value interface{}) {
	if err := a.validatePutNoteTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoteText",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutNoteUpdatedAt(value interface{}) {
	if err := a.validatePutNoteUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoteUpdatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutNoteUpdatedBy(value interface{}) {
	if err := a.validatePutNoteUpdatedByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoteUpdatedBy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutProductArn(value interface{}) {
	if err := a.validatePutProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProductArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutProductName(value interface{}) {
	if err := a.validatePutProductNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProductName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutRecordState(value interface{}) {
	if err := a.validatePutRecordStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRecordState",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutRelatedFindingsId(value interface{}) {
	if err := a.validatePutRelatedFindingsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelatedFindingsId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutRelatedFindingsProductArn(value interface{}) {
	if err := a.validatePutRelatedFindingsProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelatedFindingsProductArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutResourceApplicationArn(value interface{}) {
	if err := a.validatePutResourceApplicationArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceApplicationArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutResourceApplicationName(value interface{}) {
	if err := a.validatePutResourceApplicationNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceApplicationName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutResourceDetailsOther(value interface{}) {
	if err := a.validatePutResourceDetailsOtherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceDetailsOther",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutResourceId(value interface{}) {
	if err := a.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutResourcePartition(value interface{}) {
	if err := a.validatePutResourcePartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourcePartition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutResourceRegion(value interface{}) {
	if err := a.validatePutResourceRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceRegion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutResourceTags(value interface{}) {
	if err := a.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutResourceType(value interface{}) {
	if err := a.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutSeverityLabel(value interface{}) {
	if err := a.validatePutSeverityLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSeverityLabel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutSourceUrl(value interface{}) {
	if err := a.validatePutSourceUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceUrl",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutTitle(value interface{}) {
	if err := a.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTitle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutType(value interface{}) {
	if err := a.validatePutTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutUpdatedAt(value interface{}) {
	if err := a.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutUserDefinedFields(value interface{}) {
	if err := a.validatePutUserDefinedFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserDefinedFields",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutVerificationState(value interface{}) {
	if err := a.validatePutVerificationStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVerificationState",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) PutWorkflowStatus(value interface{}) {
	if err := a.validatePutWorkflowStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkflowStatus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetAwsAccountName() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsAccountName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetCompanyName() {
	_jsii_.InvokeVoid(
		a,
		"resetCompanyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetComplianceAssociatedStandardsId() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceAssociatedStandardsId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetComplianceSecurityControlId() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceSecurityControlId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetComplianceStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetConfidence() {
	_jsii_.InvokeVoid(
		a,
		"resetConfidence",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		a,
		"resetCriticality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetGeneratorId() {
	_jsii_.InvokeVoid(
		a,
		"resetGeneratorId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetNoteText() {
	_jsii_.InvokeVoid(
		a,
		"resetNoteText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetNoteUpdatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetNoteUpdatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetNoteUpdatedBy() {
	_jsii_.InvokeVoid(
		a,
		"resetNoteUpdatedBy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetProductArn() {
	_jsii_.InvokeVoid(
		a,
		"resetProductArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetProductName() {
	_jsii_.InvokeVoid(
		a,
		"resetProductName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetRecordState() {
	_jsii_.InvokeVoid(
		a,
		"resetRecordState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetRelatedFindingsId() {
	_jsii_.InvokeVoid(
		a,
		"resetRelatedFindingsId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetResourceApplicationArn() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceApplicationArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetResourceApplicationName() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceApplicationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetResourceDetailsOther() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceDetailsOther",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetResourcePartition() {
	_jsii_.InvokeVoid(
		a,
		"resetResourcePartition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetResourceRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetSeverityLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetSeverityLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetSourceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		a,
		"resetTitle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetUserDefinedFields() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDefinedFields",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetVerificationState() {
	_jsii_.InvokeVoid(
		a,
		"resetVerificationState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ResetWorkflowStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkflowStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_CriteriaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

