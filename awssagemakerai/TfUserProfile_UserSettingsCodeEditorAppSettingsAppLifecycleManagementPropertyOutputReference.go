package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference interface {
	cdktn.ComplexObject
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
	Fqn() *string
	// Experimental.
	IdleSettings() TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference
	// Experimental.
	IdleSettingsInput() *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty
	// Experimental.
	InternalValue() *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty
	// Experimental.
	SetInternalValue(val *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty)
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
	PutIdleSettings(value *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty)
	// Experimental.
	ResetIdleSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference
type jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) IdleSettings() TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference {
	var returns TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"idleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) IdleSettingsInput() *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty {
	var returns *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty
	_jsii_.Get(
		j,
		"idleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) InternalValue() *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty {
	var returns *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference_Override(t TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference)SetInternalValue(val *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) PutIdleSettings(value *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty) {
	if err := t.validatePutIdleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIdleSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) ResetIdleSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetIdleSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

