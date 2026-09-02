package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AcceleratorType() *string
	// Experimental.
	SetAcceleratorType(val *string)
	// Experimental.
	AcceleratorTypeInput() *string
	// Experimental.
	CapacityReservationConfig() TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference
	// Experimental.
	CapacityReservationConfigInput() *TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty
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
	CoreDumpConfig() TfEndpointConfiguration_ShadowProductionVariantsCoreDumpConfigPropertyOutputReference
	// Experimental.
	CoreDumpConfigInput() *TfEndpointConfiguration_ShadowProductionVariantsCoreDumpConfigProperty
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
	ManagedInstanceScaling() TfEndpointConfiguration_ShadowProductionVariantsManagedInstanceScalingPropertyOutputReference
	// Experimental.
	ManagedInstanceScalingInput() *TfEndpointConfiguration_ShadowProductionVariantsManagedInstanceScalingProperty
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
	RoutingConfig() TfEndpointConfiguration_ShadowProductionVariantsRoutingConfigPropertyList
	// Experimental.
	RoutingConfigInput() interface{}
	// Experimental.
	ServerlessConfig() TfEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference
	// Experimental.
	ServerlessConfigInput() *TfEndpointConfiguration_ShadowProductionVariantsServerlessConfigProperty
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
	PutCapacityReservationConfig(value *TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty)
	// Experimental.
	PutCoreDumpConfig(value *TfEndpointConfiguration_ShadowProductionVariantsCoreDumpConfigProperty)
	// Experimental.
	PutManagedInstanceScaling(value *TfEndpointConfiguration_ShadowProductionVariantsManagedInstanceScalingProperty)
	// Experimental.
	PutRoutingConfig(value interface{})
	// Experimental.
	PutServerlessConfig(value *TfEndpointConfiguration_ShadowProductionVariantsServerlessConfigProperty)
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

// The jsii proxy struct for TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference
type jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) AcceleratorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) AcceleratorTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) CapacityReservationConfig() TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference {
	var returns TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) CapacityReservationConfigInput() *TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty {
	var returns *TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty
	_jsii_.Get(
		j,
		"capacityReservationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ContainerStartupHealthCheckTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ContainerStartupHealthCheckTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) CoreDumpConfig() TfEndpointConfiguration_ShadowProductionVariantsCoreDumpConfigPropertyOutputReference {
	var returns TfEndpointConfiguration_ShadowProductionVariantsCoreDumpConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"coreDumpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) CoreDumpConfigInput() *TfEndpointConfiguration_ShadowProductionVariantsCoreDumpConfigProperty {
	var returns *TfEndpointConfiguration_ShadowProductionVariantsCoreDumpConfigProperty
	_jsii_.Get(
		j,
		"coreDumpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) EnableSsmAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsmAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) EnableSsmAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsmAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) InferenceAmiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAmiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) InferenceAmiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAmiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) InitialInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) InitialInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) InitialVariantWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialVariantWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) InitialVariantWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialVariantWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ManagedInstanceScaling() TfEndpointConfiguration_ShadowProductionVariantsManagedInstanceScalingPropertyOutputReference {
	var returns TfEndpointConfiguration_ShadowProductionVariantsManagedInstanceScalingPropertyOutputReference
	_jsii_.Get(
		j,
		"managedInstanceScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ManagedInstanceScalingInput() *TfEndpointConfiguration_ShadowProductionVariantsManagedInstanceScalingProperty {
	var returns *TfEndpointConfiguration_ShadowProductionVariantsManagedInstanceScalingProperty
	_jsii_.Get(
		j,
		"managedInstanceScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ModelDataDownloadTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ModelDataDownloadTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) RoutingConfig() TfEndpointConfiguration_ShadowProductionVariantsRoutingConfigPropertyList {
	var returns TfEndpointConfiguration_ShadowProductionVariantsRoutingConfigPropertyList
	_jsii_.Get(
		j,
		"routingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) RoutingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ServerlessConfig() TfEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference {
	var returns TfEndpointConfiguration_ShadowProductionVariantsServerlessConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"serverlessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ServerlessConfigInput() *TfEndpointConfiguration_ShadowProductionVariantsServerlessConfigProperty {
	var returns *TfEndpointConfiguration_ShadowProductionVariantsServerlessConfigProperty
	_jsii_.Get(
		j,
		"serverlessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) VariantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) VariantNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) VolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) VolumeSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGbInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpointConfiguration.ShadowProductionVariantsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference_Override(t TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpointConfiguration.ShadowProductionVariantsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetAcceleratorType(val *string) {
	if err := j.validateSetAcceleratorTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorType",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetContainerStartupHealthCheckTimeoutInSeconds(val *float64) {
	if err := j.validateSetContainerStartupHealthCheckTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetEnableSsmAccess(val interface{}) {
	if err := j.validateSetEnableSsmAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSsmAccess",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetInferenceAmiVersion(val *string) {
	if err := j.validateSetInferenceAmiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceAmiVersion",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetInitialInstanceCount(val *float64) {
	if err := j.validateSetInitialInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialInstanceCount",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetInitialVariantWeight(val *float64) {
	if err := j.validateSetInitialVariantWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialVariantWeight",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetModelDataDownloadTimeoutInSeconds(val *float64) {
	if err := j.validateSetModelDataDownloadTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataDownloadTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetVariantName(val *string) {
	if err := j.validateSetVariantNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variantName",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference)SetVolumeSizeInGb(val *float64) {
	if err := j.validateSetVolumeSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSizeInGb",
		val,
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) PutCapacityReservationConfig(value *TfEndpointConfiguration_ShadowProductionVariantsCapacityReservationConfigProperty) {
	if err := t.validatePutCapacityReservationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityReservationConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) PutCoreDumpConfig(value *TfEndpointConfiguration_ShadowProductionVariantsCoreDumpConfigProperty) {
	if err := t.validatePutCoreDumpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCoreDumpConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) PutManagedInstanceScaling(value *TfEndpointConfiguration_ShadowProductionVariantsManagedInstanceScalingProperty) {
	if err := t.validatePutManagedInstanceScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedInstanceScaling",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) PutRoutingConfig(value interface{}) {
	if err := t.validatePutRoutingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRoutingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) PutServerlessConfig(value *TfEndpointConfiguration_ShadowProductionVariantsServerlessConfigProperty) {
	if err := t.validatePutServerlessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServerlessConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetAcceleratorType() {
	_jsii_.InvokeVoid(
		t,
		"resetAcceleratorType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetCapacityReservationConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityReservationConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetContainerStartupHealthCheckTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerStartupHealthCheckTimeoutInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetCoreDumpConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCoreDumpConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetEnableSsmAccess() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableSsmAccess",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetInferenceAmiVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetInferenceAmiVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetInitialInstanceCount() {
	_jsii_.InvokeVoid(
		t,
		"resetInitialInstanceCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetInitialVariantWeight() {
	_jsii_.InvokeVoid(
		t,
		"resetInitialVariantWeight",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetManagedInstanceScaling() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedInstanceScaling",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetModelDataDownloadTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetModelDataDownloadTimeoutInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetModelName() {
	_jsii_.InvokeVoid(
		t,
		"resetModelName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetRoutingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetServerlessConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetServerlessConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetVariantName() {
	_jsii_.InvokeVoid(
		t,
		"resetVariantName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ResetVolumeSizeInGb() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeSizeInGb",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpointConfiguration_ShadowProductionVariantsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

