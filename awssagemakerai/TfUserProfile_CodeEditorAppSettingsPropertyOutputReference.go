package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserProfile_CodeEditorAppSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AppLifecycleManagement() TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference
	// Experimental.
	AppLifecycleManagementInput() *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty
	// Experimental.
	BuiltInLifecycleConfigArn() *string
	// Experimental.
	SetBuiltInLifecycleConfigArn(val *string)
	// Experimental.
	BuiltInLifecycleConfigArnInput() *string
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
	CustomImage() TfUserProfile_UserSettingsCodeEditorAppSettingsCustomImagePropertyList
	// Experimental.
	CustomImageInput() interface{}
	// Experimental.
	DefaultResourceSpec() TfUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecPropertyOutputReference
	// Experimental.
	DefaultResourceSpecInput() *TfUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfUserProfile_CodeEditorAppSettingsProperty
	// Experimental.
	SetInternalValue(val *TfUserProfile_CodeEditorAppSettingsProperty)
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
	PutAppLifecycleManagement(value *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty)
	// Experimental.
	PutCustomImage(value interface{})
	// Experimental.
	PutDefaultResourceSpec(value *TfUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty)
	// Experimental.
	ResetAppLifecycleManagement()
	// Experimental.
	ResetBuiltInLifecycleConfigArn()
	// Experimental.
	ResetCustomImage()
	// Experimental.
	ResetDefaultResourceSpec()
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

// The jsii proxy struct for TfUserProfile_CodeEditorAppSettingsPropertyOutputReference
type jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) AppLifecycleManagement() TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference {
	var returns TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference
	_jsii_.Get(
		j,
		"appLifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) AppLifecycleManagementInput() *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty {
	var returns *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty
	_jsii_.Get(
		j,
		"appLifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) BuiltInLifecycleConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"builtInLifecycleConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) CustomImage() TfUserProfile_UserSettingsCodeEditorAppSettingsCustomImagePropertyList {
	var returns TfUserProfile_UserSettingsCodeEditorAppSettingsCustomImagePropertyList
	_jsii_.Get(
		j,
		"customImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) CustomImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) DefaultResourceSpec() TfUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecPropertyOutputReference {
	var returns TfUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultResourceSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) DefaultResourceSpecInput() *TfUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty {
	var returns *TfUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty
	_jsii_.Get(
		j,
		"defaultResourceSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) InternalValue() *TfUserProfile_CodeEditorAppSettingsProperty {
	var returns *TfUserProfile_CodeEditorAppSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) LifecycleConfigArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) LifecycleConfigArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lifecycleConfigArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserProfile_CodeEditorAppSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserProfile_CodeEditorAppSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserProfile_CodeEditorAppSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.CodeEditorAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserProfile_CodeEditorAppSettingsPropertyOutputReference_Override(t TfUserProfile_CodeEditorAppSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.CodeEditorAppSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetBuiltInLifecycleConfigArn(val *string) {
	if err := j.validateSetBuiltInLifecycleConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"builtInLifecycleConfigArn",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetInternalValue(val *TfUserProfile_CodeEditorAppSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetLifecycleConfigArns(val *[]*string) {
	if err := j.validateSetLifecycleConfigArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfigArns",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) PutAppLifecycleManagement(value *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty) {
	if err := t.validatePutAppLifecycleManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAppLifecycleManagement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) PutCustomImage(value interface{}) {
	if err := t.validatePutCustomImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomImage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) PutDefaultResourceSpec(value *TfUserProfile_UserSettingsCodeEditorAppSettingsDefaultResourceSpecProperty) {
	if err := t.validatePutDefaultResourceSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultResourceSpec",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) ResetAppLifecycleManagement() {
	_jsii_.InvokeVoid(
		t,
		"resetAppLifecycleManagement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) ResetBuiltInLifecycleConfigArn() {
	_jsii_.InvokeVoid(
		t,
		"resetBuiltInLifecycleConfigArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) ResetCustomImage() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomImage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) ResetDefaultResourceSpec() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultResourceSpec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) ResetLifecycleConfigArns() {
	_jsii_.InvokeVoid(
		t,
		"resetLifecycleConfigArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserProfile_CodeEditorAppSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

