package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserPool_VerificationMessageTemplatePropertyOutputReference interface {
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
	DefaultEmailOption() *string
	// Experimental.
	SetDefaultEmailOption(val *string)
	// Experimental.
	DefaultEmailOptionInput() *string
	// Experimental.
	EmailMessage() *string
	// Experimental.
	SetEmailMessage(val *string)
	// Experimental.
	EmailMessageByLink() *string
	// Experimental.
	SetEmailMessageByLink(val *string)
	// Experimental.
	EmailMessageByLinkInput() *string
	// Experimental.
	EmailMessageInput() *string
	// Experimental.
	EmailSubject() *string
	// Experimental.
	SetEmailSubject(val *string)
	// Experimental.
	EmailSubjectByLink() *string
	// Experimental.
	SetEmailSubjectByLink(val *string)
	// Experimental.
	EmailSubjectByLinkInput() *string
	// Experimental.
	EmailSubjectInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfUserPool_VerificationMessageTemplateProperty
	// Experimental.
	SetInternalValue(val *TfUserPool_VerificationMessageTemplateProperty)
	// Experimental.
	SmsMessage() *string
	// Experimental.
	SetSmsMessage(val *string)
	// Experimental.
	SmsMessageInput() *string
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
	ResetDefaultEmailOption()
	// Experimental.
	ResetEmailMessage()
	// Experimental.
	ResetEmailMessageByLink()
	// Experimental.
	ResetEmailSubject()
	// Experimental.
	ResetEmailSubjectByLink()
	// Experimental.
	ResetSmsMessage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUserPool_VerificationMessageTemplatePropertyOutputReference
type jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) DefaultEmailOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultEmailOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) DefaultEmailOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultEmailOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) EmailMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) EmailMessageByLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageByLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) EmailMessageByLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageByLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) EmailMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) EmailSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) EmailSubjectByLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectByLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) EmailSubjectByLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectByLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) EmailSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) InternalValue() *TfUserPool_VerificationMessageTemplateProperty {
	var returns *TfUserPool_VerificationMessageTemplateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) SmsMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) SmsMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smsMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserPool_VerificationMessageTemplatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserPool_VerificationMessageTemplatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserPool_VerificationMessageTemplatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfUserPool.VerificationMessageTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserPool_VerificationMessageTemplatePropertyOutputReference_Override(t TfUserPool_VerificationMessageTemplatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfUserPool.VerificationMessageTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference)SetDefaultEmailOption(val *string) {
	if err := j.validateSetDefaultEmailOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultEmailOption",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference)SetEmailMessage(val *string) {
	if err := j.validateSetEmailMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailMessage",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference)SetEmailMessageByLink(val *string) {
	if err := j.validateSetEmailMessageByLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailMessageByLink",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference)SetEmailSubject(val *string) {
	if err := j.validateSetEmailSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSubject",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference)SetEmailSubjectByLink(val *string) {
	if err := j.validateSetEmailSubjectByLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSubjectByLink",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference)SetInternalValue(val *TfUserPool_VerificationMessageTemplateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference)SetSmsMessage(val *string) {
	if err := j.validateSetSmsMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsMessage",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) ResetDefaultEmailOption() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultEmailOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) ResetEmailMessage() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailMessage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) ResetEmailMessageByLink() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailMessageByLink",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) ResetEmailSubject() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailSubject",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) ResetEmailSubjectByLink() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailSubjectByLink",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) ResetSmsMessage() {
	_jsii_.InvokeVoid(
		t,
		"resetSmsMessage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserPool_VerificationMessageTemplatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

