package inspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/inspector/jsii"

	"github.com/cdktn-io/cdktn-aws-go/inspector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFilter_FilterCriteriaPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsAccountId() AwsFilter_AwsAccountIdPropertyList
	// Experimental.
	AwsAccountIdInput() interface{}
	// Experimental.
	CodeRepositoryProjectName() AwsFilter_CodeRepositoryProjectNamePropertyList
	// Experimental.
	CodeRepositoryProjectNameInput() interface{}
	// Experimental.
	CodeRepositoryProviderType() AwsFilter_CodeRepositoryProviderTypePropertyList
	// Experimental.
	CodeRepositoryProviderTypeInput() interface{}
	// Experimental.
	CodeVulnerabilityDetectorName() AwsFilter_CodeVulnerabilityDetectorNamePropertyList
	// Experimental.
	CodeVulnerabilityDetectorNameInput() interface{}
	// Experimental.
	CodeVulnerabilityDetectorTags() AwsFilter_CodeVulnerabilityDetectorTagsPropertyList
	// Experimental.
	CodeVulnerabilityDetectorTagsInput() interface{}
	// Experimental.
	CodeVulnerabilityFilePath() AwsFilter_CodeVulnerabilityFilePathPropertyList
	// Experimental.
	CodeVulnerabilityFilePathInput() interface{}
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
	ComponentId() AwsFilter_ComponentIdPropertyList
	// Experimental.
	ComponentIdInput() interface{}
	// Experimental.
	ComponentType() AwsFilter_ComponentTypePropertyList
	// Experimental.
	ComponentTypeInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Ec2InstanceImageId() AwsFilter_Ec2InstanceImageIdPropertyList
	// Experimental.
	Ec2InstanceImageIdInput() interface{}
	// Experimental.
	Ec2InstanceSubnetId() AwsFilter_Ec2InstanceSubnetIdPropertyList
	// Experimental.
	Ec2InstanceSubnetIdInput() interface{}
	// Experimental.
	Ec2InstanceVpcId() AwsFilter_Ec2InstanceVpcIdPropertyList
	// Experimental.
	Ec2InstanceVpcIdInput() interface{}
	// Experimental.
	EcrImageArchitecture() AwsFilter_EcrImageArchitecturePropertyList
	// Experimental.
	EcrImageArchitectureInput() interface{}
	// Experimental.
	EcrImageHash() AwsFilter_EcrImageHashPropertyList
	// Experimental.
	EcrImageHashInput() interface{}
	// Experimental.
	EcrImageInUseCount() AwsFilter_EcrImageInUseCountPropertyList
	// Experimental.
	EcrImageInUseCountInput() interface{}
	// Experimental.
	EcrImageLastInUseAt() AwsFilter_EcrImageLastInUseAtPropertyList
	// Experimental.
	EcrImageLastInUseAtInput() interface{}
	// Experimental.
	EcrImagePushedAt() AwsFilter_EcrImagePushedAtPropertyList
	// Experimental.
	EcrImagePushedAtInput() interface{}
	// Experimental.
	EcrImageRegistry() AwsFilter_EcrImageRegistryPropertyList
	// Experimental.
	EcrImageRegistryInput() interface{}
	// Experimental.
	EcrImageRepositoryName() AwsFilter_EcrImageRepositoryNamePropertyList
	// Experimental.
	EcrImageRepositoryNameInput() interface{}
	// Experimental.
	EcrImageTags() AwsFilter_EcrImageTagsPropertyList
	// Experimental.
	EcrImageTagsInput() interface{}
	// Experimental.
	EpssScore() AwsFilter_EpssScorePropertyList
	// Experimental.
	EpssScoreInput() interface{}
	// Experimental.
	ExploitAvailable() AwsFilter_ExploitAvailablePropertyList
	// Experimental.
	ExploitAvailableInput() interface{}
	// Experimental.
	FindingArn() AwsFilter_FindingArnPropertyList
	// Experimental.
	FindingArnInput() interface{}
	// Experimental.
	FindingStatus() AwsFilter_FindingStatusPropertyList
	// Experimental.
	FindingStatusInput() interface{}
	// Experimental.
	FindingType() AwsFilter_FindingTypePropertyList
	// Experimental.
	FindingTypeInput() interface{}
	// Experimental.
	FirstObservedAt() AwsFilter_FirstObservedAtPropertyList
	// Experimental.
	FirstObservedAtInput() interface{}
	// Experimental.
	FixAvailable() AwsFilter_FixAvailablePropertyList
	// Experimental.
	FixAvailableInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InspectorScore() AwsFilter_InspectorScorePropertyList
	// Experimental.
	InspectorScoreInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LambdaFunctionExecutionRoleArn() AwsFilter_LambdaFunctionExecutionRoleArnPropertyList
	// Experimental.
	LambdaFunctionExecutionRoleArnInput() interface{}
	// Experimental.
	LambdaFunctionLastModifiedAt() AwsFilter_LambdaFunctionLastModifiedAtPropertyList
	// Experimental.
	LambdaFunctionLastModifiedAtInput() interface{}
	// Experimental.
	LambdaFunctionLayers() AwsFilter_LambdaFunctionLayersPropertyList
	// Experimental.
	LambdaFunctionLayersInput() interface{}
	// Experimental.
	LambdaFunctionName() AwsFilter_LambdaFunctionNamePropertyList
	// Experimental.
	LambdaFunctionNameInput() interface{}
	// Experimental.
	LambdaFunctionRuntime() AwsFilter_LambdaFunctionRuntimePropertyList
	// Experimental.
	LambdaFunctionRuntimeInput() interface{}
	// Experimental.
	LastObservedAt() AwsFilter_LastObservedAtPropertyList
	// Experimental.
	LastObservedAtInput() interface{}
	// Experimental.
	NetworkProtocol() AwsFilter_NetworkProtocolPropertyList
	// Experimental.
	NetworkProtocolInput() interface{}
	// Experimental.
	PortRange() AwsFilter_PortRangePropertyList
	// Experimental.
	PortRangeInput() interface{}
	// Experimental.
	RelatedVulnerabilities() AwsFilter_RelatedVulnerabilitiesPropertyList
	// Experimental.
	RelatedVulnerabilitiesInput() interface{}
	// Experimental.
	ResourceId() AwsFilter_ResourceIdPropertyList
	// Experimental.
	ResourceIdInput() interface{}
	// Experimental.
	ResourceTags() AwsFilter_ResourceTagsPropertyList
	// Experimental.
	ResourceTagsInput() interface{}
	// Experimental.
	ResourceType() AwsFilter_ResourceTypePropertyList
	// Experimental.
	ResourceTypeInput() interface{}
	// Experimental.
	Severity() AwsFilter_SeverityPropertyList
	// Experimental.
	SeverityInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Title() AwsFilter_TitlePropertyList
	// Experimental.
	TitleInput() interface{}
	// Experimental.
	UpdatedAt() AwsFilter_UpdatedAtPropertyList
	// Experimental.
	UpdatedAtInput() interface{}
	// Experimental.
	VendorSeverity() AwsFilter_VendorSeverityPropertyList
	// Experimental.
	VendorSeverityInput() interface{}
	// Experimental.
	VulnerabilityId() AwsFilter_VulnerabilityIdPropertyList
	// Experimental.
	VulnerabilityIdInput() interface{}
	// Experimental.
	VulnerabilitySource() AwsFilter_VulnerabilitySourcePropertyList
	// Experimental.
	VulnerabilitySourceInput() interface{}
	// Experimental.
	VulnerablePackages() AwsFilter_VulnerablePackagesPropertyList
	// Experimental.
	VulnerablePackagesInput() interface{}
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
	PutCodeRepositoryProjectName(value interface{})
	// Experimental.
	PutCodeRepositoryProviderType(value interface{})
	// Experimental.
	PutCodeVulnerabilityDetectorName(value interface{})
	// Experimental.
	PutCodeVulnerabilityDetectorTags(value interface{})
	// Experimental.
	PutCodeVulnerabilityFilePath(value interface{})
	// Experimental.
	PutComponentId(value interface{})
	// Experimental.
	PutComponentType(value interface{})
	// Experimental.
	PutEc2InstanceImageId(value interface{})
	// Experimental.
	PutEc2InstanceSubnetId(value interface{})
	// Experimental.
	PutEc2InstanceVpcId(value interface{})
	// Experimental.
	PutEcrImageArchitecture(value interface{})
	// Experimental.
	PutEcrImageHash(value interface{})
	// Experimental.
	PutEcrImageInUseCount(value interface{})
	// Experimental.
	PutEcrImageLastInUseAt(value interface{})
	// Experimental.
	PutEcrImagePushedAt(value interface{})
	// Experimental.
	PutEcrImageRegistry(value interface{})
	// Experimental.
	PutEcrImageRepositoryName(value interface{})
	// Experimental.
	PutEcrImageTags(value interface{})
	// Experimental.
	PutEpssScore(value interface{})
	// Experimental.
	PutExploitAvailable(value interface{})
	// Experimental.
	PutFindingArn(value interface{})
	// Experimental.
	PutFindingStatus(value interface{})
	// Experimental.
	PutFindingType(value interface{})
	// Experimental.
	PutFirstObservedAt(value interface{})
	// Experimental.
	PutFixAvailable(value interface{})
	// Experimental.
	PutInspectorScore(value interface{})
	// Experimental.
	PutLambdaFunctionExecutionRoleArn(value interface{})
	// Experimental.
	PutLambdaFunctionLastModifiedAt(value interface{})
	// Experimental.
	PutLambdaFunctionLayers(value interface{})
	// Experimental.
	PutLambdaFunctionName(value interface{})
	// Experimental.
	PutLambdaFunctionRuntime(value interface{})
	// Experimental.
	PutLastObservedAt(value interface{})
	// Experimental.
	PutNetworkProtocol(value interface{})
	// Experimental.
	PutPortRange(value interface{})
	// Experimental.
	PutRelatedVulnerabilities(value interface{})
	// Experimental.
	PutResourceId(value interface{})
	// Experimental.
	PutResourceTags(value interface{})
	// Experimental.
	PutResourceType(value interface{})
	// Experimental.
	PutSeverity(value interface{})
	// Experimental.
	PutTitle(value interface{})
	// Experimental.
	PutUpdatedAt(value interface{})
	// Experimental.
	PutVendorSeverity(value interface{})
	// Experimental.
	PutVulnerabilityId(value interface{})
	// Experimental.
	PutVulnerabilitySource(value interface{})
	// Experimental.
	PutVulnerablePackages(value interface{})
	// Experimental.
	ResetAwsAccountId()
	// Experimental.
	ResetCodeRepositoryProjectName()
	// Experimental.
	ResetCodeRepositoryProviderType()
	// Experimental.
	ResetCodeVulnerabilityDetectorName()
	// Experimental.
	ResetCodeVulnerabilityDetectorTags()
	// Experimental.
	ResetCodeVulnerabilityFilePath()
	// Experimental.
	ResetComponentId()
	// Experimental.
	ResetComponentType()
	// Experimental.
	ResetEc2InstanceImageId()
	// Experimental.
	ResetEc2InstanceSubnetId()
	// Experimental.
	ResetEc2InstanceVpcId()
	// Experimental.
	ResetEcrImageArchitecture()
	// Experimental.
	ResetEcrImageHash()
	// Experimental.
	ResetEcrImageInUseCount()
	// Experimental.
	ResetEcrImageLastInUseAt()
	// Experimental.
	ResetEcrImagePushedAt()
	// Experimental.
	ResetEcrImageRegistry()
	// Experimental.
	ResetEcrImageRepositoryName()
	// Experimental.
	ResetEcrImageTags()
	// Experimental.
	ResetEpssScore()
	// Experimental.
	ResetExploitAvailable()
	// Experimental.
	ResetFindingArn()
	// Experimental.
	ResetFindingStatus()
	// Experimental.
	ResetFindingType()
	// Experimental.
	ResetFirstObservedAt()
	// Experimental.
	ResetFixAvailable()
	// Experimental.
	ResetInspectorScore()
	// Experimental.
	ResetLambdaFunctionExecutionRoleArn()
	// Experimental.
	ResetLambdaFunctionLastModifiedAt()
	// Experimental.
	ResetLambdaFunctionLayers()
	// Experimental.
	ResetLambdaFunctionName()
	// Experimental.
	ResetLambdaFunctionRuntime()
	// Experimental.
	ResetLastObservedAt()
	// Experimental.
	ResetNetworkProtocol()
	// Experimental.
	ResetPortRange()
	// Experimental.
	ResetRelatedVulnerabilities()
	// Experimental.
	ResetResourceId()
	// Experimental.
	ResetResourceTags()
	// Experimental.
	ResetResourceType()
	// Experimental.
	ResetSeverity()
	// Experimental.
	ResetTitle()
	// Experimental.
	ResetUpdatedAt()
	// Experimental.
	ResetVendorSeverity()
	// Experimental.
	ResetVulnerabilityId()
	// Experimental.
	ResetVulnerabilitySource()
	// Experimental.
	ResetVulnerablePackages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFilter_FilterCriteriaPropertyOutputReference
type jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) AwsAccountId() AwsFilter_AwsAccountIdPropertyList {
	var returns AwsFilter_AwsAccountIdPropertyList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) CodeRepositoryProjectName() AwsFilter_CodeRepositoryProjectNamePropertyList {
	var returns AwsFilter_CodeRepositoryProjectNamePropertyList
	_jsii_.Get(
		j,
		"codeRepositoryProjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) CodeRepositoryProjectNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryProjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) CodeRepositoryProviderType() AwsFilter_CodeRepositoryProviderTypePropertyList {
	var returns AwsFilter_CodeRepositoryProviderTypePropertyList
	_jsii_.Get(
		j,
		"codeRepositoryProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) CodeRepositoryProviderTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorName() AwsFilter_CodeVulnerabilityDetectorNamePropertyList {
	var returns AwsFilter_CodeVulnerabilityDetectorNamePropertyList
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorTags() AwsFilter_CodeVulnerabilityDetectorTagsPropertyList {
	var returns AwsFilter_CodeVulnerabilityDetectorTagsPropertyList
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityFilePath() AwsFilter_CodeVulnerabilityFilePathPropertyList {
	var returns AwsFilter_CodeVulnerabilityFilePathPropertyList
	_jsii_.Get(
		j,
		"codeVulnerabilityFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityFilePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ComponentId() AwsFilter_ComponentIdPropertyList {
	var returns AwsFilter_ComponentIdPropertyList
	_jsii_.Get(
		j,
		"componentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ComponentIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ComponentType() AwsFilter_ComponentTypePropertyList {
	var returns AwsFilter_ComponentTypePropertyList
	_jsii_.Get(
		j,
		"componentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ComponentTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceImageId() AwsFilter_Ec2InstanceImageIdPropertyList {
	var returns AwsFilter_Ec2InstanceImageIdPropertyList
	_jsii_.Get(
		j,
		"ec2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceSubnetId() AwsFilter_Ec2InstanceSubnetIdPropertyList {
	var returns AwsFilter_Ec2InstanceSubnetIdPropertyList
	_jsii_.Get(
		j,
		"ec2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceSubnetIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceVpcId() AwsFilter_Ec2InstanceVpcIdPropertyList {
	var returns AwsFilter_Ec2InstanceVpcIdPropertyList
	_jsii_.Get(
		j,
		"ec2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceVpcIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageArchitecture() AwsFilter_EcrImageArchitecturePropertyList {
	var returns AwsFilter_EcrImageArchitecturePropertyList
	_jsii_.Get(
		j,
		"ecrImageArchitecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageArchitectureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageArchitectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageHash() AwsFilter_EcrImageHashPropertyList {
	var returns AwsFilter_EcrImageHashPropertyList
	_jsii_.Get(
		j,
		"ecrImageHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageHashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageInUseCount() AwsFilter_EcrImageInUseCountPropertyList {
	var returns AwsFilter_EcrImageInUseCountPropertyList
	_jsii_.Get(
		j,
		"ecrImageInUseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageInUseCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageInUseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageLastInUseAt() AwsFilter_EcrImageLastInUseAtPropertyList {
	var returns AwsFilter_EcrImageLastInUseAtPropertyList
	_jsii_.Get(
		j,
		"ecrImageLastInUseAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageLastInUseAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageLastInUseAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImagePushedAt() AwsFilter_EcrImagePushedAtPropertyList {
	var returns AwsFilter_EcrImagePushedAtPropertyList
	_jsii_.Get(
		j,
		"ecrImagePushedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImagePushedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImagePushedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageRegistry() AwsFilter_EcrImageRegistryPropertyList {
	var returns AwsFilter_EcrImageRegistryPropertyList
	_jsii_.Get(
		j,
		"ecrImageRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageRegistryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageRepositoryName() AwsFilter_EcrImageRepositoryNamePropertyList {
	var returns AwsFilter_EcrImageRepositoryNamePropertyList
	_jsii_.Get(
		j,
		"ecrImageRepositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageRepositoryNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageRepositoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageTags() AwsFilter_EcrImageTagsPropertyList {
	var returns AwsFilter_EcrImageTagsPropertyList
	_jsii_.Get(
		j,
		"ecrImageTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EcrImageTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EpssScore() AwsFilter_EpssScorePropertyList {
	var returns AwsFilter_EpssScorePropertyList
	_jsii_.Get(
		j,
		"epssScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) EpssScoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"epssScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ExploitAvailable() AwsFilter_ExploitAvailablePropertyList {
	var returns AwsFilter_ExploitAvailablePropertyList
	_jsii_.Get(
		j,
		"exploitAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ExploitAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exploitAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) FindingArn() AwsFilter_FindingArnPropertyList {
	var returns AwsFilter_FindingArnPropertyList
	_jsii_.Get(
		j,
		"findingArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) FindingArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) FindingStatus() AwsFilter_FindingStatusPropertyList {
	var returns AwsFilter_FindingStatusPropertyList
	_jsii_.Get(
		j,
		"findingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) FindingStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) FindingType() AwsFilter_FindingTypePropertyList {
	var returns AwsFilter_FindingTypePropertyList
	_jsii_.Get(
		j,
		"findingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) FindingTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) FirstObservedAt() AwsFilter_FirstObservedAtPropertyList {
	var returns AwsFilter_FirstObservedAtPropertyList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) FixAvailable() AwsFilter_FixAvailablePropertyList {
	var returns AwsFilter_FixAvailablePropertyList
	_jsii_.Get(
		j,
		"fixAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) FixAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) InspectorScore() AwsFilter_InspectorScorePropertyList {
	var returns AwsFilter_InspectorScorePropertyList
	_jsii_.Get(
		j,
		"inspectorScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) InspectorScoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectorScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionExecutionRoleArn() AwsFilter_LambdaFunctionExecutionRoleArnPropertyList {
	var returns AwsFilter_LambdaFunctionExecutionRoleArnPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionExecutionRoleArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionLastModifiedAt() AwsFilter_LambdaFunctionLastModifiedAtPropertyList {
	var returns AwsFilter_LambdaFunctionLastModifiedAtPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionLastModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionLastModifiedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionLastModifiedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionLayers() AwsFilter_LambdaFunctionLayersPropertyList {
	var returns AwsFilter_LambdaFunctionLayersPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionLayers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionLayersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionLayersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionName() AwsFilter_LambdaFunctionNamePropertyList {
	var returns AwsFilter_LambdaFunctionNamePropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionRuntime() AwsFilter_LambdaFunctionRuntimePropertyList {
	var returns AwsFilter_LambdaFunctionRuntimePropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LastObservedAt() AwsFilter_LastObservedAtPropertyList {
	var returns AwsFilter_LastObservedAtPropertyList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) NetworkProtocol() AwsFilter_NetworkProtocolPropertyList {
	var returns AwsFilter_NetworkProtocolPropertyList
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) NetworkProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PortRange() AwsFilter_PortRangePropertyList {
	var returns AwsFilter_PortRangePropertyList
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PortRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) RelatedVulnerabilities() AwsFilter_RelatedVulnerabilitiesPropertyList {
	var returns AwsFilter_RelatedVulnerabilitiesPropertyList
	_jsii_.Get(
		j,
		"relatedVulnerabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) RelatedVulnerabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedVulnerabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResourceId() AwsFilter_ResourceIdPropertyList {
	var returns AwsFilter_ResourceIdPropertyList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResourceTags() AwsFilter_ResourceTagsPropertyList {
	var returns AwsFilter_ResourceTagsPropertyList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResourceType() AwsFilter_ResourceTypePropertyList {
	var returns AwsFilter_ResourceTypePropertyList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) Severity() AwsFilter_SeverityPropertyList {
	var returns AwsFilter_SeverityPropertyList
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) SeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) Title() AwsFilter_TitlePropertyList {
	var returns AwsFilter_TitlePropertyList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) UpdatedAt() AwsFilter_UpdatedAtPropertyList {
	var returns AwsFilter_UpdatedAtPropertyList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) VendorSeverity() AwsFilter_VendorSeverityPropertyList {
	var returns AwsFilter_VendorSeverityPropertyList
	_jsii_.Get(
		j,
		"vendorSeverity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) VendorSeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vendorSeverityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) VulnerabilityId() AwsFilter_VulnerabilityIdPropertyList {
	var returns AwsFilter_VulnerabilityIdPropertyList
	_jsii_.Get(
		j,
		"vulnerabilityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) VulnerabilityIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) VulnerabilitySource() AwsFilter_VulnerabilitySourcePropertyList {
	var returns AwsFilter_VulnerabilitySourcePropertyList
	_jsii_.Get(
		j,
		"vulnerabilitySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) VulnerabilitySourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilitySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) VulnerablePackages() AwsFilter_VulnerablePackagesPropertyList {
	var returns AwsFilter_VulnerablePackagesPropertyList
	_jsii_.Get(
		j,
		"vulnerablePackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) VulnerablePackagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerablePackagesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFilter_FilterCriteriaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsFilter_FilterCriteriaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFilter_FilterCriteriaPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-inspector.AwsFilter.FilterCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFilter_FilterCriteriaPropertyOutputReference_Override(a AwsFilter_FilterCriteriaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-inspector.AwsFilter.FilterCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutAwsAccountId(value interface{}) {
	if err := a.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutCodeRepositoryProjectName(value interface{}) {
	if err := a.validatePutCodeRepositoryProjectNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeRepositoryProjectName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutCodeRepositoryProviderType(value interface{}) {
	if err := a.validatePutCodeRepositoryProviderTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeRepositoryProviderType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutCodeVulnerabilityDetectorName(value interface{}) {
	if err := a.validatePutCodeVulnerabilityDetectorNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeVulnerabilityDetectorName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutCodeVulnerabilityDetectorTags(value interface{}) {
	if err := a.validatePutCodeVulnerabilityDetectorTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeVulnerabilityDetectorTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutCodeVulnerabilityFilePath(value interface{}) {
	if err := a.validatePutCodeVulnerabilityFilePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeVulnerabilityFilePath",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutComponentId(value interface{}) {
	if err := a.validatePutComponentIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComponentId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutComponentType(value interface{}) {
	if err := a.validatePutComponentTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComponentType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEc2InstanceImageId(value interface{}) {
	if err := a.validatePutEc2InstanceImageIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2InstanceImageId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEc2InstanceSubnetId(value interface{}) {
	if err := a.validatePutEc2InstanceSubnetIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2InstanceSubnetId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEc2InstanceVpcId(value interface{}) {
	if err := a.validatePutEc2InstanceVpcIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2InstanceVpcId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEcrImageArchitecture(value interface{}) {
	if err := a.validatePutEcrImageArchitectureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageArchitecture",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEcrImageHash(value interface{}) {
	if err := a.validatePutEcrImageHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageHash",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEcrImageInUseCount(value interface{}) {
	if err := a.validatePutEcrImageInUseCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageInUseCount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEcrImageLastInUseAt(value interface{}) {
	if err := a.validatePutEcrImageLastInUseAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageLastInUseAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEcrImagePushedAt(value interface{}) {
	if err := a.validatePutEcrImagePushedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImagePushedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEcrImageRegistry(value interface{}) {
	if err := a.validatePutEcrImageRegistryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageRegistry",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEcrImageRepositoryName(value interface{}) {
	if err := a.validatePutEcrImageRepositoryNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageRepositoryName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEcrImageTags(value interface{}) {
	if err := a.validatePutEcrImageTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcrImageTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutEpssScore(value interface{}) {
	if err := a.validatePutEpssScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEpssScore",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutExploitAvailable(value interface{}) {
	if err := a.validatePutExploitAvailableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExploitAvailable",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutFindingArn(value interface{}) {
	if err := a.validatePutFindingArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutFindingStatus(value interface{}) {
	if err := a.validatePutFindingStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingStatus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutFindingType(value interface{}) {
	if err := a.validatePutFindingTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFindingType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutFirstObservedAt(value interface{}) {
	if err := a.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutFixAvailable(value interface{}) {
	if err := a.validatePutFixAvailableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFixAvailable",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutInspectorScore(value interface{}) {
	if err := a.validatePutInspectorScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInspectorScore",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionExecutionRoleArn(value interface{}) {
	if err := a.validatePutLambdaFunctionExecutionRoleArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionExecutionRoleArn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionLastModifiedAt(value interface{}) {
	if err := a.validatePutLambdaFunctionLastModifiedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionLastModifiedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionLayers(value interface{}) {
	if err := a.validatePutLambdaFunctionLayersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionLayers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionName(value interface{}) {
	if err := a.validatePutLambdaFunctionNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionRuntime(value interface{}) {
	if err := a.validatePutLambdaFunctionRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionRuntime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutLastObservedAt(value interface{}) {
	if err := a.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutNetworkProtocol(value interface{}) {
	if err := a.validatePutNetworkProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkProtocol",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutPortRange(value interface{}) {
	if err := a.validatePutPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPortRange",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutRelatedVulnerabilities(value interface{}) {
	if err := a.validatePutRelatedVulnerabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelatedVulnerabilities",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutResourceId(value interface{}) {
	if err := a.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutResourceTags(value interface{}) {
	if err := a.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutResourceType(value interface{}) {
	if err := a.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutSeverity(value interface{}) {
	if err := a.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSeverity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutTitle(value interface{}) {
	if err := a.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTitle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutUpdatedAt(value interface{}) {
	if err := a.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutVendorSeverity(value interface{}) {
	if err := a.validatePutVendorSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVendorSeverity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutVulnerabilityId(value interface{}) {
	if err := a.validatePutVulnerabilityIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVulnerabilityId",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutVulnerabilitySource(value interface{}) {
	if err := a.validatePutVulnerabilitySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVulnerabilitySource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) PutVulnerablePackages(value interface{}) {
	if err := a.validatePutVulnerablePackagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVulnerablePackages",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetCodeRepositoryProjectName() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeRepositoryProjectName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetCodeRepositoryProviderType() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeRepositoryProviderType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetCodeVulnerabilityDetectorName() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeVulnerabilityDetectorName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetCodeVulnerabilityDetectorTags() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeVulnerabilityDetectorTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetCodeVulnerabilityFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeVulnerabilityFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetComponentId() {
	_jsii_.InvokeVoid(
		a,
		"resetComponentId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetComponentType() {
	_jsii_.InvokeVoid(
		a,
		"resetComponentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEc2InstanceImageId() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2InstanceImageId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEc2InstanceSubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2InstanceSubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEc2InstanceVpcId() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2InstanceVpcId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageArchitecture() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageArchitecture",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageHash() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageHash",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageInUseCount() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageInUseCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageLastInUseAt() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageLastInUseAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEcrImagePushedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImagePushedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageRegistry() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageRegistry",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageRepositoryName() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageRepositoryName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageTags() {
	_jsii_.InvokeVoid(
		a,
		"resetEcrImageTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetEpssScore() {
	_jsii_.InvokeVoid(
		a,
		"resetEpssScore",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetExploitAvailable() {
	_jsii_.InvokeVoid(
		a,
		"resetExploitAvailable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetFindingArn() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetFindingStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetFindingType() {
	_jsii_.InvokeVoid(
		a,
		"resetFindingType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetFixAvailable() {
	_jsii_.InvokeVoid(
		a,
		"resetFixAvailable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetInspectorScore() {
	_jsii_.InvokeVoid(
		a,
		"resetInspectorScore",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionExecutionRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionExecutionRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionLastModifiedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionLastModifiedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionLayers() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionLayers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionName() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionRuntime() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionRuntime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetNetworkProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetPortRange() {
	_jsii_.InvokeVoid(
		a,
		"resetPortRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetRelatedVulnerabilities() {
	_jsii_.InvokeVoid(
		a,
		"resetRelatedVulnerabilities",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		a,
		"resetSeverity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		a,
		"resetTitle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetVendorSeverity() {
	_jsii_.InvokeVoid(
		a,
		"resetVendorSeverity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetVulnerabilityId() {
	_jsii_.InvokeVoid(
		a,
		"resetVulnerabilityId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetVulnerabilitySource() {
	_jsii_.InvokeVoid(
		a,
		"resetVulnerabilitySource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ResetVulnerablePackages() {
	_jsii_.InvokeVoid(
		a,
		"resetVulnerablePackages",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFilter_FilterCriteriaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

