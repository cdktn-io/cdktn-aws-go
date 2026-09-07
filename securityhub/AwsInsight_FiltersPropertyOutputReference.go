package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/securityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/securityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsInsight_FiltersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsAccountId() AwsInsight_AwsAccountIdPropertyList
	// Experimental.
	AwsAccountIdInput() interface{}
	// Experimental.
	AwsAccountName() AwsInsight_AwsAccountNamePropertyList
	// Experimental.
	AwsAccountNameInput() interface{}
	// Experimental.
	CompanyName() AwsInsight_CompanyNamePropertyList
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
	ComplianceAssociatedStandardsId() AwsInsight_ComplianceAssociatedStandardsIdPropertyList
	// Experimental.
	ComplianceAssociatedStandardsIdInput() interface{}
	// Experimental.
	ComplianceSecurityControlId() AwsInsight_ComplianceSecurityControlIdPropertyList
	// Experimental.
	ComplianceSecurityControlIdInput() interface{}
	// Experimental.
	ComplianceSecurityControlParametersName() AwsInsight_ComplianceSecurityControlParametersNamePropertyList
	// Experimental.
	ComplianceSecurityControlParametersNameInput() interface{}
	// Experimental.
	ComplianceSecurityControlParametersValue() AwsInsight_ComplianceSecurityControlParametersValuePropertyList
	// Experimental.
	ComplianceSecurityControlParametersValueInput() interface{}
	// Experimental.
	ComplianceStatus() AwsInsight_ComplianceStatusPropertyList
	// Experimental.
	ComplianceStatusInput() interface{}
	// Experimental.
	Confidence() AwsInsight_ConfidencePropertyList
	// Experimental.
	ConfidenceInput() interface{}
	// Experimental.
	CreatedAt() AwsInsight_CreatedAtPropertyList
	// Experimental.
	CreatedAtInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Criticality() AwsInsight_CriticalityPropertyList
	// Experimental.
	CriticalityInput() interface{}
	// Experimental.
	Description() AwsInsight_DescriptionPropertyList
	// Experimental.
	DescriptionInput() interface{}
	// Experimental.
	FindingProviderFieldsConfidence() AwsInsight_FindingProviderFieldsConfidencePropertyList
	// Experimental.
	FindingProviderFieldsConfidenceInput() interface{}
	// Experimental.
	FindingProviderFieldsCriticality() AwsInsight_FindingProviderFieldsCriticalityPropertyList
	// Experimental.
	FindingProviderFieldsCriticalityInput() interface{}
	// Experimental.
	FindingProviderFieldsRelatedFindingsId() AwsInsight_FindingProviderFieldsRelatedFindingsIdPropertyList
	// Experimental.
	FindingProviderFieldsRelatedFindingsIdInput() interface{}
	// Experimental.
	FindingProviderFieldsRelatedFindingsProductArn() AwsInsight_FindingProviderFieldsRelatedFindingsProductArnPropertyList
	// Experimental.
	FindingProviderFieldsRelatedFindingsProductArnInput() interface{}
	// Experimental.
	FindingProviderFieldsSeverityLabel() AwsInsight_FindingProviderFieldsSeverityLabelPropertyList
	// Experimental.
	FindingProviderFieldsSeverityLabelInput() interface{}
	// Experimental.
	FindingProviderFieldsSeverityOriginal() AwsInsight_FindingProviderFieldsSeverityOriginalPropertyList
	// Experimental.
	FindingProviderFieldsSeverityOriginalInput() interface{}
	// Experimental.
	FindingProviderFieldsTypes() AwsInsight_FindingProviderFieldsTypesPropertyList
	// Experimental.
	FindingProviderFieldsTypesInput() interface{}
	// Experimental.
	FirstObservedAt() AwsInsight_FirstObservedAtPropertyList
	// Experimental.
	FirstObservedAtInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	GeneratorId() AwsInsight_GeneratorIdPropertyList
	// Experimental.
	GeneratorIdInput() interface{}
	// Experimental.
	Id() AwsInsight_IdPropertyList
	// Experimental.
	IdInput() interface{}
	// Experimental.
	InternalValue() *AwsInsight_FiltersProperty
	// Experimental.
	SetInternalValue(val *AwsInsight_FiltersProperty)
	// Experimental.
	Keyword() AwsInsight_KeywordPropertyList
	// Experimental.
	KeywordInput() interface{}
	// Experimental.
	LastObservedAt() AwsInsight_LastObservedAtPropertyList
	// Experimental.
	LastObservedAtInput() interface{}
	// Experimental.
	MalwareName() AwsInsight_MalwareNamePropertyList
	// Experimental.
	MalwareNameInput() interface{}
	// Experimental.
	MalwarePath() AwsInsight_MalwarePathPropertyList
	// Experimental.
	MalwarePathInput() interface{}
	// Experimental.
	MalwareState() AwsInsight_MalwareStatePropertyList
	// Experimental.
	MalwareStateInput() interface{}
	// Experimental.
	MalwareType() AwsInsight_MalwareTypePropertyList
	// Experimental.
	MalwareTypeInput() interface{}
	// Experimental.
	NetworkDestinationDomain() AwsInsight_NetworkDestinationDomainPropertyList
	// Experimental.
	NetworkDestinationDomainInput() interface{}
	// Experimental.
	NetworkDestinationIpv4() AwsInsight_NetworkDestinationIpv4PropertyList
	// Experimental.
	NetworkDestinationIpv4Input() interface{}
	// Experimental.
	NetworkDestinationIpv6() AwsInsight_NetworkDestinationIpv6PropertyList
	// Experimental.
	NetworkDestinationIpv6Input() interface{}
	// Experimental.
	NetworkDestinationPort() AwsInsight_NetworkDestinationPortPropertyList
	// Experimental.
	NetworkDestinationPortInput() interface{}
	// Experimental.
	NetworkDirection() AwsInsight_NetworkDirectionPropertyList
	// Experimental.
	NetworkDirectionInput() interface{}
	// Experimental.
	NetworkProtocol() AwsInsight_NetworkProtocolPropertyList
	// Experimental.
	NetworkProtocolInput() interface{}
	// Experimental.
	NetworkSourceDomain() AwsInsight_NetworkSourceDomainPropertyList
	// Experimental.
	NetworkSourceDomainInput() interface{}
	// Experimental.
	NetworkSourceIpv4() AwsInsight_NetworkSourceIpv4PropertyList
	// Experimental.
	NetworkSourceIpv4Input() interface{}
	// Experimental.
	NetworkSourceIpv6() AwsInsight_NetworkSourceIpv6PropertyList
	// Experimental.
	NetworkSourceIpv6Input() interface{}
	// Experimental.
	NetworkSourceMac() AwsInsight_NetworkSourceMacPropertyList
	// Experimental.
	NetworkSourceMacInput() interface{}
	// Experimental.
	NetworkSourcePort() AwsInsight_NetworkSourcePortPropertyList
	// Experimental.
	NetworkSourcePortInput() interface{}
	// Experimental.
	NoteText() AwsInsight_NoteTextPropertyList
	// Experimental.
	NoteTextInput() interface{}
	// Experimental.
	NoteUpdatedAt() AwsInsight_NoteUpdatedAtPropertyList
	// Experimental.
	NoteUpdatedAtInput() interface{}
	// Experimental.
	NoteUpdatedBy() AwsInsight_NoteUpdatedByPropertyList
	// Experimental.
	NoteUpdatedByInput() interface{}
	// Experimental.
	ProcessLaunchedAt() AwsInsight_ProcessLaunchedAtPropertyList
	// Experimental.
	ProcessLaunchedAtInput() interface{}
	// Experimental.
	ProcessName() AwsInsight_ProcessNamePropertyList
	// Experimental.
	ProcessNameInput() interface{}
	// Experimental.
	ProcessParentPid() AwsInsight_ProcessParentPidPropertyList
	// Experimental.
	ProcessParentPidInput() interface{}
	// Experimental.
	ProcessPath() AwsInsight_ProcessPathPropertyList
	// Experimental.
	ProcessPathInput() interface{}
	// Experimental.
	ProcessPid() AwsInsight_ProcessPidPropertyList
	// Experimental.
	ProcessPidInput() interface{}
	// Experimental.
	ProcessTerminatedAt() AwsInsight_ProcessTerminatedAtPropertyList
	// Experimental.
	ProcessTerminatedAtInput() interface{}
	// Experimental.
	ProductArn() AwsInsight_ProductArnPropertyList
	// Experimental.
	ProductArnInput() interface{}
	// Experimental.
	ProductFields() AwsInsight_ProductFieldsPropertyList
	// Experimental.
	ProductFieldsInput() interface{}
	// Experimental.
	ProductName() AwsInsight_ProductNamePropertyList
	// Experimental.
	ProductNameInput() interface{}
	// Experimental.
	RecommendationText() AwsInsight_RecommendationTextPropertyList
	// Experimental.
	RecommendationTextInput() interface{}
	// Experimental.
	RecordState() AwsInsight_RecordStatePropertyList
	// Experimental.
	RecordStateInput() interface{}
	// Experimental.
	RelatedFindingsId() AwsInsight_RelatedFindingsIdPropertyList
	// Experimental.
	RelatedFindingsIdInput() interface{}
	// Experimental.
	RelatedFindingsProductArn() AwsInsight_RelatedFindingsProductArnPropertyList
	// Experimental.
	RelatedFindingsProductArnInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceIamInstanceProfileArn() AwsInsight_ResourceAwsEc2InstanceIamInstanceProfileArnPropertyList
	// Experimental.
	ResourceAwsEc2InstanceIamInstanceProfileArnInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceImageId() AwsInsight_ResourceAwsEc2InstanceImageIdPropertyList
	// Experimental.
	ResourceAwsEc2InstanceImageIdInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceIpv4Addresses() AwsInsight_ResourceAwsEc2InstanceIpv4AddressesPropertyList
	// Experimental.
	ResourceAwsEc2InstanceIpv4AddressesInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceIpv6Addresses() AwsInsight_ResourceAwsEc2InstanceIpv6AddressesPropertyList
	// Experimental.
	ResourceAwsEc2InstanceIpv6AddressesInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceKeyName() AwsInsight_ResourceAwsEc2InstanceKeyNamePropertyList
	// Experimental.
	ResourceAwsEc2InstanceKeyNameInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceLaunchedAt() AwsInsight_ResourceAwsEc2InstanceLaunchedAtPropertyList
	// Experimental.
	ResourceAwsEc2InstanceLaunchedAtInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceSubnetId() AwsInsight_ResourceAwsEc2InstanceSubnetIdPropertyList
	// Experimental.
	ResourceAwsEc2InstanceSubnetIdInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceType() AwsInsight_ResourceAwsEc2InstanceTypePropertyList
	// Experimental.
	ResourceAwsEc2InstanceTypeInput() interface{}
	// Experimental.
	ResourceAwsEc2InstanceVpcId() AwsInsight_ResourceAwsEc2InstanceVpcIdPropertyList
	// Experimental.
	ResourceAwsEc2InstanceVpcIdInput() interface{}
	// Experimental.
	ResourceAwsIamAccessKeyCreatedAt() AwsInsight_ResourceAwsIamAccessKeyCreatedAtPropertyList
	// Experimental.
	ResourceAwsIamAccessKeyCreatedAtInput() interface{}
	// Experimental.
	ResourceAwsIamAccessKeyStatus() AwsInsight_ResourceAwsIamAccessKeyStatusPropertyList
	// Experimental.
	ResourceAwsIamAccessKeyStatusInput() interface{}
	// Experimental.
	ResourceAwsIamAccessKeyUserName() AwsInsight_ResourceAwsIamAccessKeyUserNamePropertyList
	// Experimental.
	ResourceAwsIamAccessKeyUserNameInput() interface{}
	// Experimental.
	ResourceAwsS3BucketOwnerId() AwsInsight_ResourceAwsS3BucketOwnerIdPropertyList
	// Experimental.
	ResourceAwsS3BucketOwnerIdInput() interface{}
	// Experimental.
	ResourceAwsS3BucketOwnerName() AwsInsight_ResourceAwsS3BucketOwnerNamePropertyList
	// Experimental.
	ResourceAwsS3BucketOwnerNameInput() interface{}
	// Experimental.
	ResourceContainerImageId() AwsInsight_ResourceContainerImageIdPropertyList
	// Experimental.
	ResourceContainerImageIdInput() interface{}
	// Experimental.
	ResourceContainerImageName() AwsInsight_ResourceContainerImageNamePropertyList
	// Experimental.
	ResourceContainerImageNameInput() interface{}
	// Experimental.
	ResourceContainerLaunchedAt() AwsInsight_ResourceContainerLaunchedAtPropertyList
	// Experimental.
	ResourceContainerLaunchedAtInput() interface{}
	// Experimental.
	ResourceContainerName() AwsInsight_ResourceContainerNamePropertyList
	// Experimental.
	ResourceContainerNameInput() interface{}
	// Experimental.
	ResourceDetailsOther() AwsInsight_ResourceDetailsOtherPropertyList
	// Experimental.
	ResourceDetailsOtherInput() interface{}
	// Experimental.
	ResourceId() AwsInsight_ResourceIdPropertyList
	// Experimental.
	ResourceIdInput() interface{}
	// Experimental.
	ResourcePartition() AwsInsight_ResourcePartitionPropertyList
	// Experimental.
	ResourcePartitionInput() interface{}
	// Experimental.
	ResourceRegion() AwsInsight_ResourceRegionPropertyList
	// Experimental.
	ResourceRegionInput() interface{}
	// Experimental.
	ResourceTags() AwsInsight_ResourceTagsPropertyList
	// Experimental.
	ResourceTagsInput() interface{}
	// Experimental.
	ResourceType() AwsInsight_ResourceTypePropertyList
	// Experimental.
	ResourceTypeInput() interface{}
	// Experimental.
	SeverityLabel() AwsInsight_SeverityLabelPropertyList
	// Experimental.
	SeverityLabelInput() interface{}
	// Experimental.
	SourceUrl() AwsInsight_SourceUrlPropertyList
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
	ThreatIntelIndicatorCategory() AwsInsight_ThreatIntelIndicatorCategoryPropertyList
	// Experimental.
	ThreatIntelIndicatorCategoryInput() interface{}
	// Experimental.
	ThreatIntelIndicatorLastObservedAt() AwsInsight_ThreatIntelIndicatorLastObservedAtPropertyList
	// Experimental.
	ThreatIntelIndicatorLastObservedAtInput() interface{}
	// Experimental.
	ThreatIntelIndicatorSource() AwsInsight_ThreatIntelIndicatorSourcePropertyList
	// Experimental.
	ThreatIntelIndicatorSourceInput() interface{}
	// Experimental.
	ThreatIntelIndicatorSourceUrl() AwsInsight_ThreatIntelIndicatorSourceUrlPropertyList
	// Experimental.
	ThreatIntelIndicatorSourceUrlInput() interface{}
	// Experimental.
	ThreatIntelIndicatorType() AwsInsight_ThreatIntelIndicatorTypePropertyList
	// Experimental.
	ThreatIntelIndicatorTypeInput() interface{}
	// Experimental.
	ThreatIntelIndicatorValue() AwsInsight_ThreatIntelIndicatorValuePropertyList
	// Experimental.
	ThreatIntelIndicatorValueInput() interface{}
	// Experimental.
	Title() AwsInsight_TitlePropertyList
	// Experimental.
	TitleInput() interface{}
	// Experimental.
	Type() AwsInsight_TypePropertyList
	// Experimental.
	TypeInput() interface{}
	// Experimental.
	UpdatedAt() AwsInsight_UpdatedAtPropertyList
	// Experimental.
	UpdatedAtInput() interface{}
	// Experimental.
	UserDefinedValues() AwsInsight_UserDefinedValuesPropertyList
	// Experimental.
	UserDefinedValuesInput() interface{}
	// Experimental.
	VerificationState() AwsInsight_VerificationStatePropertyList
	// Experimental.
	VerificationStateInput() interface{}
	// Experimental.
	WorkflowStatus() AwsInsight_WorkflowStatusPropertyList
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

// The jsii proxy struct for AwsInsight_FiltersPropertyOutputReference
type jsiiProxy_AwsInsight_FiltersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) AwsAccountId() AwsInsight_AwsAccountIdPropertyList {
	var returns AwsInsight_AwsAccountIdPropertyList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) AwsAccountName() AwsInsight_AwsAccountNamePropertyList {
	var returns AwsInsight_AwsAccountNamePropertyList
	_jsii_.Get(
		j,
		"awsAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) AwsAccountNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) CompanyName() AwsInsight_CompanyNamePropertyList {
	var returns AwsInsight_CompanyNamePropertyList
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) CompanyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplianceAssociatedStandardsId() AwsInsight_ComplianceAssociatedStandardsIdPropertyList {
	var returns AwsInsight_ComplianceAssociatedStandardsIdPropertyList
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplianceAssociatedStandardsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplianceSecurityControlId() AwsInsight_ComplianceSecurityControlIdPropertyList {
	var returns AwsInsight_ComplianceSecurityControlIdPropertyList
	_jsii_.Get(
		j,
		"complianceSecurityControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplianceSecurityControlIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityControlIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplianceSecurityControlParametersName() AwsInsight_ComplianceSecurityControlParametersNamePropertyList {
	var returns AwsInsight_ComplianceSecurityControlParametersNamePropertyList
	_jsii_.Get(
		j,
		"complianceSecurityControlParametersName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplianceSecurityControlParametersNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityControlParametersNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplianceSecurityControlParametersValue() AwsInsight_ComplianceSecurityControlParametersValuePropertyList {
	var returns AwsInsight_ComplianceSecurityControlParametersValuePropertyList
	_jsii_.Get(
		j,
		"complianceSecurityControlParametersValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplianceSecurityControlParametersValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceSecurityControlParametersValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplianceStatus() AwsInsight_ComplianceStatusPropertyList {
	var returns AwsInsight_ComplianceStatusPropertyList
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComplianceStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) Confidence() AwsInsight_ConfidencePropertyList {
	var returns AwsInsight_ConfidencePropertyList
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) CreatedAt() AwsInsight_CreatedAtPropertyList {
	var returns AwsInsight_CreatedAtPropertyList
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) CreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) Criticality() AwsInsight_CriticalityPropertyList {
	var returns AwsInsight_CriticalityPropertyList
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) CriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) Description() AwsInsight_DescriptionPropertyList {
	var returns AwsInsight_DescriptionPropertyList
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) DescriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsConfidence() AwsInsight_FindingProviderFieldsConfidencePropertyList {
	var returns AwsInsight_FindingProviderFieldsConfidencePropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsConfidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsConfidenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsConfidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsCriticality() AwsInsight_FindingProviderFieldsCriticalityPropertyList {
	var returns AwsInsight_FindingProviderFieldsCriticalityPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsCriticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsCriticalityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsCriticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsRelatedFindingsId() AwsInsight_FindingProviderFieldsRelatedFindingsIdPropertyList {
	var returns AwsInsight_FindingProviderFieldsRelatedFindingsIdPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsRelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsRelatedFindingsProductArn() AwsInsight_FindingProviderFieldsRelatedFindingsProductArnPropertyList {
	var returns AwsInsight_FindingProviderFieldsRelatedFindingsProductArnPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsRelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsSeverityLabel() AwsInsight_FindingProviderFieldsSeverityLabelPropertyList {
	var returns AwsInsight_FindingProviderFieldsSeverityLabelPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsSeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsSeverityOriginal() AwsInsight_FindingProviderFieldsSeverityOriginalPropertyList {
	var returns AwsInsight_FindingProviderFieldsSeverityOriginalPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityOriginal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsSeverityOriginalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityOriginalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsTypes() AwsInsight_FindingProviderFieldsTypesPropertyList {
	var returns AwsInsight_FindingProviderFieldsTypesPropertyList
	_jsii_.Get(
		j,
		"findingProviderFieldsTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FindingProviderFieldsTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingProviderFieldsTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FirstObservedAt() AwsInsight_FirstObservedAtPropertyList {
	var returns AwsInsight_FirstObservedAtPropertyList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) GeneratorId() AwsInsight_GeneratorIdPropertyList {
	var returns AwsInsight_GeneratorIdPropertyList
	_jsii_.Get(
		j,
		"generatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) GeneratorIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generatorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) Id() AwsInsight_IdPropertyList {
	var returns AwsInsight_IdPropertyList
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) IdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) InternalValue() *AwsInsight_FiltersProperty {
	var returns *AwsInsight_FiltersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) Keyword() AwsInsight_KeywordPropertyList {
	var returns AwsInsight_KeywordPropertyList
	_jsii_.Get(
		j,
		"keyword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) KeywordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keywordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) LastObservedAt() AwsInsight_LastObservedAtPropertyList {
	var returns AwsInsight_LastObservedAtPropertyList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) MalwareName() AwsInsight_MalwareNamePropertyList {
	var returns AwsInsight_MalwareNamePropertyList
	_jsii_.Get(
		j,
		"malwareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) MalwareNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) MalwarePath() AwsInsight_MalwarePathPropertyList {
	var returns AwsInsight_MalwarePathPropertyList
	_jsii_.Get(
		j,
		"malwarePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) MalwarePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwarePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) MalwareState() AwsInsight_MalwareStatePropertyList {
	var returns AwsInsight_MalwareStatePropertyList
	_jsii_.Get(
		j,
		"malwareState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) MalwareStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) MalwareType() AwsInsight_MalwareTypePropertyList {
	var returns AwsInsight_MalwareTypePropertyList
	_jsii_.Get(
		j,
		"malwareType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) MalwareTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkDestinationDomain() AwsInsight_NetworkDestinationDomainPropertyList {
	var returns AwsInsight_NetworkDestinationDomainPropertyList
	_jsii_.Get(
		j,
		"networkDestinationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkDestinationDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkDestinationIpv4() AwsInsight_NetworkDestinationIpv4PropertyList {
	var returns AwsInsight_NetworkDestinationIpv4PropertyList
	_jsii_.Get(
		j,
		"networkDestinationIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkDestinationIpv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkDestinationIpv6() AwsInsight_NetworkDestinationIpv6PropertyList {
	var returns AwsInsight_NetworkDestinationIpv6PropertyList
	_jsii_.Get(
		j,
		"networkDestinationIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkDestinationIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkDestinationPort() AwsInsight_NetworkDestinationPortPropertyList {
	var returns AwsInsight_NetworkDestinationPortPropertyList
	_jsii_.Get(
		j,
		"networkDestinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkDestinationPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDestinationPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkDirection() AwsInsight_NetworkDirectionPropertyList {
	var returns AwsInsight_NetworkDirectionPropertyList
	_jsii_.Get(
		j,
		"networkDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkDirectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkProtocol() AwsInsight_NetworkProtocolPropertyList {
	var returns AwsInsight_NetworkProtocolPropertyList
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkSourceDomain() AwsInsight_NetworkSourceDomainPropertyList {
	var returns AwsInsight_NetworkSourceDomainPropertyList
	_jsii_.Get(
		j,
		"networkSourceDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkSourceDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkSourceIpv4() AwsInsight_NetworkSourceIpv4PropertyList {
	var returns AwsInsight_NetworkSourceIpv4PropertyList
	_jsii_.Get(
		j,
		"networkSourceIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkSourceIpv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceIpv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkSourceIpv6() AwsInsight_NetworkSourceIpv6PropertyList {
	var returns AwsInsight_NetworkSourceIpv6PropertyList
	_jsii_.Get(
		j,
		"networkSourceIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkSourceIpv6Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceIpv6Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkSourceMac() AwsInsight_NetworkSourceMacPropertyList {
	var returns AwsInsight_NetworkSourceMacPropertyList
	_jsii_.Get(
		j,
		"networkSourceMac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkSourceMacInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourceMacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkSourcePort() AwsInsight_NetworkSourcePortPropertyList {
	var returns AwsInsight_NetworkSourcePortPropertyList
	_jsii_.Get(
		j,
		"networkSourcePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NetworkSourcePortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkSourcePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NoteText() AwsInsight_NoteTextPropertyList {
	var returns AwsInsight_NoteTextPropertyList
	_jsii_.Get(
		j,
		"noteText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NoteTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NoteUpdatedAt() AwsInsight_NoteUpdatedAtPropertyList {
	var returns AwsInsight_NoteUpdatedAtPropertyList
	_jsii_.Get(
		j,
		"noteUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NoteUpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NoteUpdatedBy() AwsInsight_NoteUpdatedByPropertyList {
	var returns AwsInsight_NoteUpdatedByPropertyList
	_jsii_.Get(
		j,
		"noteUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) NoteUpdatedByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteUpdatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessLaunchedAt() AwsInsight_ProcessLaunchedAtPropertyList {
	var returns AwsInsight_ProcessLaunchedAtPropertyList
	_jsii_.Get(
		j,
		"processLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessName() AwsInsight_ProcessNamePropertyList {
	var returns AwsInsight_ProcessNamePropertyList
	_jsii_.Get(
		j,
		"processName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessParentPid() AwsInsight_ProcessParentPidPropertyList {
	var returns AwsInsight_ProcessParentPidPropertyList
	_jsii_.Get(
		j,
		"processParentPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessParentPidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processParentPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessPath() AwsInsight_ProcessPathPropertyList {
	var returns AwsInsight_ProcessPathPropertyList
	_jsii_.Get(
		j,
		"processPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessPid() AwsInsight_ProcessPidPropertyList {
	var returns AwsInsight_ProcessPidPropertyList
	_jsii_.Get(
		j,
		"processPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessPidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessTerminatedAt() AwsInsight_ProcessTerminatedAtPropertyList {
	var returns AwsInsight_ProcessTerminatedAtPropertyList
	_jsii_.Get(
		j,
		"processTerminatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProcessTerminatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processTerminatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProductArn() AwsInsight_ProductArnPropertyList {
	var returns AwsInsight_ProductArnPropertyList
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProductFields() AwsInsight_ProductFieldsPropertyList {
	var returns AwsInsight_ProductFieldsPropertyList
	_jsii_.Get(
		j,
		"productFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProductFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProductName() AwsInsight_ProductNamePropertyList {
	var returns AwsInsight_ProductNamePropertyList
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ProductNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) RecommendationText() AwsInsight_RecommendationTextPropertyList {
	var returns AwsInsight_RecommendationTextPropertyList
	_jsii_.Get(
		j,
		"recommendationText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) RecommendationTextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recommendationTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) RecordState() AwsInsight_RecordStatePropertyList {
	var returns AwsInsight_RecordStatePropertyList
	_jsii_.Get(
		j,
		"recordState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) RecordStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) RelatedFindingsId() AwsInsight_RelatedFindingsIdPropertyList {
	var returns AwsInsight_RelatedFindingsIdPropertyList
	_jsii_.Get(
		j,
		"relatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) RelatedFindingsIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) RelatedFindingsProductArn() AwsInsight_RelatedFindingsProductArnPropertyList {
	var returns AwsInsight_RelatedFindingsProductArnPropertyList
	_jsii_.Get(
		j,
		"relatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) RelatedFindingsProductArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsProductArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIamInstanceProfileArn() AwsInsight_ResourceAwsEc2InstanceIamInstanceProfileArnPropertyList {
	var returns AwsInsight_ResourceAwsEc2InstanceIamInstanceProfileArnPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIamInstanceProfileArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceImageId() AwsInsight_ResourceAwsEc2InstanceImageIdPropertyList {
	var returns AwsInsight_ResourceAwsEc2InstanceImageIdPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIpv4Addresses() AwsInsight_ResourceAwsEc2InstanceIpv4AddressesPropertyList {
	var returns AwsInsight_ResourceAwsEc2InstanceIpv4AddressesPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIpv4AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIpv6Addresses() AwsInsight_ResourceAwsEc2InstanceIpv6AddressesPropertyList {
	var returns AwsInsight_ResourceAwsEc2InstanceIpv6AddressesPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceIpv6AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceKeyName() AwsInsight_ResourceAwsEc2InstanceKeyNamePropertyList {
	var returns AwsInsight_ResourceAwsEc2InstanceKeyNamePropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceKeyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceLaunchedAt() AwsInsight_ResourceAwsEc2InstanceLaunchedAtPropertyList {
	var returns AwsInsight_ResourceAwsEc2InstanceLaunchedAtPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceSubnetId() AwsInsight_ResourceAwsEc2InstanceSubnetIdPropertyList {
	var returns AwsInsight_ResourceAwsEc2InstanceSubnetIdPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceSubnetIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceType() AwsInsight_ResourceAwsEc2InstanceTypePropertyList {
	var returns AwsInsight_ResourceAwsEc2InstanceTypePropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceVpcId() AwsInsight_ResourceAwsEc2InstanceVpcIdPropertyList {
	var returns AwsInsight_ResourceAwsEc2InstanceVpcIdPropertyList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsEc2InstanceVpcIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyCreatedAt() AwsInsight_ResourceAwsIamAccessKeyCreatedAtPropertyList {
	var returns AwsInsight_ResourceAwsIamAccessKeyCreatedAtPropertyList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyCreatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyCreatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyStatus() AwsInsight_ResourceAwsIamAccessKeyStatusPropertyList {
	var returns AwsInsight_ResourceAwsIamAccessKeyStatusPropertyList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyUserName() AwsInsight_ResourceAwsIamAccessKeyUserNamePropertyList {
	var returns AwsInsight_ResourceAwsIamAccessKeyUserNamePropertyList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsIamAccessKeyUserNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsS3BucketOwnerId() AwsInsight_ResourceAwsS3BucketOwnerIdPropertyList {
	var returns AwsInsight_ResourceAwsS3BucketOwnerIdPropertyList
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsS3BucketOwnerIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsS3BucketOwnerName() AwsInsight_ResourceAwsS3BucketOwnerNamePropertyList {
	var returns AwsInsight_ResourceAwsS3BucketOwnerNamePropertyList
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceAwsS3BucketOwnerNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceContainerImageId() AwsInsight_ResourceContainerImageIdPropertyList {
	var returns AwsInsight_ResourceContainerImageIdPropertyList
	_jsii_.Get(
		j,
		"resourceContainerImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceContainerImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceContainerImageName() AwsInsight_ResourceContainerImageNamePropertyList {
	var returns AwsInsight_ResourceContainerImageNamePropertyList
	_jsii_.Get(
		j,
		"resourceContainerImageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceContainerImageNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerImageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceContainerLaunchedAt() AwsInsight_ResourceContainerLaunchedAtPropertyList {
	var returns AwsInsight_ResourceContainerLaunchedAtPropertyList
	_jsii_.Get(
		j,
		"resourceContainerLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceContainerLaunchedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerLaunchedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceContainerName() AwsInsight_ResourceContainerNamePropertyList {
	var returns AwsInsight_ResourceContainerNamePropertyList
	_jsii_.Get(
		j,
		"resourceContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceContainerNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceDetailsOther() AwsInsight_ResourceDetailsOtherPropertyList {
	var returns AwsInsight_ResourceDetailsOtherPropertyList
	_jsii_.Get(
		j,
		"resourceDetailsOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceDetailsOtherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceDetailsOtherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceId() AwsInsight_ResourceIdPropertyList {
	var returns AwsInsight_ResourceIdPropertyList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourcePartition() AwsInsight_ResourcePartitionPropertyList {
	var returns AwsInsight_ResourcePartitionPropertyList
	_jsii_.Get(
		j,
		"resourcePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourcePartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceRegion() AwsInsight_ResourceRegionPropertyList {
	var returns AwsInsight_ResourceRegionPropertyList
	_jsii_.Get(
		j,
		"resourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceTags() AwsInsight_ResourceTagsPropertyList {
	var returns AwsInsight_ResourceTagsPropertyList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceType() AwsInsight_ResourceTypePropertyList {
	var returns AwsInsight_ResourceTypePropertyList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) SeverityLabel() AwsInsight_SeverityLabelPropertyList {
	var returns AwsInsight_SeverityLabelPropertyList
	_jsii_.Get(
		j,
		"severityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) SeverityLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) SourceUrl() AwsInsight_SourceUrlPropertyList {
	var returns AwsInsight_SourceUrlPropertyList
	_jsii_.Get(
		j,
		"sourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) SourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorCategory() AwsInsight_ThreatIntelIndicatorCategoryPropertyList {
	var returns AwsInsight_ThreatIntelIndicatorCategoryPropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorCategoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorLastObservedAt() AwsInsight_ThreatIntelIndicatorLastObservedAtPropertyList {
	var returns AwsInsight_ThreatIntelIndicatorLastObservedAtPropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorLastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorLastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorLastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorSource() AwsInsight_ThreatIntelIndicatorSourcePropertyList {
	var returns AwsInsight_ThreatIntelIndicatorSourcePropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorSourceUrl() AwsInsight_ThreatIntelIndicatorSourceUrlPropertyList {
	var returns AwsInsight_ThreatIntelIndicatorSourceUrlPropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorSourceUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorType() AwsInsight_ThreatIntelIndicatorTypePropertyList {
	var returns AwsInsight_ThreatIntelIndicatorTypePropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorValue() AwsInsight_ThreatIntelIndicatorValuePropertyList {
	var returns AwsInsight_ThreatIntelIndicatorValuePropertyList
	_jsii_.Get(
		j,
		"threatIntelIndicatorValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ThreatIntelIndicatorValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"threatIntelIndicatorValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) Title() AwsInsight_TitlePropertyList {
	var returns AwsInsight_TitlePropertyList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) Type() AwsInsight_TypePropertyList {
	var returns AwsInsight_TypePropertyList
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) TypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) UpdatedAt() AwsInsight_UpdatedAtPropertyList {
	var returns AwsInsight_UpdatedAtPropertyList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) UserDefinedValues() AwsInsight_UserDefinedValuesPropertyList {
	var returns AwsInsight_UserDefinedValuesPropertyList
	_jsii_.Get(
		j,
		"userDefinedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) UserDefinedValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) VerificationState() AwsInsight_VerificationStatePropertyList {
	var returns AwsInsight_VerificationStatePropertyList
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) VerificationStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) WorkflowStatus() AwsInsight_WorkflowStatusPropertyList {
	var returns AwsInsight_WorkflowStatusPropertyList
	_jsii_.Get(
		j,
		"workflowStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) WorkflowStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowStatusInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsInsight_FiltersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsInsight_FiltersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsInsight_FiltersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsInsight_FiltersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsInsight.FiltersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsInsight_FiltersPropertyOutputReference_Override(a AwsInsight_FiltersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsInsight.FiltersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference)SetInternalValue(val *AwsInsight_FiltersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutAwsAccountId(value interface{}) {
	if err := a.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutAwsAccountName(value interface{}) {
	if err := a.validatePutAwsAccountNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsAccountName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutCompanyName(value interface{}) {
	if err := a.validatePutCompanyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCompanyName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutComplianceAssociatedStandardsId(value interface{}) {
	if err := a.validatePutComplianceAssociatedStandardsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComplianceAssociatedStandardsId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutComplianceSecurityControlId(value interface{}) {
	if err := a.validatePutComplianceSecurityControlIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComplianceSecurityControlId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutComplianceSecurityControlParametersName(value interface{}) {
	if err := a.validatePutComplianceSecurityControlParametersNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComplianceSecurityControlParametersName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutComplianceSecurityControlParametersValue(value interface{}) {
	if err := a.validatePutComplianceSecurityControlParametersValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComplianceSecurityControlParametersValue",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutComplianceStatus(value interface{}) {
	if err := a.validatePutComplianceStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComplianceStatus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutConfidence(value interface{}) {
	if err := a.validatePutConfidenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfidence",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutCreatedAt(value interface{}) {
	if err := a.validatePutCreatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCreatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutCriticality(value interface{}) {
	if err := a.validatePutCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCriticality",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutDescription(value interface{}) {
	if err := a.validatePutDescriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDescription",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsConfidence(value interface{}) {
	if err := a.validatePutFindingProviderFieldsConfidenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingProviderFieldsConfidence",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsCriticality(value interface{}) {
	if err := a.validatePutFindingProviderFieldsCriticalityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingProviderFieldsCriticality",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsRelatedFindingsId(value interface{}) {
	if err := a.validatePutFindingProviderFieldsRelatedFindingsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingProviderFieldsRelatedFindingsId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsRelatedFindingsProductArn(value interface{}) {
	if err := a.validatePutFindingProviderFieldsRelatedFindingsProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingProviderFieldsRelatedFindingsProductArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsSeverityLabel(value interface{}) {
	if err := a.validatePutFindingProviderFieldsSeverityLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingProviderFieldsSeverityLabel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsSeverityOriginal(value interface{}) {
	if err := a.validatePutFindingProviderFieldsSeverityOriginalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingProviderFieldsSeverityOriginal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutFindingProviderFieldsTypes(value interface{}) {
	if err := a.validatePutFindingProviderFieldsTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingProviderFieldsTypes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutFirstObservedAt(value interface{}) {
	if err := a.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutGeneratorId(value interface{}) {
	if err := a.validatePutGeneratorIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGeneratorId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutId(value interface{}) {
	if err := a.validatePutIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutKeyword(value interface{}) {
	if err := a.validatePutKeywordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKeyword",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutLastObservedAt(value interface{}) {
	if err := a.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutMalwareName(value interface{}) {
	if err := a.validatePutMalwareNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMalwareName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutMalwarePath(value interface{}) {
	if err := a.validatePutMalwarePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMalwarePath",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutMalwareState(value interface{}) {
	if err := a.validatePutMalwareStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMalwareState",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutMalwareType(value interface{}) {
	if err := a.validatePutMalwareTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMalwareType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNetworkDestinationDomain(value interface{}) {
	if err := a.validatePutNetworkDestinationDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkDestinationDomain",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNetworkDestinationIpv4(value interface{}) {
	if err := a.validatePutNetworkDestinationIpv4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkDestinationIpv4",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNetworkDestinationIpv6(value interface{}) {
	if err := a.validatePutNetworkDestinationIpv6Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkDestinationIpv6",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNetworkDestinationPort(value interface{}) {
	if err := a.validatePutNetworkDestinationPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkDestinationPort",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNetworkDirection(value interface{}) {
	if err := a.validatePutNetworkDirectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkDirection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNetworkProtocol(value interface{}) {
	if err := a.validatePutNetworkProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkProtocol",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNetworkSourceDomain(value interface{}) {
	if err := a.validatePutNetworkSourceDomainParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkSourceDomain",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNetworkSourceIpv4(value interface{}) {
	if err := a.validatePutNetworkSourceIpv4Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkSourceIpv4",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNetworkSourceIpv6(value interface{}) {
	if err := a.validatePutNetworkSourceIpv6Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkSourceIpv6",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNetworkSourceMac(value interface{}) {
	if err := a.validatePutNetworkSourceMacParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkSourceMac",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNetworkSourcePort(value interface{}) {
	if err := a.validatePutNetworkSourcePortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkSourcePort",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNoteText(value interface{}) {
	if err := a.validatePutNoteTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoteText",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNoteUpdatedAt(value interface{}) {
	if err := a.validatePutNoteUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoteUpdatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutNoteUpdatedBy(value interface{}) {
	if err := a.validatePutNoteUpdatedByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoteUpdatedBy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutProcessLaunchedAt(value interface{}) {
	if err := a.validatePutProcessLaunchedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessLaunchedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutProcessName(value interface{}) {
	if err := a.validatePutProcessNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutProcessParentPid(value interface{}) {
	if err := a.validatePutProcessParentPidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessParentPid",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutProcessPath(value interface{}) {
	if err := a.validatePutProcessPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessPath",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutProcessPid(value interface{}) {
	if err := a.validatePutProcessPidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessPid",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutProcessTerminatedAt(value interface{}) {
	if err := a.validatePutProcessTerminatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessTerminatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutProductArn(value interface{}) {
	if err := a.validatePutProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProductArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutProductFields(value interface{}) {
	if err := a.validatePutProductFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProductFields",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutProductName(value interface{}) {
	if err := a.validatePutProductNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProductName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutRecommendationText(value interface{}) {
	if err := a.validatePutRecommendationTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRecommendationText",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutRecordState(value interface{}) {
	if err := a.validatePutRecordStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRecordState",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutRelatedFindingsId(value interface{}) {
	if err := a.validatePutRelatedFindingsIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelatedFindingsId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutRelatedFindingsProductArn(value interface{}) {
	if err := a.validatePutRelatedFindingsProductArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelatedFindingsProductArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceIamInstanceProfileArn(value interface{}) {
	if err := a.validatePutResourceAwsEc2InstanceIamInstanceProfileArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsEc2InstanceIamInstanceProfileArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceImageId(value interface{}) {
	if err := a.validatePutResourceAwsEc2InstanceImageIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsEc2InstanceImageId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceIpv4Addresses(value interface{}) {
	if err := a.validatePutResourceAwsEc2InstanceIpv4AddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsEc2InstanceIpv4Addresses",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceIpv6Addresses(value interface{}) {
	if err := a.validatePutResourceAwsEc2InstanceIpv6AddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsEc2InstanceIpv6Addresses",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceKeyName(value interface{}) {
	if err := a.validatePutResourceAwsEc2InstanceKeyNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsEc2InstanceKeyName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceLaunchedAt(value interface{}) {
	if err := a.validatePutResourceAwsEc2InstanceLaunchedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsEc2InstanceLaunchedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceSubnetId(value interface{}) {
	if err := a.validatePutResourceAwsEc2InstanceSubnetIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsEc2InstanceSubnetId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceType(value interface{}) {
	if err := a.validatePutResourceAwsEc2InstanceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsEc2InstanceType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsEc2InstanceVpcId(value interface{}) {
	if err := a.validatePutResourceAwsEc2InstanceVpcIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsEc2InstanceVpcId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsIamAccessKeyCreatedAt(value interface{}) {
	if err := a.validatePutResourceAwsIamAccessKeyCreatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsIamAccessKeyCreatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsIamAccessKeyStatus(value interface{}) {
	if err := a.validatePutResourceAwsIamAccessKeyStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsIamAccessKeyStatus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsIamAccessKeyUserName(value interface{}) {
	if err := a.validatePutResourceAwsIamAccessKeyUserNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsIamAccessKeyUserName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsS3BucketOwnerId(value interface{}) {
	if err := a.validatePutResourceAwsS3BucketOwnerIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsS3BucketOwnerId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceAwsS3BucketOwnerName(value interface{}) {
	if err := a.validatePutResourceAwsS3BucketOwnerNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceAwsS3BucketOwnerName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceContainerImageId(value interface{}) {
	if err := a.validatePutResourceContainerImageIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceContainerImageId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceContainerImageName(value interface{}) {
	if err := a.validatePutResourceContainerImageNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceContainerImageName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceContainerLaunchedAt(value interface{}) {
	if err := a.validatePutResourceContainerLaunchedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceContainerLaunchedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceContainerName(value interface{}) {
	if err := a.validatePutResourceContainerNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceContainerName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceDetailsOther(value interface{}) {
	if err := a.validatePutResourceDetailsOtherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceDetailsOther",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceId(value interface{}) {
	if err := a.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourcePartition(value interface{}) {
	if err := a.validatePutResourcePartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourcePartition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceRegion(value interface{}) {
	if err := a.validatePutResourceRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceRegion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceTags(value interface{}) {
	if err := a.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutResourceType(value interface{}) {
	if err := a.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutSeverityLabel(value interface{}) {
	if err := a.validatePutSeverityLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSeverityLabel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutSourceUrl(value interface{}) {
	if err := a.validatePutSourceUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceUrl",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorCategory(value interface{}) {
	if err := a.validatePutThreatIntelIndicatorCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putThreatIntelIndicatorCategory",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorLastObservedAt(value interface{}) {
	if err := a.validatePutThreatIntelIndicatorLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putThreatIntelIndicatorLastObservedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorSource(value interface{}) {
	if err := a.validatePutThreatIntelIndicatorSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putThreatIntelIndicatorSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorSourceUrl(value interface{}) {
	if err := a.validatePutThreatIntelIndicatorSourceUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putThreatIntelIndicatorSourceUrl",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorType(value interface{}) {
	if err := a.validatePutThreatIntelIndicatorTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putThreatIntelIndicatorType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutThreatIntelIndicatorValue(value interface{}) {
	if err := a.validatePutThreatIntelIndicatorValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putThreatIntelIndicatorValue",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutTitle(value interface{}) {
	if err := a.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTitle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutType(value interface{}) {
	if err := a.validatePutTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutUpdatedAt(value interface{}) {
	if err := a.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutUserDefinedValues(value interface{}) {
	if err := a.validatePutUserDefinedValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserDefinedValues",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutVerificationState(value interface{}) {
	if err := a.validatePutVerificationStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVerificationState",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) PutWorkflowStatus(value interface{}) {
	if err := a.validatePutWorkflowStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkflowStatus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetAwsAccountName() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsAccountName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetCompanyName() {
	_jsii_.InvokeVoid(
		a,
		"resetCompanyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetComplianceAssociatedStandardsId() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceAssociatedStandardsId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetComplianceSecurityControlId() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceSecurityControlId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetComplianceSecurityControlParametersName() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceSecurityControlParametersName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetComplianceSecurityControlParametersValue() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceSecurityControlParametersValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetComplianceStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetConfidence() {
	_jsii_.InvokeVoid(
		a,
		"resetConfidence",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		a,
		"resetCriticality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsConfidence() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingProviderFieldsConfidence",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsCriticality() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingProviderFieldsCriticality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsRelatedFindingsId() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingProviderFieldsRelatedFindingsId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingProviderFieldsRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsSeverityLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingProviderFieldsSeverityLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsSeverityOriginal() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingProviderFieldsSeverityOriginal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetFindingProviderFieldsTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingProviderFieldsTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetGeneratorId() {
	_jsii_.InvokeVoid(
		a,
		"resetGeneratorId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetKeyword() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetMalwareName() {
	_jsii_.InvokeVoid(
		a,
		"resetMalwareName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetMalwarePath() {
	_jsii_.InvokeVoid(
		a,
		"resetMalwarePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetMalwareState() {
	_jsii_.InvokeVoid(
		a,
		"resetMalwareState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetMalwareType() {
	_jsii_.InvokeVoid(
		a,
		"resetMalwareType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNetworkDestinationDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkDestinationDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNetworkDestinationIpv4() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkDestinationIpv4",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNetworkDestinationIpv6() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkDestinationIpv6",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNetworkDestinationPort() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkDestinationPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNetworkDirection() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkDirection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNetworkProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNetworkSourceDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkSourceDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNetworkSourceIpv4() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkSourceIpv4",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNetworkSourceIpv6() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkSourceIpv6",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNetworkSourceMac() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkSourceMac",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNetworkSourcePort() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkSourcePort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNoteText() {
	_jsii_.InvokeVoid(
		a,
		"resetNoteText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNoteUpdatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetNoteUpdatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetNoteUpdatedBy() {
	_jsii_.InvokeVoid(
		a,
		"resetNoteUpdatedBy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetProcessLaunchedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessLaunchedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetProcessName() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetProcessParentPid() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessParentPid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetProcessPath() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetProcessPid() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessPid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetProcessTerminatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessTerminatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetProductArn() {
	_jsii_.InvokeVoid(
		a,
		"resetProductArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetProductFields() {
	_jsii_.InvokeVoid(
		a,
		"resetProductFields",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetProductName() {
	_jsii_.InvokeVoid(
		a,
		"resetProductName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetRecommendationText() {
	_jsii_.InvokeVoid(
		a,
		"resetRecommendationText",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetRecordState() {
	_jsii_.InvokeVoid(
		a,
		"resetRecordState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetRelatedFindingsId() {
	_jsii_.InvokeVoid(
		a,
		"resetRelatedFindingsId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetRelatedFindingsProductArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRelatedFindingsProductArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceIamInstanceProfileArn() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsEc2InstanceIamInstanceProfileArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceImageId() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsEc2InstanceImageId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceIpv4Addresses() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsEc2InstanceIpv4Addresses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceIpv6Addresses() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsEc2InstanceIpv6Addresses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceKeyName() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsEc2InstanceKeyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceLaunchedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsEc2InstanceLaunchedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceSubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsEc2InstanceSubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsEc2InstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsEc2InstanceVpcId() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsEc2InstanceVpcId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsIamAccessKeyCreatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsIamAccessKeyCreatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsIamAccessKeyStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsIamAccessKeyStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsIamAccessKeyUserName() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsIamAccessKeyUserName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsS3BucketOwnerId() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsS3BucketOwnerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceAwsS3BucketOwnerName() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceAwsS3BucketOwnerName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceContainerImageId() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceContainerImageId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceContainerImageName() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceContainerImageName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceContainerLaunchedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceContainerLaunchedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceContainerName() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceContainerName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceDetailsOther() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceDetailsOther",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourcePartition() {
	_jsii_.InvokeVoid(
		a,
		"resetResourcePartition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetSeverityLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetSeverityLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetSourceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorCategory() {
	_jsii_.InvokeVoid(
		a,
		"resetThreatIntelIndicatorCategory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorLastObservedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetThreatIntelIndicatorLastObservedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorSource() {
	_jsii_.InvokeVoid(
		a,
		"resetThreatIntelIndicatorSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorSourceUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetThreatIntelIndicatorSourceUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorType() {
	_jsii_.InvokeVoid(
		a,
		"resetThreatIntelIndicatorType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetThreatIntelIndicatorValue() {
	_jsii_.InvokeVoid(
		a,
		"resetThreatIntelIndicatorValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		a,
		"resetTitle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetUserDefinedValues() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDefinedValues",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetVerificationState() {
	_jsii_.InvokeVoid(
		a,
		"resetVerificationState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ResetWorkflowStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkflowStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

