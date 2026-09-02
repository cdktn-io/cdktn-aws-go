package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference interface {
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
	IdleTimeoutInMinutes() *float64
	// Experimental.
	SetIdleTimeoutInMinutes(val *float64)
	// Experimental.
	IdleTimeoutInMinutesInput() *float64
	// Experimental.
	InternalValue() *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty
	// Experimental.
	SetInternalValue(val *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty)
	// Experimental.
	LifecycleManagement() *string
	// Experimental.
	SetLifecycleManagement(val *string)
	// Experimental.
	LifecycleManagementInput() *string
	// Experimental.
	MaxIdleTimeoutInMinutes() *float64
	// Experimental.
	SetMaxIdleTimeoutInMinutes(val *float64)
	// Experimental.
	MaxIdleTimeoutInMinutesInput() *float64
	// Experimental.
	MinIdleTimeoutInMinutes() *float64
	// Experimental.
	SetMinIdleTimeoutInMinutes(val *float64)
	// Experimental.
	MinIdleTimeoutInMinutesInput() *float64
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
	ResetIdleTimeoutInMinutes()
	// Experimental.
	ResetLifecycleManagement()
	// Experimental.
	ResetMaxIdleTimeoutInMinutes()
	// Experimental.
	ResetMinIdleTimeoutInMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference
type jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) IdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) IdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) InternalValue() *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty {
	var returns *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) LifecycleManagement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) LifecycleManagementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) MaxIdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) MaxIdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) MinIdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) MinIdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference_Override(t TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetInternalValue(val *TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetLifecycleManagement(val *string) {
	if err := j.validateSetLifecycleManagementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleManagement",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetMaxIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetMaxIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetMinIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetMinIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIdleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ResetIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		t,
		"resetIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ResetLifecycleManagement() {
	_jsii_.InvokeVoid(
		t,
		"resetLifecycleManagement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ResetMaxIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ResetMinIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		t,
		"resetMinIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

