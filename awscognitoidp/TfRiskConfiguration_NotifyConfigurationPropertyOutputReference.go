package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRiskConfiguration_NotifyConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BlockEmail() TfRiskConfiguration_BlockEmailPropertyOutputReference
	// Experimental.
	BlockEmailInput() *TfRiskConfiguration_BlockEmailProperty
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
	From() *string
	// Experimental.
	SetFrom(val *string)
	// Experimental.
	FromInput() *string
	// Experimental.
	InternalValue() *TfRiskConfiguration_NotifyConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfRiskConfiguration_NotifyConfigurationProperty)
	// Experimental.
	MfaEmail() TfRiskConfiguration_MfaEmailPropertyOutputReference
	// Experimental.
	MfaEmailInput() *TfRiskConfiguration_MfaEmailProperty
	// Experimental.
	NoActionEmail() TfRiskConfiguration_NoActionEmailPropertyOutputReference
	// Experimental.
	NoActionEmailInput() *TfRiskConfiguration_NoActionEmailProperty
	// Experimental.
	ReplyTo() *string
	// Experimental.
	SetReplyTo(val *string)
	// Experimental.
	ReplyToInput() *string
	// Experimental.
	SourceArn() *string
	// Experimental.
	SetSourceArn(val *string)
	// Experimental.
	SourceArnInput() *string
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
	PutBlockEmail(value *TfRiskConfiguration_BlockEmailProperty)
	// Experimental.
	PutMfaEmail(value *TfRiskConfiguration_MfaEmailProperty)
	// Experimental.
	PutNoActionEmail(value *TfRiskConfiguration_NoActionEmailProperty)
	// Experimental.
	ResetBlockEmail()
	// Experimental.
	ResetFrom()
	// Experimental.
	ResetMfaEmail()
	// Experimental.
	ResetNoActionEmail()
	// Experimental.
	ResetReplyTo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRiskConfiguration_NotifyConfigurationPropertyOutputReference
type jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) BlockEmail() TfRiskConfiguration_BlockEmailPropertyOutputReference {
	var returns TfRiskConfiguration_BlockEmailPropertyOutputReference
	_jsii_.Get(
		j,
		"blockEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) BlockEmailInput() *TfRiskConfiguration_BlockEmailProperty {
	var returns *TfRiskConfiguration_BlockEmailProperty
	_jsii_.Get(
		j,
		"blockEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) From() *string {
	var returns *string
	_jsii_.Get(
		j,
		"from",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) FromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) InternalValue() *TfRiskConfiguration_NotifyConfigurationProperty {
	var returns *TfRiskConfiguration_NotifyConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) MfaEmail() TfRiskConfiguration_MfaEmailPropertyOutputReference {
	var returns TfRiskConfiguration_MfaEmailPropertyOutputReference
	_jsii_.Get(
		j,
		"mfaEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) MfaEmailInput() *TfRiskConfiguration_MfaEmailProperty {
	var returns *TfRiskConfiguration_MfaEmailProperty
	_jsii_.Get(
		j,
		"mfaEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) NoActionEmail() TfRiskConfiguration_NoActionEmailPropertyOutputReference {
	var returns TfRiskConfiguration_NoActionEmailPropertyOutputReference
	_jsii_.Get(
		j,
		"noActionEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) NoActionEmailInput() *TfRiskConfiguration_NoActionEmailProperty {
	var returns *TfRiskConfiguration_NoActionEmailProperty
	_jsii_.Get(
		j,
		"noActionEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) ReplyTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) ReplyToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) SourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRiskConfiguration_NotifyConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRiskConfiguration_NotifyConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRiskConfiguration_NotifyConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfRiskConfiguration.NotifyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRiskConfiguration_NotifyConfigurationPropertyOutputReference_Override(t TfRiskConfiguration_NotifyConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfRiskConfiguration.NotifyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetFrom(val *string) {
	if err := j.validateSetFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"from",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetInternalValue(val *TfRiskConfiguration_NotifyConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetReplyTo(val *string) {
	if err := j.validateSetReplyToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replyTo",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetSourceArn(val *string) {
	if err := j.validateSetSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) PutBlockEmail(value *TfRiskConfiguration_BlockEmailProperty) {
	if err := t.validatePutBlockEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBlockEmail",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) PutMfaEmail(value *TfRiskConfiguration_MfaEmailProperty) {
	if err := t.validatePutMfaEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMfaEmail",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) PutNoActionEmail(value *TfRiskConfiguration_NoActionEmailProperty) {
	if err := t.validatePutNoActionEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNoActionEmail",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetBlockEmail() {
	_jsii_.InvokeVoid(
		t,
		"resetBlockEmail",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetFrom() {
	_jsii_.InvokeVoid(
		t,
		"resetFrom",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetMfaEmail() {
	_jsii_.InvokeVoid(
		t,
		"resetMfaEmail",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetNoActionEmail() {
	_jsii_.InvokeVoid(
		t,
		"resetNoActionEmail",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) ResetReplyTo() {
	_jsii_.InvokeVoid(
		t,
		"resetReplyTo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRiskConfiguration_NotifyConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

