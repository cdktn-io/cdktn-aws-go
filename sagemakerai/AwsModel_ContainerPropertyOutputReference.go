package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsModel_ContainerPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdditionalModelDataSource() AwsModel_ContainerAdditionalModelDataSourcePropertyList
	// Experimental.
	AdditionalModelDataSourceInput() interface{}
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
	ContainerHostname() *string
	// Experimental.
	SetContainerHostname(val *string)
	// Experimental.
	ContainerHostnameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Environment() *map[string]*string
	// Experimental.
	SetEnvironment(val *map[string]*string)
	// Experimental.
	EnvironmentInput() *map[string]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	Image() *string
	// Experimental.
	SetImage(val *string)
	// Experimental.
	ImageConfig() AwsModel_ContainerImageConfigPropertyOutputReference
	// Experimental.
	ImageConfigInput() *AwsModel_ContainerImageConfigProperty
	// Experimental.
	ImageInput() *string
	// Experimental.
	InferenceSpecificationName() *string
	// Experimental.
	SetInferenceSpecificationName(val *string)
	// Experimental.
	InferenceSpecificationNameInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Mode() *string
	// Experimental.
	SetMode(val *string)
	// Experimental.
	ModeInput() *string
	// Experimental.
	ModelDataSource() AwsModel_ContainerModelDataSourcePropertyOutputReference
	// Experimental.
	ModelDataSourceInput() *AwsModel_ContainerModelDataSourceProperty
	// Experimental.
	ModelDataUrl() *string
	// Experimental.
	SetModelDataUrl(val *string)
	// Experimental.
	ModelDataUrlInput() *string
	// Experimental.
	ModelPackageName() *string
	// Experimental.
	SetModelPackageName(val *string)
	// Experimental.
	ModelPackageNameInput() *string
	// Experimental.
	MultiModelConfig() AwsModel_ContainerMultiModelConfigPropertyOutputReference
	// Experimental.
	MultiModelConfigInput() *AwsModel_ContainerMultiModelConfigProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutAdditionalModelDataSource(value interface{})
	// Experimental.
	PutImageConfig(value *AwsModel_ContainerImageConfigProperty)
	// Experimental.
	PutModelDataSource(value *AwsModel_ContainerModelDataSourceProperty)
	// Experimental.
	PutMultiModelConfig(value *AwsModel_ContainerMultiModelConfigProperty)
	// Experimental.
	ResetAdditionalModelDataSource()
	// Experimental.
	ResetContainerHostname()
	// Experimental.
	ResetEnvironment()
	// Experimental.
	ResetImage()
	// Experimental.
	ResetImageConfig()
	// Experimental.
	ResetInferenceSpecificationName()
	// Experimental.
	ResetMode()
	// Experimental.
	ResetModelDataSource()
	// Experimental.
	ResetModelDataUrl()
	// Experimental.
	ResetModelPackageName()
	// Experimental.
	ResetMultiModelConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsModel_ContainerPropertyOutputReference
type jsiiProxy_AwsModel_ContainerPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) AdditionalModelDataSource() AwsModel_ContainerAdditionalModelDataSourcePropertyList {
	var returns AwsModel_ContainerAdditionalModelDataSourcePropertyList
	_jsii_.Get(
		j,
		"additionalModelDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) AdditionalModelDataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalModelDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ContainerHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ContainerHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ImageConfig() AwsModel_ContainerImageConfigPropertyOutputReference {
	var returns AwsModel_ContainerImageConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ImageConfigInput() *AwsModel_ContainerImageConfigProperty {
	var returns *AwsModel_ContainerImageConfigProperty
	_jsii_.Get(
		j,
		"imageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) InferenceSpecificationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceSpecificationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) InferenceSpecificationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceSpecificationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ModelDataSource() AwsModel_ContainerModelDataSourcePropertyOutputReference {
	var returns AwsModel_ContainerModelDataSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"modelDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ModelDataSourceInput() *AwsModel_ContainerModelDataSourceProperty {
	var returns *AwsModel_ContainerModelDataSourceProperty
	_jsii_.Get(
		j,
		"modelDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ModelDataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ModelDataUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ModelPackageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ModelPackageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) MultiModelConfig() AwsModel_ContainerMultiModelConfigPropertyOutputReference {
	var returns AwsModel_ContainerMultiModelConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"multiModelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) MultiModelConfigInput() *AwsModel_ContainerMultiModelConfigProperty {
	var returns *AwsModel_ContainerMultiModelConfigProperty
	_jsii_.Get(
		j,
		"multiModelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsModel_ContainerPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsModel_ContainerPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsModel_ContainerPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsModel_ContainerPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsModel.ContainerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsModel_ContainerPropertyOutputReference_Override(a AwsModel_ContainerPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsModel.ContainerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetContainerHostname(val *string) {
	if err := j.validateSetContainerHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerHostname",
		val,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetInferenceSpecificationName(val *string) {
	if err := j.validateSetInferenceSpecificationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceSpecificationName",
		val,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetModelDataUrl(val *string) {
	if err := j.validateSetModelDataUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataUrl",
		val,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetModelPackageName(val *string) {
	if err := j.validateSetModelPackageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelPackageName",
		val,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsModel_ContainerPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) PutAdditionalModelDataSource(value interface{}) {
	if err := a.validatePutAdditionalModelDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdditionalModelDataSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) PutImageConfig(value *AwsModel_ContainerImageConfigProperty) {
	if err := a.validatePutImageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImageConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) PutModelDataSource(value *AwsModel_ContainerModelDataSourceProperty) {
	if err := a.validatePutModelDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putModelDataSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) PutMultiModelConfig(value *AwsModel_ContainerMultiModelConfigProperty) {
	if err := a.validatePutMultiModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMultiModelConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ResetAdditionalModelDataSource() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalModelDataSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ResetContainerHostname() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerHostname",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		a,
		"resetImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ResetImageConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetImageConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ResetInferenceSpecificationName() {
	_jsii_.InvokeVoid(
		a,
		"resetInferenceSpecificationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		a,
		"resetMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ResetModelDataSource() {
	_jsii_.InvokeVoid(
		a,
		"resetModelDataSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ResetModelDataUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetModelDataUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ResetModelPackageName() {
	_jsii_.InvokeVoid(
		a,
		"resetModelPackageName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ResetMultiModelConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiModelConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsModel_ContainerPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

