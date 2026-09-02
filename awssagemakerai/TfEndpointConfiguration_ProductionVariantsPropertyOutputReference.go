package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpointConfiguration_ProductionVariantsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AcceleratorType() *string
	// Experimental.
	SetAcceleratorType(val *string)
	// Experimental.
	AcceleratorTypeInput() *string
	// Experimental.
	CapacityReservationConfig() TfEndpointConfiguration_ProductionVariantsCapacityReservationConfigPropertyOutputReference
	// Experimental.
	CapacityReservationConfigInput() *TfEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty
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
	CoreDumpConfig() TfEndpointConfiguration_ProductionVariantsCoreDumpConfigPropertyOutputReference
	// Experimental.
	CoreDumpConfigInput() *TfEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty
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
	ManagedInstanceScaling() TfEndpointConfiguration_ProductionVariantsManagedInstanceScalingPropertyOutputReference
	// Experimental.
	ManagedInstanceScalingInput() *TfEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty
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
	RoutingConfig() TfEndpointConfiguration_ProductionVariantsRoutingConfigPropertyList
	// Experimental.
	RoutingConfigInput() interface{}
	// Experimental.
	ServerlessConfig() TfEndpointConfiguration_ProductionVariantsServerlessConfigPropertyOutputReference
	// Experimental.
	ServerlessConfigInput() *TfEndpointConfiguration_ProductionVariantsServerlessConfigProperty
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
	PutCapacityReservationConfig(value *TfEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty)
	// Experimental.
	PutCoreDumpConfig(value *TfEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty)
	// Experimental.
	PutManagedInstanceScaling(value *TfEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty)
	// Experimental.
	PutRoutingConfig(value interface{})
	// Experimental.
	PutServerlessConfig(value *TfEndpointConfiguration_ProductionVariantsServerlessConfigProperty)
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

// The jsii proxy struct for TfEndpointConfiguration_ProductionVariantsPropertyOutputReference
type jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) AcceleratorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) AcceleratorTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) CapacityReservationConfig() TfEndpointConfiguration_ProductionVariantsCapacityReservationConfigPropertyOutputReference {
	var returns TfEndpointConfiguration_ProductionVariantsCapacityReservationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) CapacityReservationConfigInput() *TfEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty {
	var returns *TfEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty
	_jsii_.Get(
		j,
		"capacityReservationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ContainerStartupHealthCheckTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ContainerStartupHealthCheckTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) CoreDumpConfig() TfEndpointConfiguration_ProductionVariantsCoreDumpConfigPropertyOutputReference {
	var returns TfEndpointConfiguration_ProductionVariantsCoreDumpConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"coreDumpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) CoreDumpConfigInput() *TfEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty {
	var returns *TfEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty
	_jsii_.Get(
		j,
		"coreDumpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) EnableSsmAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsmAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) EnableSsmAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsmAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) InferenceAmiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAmiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) InferenceAmiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAmiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) InitialInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) InitialInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) InitialVariantWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialVariantWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) InitialVariantWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialVariantWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ManagedInstanceScaling() TfEndpointConfiguration_ProductionVariantsManagedInstanceScalingPropertyOutputReference {
	var returns TfEndpointConfiguration_ProductionVariantsManagedInstanceScalingPropertyOutputReference
	_jsii_.Get(
		j,
		"managedInstanceScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ManagedInstanceScalingInput() *TfEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty {
	var returns *TfEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty
	_jsii_.Get(
		j,
		"managedInstanceScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ModelDataDownloadTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ModelDataDownloadTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) RoutingConfig() TfEndpointConfiguration_ProductionVariantsRoutingConfigPropertyList {
	var returns TfEndpointConfiguration_ProductionVariantsRoutingConfigPropertyList
	_jsii_.Get(
		j,
		"routingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) RoutingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ServerlessConfig() TfEndpointConfiguration_ProductionVariantsServerlessConfigPropertyOutputReference {
	var returns TfEndpointConfiguration_ProductionVariantsServerlessConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"serverlessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ServerlessConfigInput() *TfEndpointConfiguration_ProductionVariantsServerlessConfigProperty {
	var returns *TfEndpointConfiguration_ProductionVariantsServerlessConfigProperty
	_jsii_.Get(
		j,
		"serverlessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) VariantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) VariantNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) VolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) VolumeSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGbInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpointConfiguration_ProductionVariantsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfEndpointConfiguration_ProductionVariantsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpointConfiguration_ProductionVariantsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpointConfiguration.ProductionVariantsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpointConfiguration_ProductionVariantsPropertyOutputReference_Override(t TfEndpointConfiguration_ProductionVariantsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpointConfiguration.ProductionVariantsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetAcceleratorType(val *string) {
	if err := j.validateSetAcceleratorTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorType",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetContainerStartupHealthCheckTimeoutInSeconds(val *float64) {
	if err := j.validateSetContainerStartupHealthCheckTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetEnableSsmAccess(val interface{}) {
	if err := j.validateSetEnableSsmAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSsmAccess",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetInferenceAmiVersion(val *string) {
	if err := j.validateSetInferenceAmiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceAmiVersion",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetInitialInstanceCount(val *float64) {
	if err := j.validateSetInitialInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialInstanceCount",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetInitialVariantWeight(val *float64) {
	if err := j.validateSetInitialVariantWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialVariantWeight",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetModelDataDownloadTimeoutInSeconds(val *float64) {
	if err := j.validateSetModelDataDownloadTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataDownloadTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetVariantName(val *string) {
	if err := j.validateSetVariantNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variantName",
		val,
	)
}

func (j *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference)SetVolumeSizeInGb(val *float64) {
	if err := j.validateSetVolumeSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSizeInGb",
		val,
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) PutCapacityReservationConfig(value *TfEndpointConfiguration_ProductionVariantsCapacityReservationConfigProperty) {
	if err := t.validatePutCapacityReservationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityReservationConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) PutCoreDumpConfig(value *TfEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty) {
	if err := t.validatePutCoreDumpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCoreDumpConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) PutManagedInstanceScaling(value *TfEndpointConfiguration_ProductionVariantsManagedInstanceScalingProperty) {
	if err := t.validatePutManagedInstanceScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedInstanceScaling",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) PutRoutingConfig(value interface{}) {
	if err := t.validatePutRoutingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRoutingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) PutServerlessConfig(value *TfEndpointConfiguration_ProductionVariantsServerlessConfigProperty) {
	if err := t.validatePutServerlessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServerlessConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetAcceleratorType() {
	_jsii_.InvokeVoid(
		t,
		"resetAcceleratorType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetCapacityReservationConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityReservationConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetContainerStartupHealthCheckTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerStartupHealthCheckTimeoutInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetCoreDumpConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCoreDumpConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetEnableSsmAccess() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableSsmAccess",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetInferenceAmiVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetInferenceAmiVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetInitialInstanceCount() {
	_jsii_.InvokeVoid(
		t,
		"resetInitialInstanceCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetInitialVariantWeight() {
	_jsii_.InvokeVoid(
		t,
		"resetInitialVariantWeight",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetManagedInstanceScaling() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedInstanceScaling",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetModelDataDownloadTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetModelDataDownloadTimeoutInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetModelName() {
	_jsii_.InvokeVoid(
		t,
		"resetModelName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetRoutingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetServerlessConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetServerlessConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetVariantName() {
	_jsii_.InvokeVoid(
		t,
		"resetVariantName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ResetVolumeSizeInGb() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeSizeInGb",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpointConfiguration_ProductionVariantsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

