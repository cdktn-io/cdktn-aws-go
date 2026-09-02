package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfModel_PrimaryContainerPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdditionalModelDataSource() TfModel_PrimaryContainerAdditionalModelDataSourcePropertyList
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
	ImageConfig() TfModel_PrimaryContainerImageConfigPropertyOutputReference
	// Experimental.
	ImageConfigInput() *TfModel_PrimaryContainerImageConfigProperty
	// Experimental.
	ImageInput() *string
	// Experimental.
	InferenceSpecificationName() *string
	// Experimental.
	SetInferenceSpecificationName(val *string)
	// Experimental.
	InferenceSpecificationNameInput() *string
	// Experimental.
	InternalValue() *TfModel_PrimaryContainerProperty
	// Experimental.
	SetInternalValue(val *TfModel_PrimaryContainerProperty)
	// Experimental.
	Mode() *string
	// Experimental.
	SetMode(val *string)
	// Experimental.
	ModeInput() *string
	// Experimental.
	ModelDataSource() TfModel_PrimaryContainerModelDataSourcePropertyOutputReference
	// Experimental.
	ModelDataSourceInput() *TfModel_PrimaryContainerModelDataSourceProperty
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
	MultiModelConfig() TfModel_PrimaryContainerMultiModelConfigPropertyOutputReference
	// Experimental.
	MultiModelConfigInput() *TfModel_PrimaryContainerMultiModelConfigProperty
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
	PutImageConfig(value *TfModel_PrimaryContainerImageConfigProperty)
	// Experimental.
	PutModelDataSource(value *TfModel_PrimaryContainerModelDataSourceProperty)
	// Experimental.
	PutMultiModelConfig(value *TfModel_PrimaryContainerMultiModelConfigProperty)
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

// The jsii proxy struct for TfModel_PrimaryContainerPropertyOutputReference
type jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) AdditionalModelDataSource() TfModel_PrimaryContainerAdditionalModelDataSourcePropertyList {
	var returns TfModel_PrimaryContainerAdditionalModelDataSourcePropertyList
	_jsii_.Get(
		j,
		"additionalModelDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) AdditionalModelDataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalModelDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ContainerHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ContainerHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ImageConfig() TfModel_PrimaryContainerImageConfigPropertyOutputReference {
	var returns TfModel_PrimaryContainerImageConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ImageConfigInput() *TfModel_PrimaryContainerImageConfigProperty {
	var returns *TfModel_PrimaryContainerImageConfigProperty
	_jsii_.Get(
		j,
		"imageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) InferenceSpecificationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceSpecificationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) InferenceSpecificationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceSpecificationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) InternalValue() *TfModel_PrimaryContainerProperty {
	var returns *TfModel_PrimaryContainerProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ModelDataSource() TfModel_PrimaryContainerModelDataSourcePropertyOutputReference {
	var returns TfModel_PrimaryContainerModelDataSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"modelDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ModelDataSourceInput() *TfModel_PrimaryContainerModelDataSourceProperty {
	var returns *TfModel_PrimaryContainerModelDataSourceProperty
	_jsii_.Get(
		j,
		"modelDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ModelDataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ModelDataUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ModelPackageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ModelPackageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) MultiModelConfig() TfModel_PrimaryContainerMultiModelConfigPropertyOutputReference {
	var returns TfModel_PrimaryContainerMultiModelConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"multiModelConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) MultiModelConfigInput() *TfModel_PrimaryContainerMultiModelConfigProperty {
	var returns *TfModel_PrimaryContainerMultiModelConfigProperty
	_jsii_.Get(
		j,
		"multiModelConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfModel_PrimaryContainerPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfModel_PrimaryContainerPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfModel_PrimaryContainerPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfModel.PrimaryContainerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfModel_PrimaryContainerPropertyOutputReference_Override(t TfModel_PrimaryContainerPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfModel.PrimaryContainerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetContainerHostname(val *string) {
	if err := j.validateSetContainerHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerHostname",
		val,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetInferenceSpecificationName(val *string) {
	if err := j.validateSetInferenceSpecificationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceSpecificationName",
		val,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetInternalValue(val *TfModel_PrimaryContainerProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetModelDataUrl(val *string) {
	if err := j.validateSetModelDataUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataUrl",
		val,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetModelPackageName(val *string) {
	if err := j.validateSetModelPackageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelPackageName",
		val,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) PutAdditionalModelDataSource(value interface{}) {
	if err := t.validatePutAdditionalModelDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdditionalModelDataSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) PutImageConfig(value *TfModel_PrimaryContainerImageConfigProperty) {
	if err := t.validatePutImageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImageConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) PutModelDataSource(value *TfModel_PrimaryContainerModelDataSourceProperty) {
	if err := t.validatePutModelDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putModelDataSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) PutMultiModelConfig(value *TfModel_PrimaryContainerMultiModelConfigProperty) {
	if err := t.validatePutMultiModelConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMultiModelConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ResetAdditionalModelDataSource() {
	_jsii_.InvokeVoid(
		t,
		"resetAdditionalModelDataSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ResetContainerHostname() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerHostname",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		t,
		"resetImage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ResetImageConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetImageConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ResetInferenceSpecificationName() {
	_jsii_.InvokeVoid(
		t,
		"resetInferenceSpecificationName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		t,
		"resetMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ResetModelDataSource() {
	_jsii_.InvokeVoid(
		t,
		"resetModelDataSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ResetModelDataUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetModelDataUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ResetModelPackageName() {
	_jsii_.InvokeVoid(
		t,
		"resetModelPackageName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ResetMultiModelConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetMultiModelConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfModel_PrimaryContainerPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

