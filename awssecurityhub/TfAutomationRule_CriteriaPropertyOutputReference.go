package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssecurityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssecurityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAutomationRule_CriteriaPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsAccountId() TfAutomationRule_AwsAccountIdPropertyList
	// Experimental.
	AwsAccountIdInput() interface{}
	// Experimental.
	AwsAccountName() TfAutomationRule_AwsAccountNamePropertyList
	// Experimental.
	AwsAccountNameInput() interface{}
	// Experimental.
	CompanyName() TfAutomationRule_CompanyNamePropertyList
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
	ComplianceAssociatedStandardsId() TfAutomationRule_ComplianceAssociatedStandardsIdPropertyList
	// Experimental.
	ComplianceAssociatedStandardsIdInput() interface{}
	// Experimental.
	ComplianceSecurityControlId() TfAutomationRule_ComplianceSecurityControlIdPropertyList
	// Experimental.
	ComplianceSecurityControlIdInput() interface{}
	// Experimental.
	ComplianceStatus() TfAutomationRule_ComplianceStatusPropertyList
	// Experimental.
	ComplianceStatusInput() interface{}
	// Experimental.
	Confidence() TfAutomationRule_ConfidencePropertyList
	// Experimental.
	ConfidenceInput() interface{}
	// Experimental.
	CreatedAt() TfAutomationRule_CreatedAtPropertyList
	// Experimental.
	CreatedAtInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Criticality() TfAutomationRule_CriticalityPropertyList
	// Experimental.
	CriticalityInput() interface{}
	// Experimental.
	Description() TfAutomationRule_DescriptionPropertyList
	// Experimental.
	DescriptionInput() interface{}
	// Experimental.
	FirstObservedAt() TfAutomationRule_FirstObservedAtPropertyList
	// Experimental.
	FirstObservedAtInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	GeneratorId() TfAutomationRule_GeneratorIdPropertyList
	// Experimental.
	GeneratorIdInput() interface{}
	// Experimental.
	Id() TfAutomationRule_IdPropertyList
	// Experimental.
	IdInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LastObservedAt() TfAutomationRule_LastObservedAtPropertyList
	// Experimental.
	LastObservedAtInput() interface{}
	// Experimental.
	NoteText() TfAutomationRule_NoteTextPropertyList
	// Experimental.
	NoteTextInput() interface{}
	// Experimental.
	NoteUpdatedAt() TfAutomationRule_NoteUpdatedAtPropertyList
	// Experimental.
	NoteUpdatedAtInput() interface{}
	// Experimental.
	NoteUpdatedBy() TfAutomationRule_NoteUpdatedByPropertyList
	// Experimental.
	NoteUpdatedByInput() interface{}
	// Experimental.
	ProductArn() TfAutomationRule_ProductArnPropertyList
	// Experimental.
	ProductArnInput() interface{}
	// Experimental.
	ProductName() TfAutomationRule_ProductNamePropertyList
	// Experimental.
	ProductNameInput() interface{}
	// Experimental.
	RecordState() TfAutomationRule_RecordStatePropertyList
	// Experimental.
	RecordStateInput() interface{}
	// Experimental.
	RelatedFindingsId() TfAutomationRule_RelatedFindingsIdPropertyList
	// Experimental.
	RelatedFindingsIdInput() interface{}
	// Experimental.
	RelatedFindingsProductArn() TfAutomationRule_RelatedFindingsProductArnPropertyList
	// Experimental.
	RelatedFindingsProductArnInput() interface{}
	// Experimental.
	ResourceApplicationArn() TfAutomationRule_ResourceApplicationArnPropertyList
	// Experimental.
	ResourceApplicationArnInput() interface{}
	// Experimental.
	ResourceApplicationName() TfAutomationRule_ResourceApplicationNamePropertyList
	// Experimental.
	ResourceApplicationNameInput() interface{}
	// Experimental.
	ResourceDetailsOther() TfAutomationRule_ResourceDetailsOtherPropertyList
	// Experimental.
	ResourceDetailsOtherInput() interface{}
	// Experimental.
	ResourceId() TfAutomationRule_ResourceIdPropertyList
	// Experimental.
	ResourceIdInput() interface{}
	// Experimental.
	ResourcePartition() TfAutomationRule_ResourcePartitionPropertyList
	// Experimental.
	ResourcePartitionInput() interface{}
	// Experimental.
	ResourceRegion() TfAutomationRule_ResourceRegionPropertyList
	// Experimental.
	ResourceRegionInput() interface{}
	// Experimental.
	ResourceTags() TfAutomationRule_ResourceTagsPropertyList
	// Experimental.
	ResourceTagsInput() interface{}
	// Experimental.
	ResourceType() TfAutomationRule_ResourceTypePropertyList
	// Experimental.
	ResourceTypeInput() interface{}
	// Experimental.
	SeverityLabel() TfAutomationRule_SeverityLabelPropertyList
	// Experimental.
	SeverityLabelInput() interface{}
	// Experimental.
	SourceUrl() TfAutomationRule_SourceUrlPropertyList
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
	Title() TfAutomationRule_TitlePropertyList
	// Experimental.
	TitleInput() interface{}
	// Experimental.
	Type() TfAutomationRule_TypePropertyList
	// Experimental.
	TypeInput() interface{}
	// Experimental.
	UpdatedAt() TfAutomationRule_UpdatedAtPropertyList
	// Experimental.
	UpdatedAtInput() interface{}
	// Experimental.
	UserDefinedFields() TfAutomationRule_UserDefinedFieldsPropertyList
	// Experimental.
	UserDefinedFieldsInput() interface{}
	// Experimental.
	VerificationState() TfAutomationRule_VerificationStatePropertyList
	// Experimental.
	VerificationStateInput() interface{}
	// Experimental.
	WorkflowStatus() TfAutomationRule_WorkflowStatusPropertyList
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

// The jsii proxy struct for TfAutomationRule_CriteriaPropertyOutputReference
type jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) AwsAccountId() TfAutomationRule_AwsAccountIdPropertyList {
	var returns TfAutomationRule_AwsAccountIdPropertyList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) AwsAccountName() TfAutomationRule_AwsAccountNamePropertyList {
	var returns TfAutomationRule_AwsAccountNamePropertyList
	_jsii_.Get(
		j,
		"awsAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) AwsAccountNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) CompanyName() TfAutomationRule_CompanyNamePropertyList {
	var returns TfAutomationRule_CompanyNamePropertyList
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) CompanyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ComplianceAssociatedStandardsId() TfAutomationRule_ComplianceAssociatedStandardsIdPropertyList {
	var returns TfAutomationRule_ComplianceAssociatedStandardsIdPropertyList
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ComplianceAssociatedStandardsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ComplianceSecurityControlId() TfAutomationRule_ComplianceSecurityControlIdPropertyList {
	var returns TfAutomationRule_ComplianceSecurityControlIdPropertyList
	_jsii_.Get(
		j,
		"complianceSecurityControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ComplianceSecurityControlIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityControlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ComplianceStatus() TfAutomationRule_ComplianceStatusPropertyList {
	var returns TfAutomationRule_ComplianceStatusPropertyList
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ComplianceStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) Confidence() TfAutomationRule_ConfidencePropertyList {
	var returns TfAutomationRule_ConfidencePropertyList
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) CreatedAt() TfAutomationRule_CreatedAtPropertyList {
	var returns TfAutomationRule_CreatedAtPropertyList
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) CreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) Criticality() TfAutomationRule_CriticalityPropertyList {
	var returns TfAutomationRule_CriticalityPropertyList
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) CriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) Description() TfAutomationRule_DescriptionPropertyList {
	var returns TfAutomationRule_DescriptionPropertyList
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) DescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) FirstObservedAt() TfAutomationRule_FirstObservedAtPropertyList {
	var returns TfAutomationRule_FirstObservedAtPropertyList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) GeneratorId() TfAutomationRule_GeneratorIdPropertyList {
	var returns TfAutomationRule_GeneratorIdPropertyList
	_jsii_.Get(
		j,
		"generatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) GeneratorIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) Id() TfAutomationRule_IdPropertyList {
	var returns TfAutomationRule_IdPropertyList
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) IdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) LastObservedAt() TfAutomationRule_LastObservedAtPropertyList {
	var returns TfAutomationRule_LastObservedAtPropertyList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) NoteText() TfAutomationRule_NoteTextPropertyList {
	var returns TfAutomationRule_NoteTextPropertyList
	_jsii_.Get(
		j,
		"noteText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) NoteTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedAt() TfAutomationRule_NoteUpdatedAtPropertyList {
	var returns TfAutomationRule_NoteUpdatedAtPropertyList
	_jsii_.Get(
		j,
		"noteUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedBy() TfAutomationRule_NoteUpdatedByPropertyList {
	var returns TfAutomationRule_NoteUpdatedByPropertyList
	_jsii_.Get(
		j,
		"noteUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) NoteUpdatedByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ProductArn() TfAutomationRule_ProductArnPropertyList {
	var returns TfAutomationRule_ProductArnPropertyList
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ProductName() TfAutomationRule_ProductNamePropertyList {
	var returns TfAutomationRule_ProductNamePropertyList
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ProductNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) RecordState() TfAutomationRule_RecordStatePropertyList {
	var returns TfAutomationRule_RecordStatePropertyList
	_jsii_.Get(
		j,
		"recordState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) RecordStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsId() TfAutomationRule_RelatedFindingsIdPropertyList {
	var returns TfAutomationRule_RelatedFindingsIdPropertyList
	_jsii_.Get(
		j,
		"relatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsProductArn() TfAutomationRule_RelatedFindingsProductArnPropertyList {
	var returns TfAutomationRule_RelatedFindingsProductArnPropertyList
	_jsii_.Get(
		j,
		"relatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) RelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationArn() TfAutomationRule_ResourceApplicationArnPropertyList {
	var returns TfAutomationRule_ResourceApplicationArnPropertyList
	_jsii_.Get(
		j,
		"resourceApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceApplicationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationName() TfAutomationRule_ResourceApplicationNamePropertyList {
	var returns TfAutomationRule_ResourceApplicationNamePropertyList
	_jsii_.Get(
		j,
		"resourceApplicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceApplicationNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceApplicationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceDetailsOther() TfAutomationRule_ResourceDetailsOtherPropertyList {
	var returns TfAutomationRule_ResourceDetailsOtherPropertyList
	_jsii_.Get(
		j,
		"resourceDetailsOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceDetailsOtherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceDetailsOtherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceId() TfAutomationRule_ResourceIdPropertyList {
	var returns TfAutomationRule_ResourceIdPropertyList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourcePartition() TfAutomationRule_ResourcePartitionPropertyList {
	var returns TfAutomationRule_ResourcePartitionPropertyList
	_jsii_.Get(
		j,
		"resourcePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourcePartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceRegion() TfAutomationRule_ResourceRegionPropertyList {
	var returns TfAutomationRule_ResourceRegionPropertyList
	_jsii_.Get(
		j,
		"resourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceTags() TfAutomationRule_ResourceTagsPropertyList {
	var returns TfAutomationRule_ResourceTagsPropertyList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceType() TfAutomationRule_ResourceTypePropertyList {
	var returns TfAutomationRule_ResourceTypePropertyList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) SeverityLabel() TfAutomationRule_SeverityLabelPropertyList {
	var returns TfAutomationRule_SeverityLabelPropertyList
	_jsii_.Get(
		j,
		"severityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) SeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) SourceUrl() TfAutomationRule_SourceUrlPropertyList {
	var returns TfAutomationRule_SourceUrlPropertyList
	_jsii_.Get(
		j,
		"sourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) SourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) Title() TfAutomationRule_TitlePropertyList {
	var returns TfAutomationRule_TitlePropertyList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) Type() TfAutomationRule_TypePropertyList {
	var returns TfAutomationRule_TypePropertyList
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) TypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) UpdatedAt() TfAutomationRule_UpdatedAtPropertyList {
	var returns TfAutomationRule_UpdatedAtPropertyList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) UserDefinedFields() TfAutomationRule_UserDefinedFieldsPropertyList {
	var returns TfAutomationRule_UserDefinedFieldsPropertyList
	_jsii_.Get(
		j,
		"userDefinedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) UserDefinedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) VerificationState() TfAutomationRule_VerificationStatePropertyList {
	var returns TfAutomationRule_VerificationStatePropertyList
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) VerificationStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) WorkflowStatus() TfAutomationRule_WorkflowStatusPropertyList {
	var returns TfAutomationRule_WorkflowStatusPropertyList
	_jsii_.Get(
		j,
		"workflowStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) WorkflowStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowStatusInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAutomationRule_CriteriaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfAutomationRule_CriteriaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAutomationRule_CriteriaPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfAutomationRule.CriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAutomationRule_CriteriaPropertyOutputReference_Override(t TfAutomationRule_CriteriaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfAutomationRule.CriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutAwsAccountId(value interface{}) {
	if err := t.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutAwsAccountName(value interface{}) {
	if err := t.validatePutAwsAccountNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsAccountName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutCompanyName(value interface{}) {
	if err := t.validatePutCompanyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCompanyName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutComplianceAssociatedStandardsId(value interface{}) {
	if err := t.validatePutComplianceAssociatedStandardsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putComplianceAssociatedStandardsId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutComplianceSecurityControlId(value interface{}) {
	if err := t.validatePutComplianceSecurityControlIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putComplianceSecurityControlId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutComplianceStatus(value interface{}) {
	if err := t.validatePutComplianceStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putComplianceStatus",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutConfidence(value interface{}) {
	if err := t.validatePutConfidenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConfidence",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutCreatedAt(value interface{}) {
	if err := t.validatePutCreatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCreatedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutCriticality(value interface{}) {
	if err := t.validatePutCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCriticality",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutDescription(value interface{}) {
	if err := t.validatePutDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDescription",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutFirstObservedAt(value interface{}) {
	if err := t.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutGeneratorId(value interface{}) {
	if err := t.validatePutGeneratorIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGeneratorId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutId(value interface{}) {
	if err := t.validatePutIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutLastObservedAt(value interface{}) {
	if err := t.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutNoteText(value interface{}) {
	if err := t.validatePutNoteTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNoteText",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutNoteUpdatedAt(value interface{}) {
	if err := t.validatePutNoteUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNoteUpdatedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutNoteUpdatedBy(value interface{}) {
	if err := t.validatePutNoteUpdatedByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNoteUpdatedBy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutProductArn(value interface{}) {
	if err := t.validatePutProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProductArn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutProductName(value interface{}) {
	if err := t.validatePutProductNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProductName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutRecordState(value interface{}) {
	if err := t.validatePutRecordStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRecordState",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutRelatedFindingsId(value interface{}) {
	if err := t.validatePutRelatedFindingsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRelatedFindingsId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutRelatedFindingsProductArn(value interface{}) {
	if err := t.validatePutRelatedFindingsProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRelatedFindingsProductArn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutResourceApplicationArn(value interface{}) {
	if err := t.validatePutResourceApplicationArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceApplicationArn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutResourceApplicationName(value interface{}) {
	if err := t.validatePutResourceApplicationNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceApplicationName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutResourceDetailsOther(value interface{}) {
	if err := t.validatePutResourceDetailsOtherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceDetailsOther",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutResourceId(value interface{}) {
	if err := t.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutResourcePartition(value interface{}) {
	if err := t.validatePutResourcePartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourcePartition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutResourceRegion(value interface{}) {
	if err := t.validatePutResourceRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceRegion",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutResourceTags(value interface{}) {
	if err := t.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutResourceType(value interface{}) {
	if err := t.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutSeverityLabel(value interface{}) {
	if err := t.validatePutSeverityLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSeverityLabel",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutSourceUrl(value interface{}) {
	if err := t.validatePutSourceUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourceUrl",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutTitle(value interface{}) {
	if err := t.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTitle",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutType(value interface{}) {
	if err := t.validatePutTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutUpdatedAt(value interface{}) {
	if err := t.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutUserDefinedFields(value interface{}) {
	if err := t.validatePutUserDefinedFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUserDefinedFields",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutVerificationState(value interface{}) {
	if err := t.validatePutVerificationStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVerificationState",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) PutWorkflowStatus(value interface{}) {
	if err := t.validatePutWorkflowStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWorkflowStatus",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetAwsAccountName() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsAccountName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetCompanyName() {
	_jsii_.InvokeVoid(
		t,
		"resetCompanyName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetComplianceAssociatedStandardsId() {
	_jsii_.InvokeVoid(
		t,
		"resetComplianceAssociatedStandardsId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetComplianceSecurityControlId() {
	_jsii_.InvokeVoid(
		t,
		"resetComplianceSecurityControlId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetComplianceStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetComplianceStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetConfidence() {
	_jsii_.InvokeVoid(
		t,
		"resetConfidence",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		t,
		"resetCriticality",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetGeneratorId() {
	_jsii_.InvokeVoid(
		t,
		"resetGeneratorId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetNoteText() {
	_jsii_.InvokeVoid(
		t,
		"resetNoteText",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetNoteUpdatedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetNoteUpdatedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetNoteUpdatedBy() {
	_jsii_.InvokeVoid(
		t,
		"resetNoteUpdatedBy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetProductArn() {
	_jsii_.InvokeVoid(
		t,
		"resetProductArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetProductName() {
	_jsii_.InvokeVoid(
		t,
		"resetProductName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetRecordState() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetRelatedFindingsId() {
	_jsii_.InvokeVoid(
		t,
		"resetRelatedFindingsId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetResourceApplicationArn() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceApplicationArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetResourceApplicationName() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceApplicationName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetResourceDetailsOther() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceDetailsOther",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetResourcePartition() {
	_jsii_.InvokeVoid(
		t,
		"resetResourcePartition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetResourceRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetSeverityLabel() {
	_jsii_.InvokeVoid(
		t,
		"resetSeverityLabel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetSourceUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		t,
		"resetTitle",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		t,
		"resetType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetUserDefinedFields() {
	_jsii_.InvokeVoid(
		t,
		"resetUserDefinedFields",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetVerificationState() {
	_jsii_.InvokeVoid(
		t,
		"resetVerificationState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ResetWorkflowStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkflowStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriteriaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

