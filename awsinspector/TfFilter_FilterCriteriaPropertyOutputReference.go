package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsinspector/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsinspector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFilter_FilterCriteriaPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsAccountId() TfFilter_AwsAccountIdPropertyList
	// Experimental.
	AwsAccountIdInput() interface{}
	// Experimental.
	CodeRepositoryProjectName() TfFilter_CodeRepositoryProjectNamePropertyList
	// Experimental.
	CodeRepositoryProjectNameInput() interface{}
	// Experimental.
	CodeRepositoryProviderType() TfFilter_CodeRepositoryProviderTypePropertyList
	// Experimental.
	CodeRepositoryProviderTypeInput() interface{}
	// Experimental.
	CodeVulnerabilityDetectorName() TfFilter_CodeVulnerabilityDetectorNamePropertyList
	// Experimental.
	CodeVulnerabilityDetectorNameInput() interface{}
	// Experimental.
	CodeVulnerabilityDetectorTags() TfFilter_CodeVulnerabilityDetectorTagsPropertyList
	// Experimental.
	CodeVulnerabilityDetectorTagsInput() interface{}
	// Experimental.
	CodeVulnerabilityFilePath() TfFilter_CodeVulnerabilityFilePathPropertyList
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
	ComponentId() TfFilter_ComponentIdPropertyList
	// Experimental.
	ComponentIdInput() interface{}
	// Experimental.
	ComponentType() TfFilter_ComponentTypePropertyList
	// Experimental.
	ComponentTypeInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Ec2InstanceImageId() TfFilter_Ec2InstanceImageIdPropertyList
	// Experimental.
	Ec2InstanceImageIdInput() interface{}
	// Experimental.
	Ec2InstanceSubnetId() TfFilter_Ec2InstanceSubnetIdPropertyList
	// Experimental.
	Ec2InstanceSubnetIdInput() interface{}
	// Experimental.
	Ec2InstanceVpcId() TfFilter_Ec2InstanceVpcIdPropertyList
	// Experimental.
	Ec2InstanceVpcIdInput() interface{}
	// Experimental.
	EcrImageArchitecture() TfFilter_EcrImageArchitecturePropertyList
	// Experimental.
	EcrImageArchitectureInput() interface{}
	// Experimental.
	EcrImageHash() TfFilter_EcrImageHashPropertyList
	// Experimental.
	EcrImageHashInput() interface{}
	// Experimental.
	EcrImageInUseCount() TfFilter_EcrImageInUseCountPropertyList
	// Experimental.
	EcrImageInUseCountInput() interface{}
	// Experimental.
	EcrImageLastInUseAt() TfFilter_EcrImageLastInUseAtPropertyList
	// Experimental.
	EcrImageLastInUseAtInput() interface{}
	// Experimental.
	EcrImagePushedAt() TfFilter_EcrImagePushedAtPropertyList
	// Experimental.
	EcrImagePushedAtInput() interface{}
	// Experimental.
	EcrImageRegistry() TfFilter_EcrImageRegistryPropertyList
	// Experimental.
	EcrImageRegistryInput() interface{}
	// Experimental.
	EcrImageRepositoryName() TfFilter_EcrImageRepositoryNamePropertyList
	// Experimental.
	EcrImageRepositoryNameInput() interface{}
	// Experimental.
	EcrImageTags() TfFilter_EcrImageTagsPropertyList
	// Experimental.
	EcrImageTagsInput() interface{}
	// Experimental.
	EpssScore() TfFilter_EpssScorePropertyList
	// Experimental.
	EpssScoreInput() interface{}
	// Experimental.
	ExploitAvailable() TfFilter_ExploitAvailablePropertyList
	// Experimental.
	ExploitAvailableInput() interface{}
	// Experimental.
	FindingArn() TfFilter_FindingArnPropertyList
	// Experimental.
	FindingArnInput() interface{}
	// Experimental.
	FindingStatus() TfFilter_FindingStatusPropertyList
	// Experimental.
	FindingStatusInput() interface{}
	// Experimental.
	FindingType() TfFilter_FindingTypePropertyList
	// Experimental.
	FindingTypeInput() interface{}
	// Experimental.
	FirstObservedAt() TfFilter_FirstObservedAtPropertyList
	// Experimental.
	FirstObservedAtInput() interface{}
	// Experimental.
	FixAvailable() TfFilter_FixAvailablePropertyList
	// Experimental.
	FixAvailableInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InspectorScore() TfFilter_InspectorScorePropertyList
	// Experimental.
	InspectorScoreInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LambdaFunctionExecutionRoleArn() TfFilter_LambdaFunctionExecutionRoleArnPropertyList
	// Experimental.
	LambdaFunctionExecutionRoleArnInput() interface{}
	// Experimental.
	LambdaFunctionLastModifiedAt() TfFilter_LambdaFunctionLastModifiedAtPropertyList
	// Experimental.
	LambdaFunctionLastModifiedAtInput() interface{}
	// Experimental.
	LambdaFunctionLayers() TfFilter_LambdaFunctionLayersPropertyList
	// Experimental.
	LambdaFunctionLayersInput() interface{}
	// Experimental.
	LambdaFunctionName() TfFilter_LambdaFunctionNamePropertyList
	// Experimental.
	LambdaFunctionNameInput() interface{}
	// Experimental.
	LambdaFunctionRuntime() TfFilter_LambdaFunctionRuntimePropertyList
	// Experimental.
	LambdaFunctionRuntimeInput() interface{}
	// Experimental.
	LastObservedAt() TfFilter_LastObservedAtPropertyList
	// Experimental.
	LastObservedAtInput() interface{}
	// Experimental.
	NetworkProtocol() TfFilter_NetworkProtocolPropertyList
	// Experimental.
	NetworkProtocolInput() interface{}
	// Experimental.
	PortRange() TfFilter_PortRangePropertyList
	// Experimental.
	PortRangeInput() interface{}
	// Experimental.
	RelatedVulnerabilities() TfFilter_RelatedVulnerabilitiesPropertyList
	// Experimental.
	RelatedVulnerabilitiesInput() interface{}
	// Experimental.
	ResourceId() TfFilter_ResourceIdPropertyList
	// Experimental.
	ResourceIdInput() interface{}
	// Experimental.
	ResourceTags() TfFilter_ResourceTagsPropertyList
	// Experimental.
	ResourceTagsInput() interface{}
	// Experimental.
	ResourceType() TfFilter_ResourceTypePropertyList
	// Experimental.
	ResourceTypeInput() interface{}
	// Experimental.
	Severity() TfFilter_SeverityPropertyList
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
	Title() TfFilter_TitlePropertyList
	// Experimental.
	TitleInput() interface{}
	// Experimental.
	UpdatedAt() TfFilter_UpdatedAtPropertyList
	// Experimental.
	UpdatedAtInput() interface{}
	// Experimental.
	VendorSeverity() TfFilter_VendorSeverityPropertyList
	// Experimental.
	VendorSeverityInput() interface{}
	// Experimental.
	VulnerabilityId() TfFilter_VulnerabilityIdPropertyList
	// Experimental.
	VulnerabilityIdInput() interface{}
	// Experimental.
	VulnerabilitySource() TfFilter_VulnerabilitySourcePropertyList
	// Experimental.
	VulnerabilitySourceInput() interface{}
	// Experimental.
	VulnerablePackages() TfFilter_VulnerablePackagesPropertyList
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

// The jsii proxy struct for TfFilter_FilterCriteriaPropertyOutputReference
type jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) AwsAccountId() TfFilter_AwsAccountIdPropertyList {
	var returns TfFilter_AwsAccountIdPropertyList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) CodeRepositoryProjectName() TfFilter_CodeRepositoryProjectNamePropertyList {
	var returns TfFilter_CodeRepositoryProjectNamePropertyList
	_jsii_.Get(
		j,
		"codeRepositoryProjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) CodeRepositoryProjectNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryProjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) CodeRepositoryProviderType() TfFilter_CodeRepositoryProviderTypePropertyList {
	var returns TfFilter_CodeRepositoryProviderTypePropertyList
	_jsii_.Get(
		j,
		"codeRepositoryProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) CodeRepositoryProviderTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorName() TfFilter_CodeVulnerabilityDetectorNamePropertyList {
	var returns TfFilter_CodeVulnerabilityDetectorNamePropertyList
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorTags() TfFilter_CodeVulnerabilityDetectorTagsPropertyList {
	var returns TfFilter_CodeVulnerabilityDetectorTagsPropertyList
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityDetectorTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityDetectorTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityFilePath() TfFilter_CodeVulnerabilityFilePathPropertyList {
	var returns TfFilter_CodeVulnerabilityFilePathPropertyList
	_jsii_.Get(
		j,
		"codeVulnerabilityFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) CodeVulnerabilityFilePathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeVulnerabilityFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ComponentId() TfFilter_ComponentIdPropertyList {
	var returns TfFilter_ComponentIdPropertyList
	_jsii_.Get(
		j,
		"componentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ComponentIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ComponentType() TfFilter_ComponentTypePropertyList {
	var returns TfFilter_ComponentTypePropertyList
	_jsii_.Get(
		j,
		"componentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ComponentTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceImageId() TfFilter_Ec2InstanceImageIdPropertyList {
	var returns TfFilter_Ec2InstanceImageIdPropertyList
	_jsii_.Get(
		j,
		"ec2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceSubnetId() TfFilter_Ec2InstanceSubnetIdPropertyList {
	var returns TfFilter_Ec2InstanceSubnetIdPropertyList
	_jsii_.Get(
		j,
		"ec2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceSubnetIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceVpcId() TfFilter_Ec2InstanceVpcIdPropertyList {
	var returns TfFilter_Ec2InstanceVpcIdPropertyList
	_jsii_.Get(
		j,
		"ec2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) Ec2InstanceVpcIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageArchitecture() TfFilter_EcrImageArchitecturePropertyList {
	var returns TfFilter_EcrImageArchitecturePropertyList
	_jsii_.Get(
		j,
		"ecrImageArchitecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageArchitectureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageArchitectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageHash() TfFilter_EcrImageHashPropertyList {
	var returns TfFilter_EcrImageHashPropertyList
	_jsii_.Get(
		j,
		"ecrImageHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageHashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageInUseCount() TfFilter_EcrImageInUseCountPropertyList {
	var returns TfFilter_EcrImageInUseCountPropertyList
	_jsii_.Get(
		j,
		"ecrImageInUseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageInUseCountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageInUseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageLastInUseAt() TfFilter_EcrImageLastInUseAtPropertyList {
	var returns TfFilter_EcrImageLastInUseAtPropertyList
	_jsii_.Get(
		j,
		"ecrImageLastInUseAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageLastInUseAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageLastInUseAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImagePushedAt() TfFilter_EcrImagePushedAtPropertyList {
	var returns TfFilter_EcrImagePushedAtPropertyList
	_jsii_.Get(
		j,
		"ecrImagePushedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImagePushedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImagePushedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageRegistry() TfFilter_EcrImageRegistryPropertyList {
	var returns TfFilter_EcrImageRegistryPropertyList
	_jsii_.Get(
		j,
		"ecrImageRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageRegistryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageRepositoryName() TfFilter_EcrImageRepositoryNamePropertyList {
	var returns TfFilter_EcrImageRepositoryNamePropertyList
	_jsii_.Get(
		j,
		"ecrImageRepositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageRepositoryNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageRepositoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageTags() TfFilter_EcrImageTagsPropertyList {
	var returns TfFilter_EcrImageTagsPropertyList
	_jsii_.Get(
		j,
		"ecrImageTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EcrImageTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EpssScore() TfFilter_EpssScorePropertyList {
	var returns TfFilter_EpssScorePropertyList
	_jsii_.Get(
		j,
		"epssScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) EpssScoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"epssScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ExploitAvailable() TfFilter_ExploitAvailablePropertyList {
	var returns TfFilter_ExploitAvailablePropertyList
	_jsii_.Get(
		j,
		"exploitAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ExploitAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exploitAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) FindingArn() TfFilter_FindingArnPropertyList {
	var returns TfFilter_FindingArnPropertyList
	_jsii_.Get(
		j,
		"findingArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) FindingArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) FindingStatus() TfFilter_FindingStatusPropertyList {
	var returns TfFilter_FindingStatusPropertyList
	_jsii_.Get(
		j,
		"findingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) FindingStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) FindingType() TfFilter_FindingTypePropertyList {
	var returns TfFilter_FindingTypePropertyList
	_jsii_.Get(
		j,
		"findingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) FindingTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) FirstObservedAt() TfFilter_FirstObservedAtPropertyList {
	var returns TfFilter_FirstObservedAtPropertyList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) FixAvailable() TfFilter_FixAvailablePropertyList {
	var returns TfFilter_FixAvailablePropertyList
	_jsii_.Get(
		j,
		"fixAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) FixAvailableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) InspectorScore() TfFilter_InspectorScorePropertyList {
	var returns TfFilter_InspectorScorePropertyList
	_jsii_.Get(
		j,
		"inspectorScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) InspectorScoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectorScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionExecutionRoleArn() TfFilter_LambdaFunctionExecutionRoleArnPropertyList {
	var returns TfFilter_LambdaFunctionExecutionRoleArnPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionExecutionRoleArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionLastModifiedAt() TfFilter_LambdaFunctionLastModifiedAtPropertyList {
	var returns TfFilter_LambdaFunctionLastModifiedAtPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionLastModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionLastModifiedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionLastModifiedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionLayers() TfFilter_LambdaFunctionLayersPropertyList {
	var returns TfFilter_LambdaFunctionLayersPropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionLayers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionLayersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionLayersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionName() TfFilter_LambdaFunctionNamePropertyList {
	var returns TfFilter_LambdaFunctionNamePropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionRuntime() TfFilter_LambdaFunctionRuntimePropertyList {
	var returns TfFilter_LambdaFunctionRuntimePropertyList
	_jsii_.Get(
		j,
		"lambdaFunctionRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LambdaFunctionRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LastObservedAt() TfFilter_LastObservedAtPropertyList {
	var returns TfFilter_LastObservedAtPropertyList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) NetworkProtocol() TfFilter_NetworkProtocolPropertyList {
	var returns TfFilter_NetworkProtocolPropertyList
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) NetworkProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PortRange() TfFilter_PortRangePropertyList {
	var returns TfFilter_PortRangePropertyList
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PortRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) RelatedVulnerabilities() TfFilter_RelatedVulnerabilitiesPropertyList {
	var returns TfFilter_RelatedVulnerabilitiesPropertyList
	_jsii_.Get(
		j,
		"relatedVulnerabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) RelatedVulnerabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedVulnerabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResourceId() TfFilter_ResourceIdPropertyList {
	var returns TfFilter_ResourceIdPropertyList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResourceTags() TfFilter_ResourceTagsPropertyList {
	var returns TfFilter_ResourceTagsPropertyList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResourceType() TfFilter_ResourceTypePropertyList {
	var returns TfFilter_ResourceTypePropertyList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) Severity() TfFilter_SeverityPropertyList {
	var returns TfFilter_SeverityPropertyList
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) SeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) Title() TfFilter_TitlePropertyList {
	var returns TfFilter_TitlePropertyList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) UpdatedAt() TfFilter_UpdatedAtPropertyList {
	var returns TfFilter_UpdatedAtPropertyList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) VendorSeverity() TfFilter_VendorSeverityPropertyList {
	var returns TfFilter_VendorSeverityPropertyList
	_jsii_.Get(
		j,
		"vendorSeverity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) VendorSeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vendorSeverityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) VulnerabilityId() TfFilter_VulnerabilityIdPropertyList {
	var returns TfFilter_VulnerabilityIdPropertyList
	_jsii_.Get(
		j,
		"vulnerabilityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) VulnerabilityIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) VulnerabilitySource() TfFilter_VulnerabilitySourcePropertyList {
	var returns TfFilter_VulnerabilitySourcePropertyList
	_jsii_.Get(
		j,
		"vulnerabilitySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) VulnerabilitySourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilitySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) VulnerablePackages() TfFilter_VulnerablePackagesPropertyList {
	var returns TfFilter_VulnerablePackagesPropertyList
	_jsii_.Get(
		j,
		"vulnerablePackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) VulnerablePackagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerablePackagesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFilter_FilterCriteriaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfFilter_FilterCriteriaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFilter_FilterCriteriaPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-inspector.TfFilter.FilterCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFilter_FilterCriteriaPropertyOutputReference_Override(t TfFilter_FilterCriteriaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-inspector.TfFilter.FilterCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutAwsAccountId(value interface{}) {
	if err := t.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutCodeRepositoryProjectName(value interface{}) {
	if err := t.validatePutCodeRepositoryProjectNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeRepositoryProjectName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutCodeRepositoryProviderType(value interface{}) {
	if err := t.validatePutCodeRepositoryProviderTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeRepositoryProviderType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutCodeVulnerabilityDetectorName(value interface{}) {
	if err := t.validatePutCodeVulnerabilityDetectorNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeVulnerabilityDetectorName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutCodeVulnerabilityDetectorTags(value interface{}) {
	if err := t.validatePutCodeVulnerabilityDetectorTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeVulnerabilityDetectorTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutCodeVulnerabilityFilePath(value interface{}) {
	if err := t.validatePutCodeVulnerabilityFilePathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeVulnerabilityFilePath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutComponentId(value interface{}) {
	if err := t.validatePutComponentIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putComponentId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutComponentType(value interface{}) {
	if err := t.validatePutComponentTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putComponentType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEc2InstanceImageId(value interface{}) {
	if err := t.validatePutEc2InstanceImageIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEc2InstanceImageId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEc2InstanceSubnetId(value interface{}) {
	if err := t.validatePutEc2InstanceSubnetIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEc2InstanceSubnetId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEc2InstanceVpcId(value interface{}) {
	if err := t.validatePutEc2InstanceVpcIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEc2InstanceVpcId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEcrImageArchitecture(value interface{}) {
	if err := t.validatePutEcrImageArchitectureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcrImageArchitecture",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEcrImageHash(value interface{}) {
	if err := t.validatePutEcrImageHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcrImageHash",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEcrImageInUseCount(value interface{}) {
	if err := t.validatePutEcrImageInUseCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcrImageInUseCount",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEcrImageLastInUseAt(value interface{}) {
	if err := t.validatePutEcrImageLastInUseAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcrImageLastInUseAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEcrImagePushedAt(value interface{}) {
	if err := t.validatePutEcrImagePushedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcrImagePushedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEcrImageRegistry(value interface{}) {
	if err := t.validatePutEcrImageRegistryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcrImageRegistry",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEcrImageRepositoryName(value interface{}) {
	if err := t.validatePutEcrImageRepositoryNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcrImageRepositoryName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEcrImageTags(value interface{}) {
	if err := t.validatePutEcrImageTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcrImageTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutEpssScore(value interface{}) {
	if err := t.validatePutEpssScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEpssScore",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutExploitAvailable(value interface{}) {
	if err := t.validatePutExploitAvailableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExploitAvailable",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutFindingArn(value interface{}) {
	if err := t.validatePutFindingArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFindingArn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutFindingStatus(value interface{}) {
	if err := t.validatePutFindingStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFindingStatus",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutFindingType(value interface{}) {
	if err := t.validatePutFindingTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFindingType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutFirstObservedAt(value interface{}) {
	if err := t.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutFixAvailable(value interface{}) {
	if err := t.validatePutFixAvailableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFixAvailable",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutInspectorScore(value interface{}) {
	if err := t.validatePutInspectorScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInspectorScore",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionExecutionRoleArn(value interface{}) {
	if err := t.validatePutLambdaFunctionExecutionRoleArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaFunctionExecutionRoleArn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionLastModifiedAt(value interface{}) {
	if err := t.validatePutLambdaFunctionLastModifiedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaFunctionLastModifiedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionLayers(value interface{}) {
	if err := t.validatePutLambdaFunctionLayersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaFunctionLayers",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionName(value interface{}) {
	if err := t.validatePutLambdaFunctionNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaFunctionName",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutLambdaFunctionRuntime(value interface{}) {
	if err := t.validatePutLambdaFunctionRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaFunctionRuntime",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutLastObservedAt(value interface{}) {
	if err := t.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutNetworkProtocol(value interface{}) {
	if err := t.validatePutNetworkProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkProtocol",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutPortRange(value interface{}) {
	if err := t.validatePutPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPortRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutRelatedVulnerabilities(value interface{}) {
	if err := t.validatePutRelatedVulnerabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRelatedVulnerabilities",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutResourceId(value interface{}) {
	if err := t.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutResourceTags(value interface{}) {
	if err := t.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutResourceType(value interface{}) {
	if err := t.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceType",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutSeverity(value interface{}) {
	if err := t.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSeverity",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutTitle(value interface{}) {
	if err := t.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTitle",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutUpdatedAt(value interface{}) {
	if err := t.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutVendorSeverity(value interface{}) {
	if err := t.validatePutVendorSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVendorSeverity",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutVulnerabilityId(value interface{}) {
	if err := t.validatePutVulnerabilityIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVulnerabilityId",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutVulnerabilitySource(value interface{}) {
	if err := t.validatePutVulnerabilitySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVulnerabilitySource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) PutVulnerablePackages(value interface{}) {
	if err := t.validatePutVulnerablePackagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVulnerablePackages",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetCodeRepositoryProjectName() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeRepositoryProjectName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetCodeRepositoryProviderType() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeRepositoryProviderType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetCodeVulnerabilityDetectorName() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeVulnerabilityDetectorName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetCodeVulnerabilityDetectorTags() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeVulnerabilityDetectorTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetCodeVulnerabilityFilePath() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeVulnerabilityFilePath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetComponentId() {
	_jsii_.InvokeVoid(
		t,
		"resetComponentId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetComponentType() {
	_jsii_.InvokeVoid(
		t,
		"resetComponentType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEc2InstanceImageId() {
	_jsii_.InvokeVoid(
		t,
		"resetEc2InstanceImageId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEc2InstanceSubnetId() {
	_jsii_.InvokeVoid(
		t,
		"resetEc2InstanceSubnetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEc2InstanceVpcId() {
	_jsii_.InvokeVoid(
		t,
		"resetEc2InstanceVpcId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageArchitecture() {
	_jsii_.InvokeVoid(
		t,
		"resetEcrImageArchitecture",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageHash() {
	_jsii_.InvokeVoid(
		t,
		"resetEcrImageHash",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageInUseCount() {
	_jsii_.InvokeVoid(
		t,
		"resetEcrImageInUseCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageLastInUseAt() {
	_jsii_.InvokeVoid(
		t,
		"resetEcrImageLastInUseAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEcrImagePushedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetEcrImagePushedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageRegistry() {
	_jsii_.InvokeVoid(
		t,
		"resetEcrImageRegistry",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageRepositoryName() {
	_jsii_.InvokeVoid(
		t,
		"resetEcrImageRepositoryName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEcrImageTags() {
	_jsii_.InvokeVoid(
		t,
		"resetEcrImageTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetEpssScore() {
	_jsii_.InvokeVoid(
		t,
		"resetEpssScore",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetExploitAvailable() {
	_jsii_.InvokeVoid(
		t,
		"resetExploitAvailable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetFindingArn() {
	_jsii_.InvokeVoid(
		t,
		"resetFindingArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetFindingStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetFindingStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetFindingType() {
	_jsii_.InvokeVoid(
		t,
		"resetFindingType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetFixAvailable() {
	_jsii_.InvokeVoid(
		t,
		"resetFixAvailable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetInspectorScore() {
	_jsii_.InvokeVoid(
		t,
		"resetInspectorScore",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionExecutionRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaFunctionExecutionRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionLastModifiedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaFunctionLastModifiedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionLayers() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaFunctionLayers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionName() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaFunctionName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetLambdaFunctionRuntime() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaFunctionRuntime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetNetworkProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetPortRange() {
	_jsii_.InvokeVoid(
		t,
		"resetPortRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetRelatedVulnerabilities() {
	_jsii_.InvokeVoid(
		t,
		"resetRelatedVulnerabilities",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		t,
		"resetSeverity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		t,
		"resetTitle",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		t,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetVendorSeverity() {
	_jsii_.InvokeVoid(
		t,
		"resetVendorSeverity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetVulnerabilityId() {
	_jsii_.InvokeVoid(
		t,
		"resetVulnerabilityId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetVulnerabilitySource() {
	_jsii_.InvokeVoid(
		t,
		"resetVulnerabilitySource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ResetVulnerablePackages() {
	_jsii_.InvokeVoid(
		t,
		"resetVulnerablePackages",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFilter_FilterCriteriaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

