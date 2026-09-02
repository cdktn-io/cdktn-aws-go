package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AppLifecycleManagement() TfDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
	// Experimental.
	AppLifecycleManagementInput() *TfDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	// Experimental.
	BuiltInLifecycleConfigArn() *string
	// Experimental.
	SetBuiltInLifecycleConfigArn(val *string)
	// Experimental.
	BuiltInLifecycleConfigArnInput() *string
	// Experimental.
	CodeRepository() TfDomain_DefaultUserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList
	// Experimental.
	CodeRepositoryInput() interface{}
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomImage() TfDomain_DefaultUserSettingsJupyterLabAppSettingsCustomImagePropertyList
	// Experimental.
	CustomImageInput() interface{}
	// Experimental.
	DefaultResourceSpec() TfDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference
	// Experimental.
	DefaultResourceSpecInput() *TfDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	// Experimental.
	EmrSettings() TfDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference
	// Experimental.
	EmrSettingsInput() *TfDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDomain_DefaultUserSettingsJupyterLabAppSettingsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_DefaultUserSettingsJupyterLabAppSettingsProperty)
	// Experimental.
	LifecycleConfigArns() *[]*string
	// Experimental.
	SetLifecycleConfigArns(val *[]*string)
	// Experimental.
	LifecycleConfigArnsInput() *[]*string
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
	PutAppLifecycleManagement(value *TfDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty)
	// Experimental.
	PutCodeRepository(value interface{})
	// Experimental.
	PutCustomImage(value interface{})
	// Experimental.
	PutDefaultResourceSpec(value *TfDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty)
	// Experimental.
	PutEmrSettings(value *TfDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty)
	// Experimental.
	ResetAppLifecycleManagement()
	// Experimental.
	ResetBuiltInLifecycleConfigArn()
	// Experimental.
	ResetCodeRepository()
	// Experimental.
	ResetCustomImage()
	// Experimental.
	ResetDefaultResourceSpec()
	// Experimental.
	ResetEmrSettings()
	// Experimental.
	ResetLifecycleConfigArns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference
type jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) AppLifecycleManagement() TfDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference {
	var returns TfDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
	_jsii_.Get(
		j,
		"appLifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) AppLifecycleManagementInput() *TfDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty {
	var returns *TfDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	_jsii_.Get(
		j,
		"appLifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) CodeRepository() TfDomain_DefaultUserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList {
	var returns TfDomain_DefaultUserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) CodeRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) CustomImage() TfDomain_DefaultUserSettingsJupyterLabAppSettingsCustomImagePropertyList {
	var returns TfDomain_DefaultUserSettingsJupyterLabAppSettingsCustomImagePropertyList
	_jsii_.Get(
		j,
		"customImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) CustomImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) DefaultResourceSpec() TfDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference {
	var returns TfDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) DefaultResourceSpecInput() *TfDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty {
	var returns *TfDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) EmrSettings() TfDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference {
	var returns TfDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"emrSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) EmrSettingsInput() *TfDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty {
	var returns *TfDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty
	_jsii_.Get(
		j,
		"emrSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) InternalValue() *TfDomain_DefaultUserSettingsJupyterLabAppSettingsProperty {
	var returns *TfDomain_DefaultUserSettingsJupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) LifecycleConfigArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) LifecycleConfigArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference_Override(t TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetBuiltInLifecycleConfigArn(val *string) {
	if err := j.validateSetBuiltInLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInLifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetInternalValue(val *TfDomain_DefaultUserSettingsJupyterLabAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetLifecycleConfigArns(val *[]*string) {
	if err := j.validateSetLifecycleConfigArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArns",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) PutAppLifecycleManagement(value *TfDomain_DefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty) {
	if err := t.validatePutAppLifecycleManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAppLifecycleManagement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) PutCodeRepository(value interface{}) {
	if err := t.validatePutCodeRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeRepository",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) PutCustomImage(value interface{}) {
	if err := t.validatePutCustomImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomImage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) PutDefaultResourceSpec(value *TfDomain_DefaultUserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty) {
	if err := t.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) PutEmrSettings(value *TfDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty) {
	if err := t.validatePutEmrSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEmrSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetAppLifecycleManagement() {
	_jsii_.InvokeVoid(
		t,
		"resetAppLifecycleManagement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetBuiltInLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		t,
		"resetBuiltInLifecycleConfigArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetCustomImage() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomImage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetEmrSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetEmrSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ResetLifecycleConfigArns() {
	_jsii_.InvokeVoid(
		t,
		"resetLifecycleConfigArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

