package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssecurityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssecurityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfInsight_FiltersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsAccountId() TfInsight_AwsAccountIdPropertyList
	// Experimental.
	AwsAccountIdInput() interface{}
	// Experimental.
	AwsAccountName() TfInsight_AwsAccountNamePropertyList
	// Experimental.
	AwsAccountNameInput() interface{}
	// Experimental.
	CompanyName() TfInsight_CompanyNamePropertyList
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
	ComplianceAssociatedStandardsId() TfInsight_ComplianceAssociatedStandardsIdPropertyList
	// Experimental.
	ComplianceAssociatedStandardsIdInput() interface{}
	// Experimental.
	ComplianceSecurityControlId() TfInsight_ComplianceSecurityControlIdPropertyList
	// Experimental.
	ComplianceSecurityControlIdInput() interface{}
	// Experimental.
	ComplianceSecurityControlParametersName() TfInsight_ComplianceSecurityControlParametersNamePropertyList
	// Experimental.
	ComplianceSecurityControlParametersNameInput() interface{}
	// Experimental.
	ComplianceSecurityControlParametersValue() TfInsight_ComplianceSecurityControlParametersValuePropertyList
	// Experimental.
	ComplianceSecurityControlParametersValueInput() interface{}
	// Experimental.
	ComplianceStatus() TfInsight_ComplianceStatusPropertyList
	// Experimental.
	ComplianceStatusInput() interface{}
	// Experimental.
	Confidence() TfInsight_ConfidencePropertyList
	// Experimental.
	ConfidenceInput() interface{}
	// Experimental.
	CreatedAt() TfInsight_CreatedAtPropertyList
	// Experimental.
	CreatedAtInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Criticality() TfInsight_CriticalityPropertyList
	// Experimental.
	CriticalityInput() interface{}
	// Experimental.
	Description() TfInsight_DescriptionPropertyList
	// Experimental.
	DescriptionInput() interface{}
	// Experimental.
	FindingProviderFieldsConfidence() TfInsight_FindingProviderFieldsConfidencePropertyList
	// Experimental.
	FindingProviderFieldsConfidenceInput() interface{}
	// Experimental.
	FindingProviderFieldsCriticality() TfInsight_FindingProviderFieldsCriticalityPropertyList
	// Experimental.
	FindingProviderFieldsCriticalityInput() interface{}
	// Experimental.
	FindingProviderFieldsRelatedFindingsId() TfInsight_FindingProviderFieldsRelatedFindingsIdPropertyList
	// Experimental.
	FindingProviderFieldsRelatedFindingsIdInput() interface{}
	// Experimental.
	FindingProviderFieldsRelatedFindingsProductArn() TfInsight_FindingProviderFieldsRelatedFindingsProductArnPropertyList
	// Experimental.
	FindingProviderFieldsRelatedFindingsProductArnInput() interface{}
	// Experimental.
	FindingProviderFieldsSeverityLabel() TfInsight_FindingProviderFieldsSeverityLabelPropertyList
	// Experimental.
	FindingProviderFieldsSeverityLabelInput() interface{}
	// Experimental.
	FindingProviderFieldsSeverityOriginal() TfInsight_FindingProviderFieldsSeverityOriginalPropertyList
	// Experimental.
	FindingProviderFieldsSeverityOriginalInput() interface{}
	// Experimental.
	FindingProviderFieldsTypes() TfInsight_FindingProviderFieldsTypesPropertyList
	// Experimental.
	FindingProviderFieldsTypesInput() interface{}
	// Experimental.
	FirstObservedAt() TfInsight_FirstObservedAtPropertyList
	// Experimental.
	FirstObservedAtInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	GeneratorId() TfInsight_GeneratorIdPropertyList
	// Experimental.
	GeneratorIdInput() interface{}
	// Experimental.
	Id() TfInsight_IdPropertyList
	// Experimental.
	IdInput() interface{}
	// Experimental.
	InternalValue() *TfInsight_FiltersProperty
	// Experimental.
	SetInternalValue(val *TfInsight_FiltersProperty)
	// Experimental.
	Keyword() TfInsight_KeywordPropertyList
	// Experimental.
	KeywordInput() interface{}
	// Experimental.
	LastObservedAt() TfInsight_LastObservedAtPropertyList
	// Experimental.
	LastObservedAtInput() interface{}
	// Experimental.
	MalwareName() TfInsight_MalwareNamePropertyList
	// Experimental.
	MalwareNameInput() interface{}
	// Experimental.
	MalwarePath() TfInsight_MalwarePathPropertyList
	// Experimental.
	MalwarePathInput() interface{}
	// Experimental.
	MalwareState() TfInsight_MalwareStatePropertyList
	// Experimental.
	MalwareStateInput() interface{}
	// Experimental.
	MalwareType() TfInsight_MalwareTypePropertyList
	// Experimental.
	MalwareTypeInput() interface{}
	// Experimental.
	NetworkDestinationDomain() TfInsight_NetworkDestinationDomainPropertyList
	// Experimental.
	NetworkDestinationDomainInput() interface{}
	// Experimental.
	NetworkDestinationIpv4() TfInsight_NetworkDestinationIpv4PropertyList
	// Experimental.
	NetworkDestinationIpv4Input() interface{}
	// Experimental.
	NetworkDestinationIpv6() TfInsight_NetworkDestinationIpv6PropertyList
	// Experimental.
	NetworkDestinationIpv6Input() interface{}
	// Experimental.
	NetworkDestinationPort() TfInsight_NetworkDestinationPortPropertyList
	// Experimental.
	NetworkDestinationPortInput() interface{}
	// Experimental.
	NetworkDirection() TfInsight_NetworkDirectionPropertyList
	// Experimental.
	NetworkDirectionInput() interface{}
	// Experimental.
	NetworkProtocol() TfInsight_NetworkProtocolPropertyList
	// Experimental.
	NetworkProtocolInput() interface{}
	// Experimental.
	NetworkSourceDomain() TfInsight_NetworkSourceDomainPropertyList
	// Experimental.
	NetworkSourceDomainInput() interface{}
	// Experimental.
	NetworkSourceIpv4() TfInsight_NetworkSourceIpv4PropertyList
	// Experimental.
	NetworkSourceIpv4Input() interface{}
	// Experimental.
	NetworkSourceIpv6() TfInsight_NetworkSourceIpv6PropertyList
	// Experimental.
	NetworkSourceIpv6Input() interface{}
	// Experimental.
	NetworkSourceMac() TfInsight_NetworkSourceMacPropertyList
	// Experimental.
	NetworkSourceMacInput() interface{}
	// Experimental.
	NetworkSourcePort() TfInsight_NetworkSourcePortPropertyList
	// Experimental.
	NetworkSourcePortInput() interface{}
	// Experimental.
	NoteText() TfInsight_NoteTextPropertyList
	// Experimental.
	NoteTextInput() interface{}
	// Experimental.
	NoteUpdatedAt() TfInsight_NoteUpdatedAtPropertyList
	// Experimental.
	NoteUpdatedAtInput() interface{}
	// Experimental.
	NoteUpdatedBy() TfInsight_NoteUpdatedByPropertyList
	// Experimental.
	NoteUpdatedByInput() interface{}
	// Experimental.
	ProcessLaunchedAt() TfInsight_ProcessLaunchedAtPropertyList
	// Experimental.
	ProcessLaunchedAtInput() interface{}
	// Experimental.
	ProcessName() TfInsight_ProcessNamePropertyList
	// Experimental.
	ProcessNameInput() interface{}
	// Experimental.
	ProcessParentPid() TfInsight_ProcessParentPidPropertyList
	// Experimental.
	ProcessParentPidInput() interface{}
	// Experimental.
	ProcessPath() TfInsight_ProcessPathPropertyList
	// Experimental.
	ProcessPathInput() interface{}
	// Experimental.
	ProcessPid() TfInsight_ProcessPidPropertyList
	// Experimental.
	ProcessPidInput() interface{}
	// Experimental.
	ProcessTerminatedAt() TfInsight_ProcessTerminatedAtPropertyList
	// Experimental.
	ProcessTerminatedAtInput() interface{}
	// Experimental.
	ProductArn() TfInsight_ProductArnPropertyList
	// Experimental.
	ProductArnInput() interface{}
	// Experimental.
	ProductFields() TfInsight_ProductFieldsPropertyList
	// Experimental.
	ProductFieldsInput() interface{}
	// Experimental.
	ProductName() TfInsight_ProductNamePropertyList
	// Experimental.
	ProductNameInput() interface{}
	// Experimental.
	RecommendationText() TfInsight_RecommendationTextPropertyList
	// Experimental.
	RecommendationTextInput() interface{}
	// Experimental.
	RecordState() TfInsight_RecordStatePropertyList
	// Experimental.
	RecordStateInput() interface{}
	// Experimental.
	RelatedFindingsId() TfInsight_RelatedFindingsIdPropertyList
	// Experimental.
	RelatedFindingsIdInput() interface{}
	// Experimental.
	RelatedFindingsProductArn() TfInsight_RelatedFindingsProductArnPropertyList
	// Experimental.
	RelatedFindingsProductArnInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceIamInstanceProfileArn() TfInsight_ResourceAwsEc2InstanceIamInstanceProfileArnPropertyList
	// Experimental.
	ResourceAwsEc2InstanceIamInstanceProfileArnInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceImageId() TfInsight_ResourceAwsEc2InstanceImageIdPropertyList
	// Experimental.
	ResourceAwsEc2InstanceImageIdInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceIpv4Addresses() TfInsight_ResourceAwsEc2InstanceIpv4AddressesPropertyList
	// Experimental.
	ResourceAwsEc2InstanceIpv4AddressesInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceIpv6Addresses() TfInsight_ResourceAwsEc2InstanceIpv6AddressesPropertyList
	// Experimental.
	ResourceAwsEc2InstanceIpv6AddressesInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceKeyName() TfInsight_ResourceAwsEc2InstanceKeyNamePropertyList
	// Experimental.
	ResourceAwsEc2InstanceKeyNameInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceLaunchedAt() TfInsight_ResourceAwsEc2InstanceLaunchedAtPropertyList
	// Experimental.
	ResourceAwsEc2InstanceLaunchedAtInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceSubnetId() TfInsight_ResourceAwsEc2InstanceSubnetIdPropertyList
	// Experimental.
	ResourceAwsEc2InstanceSubnetIdInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceType() TfInsight_ResourceAwsEc2InstanceTypePropertyList
	// Experimental.
	ResourceAwsEc2InstanceTypeInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceVpcId() TfInsight_ResourceAwsEc2InstanceVpcIdPropertyList
	// Experimental.
	ResourceAwsEc2InstanceVpcIdInput() interface{}
	// Experimental.
	ResourceAwsIamAccessKeyCreatedAt() TfInsight_ResourceAwsIamAccessKeyCreatedAtPropertyList
	// Experimental.
	ResourceAwsIamAccessKeyCreatedAtInput() interface{}
	// Experimental.
	ResourceAwsIamAccessKeyStatus() TfInsight_ResourceAwsIamAccessKeyStatusPropertyList
	// Experimental.
	ResourceAwsIamAccessKeyStatusInput() interface{}
	// Experimental.
	ResourceAwsIamAccessKeyUserName() TfInsight_ResourceAwsIamAccessKeyUserNamePropertyList
	// Experimental.
	ResourceAwsIamAccessKeyUserNameInput() interface{}
	// Experimental.
	ResourceAwsS3BucketOwnerId() TfInsight_ResourceAwsS3BucketOwnerIdPropertyList
	// Experimental.
	ResourceAwsS3BucketOwnerIdInput() interface{}
	// Experimental.
	ResourceAwsS3BucketOwnerName() TfInsight_ResourceAwsS3BucketOwnerNamePropertyList
	// Experimental.
	ResourceAwsS3BucketOwnerNameInput() interface{}
	// Experimental.
	ResourceContainerImageId() TfInsight_ResourceContainerImageIdPropertyList
	// Experimental.
	ResourceContainerImageIdInput() interface{}
	// Experimental.
	ResourceContainerImageName() TfInsight_ResourceContainerImageNamePropertyList
	// Experimental.
	ResourceContainerImageNameInput() interface{}
	// Experimental.
	ResourceContainerLaunchedAt() TfInsight_ResourceContainerLaunchedAtPropertyList
	// Experimental.
	ResourceContainerLaunchedAtInput() interface{}
	// Experimental.
	ResourceContainerName() TfInsight_ResourceContainerNamePropertyList
	// Experimental.
	ResourceContainerNameInput() interface{}
	// Experimental.
	ResourceDetailsOther() TfInsight_ResourceDetailsOtherPropertyList
	// Experimental.
	ResourceDetailsOtherInput() interface{}
	// Experimental.
	ResourceId() TfInsight_ResourceIdPropertyList
	// Experimental.
	ResourceIdInput() interface{}
	// Experimental.
	ResourcePartition() TfInsight_ResourcePartitionPropertyList
	// Experimental.
	ResourcePartitionInput() interface{}
	// Experimental.
	ResourceRegion() TfInsight_ResourceRegionPropertyList
	// Experimental.
	ResourceRegionInput() interface{}
	// Experimental.
	ResourceTags() TfInsight_ResourceTagsPropertyList
	// Experimental.
	ResourceTagsInput() interface{}
	// Experimental.
	ResourceType() TfInsight_ResourceTypePropertyList
	// Experimental.
	ResourceTypeInput() interface{}
	// Experimental.
	SeverityLabel() TfInsight_SeverityLabelPropertyList
	// Experimental.
	SeverityLabelInput() interface{}
	// Experimental.
	SourceUrl() TfInsight_SourceUrlPropertyList
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
	ThreatIntelIndicatorCategory() TfInsight_ThreatIntelIndicatorCategoryPropertyList
	// Experimental.
	ThreatIntelIndicatorCategoryInput() interface{}
	// Experimental.
	ThreatIntelIndicatorLastObservedAt() TfInsight_ThreatIntelIndicatorLastObservedAtPropertyList
	// Experimental.
	ThreatIntelIndicatorLastObservedAtInput() interface{}
	// Experimental.
	ThreatIntelIndicatorSource() TfInsight_ThreatIntelIndicatorSourcePropertyList
	// Experimental.
	ThreatIntelIndicatorSourceInput() interface{}
	// Experimental.
	ThreatIntelIndicatorSourceUrl() TfInsight_ThreatIntelIndicatorSourceUrlPropertyList
	// Experimental.
	ThreatIntelIndicatorSourceUrlInput() interface{}
	// Experimental.
	ThreatIntelIndicatorType() TfInsight_ThreatIntelIndicatorTypePropertyList
	// Experimental.
	ThreatIntelIndicatorTypeInput() interface{}
	// Experimental.
	ThreatIntelIndicatorValue() TfInsight_ThreatIntelIndicatorValuePropertyList
	// Experimental.
	ThreatIntelIndicatorValueInput() interface{}
	// Experimental.
	Title() TfInsight_TitlePropertyList
	// Experimental.
	TitleInput() interface{}
	// Experimental.
	Type() TfInsight_TypePropertyList
	// Experimental.
	TypeInput() interface{}
	// Experimental.
	UpdatedAt() TfInsight_UpdatedAtPropertyList
	// Experimental.
	UpdatedAtInput() interface{}
	// Experimental.
	UserDefinedValues() TfInsight_UserDefinedValuesPropertyList
	// Experimental.
	UserDefinedValuesInput() interface{}
	// Experimental.
	VerificationState() TfInsight_VerificationStatePropertyList
	// Experimental.
	VerificationStateInput() interface{}
	// Experimental.
	WorkflowStatus() TfInsight_WorkflowStatusPropertyList
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
	PutComplianceSecurityControlParametersName(value interface{})
	// Experimental.
	PutComplianceSecurityControlParametersValue(value interface{})
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
	PutFindingProviderFieldsConfidence(value interface{})
	// Experimental.
	PutFindingProviderFieldsCriticality(value interface{})
	// Experimental.
	PutFindingProviderFieldsRelatedFindingsId(value interface{})
	// Experimental.
	PutFindingProviderFieldsRelatedFindingsProductArn(value interface{})
	// Experimental.
	PutFindingProviderFieldsSeverityLabel(value interface{})
	// Experimental.
	PutFindingProviderFieldsSeverityOriginal(value interface{})
	// Experimental.
	PutFindingProviderFieldsTypes(value interface{})
	// Experimental.
	PutFirstObservedAt(value interface{})
	// Experimental.
	PutGeneratorId(value interface{})
	// Experimental.
	PutId(value interface{})
	// Experimental.
	PutKeyword(value interface{})
	// Experimental.
	PutLastObservedAt(value interface{})
	// Experimental.
	PutMalwareName(value interface{})
	// Experimental.
	PutMalwarePath(value interface{})
	// Experimental.
	PutMalwareState(value interface{})
	// Experimental.
	PutMalwareType(value interface{})
	// Experimental.
	PutNetworkDestinationDomain(value interface{})
	// Experimental.
	PutNetworkDestinationIpv4(value interface{})
	// Experimental.
	PutNetworkDestinationIpv6(value interface{})
	// Experimental.
	PutNetworkDestinationPort(value interface{})
	// Experimental.
	PutNetworkDirection(value interface{})
	// Experimental.
	PutNetworkProtocol(value interface{})
	// Experimental.
	PutNetworkSourceDomain(value interface{})
	// Experimental.
	PutNetworkSourceIpv4(value interface{})
	// Experimental.
	PutNetworkSourceIpv6(value interface{})
	// Experimental.
	PutNetworkSourceMac(value interface{})
	// Experimental.
	PutNetworkSourcePort(value interface{})
	// Experimental.
	PutNoteText(value interface{})
	// Experimental.
	PutNoteUpdatedAt(value interface{})
	// Experimental.
	PutNoteUpdatedBy(value interface{})
	// Experimental.
	PutProcessLaunchedAt(value interface{})
	// Experimental.
	PutProcessName(value interface{})
	// Experimental.
	PutProcessParentPid(value interface{})
	// Experimental.
	PutProcessPath(value interface{})
	// Experimental.
	PutProcessPid(value interface{})
	// Experimental.
	PutProcessTerminatedAt(value interface{})
	// Experimental.
	PutProductArn(value interface{})
	// Experimental.
	PutProductFields(value interface{})
	// Experimental.
	PutProductName(value interface{})
	// Experimental.
	PutRecommendationText(value interface{})
	// Experimental.
	PutRecordState(value interface{})
	// Experimental.
	PutRelatedFindingsId(value interface{})
	// Experimental.
	PutRelatedFindingsProductArn(value interface{})
	// Experimental.
	PutResourceAwsEc2InstanceIamInstanceProfileArn(value interface{})
	// Experimental.
	PutResourceAwsEc2InstanceImageId(value interface{})
	// Experimental.
	PutResourceAwsEc2InstanceIpv4Addresses(value interface{})
	// Experimental.
	PutResourceAwsEc2InstanceIpv6Addresses(value interface{})
	// Experimental.
	PutResourceAwsEc2InstanceKeyName(value interface{})
	// Experimental.
	PutResourceAwsEc2InstanceLaunchedAt(value interface{})
	// Experimental.
	PutResourceAwsEc2InstanceSubnetId(value interface{})
	// Experimental.
	PutResourceAwsEc2InstanceType(value interface{})
	// Experimental.
	PutResourceAwsEc2InstanceVpcId(value interface{})
	// Experimental.
	PutResourceAwsIamAccessKeyCreatedAt(value interface{})
	// Experimental.
	PutResourceAwsIamAccessKeyStatus(value interface{})
	// Experimental.
	PutResourceAwsIamAccessKeyUserName(value interface{})
	// Experimental.
	PutResourceAwsS3BucketOwnerId(value interface{})
	// Experimental.
	PutResourceAwsS3BucketOwnerName(value interface{})
	// Experimental.
	PutResourceContainerImageId(value interface{})
	// Experimental.
	PutResourceContainerImageName(value interface{})
	// Experimental.
	PutResourceContainerLaunchedAt(value interface{})
	// Experimental.
	PutResourceContainerName(value interface{})
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
	PutThreatIntelIndicatorCategory(value interface{})
	// Experimental.
	PutThreatIntelIndicatorLastObservedAt(value interface{})
	// Experimental.
	PutThreatIntelIndicatorSource(value interface{})
	// Experimental.
	PutThreatIntelIndicatorSourceUrl(value interface{})
	// Experimental.
	PutThreatIntelIndicatorType(value interface{})
	// Experimental.
	PutThreatIntelIndicatorValue(value interface{})
	// Experimental.
	PutTitle(value interface{})
	// Experimental.
	PutType(value interface{})
	// Experimental.
	PutUpdatedAt(value interface{})
	// Experimental.
	PutUserDefinedValues(value interface{})
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
	ResetComplianceSecurityControlParametersName()
	// Experimental.
	ResetComplianceSecurityControlParametersValue()
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
	ResetFindingProviderFieldsConfidence()
	// Experimental.
	ResetFindingProviderFieldsCriticality()
	// Experimental.
	ResetFindingProviderFieldsRelatedFindingsId()
	// Experimental.
	ResetFindingProviderFieldsRelatedFindingsProductArn()
	// Experimental.
	ResetFindingProviderFieldsSeverityLabel()
	// Experimental.
	ResetFindingProviderFieldsSeverityOriginal()
	// Experimental.
	ResetFindingProviderFieldsTypes()
	// Experimental.
	ResetFirstObservedAt()
	// Experimental.
	ResetGeneratorId()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKeyword()
	// Experimental.
	ResetLastObservedAt()
	// Experimental.
	ResetMalwareName()
	// Experimental.
	ResetMalwarePath()
	// Experimental.
	ResetMalwareState()
	// Experimental.
	ResetMalwareType()
	// Experimental.
	ResetNetworkDestinationDomain()
	// Experimental.
	ResetNetworkDestinationIpv4()
	// Experimental.
	ResetNetworkDestinationIpv6()
	// Experimental.
	ResetNetworkDestinationPort()
	// Experimental.
	ResetNetworkDirection()
	// Experimental.
	ResetNetworkProtocol()
	// Experimental.
	ResetNetworkSourceDomain()
	// Experimental.
	ResetNetworkSourceIpv4()
	// Experimental.
	ResetNetworkSourceIpv6()
	// Experimental.
	ResetNetworkSourceMac()
	// Experimental.
	ResetNetworkSourcePort()
	// Experimental.
	ResetNoteText()
	// Experimental.
	ResetNoteUpdatedAt()
	// Experimental.
	ResetNoteUpdatedBy()
	// Experimental.
	ResetProcessLaunchedAt()
	// Experimental.
	ResetProcessName()
	// Experimental.
	ResetProcessParentPid()
	// Experimental.
	ResetProcessPath()
	// Experimental.
	ResetProcessPid()
	// Experimental.
	ResetProcessTerminatedAt()
	// Experimental.
	ResetProductArn()
	// Experimental.
	ResetProductFields()
	// Experimental.
	ResetProductName()
	// Experimental.
	ResetRecommendationText()
	// Experimental.
	ResetRecordState()
	// Experimental.
	ResetRelatedFindingsId()
	// Experimental.
	ResetRelatedFindingsProductArn()
	// Experimental.
	ResetResourceAwsEc2InstanceIamInstanceProfileArn()
	// Experimental.
	ResetResourceAwsEc2InstanceImageId()
	// Experimental.
	ResetResourceAwsEc2InstanceIpv4Addresses()
	// Experimental.
	ResetResourceAwsEc2InstanceIpv6Addresses()
	// Experimental.
	ResetResourceAwsEc2InstanceKeyName()
	// Experimental.
	ResetResourceAwsEc2InstanceLaunchedAt()
	// Experimental.
	ResetResourceAwsEc2InstanceSubnetId()
	// Experimental.
	ResetResourceAwsEc2InstanceType()
	// Experimental.
	ResetResourceAwsEc2InstanceVpcId()
	// Experimental.
	ResetResourceAwsIamAccessKeyCreatedAt()
	// Experimental.
	ResetResourceAwsIamAccessKeyStatus()
	// Experimental.
	ResetResourceAwsIamAccessKeyUserName()
	// Experimental.
	ResetResourceAwsS3BucketOwnerId()
	// Experimental.
	ResetResourceAwsS3BucketOwnerName()
	// Experimental.
	ResetResourceContainerImageId()
	// Experimental.
	ResetResourceContainerImageName()
	// Experimental.
	ResetResourceContainerLaunchedAt()
	// Experimental.
	ResetResourceContainerName()
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
	ResetThreatIntelIndicatorCategory()
	// Experimental.
	ResetThreatIntelIndicatorLastObservedAt()
	// Experimental.
	ResetThreatIntelIndicatorSource()
	// Experimental.
	ResetThreatIntelIndicatorSourceUrl()
	// Experimental.
	ResetThreatIntelIndicatorType()
	// Experimental.
	ResetThreatIntelIndicatorValue()
	// Experimental.
	ResetTitle()
	// Experimental.
	ResetType()
	// Experimental.
	ResetUpdatedAt()
	// Experimental.
	ResetUserDefinedValues()
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

// The jsii proxy struct for TfInsight_FiltersPropertyOutputReference
type jsiiProxy_TfInsight_FiltersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) AwsAccountId() TfInsight_AwsAccountIdPropertyList {
	var returns TfInsight_AwsAccountIdPropertyList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) AwsAccountName() TfInsight_AwsAccountNamePropertyList {
	var returns TfInsight_AwsAccountNamePropertyList
	_jsii_.Get(
		j,
		"awsAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) AwsAccountNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) CompanyName() TfInsight_CompanyNamePropertyList {
	var returns TfInsight_CompanyNamePropertyList
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) CompanyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplianceAssociatedStandardsId() TfInsight_ComplianceAssociatedStandardsIdPropertyList {
	var returns TfInsight_ComplianceAssociatedStandardsIdPropertyList
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplianceAssociatedStandardsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplianceSecurityControlId() TfInsight_ComplianceSecurityControlIdPropertyList {
	var returns TfInsight_ComplianceSecurityControlIdPropertyList
	_jsii_.Get(
		j,
		"complianceSecurityControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplianceSecurityControlIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityControlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplianceSecurityControlParametersName() TfInsight_ComplianceSecurityControlParametersNamePropertyList {
	var returns TfInsight_ComplianceSecurityControlParametersNamePropertyList
	_jsii_.Get(
		j,
		"complianceSecurityControlParametersName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplianceSecurityControlParametersNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityControlParametersNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplianceSecurityControlParametersValue() TfInsight_ComplianceSecurityControlParametersValuePropertyList {
	var returns TfInsight_ComplianceSecurityControlParametersValuePropertyList
	_jsii_.Get(
		j,
		"complianceSecurityControlParametersValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplianceSecurityControlParametersValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityControlParametersValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplianceStatus() TfInsight_ComplianceStatusPropertyList {
	var returns TfInsight_ComplianceStatusPropertyList
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComplianceStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) Confidence() TfInsight_ConfidencePropertyList {
	var returns TfInsight_ConfidencePropertyList
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) CreatedAt() TfInsight_CreatedAtPropertyList {
	var returns TfInsight_CreatedAtPropertyList
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) CreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) Criticality() TfInsight_CriticalityPropertyList {
	var returns TfInsight_CriticalityPropertyList
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) CriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) Description() TfInsight_DescriptionPropertyList {
	var returns TfInsight_DescriptionPropertyList
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) DescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsConfidence() TfInsight_FindingProviderFieldsConfidencePropertyList {
	var returns TfInsight_FindingProviderFieldsConfidencePropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsConfidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsConfidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsCriticality() TfInsight_FindingProviderFieldsCriticalityPropertyList {
	var returns TfInsight_FindingProviderFieldsCriticalityPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsCriticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsCriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsCriticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsRelatedFindingsId() TfInsight_FindingProviderFieldsRelatedFindingsIdPropertyList {
	var returns TfInsight_FindingProviderFieldsRelatedFindingsIdPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsRelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsRelatedFindingsProductArn() TfInsight_FindingProviderFieldsRelatedFindingsProductArnPropertyList {
	var returns TfInsight_FindingProviderFieldsRelatedFindingsProductArnPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsRelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsSeverityLabel() TfInsight_FindingProviderFieldsSeverityLabelPropertyList {
	var returns TfInsight_FindingProviderFieldsSeverityLabelPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsSeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsSeverityOriginal() TfInsight_FindingProviderFieldsSeverityOriginalPropertyList {
	var returns TfInsight_FindingProviderFieldsSeverityOriginalPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityOriginal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsSeverityOriginalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityOriginalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsTypes() TfInsight_FindingProviderFieldsTypesPropertyList {
	var returns TfInsight_FindingProviderFieldsTypesPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FindingProviderFieldsTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FirstObservedAt() TfInsight_FirstObservedAtPropertyList {
	var returns TfInsight_FirstObservedAtPropertyList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) GeneratorId() TfInsight_GeneratorIdPropertyList {
	var returns TfInsight_GeneratorIdPropertyList
	_jsii_.Get(
		j,
		"generatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) GeneratorIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) Id() TfInsight_IdPropertyList {
	var returns TfInsight_IdPropertyList
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) IdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) InternalValue() *TfInsight_FiltersProperty {
	var returns *TfInsight_FiltersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) Keyword() TfInsight_KeywordPropertyList {
	var returns TfInsight_KeywordPropertyList
	_jsii_.Get(
		j,
		"keyword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) KeywordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keywordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) LastObservedAt() TfInsight_LastObservedAtPropertyList {
	var returns TfInsight_LastObservedAtPropertyList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) MalwareName() TfInsight_MalwareNamePropertyList {
	var returns TfInsight_MalwareNamePropertyList
	_jsii_.Get(
		j,
		"malwareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) MalwareNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) MalwarePath() TfInsight_MalwarePathPropertyList {
	var returns TfInsight_MalwarePathPropertyList
	_jsii_.Get(
		j,
		"malwarePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) MalwarePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwarePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) MalwareState() TfInsight_MalwareStatePropertyList {
	var returns TfInsight_MalwareStatePropertyList
	_jsii_.Get(
		j,
		"malwareState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) MalwareStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) MalwareType() TfInsight_MalwareTypePropertyList {
	var returns TfInsight_MalwareTypePropertyList
	_jsii_.Get(
		j,
		"malwareType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) MalwareTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkDestinationDomain() TfInsight_NetworkDestinationDomainPropertyList {
	var returns TfInsight_NetworkDestinationDomainPropertyList
	_jsii_.Get(
		j,
		"networkDestinationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkDestinationDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkDestinationIpv4() TfInsight_NetworkDestinationIpv4PropertyList {
	var returns TfInsight_NetworkDestinationIpv4PropertyList
	_jsii_.Get(
		j,
		"networkDestinationIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkDestinationIpv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkDestinationIpv6() TfInsight_NetworkDestinationIpv6PropertyList {
	var returns TfInsight_NetworkDestinationIpv6PropertyList
	_jsii_.Get(
		j,
		"networkDestinationIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkDestinationIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkDestinationPort() TfInsight_NetworkDestinationPortPropertyList {
	var returns TfInsight_NetworkDestinationPortPropertyList
	_jsii_.Get(
		j,
		"networkDestinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkDestinationPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkDirection() TfInsight_NetworkDirectionPropertyList {
	var returns TfInsight_NetworkDirectionPropertyList
	_jsii_.Get(
		j,
		"networkDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkDirectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkProtocol() TfInsight_NetworkProtocolPropertyList {
	var returns TfInsight_NetworkProtocolPropertyList
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkSourceDomain() TfInsight_NetworkSourceDomainPropertyList {
	var returns TfInsight_NetworkSourceDomainPropertyList
	_jsii_.Get(
		j,
		"networkSourceDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkSourceDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkSourceIpv4() TfInsight_NetworkSourceIpv4PropertyList {
	var returns TfInsight_NetworkSourceIpv4PropertyList
	_jsii_.Get(
		j,
		"networkSourceIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkSourceIpv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkSourceIpv6() TfInsight_NetworkSourceIpv6PropertyList {
	var returns TfInsight_NetworkSourceIpv6PropertyList
	_jsii_.Get(
		j,
		"networkSourceIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkSourceIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkSourceMac() TfInsight_NetworkSourceMacPropertyList {
	var returns TfInsight_NetworkSourceMacPropertyList
	_jsii_.Get(
		j,
		"networkSourceMac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkSourceMacInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceMacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkSourcePort() TfInsight_NetworkSourcePortPropertyList {
	var returns TfInsight_NetworkSourcePortPropertyList
	_jsii_.Get(
		j,
		"networkSourcePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NetworkSourcePortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourcePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NoteText() TfInsight_NoteTextPropertyList {
	var returns TfInsight_NoteTextPropertyList
	_jsii_.Get(
		j,
		"noteText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NoteTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NoteUpdatedAt() TfInsight_NoteUpdatedAtPropertyList {
	var returns TfInsight_NoteUpdatedAtPropertyList
	_jsii_.Get(
		j,
		"noteUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NoteUpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NoteUpdatedBy() TfInsight_NoteUpdatedByPropertyList {
	var returns TfInsight_NoteUpdatedByPropertyList
	_jsii_.Get(
		j,
		"noteUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) NoteUpdatedByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessLaunchedAt() TfInsight_ProcessLaunchedAtPropertyList {
	var returns TfInsight_ProcessLaunchedAtPropertyList
	_jsii_.Get(
		j,
		"processLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessName() TfInsight_ProcessNamePropertyList {
	var returns TfInsight_ProcessNamePropertyList
	_jsii_.Get(
		j,
		"processName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessParentPid() TfInsight_ProcessParentPidPropertyList {
	var returns TfInsight_ProcessParentPidPropertyList
	_jsii_.Get(
		j,
		"processParentPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessParentPidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processParentPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessPath() TfInsight_ProcessPathPropertyList {
	var returns TfInsight_ProcessPathPropertyList
	_jsii_.Get(
		j,
		"processPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessPid() TfInsight_ProcessPidPropertyList {
	var returns TfInsight_ProcessPidPropertyList
	_jsii_.Get(
		j,
		"processPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessPidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessTerminatedAt() TfInsight_ProcessTerminatedAtPropertyList {
	var returns TfInsight_ProcessTerminatedAtPropertyList
	_jsii_.Get(
		j,
		"processTerminatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProcessTerminatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processTerminatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProductArn() TfInsight_ProductArnPropertyList {
	var returns TfInsight_ProductArnPropertyList
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProductFields() TfInsight_ProductFieldsPropertyList {
	var returns TfInsight_ProductFieldsPropertyList
	_jsii_.Get(
		j,
		"productFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProductFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProductName() TfInsight_ProductNamePropertyList {
	var returns TfInsight_ProductNamePropertyList
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ProductNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) RecommendationText() TfInsight_RecommendationTextPropertyList {
	var returns TfInsight_RecommendationTextPropertyList
	_jsii_.Get(
		j,
		"recommendationText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) RecommendationTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recommendationTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) RecordState() TfInsight_RecordStatePropertyList {
	var returns TfInsight_RecordStatePropertyList
	_jsii_.Get(
		j,
		"recordState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) RecordStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) RelatedFindingsId() TfInsight_RelatedFindingsIdPropertyList {
	var returns TfInsight_RelatedFindingsIdPropertyList
	_jsii_.Get(
		j,
		"relatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) RelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) RelatedFindingsProductArn() TfInsight_RelatedFindingsProductArnPropertyList {
	var returns TfInsight_RelatedFindingsProductArnPropertyList
	_jsii_.Get(
		j,
		"relatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) RelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIamInstanceProfileArn() TfInsight_ResourceAwsEc2InstanceIamInstanceProfileArnPropertyList {
	var returns TfInsight_ResourceAwsEc2InstanceIamInstanceProfileArnPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIamInstanceProfileArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceImageId() TfInsight_ResourceAwsEc2InstanceImageIdPropertyList {
	var returns TfInsight_ResourceAwsEc2InstanceImageIdPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIpv4Addresses() TfInsight_ResourceAwsEc2InstanceIpv4AddressesPropertyList {
	var returns TfInsight_ResourceAwsEc2InstanceIpv4AddressesPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIpv4AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIpv6Addresses() TfInsight_ResourceAwsEc2InstanceIpv6AddressesPropertyList {
	var returns TfInsight_ResourceAwsEc2InstanceIpv6AddressesPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIpv6AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceKeyName() TfInsight_ResourceAwsEc2InstanceKeyNamePropertyList {
	var returns TfInsight_ResourceAwsEc2InstanceKeyNamePropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceKeyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceLaunchedAt() TfInsight_ResourceAwsEc2InstanceLaunchedAtPropertyList {
	var returns TfInsight_ResourceAwsEc2InstanceLaunchedAtPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceSubnetId() TfInsight_ResourceAwsEc2InstanceSubnetIdPropertyList {
	var returns TfInsight_ResourceAwsEc2InstanceSubnetIdPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceSubnetIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceType() TfInsight_ResourceAwsEc2InstanceTypePropertyList {
	var returns TfInsight_ResourceAwsEc2InstanceTypePropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceVpcId() TfInsight_ResourceAwsEc2InstanceVpcIdPropertyList {
	var returns TfInsight_ResourceAwsEc2InstanceVpcIdPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceVpcIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyCreatedAt() TfInsight_ResourceAwsIamAccessKeyCreatedAtPropertyList {
	var returns TfInsight_ResourceAwsIamAccessKeyCreatedAtPropertyList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyCreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyCreatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyStatus() TfInsight_ResourceAwsIamAccessKeyStatusPropertyList {
	var returns TfInsight_ResourceAwsIamAccessKeyStatusPropertyList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyUserName() TfInsight_ResourceAwsIamAccessKeyUserNamePropertyList {
	var returns TfInsight_ResourceAwsIamAccessKeyUserNamePropertyList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyUserNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsS3BucketOwnerId() TfInsight_ResourceAwsS3BucketOwnerIdPropertyList {
	var returns TfInsight_ResourceAwsS3BucketOwnerIdPropertyList
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsS3BucketOwnerIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsS3BucketOwnerName() TfInsight_ResourceAwsS3BucketOwnerNamePropertyList {
	var returns TfInsight_ResourceAwsS3BucketOwnerNamePropertyList
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceAwsS3BucketOwnerNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceContainerImageId() TfInsight_ResourceContainerImageIdPropertyList {
	var returns TfInsight_ResourceContainerImageIdPropertyList
	_jsii_.Get(
		j,
		"resourceContainerImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceContainerImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceContainerImageName() TfInsight_ResourceContainerImageNamePropertyList {
	var returns TfInsight_ResourceContainerImageNamePropertyList
	_jsii_.Get(
		j,
		"resourceContainerImageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceContainerImageNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerImageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceContainerLaunchedAt() TfInsight_ResourceContainerLaunchedAtPropertyList {
	var returns TfInsight_ResourceContainerLaunchedAtPropertyList
	_jsii_.Get(
		j,
		"resourceContainerLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceContainerLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceContainerName() TfInsight_ResourceContainerNamePropertyList {
	var returns TfInsight_ResourceContainerNamePropertyList
	_jsii_.Get(
		j,
		"resourceContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceContainerNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceDetailsOther() TfInsight_ResourceDetailsOtherPropertyList {
	var returns TfInsight_ResourceDetailsOtherPropertyList
	_jsii_.Get(
		j,
		"resourceDetailsOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceDetailsOtherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceDetailsOtherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceId() TfInsight_ResourceIdPropertyList {
	var returns TfInsight_ResourceIdPropertyList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourcePartition() TfInsight_ResourcePartitionPropertyList {
	var returns TfInsight_ResourcePartitionPropertyList
	_jsii_.Get(
		j,
		"resourcePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourcePartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceRegion() TfInsight_ResourceRegionPropertyList {
	var returns TfInsight_ResourceRegionPropertyList
	_jsii_.Get(
		j,
		"resourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceTags() TfInsight_ResourceTagsPropertyList {
	var returns TfInsight_ResourceTagsPropertyList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceType() TfInsight_ResourceTypePropertyList {
	var returns TfInsight_ResourceTypePropertyList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) SeverityLabel() TfInsight_SeverityLabelPropertyList {
	var returns TfInsight_SeverityLabelPropertyList
	_jsii_.Get(
		j,
		"severityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) SeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) SourceUrl() TfInsight_SourceUrlPropertyList {
	var returns TfInsight_SourceUrlPropertyList
	_jsii_.Get(
		j,
		"sourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) SourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorCategory() TfInsight_ThreatIntelIndicatorCategoryPropertyList {
	var returns TfInsight_ThreatIntelIndicatorCategoryPropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorCategoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorLastObservedAt() TfInsight_ThreatIntelIndicatorLastObservedAtPropertyList {
	var returns TfInsight_ThreatIntelIndicatorLastObservedAtPropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorLastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorLastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorLastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorSource() TfInsight_ThreatIntelIndicatorSourcePropertyList {
	var returns TfInsight_ThreatIntelIndicatorSourcePropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorSourceUrl() TfInsight_ThreatIntelIndicatorSourceUrlPropertyList {
	var returns TfInsight_ThreatIntelIndicatorSourceUrlPropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorSourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorType() TfInsight_ThreatIntelIndicatorTypePropertyList {
	var returns TfInsight_ThreatIntelIndicatorTypePropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorValue() TfInsight_ThreatIntelIndicatorValuePropertyList {
	var returns TfInsight_ThreatIntelIndicatorValuePropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) Title() TfInsight_TitlePropertyList {
	var returns TfInsight_TitlePropertyList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) Type() TfInsight_TypePropertyList {
	var returns TfInsight_TypePropertyList
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) TypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) UpdatedAt() TfInsight_UpdatedAtPropertyList {
	var returns TfInsight_UpdatedAtPropertyList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) UserDefinedValues() TfInsight_UserDefinedValuesPropertyList {
	var returns TfInsight_UserDefinedValuesPropertyList
	_jsii_.Get(
		j,
		"userDefinedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) UserDefinedValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) VerificationState() TfInsight_VerificationStatePropertyList {
	var returns TfInsight_VerificationStatePropertyList
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) VerificationStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) WorkflowStatus() TfInsight_WorkflowStatusPropertyList {
	var returns TfInsight_WorkflowStatusPropertyList
	_jsii_.Get(
		j,
		"workflowStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference) WorkflowStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowStatusInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfInsight_FiltersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfInsight_FiltersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfInsight_FiltersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfInsight_FiltersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfInsight.FiltersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfInsight_FiltersPropertyOutputReference_Override(t TfInsight_FiltersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfInsight.FiltersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference)SetInternalValue(val *TfInsight_FiltersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfInsight_FiltersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutAwsAccountId(value interface{}) {
	if err := t.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutAwsAccountName(value interface{}) {
	if err := t.validatePutAwsAccountNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsAccountName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutCompanyName(value interface{}) {
	if err := t.validatePutCompanyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCompanyName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutComplianceAssociatedStandardsId(value interface{}) {
	if err := t.validatePutComplianceAssociatedStandardsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putComplianceAssociatedStandardsId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutComplianceSecurityControlId(value interface{}) {
	if err := t.validatePutComplianceSecurityControlIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putComplianceSecurityControlId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutComplianceSecurityControlParametersName(value interface{}) {
	if err := t.validatePutComplianceSecurityControlParametersNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putComplianceSecurityControlParametersName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutComplianceSecurityControlParametersValue(value interface{}) {
	if err := t.validatePutComplianceSecurityControlParametersValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putComplianceSecurityControlParametersValue",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutComplianceStatus(value interface{}) {
	if err := t.validatePutComplianceStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putComplianceStatus",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutConfidence(value interface{}) {
	if err := t.validatePutConfidenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConfidence",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutCreatedAt(value interface{}) {
	if err := t.validatePutCreatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCreatedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutCriticality(value interface{}) {
	if err := t.validatePutCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCriticality",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutDescription(value interface{}) {
	if err := t.validatePutDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDescription",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsConfidence(value interface{}) {
	if err := t.validatePutFindingProviderFieldsConfidenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFindingProviderFieldsConfidence",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsCriticality(value interface{}) {
	if err := t.validatePutFindingProviderFieldsCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFindingProviderFieldsCriticality",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsRelatedFindingsId(value interface{}) {
	if err := t.validatePutFindingProviderFieldsRelatedFindingsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFindingProviderFieldsRelatedFindingsId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsRelatedFindingsProductArn(value interface{}) {
	if err := t.validatePutFindingProviderFieldsRelatedFindingsProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFindingProviderFieldsRelatedFindingsProductArn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsSeverityLabel(value interface{}) {
	if err := t.validatePutFindingProviderFieldsSeverityLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFindingProviderFieldsSeverityLabel",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsSeverityOriginal(value interface{}) {
	if err := t.validatePutFindingProviderFieldsSeverityOriginalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFindingProviderFieldsSeverityOriginal",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsTypes(value interface{}) {
	if err := t.validatePutFindingProviderFieldsTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFindingProviderFieldsTypes",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutFirstObservedAt(value interface{}) {
	if err := t.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutGeneratorId(value interface{}) {
	if err := t.validatePutGeneratorIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGeneratorId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutId(value interface{}) {
	if err := t.validatePutIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutKeyword(value interface{}) {
	if err := t.validatePutKeywordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKeyword",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutLastObservedAt(value interface{}) {
	if err := t.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutMalwareName(value interface{}) {
	if err := t.validatePutMalwareNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMalwareName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutMalwarePath(value interface{}) {
	if err := t.validatePutMalwarePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMalwarePath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutMalwareState(value interface{}) {
	if err := t.validatePutMalwareStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMalwareState",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutMalwareType(value interface{}) {
	if err := t.validatePutMalwareTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMalwareType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNetworkDestinationDomain(value interface{}) {
	if err := t.validatePutNetworkDestinationDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkDestinationDomain",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNetworkDestinationIpv4(value interface{}) {
	if err := t.validatePutNetworkDestinationIpv4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkDestinationIpv4",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNetworkDestinationIpv6(value interface{}) {
	if err := t.validatePutNetworkDestinationIpv6Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkDestinationIpv6",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNetworkDestinationPort(value interface{}) {
	if err := t.validatePutNetworkDestinationPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkDestinationPort",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNetworkDirection(value interface{}) {
	if err := t.validatePutNetworkDirectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkDirection",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNetworkProtocol(value interface{}) {
	if err := t.validatePutNetworkProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkProtocol",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNetworkSourceDomain(value interface{}) {
	if err := t.validatePutNetworkSourceDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkSourceDomain",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNetworkSourceIpv4(value interface{}) {
	if err := t.validatePutNetworkSourceIpv4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkSourceIpv4",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNetworkSourceIpv6(value interface{}) {
	if err := t.validatePutNetworkSourceIpv6Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkSourceIpv6",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNetworkSourceMac(value interface{}) {
	if err := t.validatePutNetworkSourceMacParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkSourceMac",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNetworkSourcePort(value interface{}) {
	if err := t.validatePutNetworkSourcePortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkSourcePort",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNoteText(value interface{}) {
	if err := t.validatePutNoteTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNoteText",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNoteUpdatedAt(value interface{}) {
	if err := t.validatePutNoteUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNoteUpdatedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutNoteUpdatedBy(value interface{}) {
	if err := t.validatePutNoteUpdatedByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNoteUpdatedBy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutProcessLaunchedAt(value interface{}) {
	if err := t.validatePutProcessLaunchedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProcessLaunchedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutProcessName(value interface{}) {
	if err := t.validatePutProcessNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProcessName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutProcessParentPid(value interface{}) {
	if err := t.validatePutProcessParentPidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProcessParentPid",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutProcessPath(value interface{}) {
	if err := t.validatePutProcessPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProcessPath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutProcessPid(value interface{}) {
	if err := t.validatePutProcessPidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProcessPid",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutProcessTerminatedAt(value interface{}) {
	if err := t.validatePutProcessTerminatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProcessTerminatedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutProductArn(value interface{}) {
	if err := t.validatePutProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProductArn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutProductFields(value interface{}) {
	if err := t.validatePutProductFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProductFields",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutProductName(value interface{}) {
	if err := t.validatePutProductNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProductName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutRecommendationText(value interface{}) {
	if err := t.validatePutRecommendationTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRecommendationText",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutRecordState(value interface{}) {
	if err := t.validatePutRecordStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRecordState",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutRelatedFindingsId(value interface{}) {
	if err := t.validatePutRelatedFindingsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRelatedFindingsId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutRelatedFindingsProductArn(value interface{}) {
	if err := t.validatePutRelatedFindingsProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRelatedFindingsProductArn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceIamInstanceProfileArn(value interface{}) {
	if err := t.validatePutResourceAwsEc2InstanceIamInstanceProfileArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsEc2InstanceIamInstanceProfileArn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceImageId(value interface{}) {
	if err := t.validatePutResourceAwsEc2InstanceImageIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsEc2InstanceImageId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceIpv4Addresses(value interface{}) {
	if err := t.validatePutResourceAwsEc2InstanceIpv4AddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsEc2InstanceIpv4Addresses",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceIpv6Addresses(value interface{}) {
	if err := t.validatePutResourceAwsEc2InstanceIpv6AddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsEc2InstanceIpv6Addresses",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceKeyName(value interface{}) {
	if err := t.validatePutResourceAwsEc2InstanceKeyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsEc2InstanceKeyName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceLaunchedAt(value interface{}) {
	if err := t.validatePutResourceAwsEc2InstanceLaunchedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsEc2InstanceLaunchedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceSubnetId(value interface{}) {
	if err := t.validatePutResourceAwsEc2InstanceSubnetIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsEc2InstanceSubnetId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceType(value interface{}) {
	if err := t.validatePutResourceAwsEc2InstanceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsEc2InstanceType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceVpcId(value interface{}) {
	if err := t.validatePutResourceAwsEc2InstanceVpcIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsEc2InstanceVpcId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsIamAccessKeyCreatedAt(value interface{}) {
	if err := t.validatePutResourceAwsIamAccessKeyCreatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsIamAccessKeyCreatedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsIamAccessKeyStatus(value interface{}) {
	if err := t.validatePutResourceAwsIamAccessKeyStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsIamAccessKeyStatus",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsIamAccessKeyUserName(value interface{}) {
	if err := t.validatePutResourceAwsIamAccessKeyUserNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsIamAccessKeyUserName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsS3BucketOwnerId(value interface{}) {
	if err := t.validatePutResourceAwsS3BucketOwnerIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsS3BucketOwnerId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceAwsS3BucketOwnerName(value interface{}) {
	if err := t.validatePutResourceAwsS3BucketOwnerNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceAwsS3BucketOwnerName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceContainerImageId(value interface{}) {
	if err := t.validatePutResourceContainerImageIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceContainerImageId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceContainerImageName(value interface{}) {
	if err := t.validatePutResourceContainerImageNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceContainerImageName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceContainerLaunchedAt(value interface{}) {
	if err := t.validatePutResourceContainerLaunchedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceContainerLaunchedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceContainerName(value interface{}) {
	if err := t.validatePutResourceContainerNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceContainerName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceDetailsOther(value interface{}) {
	if err := t.validatePutResourceDetailsOtherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceDetailsOther",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceId(value interface{}) {
	if err := t.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourcePartition(value interface{}) {
	if err := t.validatePutResourcePartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourcePartition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceRegion(value interface{}) {
	if err := t.validatePutResourceRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceRegion",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceTags(value interface{}) {
	if err := t.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutResourceType(value interface{}) {
	if err := t.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutSeverityLabel(value interface{}) {
	if err := t.validatePutSeverityLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSeverityLabel",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutSourceUrl(value interface{}) {
	if err := t.validatePutSourceUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourceUrl",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorCategory(value interface{}) {
	if err := t.validatePutThreatIntelIndicatorCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putThreatIntelIndicatorCategory",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorLastObservedAt(value interface{}) {
	if err := t.validatePutThreatIntelIndicatorLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putThreatIntelIndicatorLastObservedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorSource(value interface{}) {
	if err := t.validatePutThreatIntelIndicatorSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putThreatIntelIndicatorSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorSourceUrl(value interface{}) {
	if err := t.validatePutThreatIntelIndicatorSourceUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putThreatIntelIndicatorSourceUrl",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorType(value interface{}) {
	if err := t.validatePutThreatIntelIndicatorTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putThreatIntelIndicatorType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorValue(value interface{}) {
	if err := t.validatePutThreatIntelIndicatorValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putThreatIntelIndicatorValue",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutTitle(value interface{}) {
	if err := t.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTitle",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutType(value interface{}) {
	if err := t.validatePutTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutUpdatedAt(value interface{}) {
	if err := t.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutUserDefinedValues(value interface{}) {
	if err := t.validatePutUserDefinedValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUserDefinedValues",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutVerificationState(value interface{}) {
	if err := t.validatePutVerificationStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVerificationState",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) PutWorkflowStatus(value interface{}) {
	if err := t.validatePutWorkflowStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWorkflowStatus",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetAwsAccountName() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsAccountName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetCompanyName() {
	_jsii_.InvokeVoid(
		t,
		"resetCompanyName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetComplianceAssociatedStandardsId() {
	_jsii_.InvokeVoid(
		t,
		"resetComplianceAssociatedStandardsId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetComplianceSecurityControlId() {
	_jsii_.InvokeVoid(
		t,
		"resetComplianceSecurityControlId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetComplianceSecurityControlParametersName() {
	_jsii_.InvokeVoid(
		t,
		"resetComplianceSecurityControlParametersName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetComplianceSecurityControlParametersValue() {
	_jsii_.InvokeVoid(
		t,
		"resetComplianceSecurityControlParametersValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetComplianceStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetComplianceStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetConfidence() {
	_jsii_.InvokeVoid(
		t,
		"resetConfidence",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		t,
		"resetCriticality",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsConfidence() {
	_jsii_.InvokeVoid(
		t,
		"resetFindingProviderFieldsConfidence",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsCriticality() {
	_jsii_.InvokeVoid(
		t,
		"resetFindingProviderFieldsCriticality",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsRelatedFindingsId() {
	_jsii_.InvokeVoid(
		t,
		"resetFindingProviderFieldsRelatedFindingsId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		t,
		"resetFindingProviderFieldsRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsSeverityLabel() {
	_jsii_.InvokeVoid(
		t,
		"resetFindingProviderFieldsSeverityLabel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsSeverityOriginal() {
	_jsii_.InvokeVoid(
		t,
		"resetFindingProviderFieldsSeverityOriginal",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetFindingProviderFieldsTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetGeneratorId() {
	_jsii_.InvokeVoid(
		t,
		"resetGeneratorId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetKeyword() {
	_jsii_.InvokeVoid(
		t,
		"resetKeyword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetMalwareName() {
	_jsii_.InvokeVoid(
		t,
		"resetMalwareName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetMalwarePath() {
	_jsii_.InvokeVoid(
		t,
		"resetMalwarePath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetMalwareState() {
	_jsii_.InvokeVoid(
		t,
		"resetMalwareState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetMalwareType() {
	_jsii_.InvokeVoid(
		t,
		"resetMalwareType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNetworkDestinationDomain() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkDestinationDomain",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNetworkDestinationIpv4() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkDestinationIpv4",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNetworkDestinationIpv6() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkDestinationIpv6",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNetworkDestinationPort() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkDestinationPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNetworkDirection() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkDirection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNetworkProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNetworkSourceDomain() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkSourceDomain",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNetworkSourceIpv4() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkSourceIpv4",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNetworkSourceIpv6() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkSourceIpv6",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNetworkSourceMac() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkSourceMac",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNetworkSourcePort() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkSourcePort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNoteText() {
	_jsii_.InvokeVoid(
		t,
		"resetNoteText",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNoteUpdatedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetNoteUpdatedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetNoteUpdatedBy() {
	_jsii_.InvokeVoid(
		t,
		"resetNoteUpdatedBy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetProcessLaunchedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessLaunchedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetProcessName() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetProcessParentPid() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessParentPid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetProcessPath() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetProcessPid() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessPid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetProcessTerminatedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessTerminatedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetProductArn() {
	_jsii_.InvokeVoid(
		t,
		"resetProductArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetProductFields() {
	_jsii_.InvokeVoid(
		t,
		"resetProductFields",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetProductName() {
	_jsii_.InvokeVoid(
		t,
		"resetProductName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetRecommendationText() {
	_jsii_.InvokeVoid(
		t,
		"resetRecommendationText",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetRecordState() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetRelatedFindingsId() {
	_jsii_.InvokeVoid(
		t,
		"resetRelatedFindingsId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceIamInstanceProfileArn() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsEc2InstanceIamInstanceProfileArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceImageId() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsEc2InstanceImageId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceIpv4Addresses() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsEc2InstanceIpv4Addresses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceIpv6Addresses() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsEc2InstanceIpv6Addresses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceKeyName() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsEc2InstanceKeyName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceLaunchedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsEc2InstanceLaunchedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceSubnetId() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsEc2InstanceSubnetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceType() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsEc2InstanceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceVpcId() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsEc2InstanceVpcId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsIamAccessKeyCreatedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsIamAccessKeyCreatedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsIamAccessKeyStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsIamAccessKeyStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsIamAccessKeyUserName() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsIamAccessKeyUserName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsS3BucketOwnerId() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsS3BucketOwnerId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceAwsS3BucketOwnerName() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAwsS3BucketOwnerName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceContainerImageId() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceContainerImageId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceContainerImageName() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceContainerImageName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceContainerLaunchedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceContainerLaunchedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceContainerName() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceContainerName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceDetailsOther() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceDetailsOther",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourcePartition() {
	_jsii_.InvokeVoid(
		t,
		"resetResourcePartition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetSeverityLabel() {
	_jsii_.InvokeVoid(
		t,
		"resetSeverityLabel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetSourceUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorCategory() {
	_jsii_.InvokeVoid(
		t,
		"resetThreatIntelIndicatorCategory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorLastObservedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetThreatIntelIndicatorLastObservedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorSource() {
	_jsii_.InvokeVoid(
		t,
		"resetThreatIntelIndicatorSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorSourceUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetThreatIntelIndicatorSourceUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorType() {
	_jsii_.InvokeVoid(
		t,
		"resetThreatIntelIndicatorType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorValue() {
	_jsii_.InvokeVoid(
		t,
		"resetThreatIntelIndicatorValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		t,
		"resetTitle",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		t,
		"resetType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetUserDefinedValues() {
	_jsii_.InvokeVoid(
		t,
		"resetUserDefinedValues",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetVerificationState() {
	_jsii_.InvokeVoid(
		t,
		"resetVerificationState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ResetWorkflowStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkflowStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfInsight_FiltersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

