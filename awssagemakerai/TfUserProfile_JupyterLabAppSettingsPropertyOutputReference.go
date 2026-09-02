package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserProfile_JupyterLabAppSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AppLifecycleManagement() TfUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
	// Experimental.
	AppLifecycleManagementInput() *TfUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	// Experimental.
	BuiltInLifecycleConfigArn() *string
	// Experimental.
	SetBuiltInLifecycleConfigArn(val *string)
	// Experimental.
	BuiltInLifecycleConfigArnInput() *string
	// Experimental.
	CodeRepository() TfUserProfile_UserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList
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
	CustomImage() TfUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyList
	// Experimental.
	CustomImageInput() interface{}
	// Experimental.
	DefaultResourceSpec() TfUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference
	// Experimental.
	DefaultResourceSpecInput() *TfUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	// Experimental.
	EmrSettings() TfUserProfile_EmrSettingsPropertyOutputReference
	// Experimental.
	EmrSettingsInput() *TfUserProfile_EmrSettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfUserProfile_JupyterLabAppSettingsProperty
	// Experimental.
	SetInternalValue(val *TfUserProfile_JupyterLabAppSettingsProperty)
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
	PutAppLifecycleManagement(value *TfUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty)
	// Experimental.
	PutCodeRepository(value interface{})
	// Experimental.
	PutCustomImage(value interface{})
	// Experimental.
	PutDefaultResourceSpec(value *TfUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty)
	// Experimental.
	PutEmrSettings(value *TfUserProfile_EmrSettingsProperty)
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

// The jsii proxy struct for TfUserProfile_JupyterLabAppSettingsPropertyOutputReference
type jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) AppLifecycleManagement() TfUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference {
	var returns TfUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementPropertyOutputReference
	_jsii_.Get(
		j,
		"appLifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) AppLifecycleManagementInput() *TfUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty {
	var returns *TfUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty
	_jsii_.Get(
		j,
		"appLifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) CodeRepository() TfUserProfile_UserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList {
	var returns TfUserProfile_UserSettingsJupyterLabAppSettingsCodeRepositoryPropertyList
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) CodeRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) CustomImage() TfUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyList {
	var returns TfUserProfile_UserSettingsJupyterLabAppSettingsCustomImagePropertyList
	_jsii_.Get(
		j,
		"customImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) CustomImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) DefaultResourceSpec() TfUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference {
	var returns TfUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) DefaultResourceSpecInput() *TfUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty {
	var returns *TfUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) EmrSettings() TfUserProfile_EmrSettingsPropertyOutputReference {
	var returns TfUserProfile_EmrSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"emrSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) EmrSettingsInput() *TfUserProfile_EmrSettingsProperty {
	var returns *TfUserProfile_EmrSettingsProperty
	_jsii_.Get(
		j,
		"emrSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) InternalValue() *TfUserProfile_JupyterLabAppSettingsProperty {
	var returns *TfUserProfile_JupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) LifecycleConfigArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) LifecycleConfigArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserProfile_JupyterLabAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserProfile_JupyterLabAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserProfile_JupyterLabAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.JupyterLabAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserProfile_JupyterLabAppSettingsPropertyOutputReference_Override(t TfUserProfile_JupyterLabAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.JupyterLabAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetBuiltInLifecycleConfigArn(val *string) {
	if err := j.validateSetBuiltInLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInLifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetInternalValue(val *TfUserProfile_JupyterLabAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetLifecycleConfigArns(val *[]*string) {
	if err := j.validateSetLifecycleConfigArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArns",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) PutAppLifecycleManagement(value *TfUserProfile_UserSettingsJupyterLabAppSettingsAppLifecycleManagementProperty) {
	if err := t.validatePutAppLifecycleManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAppLifecycleManagement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) PutCodeRepository(value interface{}) {
	if err := t.validatePutCodeRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeRepository",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) PutCustomImage(value interface{}) {
	if err := t.validatePutCustomImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomImage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) PutDefaultResourceSpec(value *TfUserProfile_UserSettingsJupyterLabAppSettingsDefaultResourceSpecProperty) {
	if err := t.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) PutEmrSettings(value *TfUserProfile_EmrSettingsProperty) {
	if err := t.validatePutEmrSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEmrSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetAppLifecycleManagement() {
	_jsii_.InvokeVoid(
		t,
		"resetAppLifecycleManagement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetBuiltInLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		t,
		"resetBuiltInLifecycleConfigArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetCustomImage() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomImage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetEmrSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetEmrSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) ResetLifecycleConfigArns() {
	_jsii_.InvokeVoid(
		t,
		"resetLifecycleConfigArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserProfile_JupyterLabAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

