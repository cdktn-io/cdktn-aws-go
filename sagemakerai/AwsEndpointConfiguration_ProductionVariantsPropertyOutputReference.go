package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AcceleratorType() *string
	// Experimental.
	SetAcceleratorType(val *string)
	// Experimental.
	AcceleratorTypeInput() *string
	// Experimental.
	CapacityReservationConfig() AwsEndpointConfiguration_ProductionVariantsCapacityReservationConfigPropertyOutputReference
	// Experimental.
	CapacityReservationConfigInput() *AwsEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty
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
	ContainerStartupHealthCheckTimeoutInSeconds() *float64
	// Experimental.
	SetContainerStartupHealthCheckTimeoutInSeconds(val *float64)
	// Experimental.
	ContainerStartupHealthCheckTimeoutInSecondsInput() *float64
	// Experimental.
	CoreDumpConfig() AwsEndpointConfiguration_ProductionVariantsCoreDumpConfigPropertyOutputReference
	// Experimental.
	CoreDumpConfigInput() *AwsEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EnableSsmAccess() interface{}
	// Experimental.
	SetEnableSsmAccess(val interface{})
	// Experimental.
	EnableSsmAccessInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InferenceAmiVersion() *string
	// Experimental.
	SetInferenceAmiVersion(val *string)
	// Experimental.
	InferenceAmiVersionInput() *string
	// Experimental.
	InitialInstanceCount() *float64
	// Experimental.
	SetInitialInstanceCount(val *float64)
	// Experimental.
	InitialInstanceCountInput() *float64
	// Experimental.
	InitialVariantWeight() *float64
	// Experimental.
	SetInitialVariantWeight(val *float64)
	// Experimental.
	InitialVariantWeightInput() *float64
	// Experimental.
	InstanceType() *string
	// Experimental.
	SetInstanceType(val *string)
	// Experimental.
	InstanceTypeInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ManagedInstanceScaling() AwsEndpointConfiguration_ProductionVariantsManagedInstanceScalingPropertyOutputReference
	// Experimental.
	ManagedInstanceScalingInput() *AwsEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty
	// Experimental.
	ModelDataDownloadTimeoutInSeconds() *float64
	// Experimental.
	SetModelDataDownloadTimeoutInSeconds(val *float64)
	// Experimental.
	ModelDataDownloadTimeoutInSecondsInput() *float64
	// Experimental.
	ModelName() *string
	// Experimental.
	SetModelName(val *string)
	// Experimental.
	ModelNameInput() *string
	// Experimental.
	RoutingConfig() AwsEndpointConfiguration_ProductionVariantsRoutingConfigPropertyList
	// Experimental.
	RoutingConfigInput() interface{}
	// Experimental.
	ServerlessConfig() AwsEndpointConfiguration_ProductionVariantsServerlessConfigPropertyOutputReference
	// Experimental.
	ServerlessConfigInput() *AwsEndpointConfiguration_ProductionVariantsServerlessConfigProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VariantName() *string
	// Experimental.
	SetVariantName(val *string)
	// Experimental.
	VariantNameInput() *string
	// Experimental.
	VolumeSizeInGb() *float64
	// Experimental.
	SetVolumeSizeInGb(val *float64)
	// Experimental.
	VolumeSizeInGbInput() *float64
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
	PutCapacityReservationConfig(value *AwsEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty)
	// Experimental.
	PutCoreDumpConfig(value *AwsEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty)
	// Experimental.
	PutManagedInstanceScaling(value *AwsEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty)
	// Experimental.
	PutRoutingConfig(value interface{})
	// Experimental.
	PutServerlessConfig(value *AwsEndpointConfiguration_ProductionVariantsServerlessConfigProperty)
	// Experimental.
	ResetAcceleratorType()
	// Experimental.
	ResetCapacityReservationConfig()
	// Experimental.
	ResetContainerStartupHealthCheckTimeoutInSeconds()
	// Experimental.
	ResetCoreDumpConfig()
	// Experimental.
	ResetEnableSsmAccess()
	// Experimental.
	ResetInferenceAmiVersion()
	// Experimental.
	ResetInitialInstanceCount()
	// Experimental.
	ResetInitialVariantWeight()
	// Experimental.
	ResetInstanceType()
	// Experimental.
	ResetManagedInstanceScaling()
	// Experimental.
	ResetModelDataDownloadTimeoutInSeconds()
	// Experimental.
	ResetModelName()
	// Experimental.
	ResetRoutingConfig()
	// Experimental.
	ResetServerlessConfig()
	// Experimental.
	ResetVariantName()
	// Experimental.
	ResetVolumeSizeInGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference
type jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) AcceleratorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) AcceleratorTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) CapacityReservationConfig() AwsEndpointConfiguration_ProductionVariantsCapacityReservationConfigPropertyOutputReference {
	var returns AwsEndpointConfiguration_ProductionVariantsCapacityReservationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) CapacityReservationConfigInput() *AwsEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty {
	var returns *AwsEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty
	_jsii_.Get(
		j,
		"capacityReservationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ContainerStartupHealthCheckTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ContainerStartupHealthCheckTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) CoreDumpConfig() AwsEndpointConfiguration_ProductionVariantsCoreDumpConfigPropertyOutputReference {
	var returns AwsEndpointConfiguration_ProductionVariantsCoreDumpConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"coreDumpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) CoreDumpConfigInput() *AwsEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty {
	var returns *AwsEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty
	_jsii_.Get(
		j,
		"coreDumpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) EnableSsmAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsmAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) EnableSsmAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsmAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) InferenceAmiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAmiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) InferenceAmiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAmiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) InitialInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) InitialInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) InitialVariantWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialVariantWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) InitialVariantWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialVariantWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ManagedInstanceScaling() AwsEndpointConfiguration_ProductionVariantsManagedInstanceScalingPropertyOutputReference {
	var returns AwsEndpointConfiguration_ProductionVariantsManagedInstanceScalingPropertyOutputReference
	_jsii_.Get(
		j,
		"managedInstanceScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ManagedInstanceScalingInput() *AwsEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty {
	var returns *AwsEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty
	_jsii_.Get(
		j,
		"managedInstanceScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ModelDataDownloadTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ModelDataDownloadTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) RoutingConfig() AwsEndpointConfiguration_ProductionVariantsRoutingConfigPropertyList {
	var returns AwsEndpointConfiguration_ProductionVariantsRoutingConfigPropertyList
	_jsii_.Get(
		j,
		"routingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) RoutingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ServerlessConfig() AwsEndpointConfiguration_ProductionVariantsServerlessConfigPropertyOutputReference {
	var returns AwsEndpointConfiguration_ProductionVariantsServerlessConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"serverlessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ServerlessConfigInput() *AwsEndpointConfiguration_ProductionVariantsServerlessConfigProperty {
	var returns *AwsEndpointConfiguration_ProductionVariantsServerlessConfigProperty
	_jsii_.Get(
		j,
		"serverlessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) VariantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) VariantNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) VolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) VolumeSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGbInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEndpointConfiguration_ProductionVariantsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEndpointConfiguration_ProductionVariantsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsEndpointConfiguration.ProductionVariantsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEndpointConfiguration_ProductionVariantsPropertyOutputReference_Override(a AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsEndpointConfiguration.ProductionVariantsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetAcceleratorType(val *string) {
	if err := j.validateSetAcceleratorTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorType",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetContainerStartupHealthCheckTimeoutInSeconds(val *float64) {
	if err := j.validateSetContainerStartupHealthCheckTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetEnableSsmAccess(val interface{}) {
	if err := j.validateSetEnableSsmAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSsmAccess",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetInferenceAmiVersion(val *string) {
	if err := j.validateSetInferenceAmiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceAmiVersion",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetInitialInstanceCount(val *float64) {
	if err := j.validateSetInitialInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialInstanceCount",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetInitialVariantWeight(val *float64) {
	if err := j.validateSetInitialVariantWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialVariantWeight",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetModelDataDownloadTimeoutInSeconds(val *float64) {
	if err := j.validateSetModelDataDownloadTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataDownloadTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetVariantName(val *string) {
	if err := j.validateSetVariantNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variantName",
		val,
	)
}

func (j *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetVolumeSizeInGb(val *float64) {
	if err := j.validateSetVolumeSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSizeInGb",
		val,
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) PutCapacityReservationConfig(value *AwsEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty) {
	if err := a.validatePutCapacityReservationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCapacityReservationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) PutCoreDumpConfig(value *AwsEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty) {
	if err := a.validatePutCoreDumpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCoreDumpConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) PutManagedInstanceScaling(value *AwsEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty) {
	if err := a.validatePutManagedInstanceScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedInstanceScaling",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) PutRoutingConfig(value interface{}) {
	if err := a.validatePutRoutingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRoutingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) PutServerlessConfig(value *AwsEndpointConfiguration_ProductionVariantsServerlessConfigProperty) {
	if err := a.validatePutServerlessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerlessConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetAcceleratorType() {
	_jsii_.InvokeVoid(
		a,
		"resetAcceleratorType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetCapacityReservationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityReservationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetContainerStartupHealthCheckTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerStartupHealthCheckTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetCoreDumpConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCoreDumpConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetEnableSsmAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableSsmAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetInferenceAmiVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetInferenceAmiVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetInitialInstanceCount() {
	_jsii_.InvokeVoid(
		a,
		"resetInitialInstanceCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetInitialVariantWeight() {
	_jsii_.InvokeVoid(
		a,
		"resetInitialVariantWeight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetManagedInstanceScaling() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedInstanceScaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetModelDataDownloadTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetModelDataDownloadTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetModelName() {
	_jsii_.InvokeVoid(
		a,
		"resetModelName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetRoutingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetServerlessConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetServerlessConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetVariantName() {
	_jsii_.InvokeVoid(
		a,
		"resetVariantName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetVolumeSizeInGb() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumeSizeInGb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEndpointConfiguration_ProductionVariantsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

