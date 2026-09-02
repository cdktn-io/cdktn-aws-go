package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnect/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfQuickConnect_QuickConnectConfigPropertyOutputReference interface {
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
	InternalValue() *TfQuickConnect_QuickConnectConfigProperty
	// Experimental.
	SetInternalValue(val *TfQuickConnect_QuickConnectConfigProperty)
	// Experimental.
	PhoneConfig() TfQuickConnect_PhoneConfigPropertyList
	// Experimental.
	PhoneConfigInput() interface{}
	// Experimental.
	QueueConfig() TfQuickConnect_QueueConfigPropertyList
	// Experimental.
	QueueConfigInput() interface{}
	// Experimental.
	QuickConnectType() *string
	// Experimental.
	SetQuickConnectType(val *string)
	// Experimental.
	QuickConnectTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserConfig() TfQuickConnect_UserConfigPropertyList
	// Experimental.
	UserConfigInput() interface{}
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
	PutPhoneConfig(value interface{})
	// Experimental.
	PutQueueConfig(value interface{})
	// Experimental.
	PutUserConfig(value interface{})
	// Experimental.
	ResetPhoneConfig()
	// Experimental.
	ResetQueueConfig()
	// Experimental.
	ResetUserConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfQuickConnect_QuickConnectConfigPropertyOutputReference
type jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) InternalValue() *TfQuickConnect_QuickConnectConfigProperty {
	var returns *TfQuickConnect_QuickConnectConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) PhoneConfig() TfQuickConnect_PhoneConfigPropertyList {
	var returns TfQuickConnect_PhoneConfigPropertyList
	_jsii_.Get(
		j,
		"phoneConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) PhoneConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phoneConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) QueueConfig() TfQuickConnect_QueueConfigPropertyList {
	var returns TfQuickConnect_QueueConfigPropertyList
	_jsii_.Get(
		j,
		"queueConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) QueueConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queueConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) QuickConnectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quickConnectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) QuickConnectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quickConnectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) UserConfig() TfQuickConnect_UserConfigPropertyList {
	var returns TfQuickConnect_UserConfigPropertyList
	_jsii_.Get(
		j,
		"userConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) UserConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfQuickConnect_QuickConnectConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfQuickConnect_QuickConnectConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfQuickConnect_QuickConnectConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect.TfQuickConnect.QuickConnectConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfQuickConnect_QuickConnectConfigPropertyOutputReference_Override(t TfQuickConnect_QuickConnectConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect.TfQuickConnect.QuickConnectConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference)SetInternalValue(val *TfQuickConnect_QuickConnectConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference)SetQuickConnectType(val *string) {
	if err := j.validateSetQuickConnectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quickConnectType",
		val,
	)
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) PutPhoneConfig(value interface{}) {
	if err := t.validatePutPhoneConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPhoneConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) PutQueueConfig(value interface{}) {
	if err := t.validatePutQueueConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueueConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) PutUserConfig(value interface{}) {
	if err := t.validatePutUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUserConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) ResetPhoneConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetPhoneConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) ResetQueueConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetQueueConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) ResetUserConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetUserConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfQuickConnect_QuickConnectConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

