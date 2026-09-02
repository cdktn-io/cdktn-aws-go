package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsManagedRulesAcfpRuleSet() TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyList
	// Experimental.
	AwsManagedRulesAcfpRuleSetInput() interface{}
	// Experimental.
	AwsManagedRulesAntiDdosRuleSet() TfWebAclRule_AwsManagedRulesAntiDdosRuleSetPropertyList
	// Experimental.
	AwsManagedRulesAntiDdosRuleSetInput() interface{}
	// Experimental.
	AwsManagedRulesAtpRuleSet() TfWebAclRule_AwsManagedRulesAtpRuleSetPropertyList
	// Experimental.
	AwsManagedRulesAtpRuleSetInput() interface{}
	// Experimental.
	AwsManagedRulesBotControlRuleSet() TfWebAclRule_AwsManagedRulesBotControlRuleSetPropertyList
	// Experimental.
	AwsManagedRulesBotControlRuleSetInput() interface{}
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LoginPath() *string
	// Experimental.
	SetLoginPath(val *string)
	// Experimental.
	LoginPathInput() *string
	// Experimental.
	PasswordField() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsPasswordFieldPropertyList
	// Experimental.
	PasswordFieldInput() interface{}
	// Experimental.
	PayloadType() *string
	// Experimental.
	SetPayloadType(val *string)
	// Experimental.
	PayloadTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UsernameField() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsUsernameFieldPropertyList
	// Experimental.
	UsernameFieldInput() interface{}
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
	PutAwsManagedRulesAcfpRuleSet(value interface{})
	// Experimental.
	PutAwsManagedRulesAntiDdosRuleSet(value interface{})
	// Experimental.
	PutAwsManagedRulesAtpRuleSet(value interface{})
	// Experimental.
	PutAwsManagedRulesBotControlRuleSet(value interface{})
	// Experimental.
	PutPasswordField(value interface{})
	// Experimental.
	PutUsernameField(value interface{})
	// Experimental.
	ResetAwsManagedRulesAcfpRuleSet()
	// Experimental.
	ResetAwsManagedRulesAntiDdosRuleSet()
	// Experimental.
	ResetAwsManagedRulesAtpRuleSet()
	// Experimental.
	ResetAwsManagedRulesBotControlRuleSet()
	// Experimental.
	ResetLoginPath()
	// Experimental.
	ResetPasswordField()
	// Experimental.
	ResetPayloadType()
	// Experimental.
	ResetUsernameField()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference
type jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAcfpRuleSet() TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyList {
	var returns TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesAcfpRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAcfpRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesAcfpRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAntiDdosRuleSet() TfWebAclRule_AwsManagedRulesAntiDdosRuleSetPropertyList {
	var returns TfWebAclRule_AwsManagedRulesAntiDdosRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesAntiDdosRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAntiDdosRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesAntiDdosRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAtpRuleSet() TfWebAclRule_AwsManagedRulesAtpRuleSetPropertyList {
	var returns TfWebAclRule_AwsManagedRulesAtpRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesAtpRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAtpRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesAtpRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesBotControlRuleSet() TfWebAclRule_AwsManagedRulesBotControlRuleSetPropertyList {
	var returns TfWebAclRule_AwsManagedRulesBotControlRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesBotControlRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesBotControlRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesBotControlRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) LoginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) LoginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PasswordField() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsPasswordFieldPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsPasswordFieldPropertyList
	_jsii_.Get(
		j,
		"passwordField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PasswordFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PayloadType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PayloadTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) UsernameField() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsUsernameFieldPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsUsernameFieldPropertyList
	_jsii_.Get(
		j,
		"usernameField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) UsernameFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usernameFieldInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.ManagedRuleGroupConfigsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference_Override(t TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.ManagedRuleGroupConfigsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetLoginPath(val *string) {
	if err := j.validateSetLoginPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginPath",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetPayloadType(val *string) {
	if err := j.validateSetPayloadTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadType",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesAcfpRuleSet(value interface{}) {
	if err := t.validatePutAwsManagedRulesAcfpRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsManagedRulesAcfpRuleSet",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesAntiDdosRuleSet(value interface{}) {
	if err := t.validatePutAwsManagedRulesAntiDdosRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsManagedRulesAntiDdosRuleSet",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesAtpRuleSet(value interface{}) {
	if err := t.validatePutAwsManagedRulesAtpRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsManagedRulesAtpRuleSet",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesBotControlRuleSet(value interface{}) {
	if err := t.validatePutAwsManagedRulesBotControlRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsManagedRulesBotControlRuleSet",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutPasswordField(value interface{}) {
	if err := t.validatePutPasswordFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPasswordField",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutUsernameField(value interface{}) {
	if err := t.validatePutUsernameFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUsernameField",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesAcfpRuleSet() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsManagedRulesAcfpRuleSet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesAntiDdosRuleSet() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsManagedRulesAntiDdosRuleSet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesAtpRuleSet() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsManagedRulesAtpRuleSet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesBotControlRuleSet() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsManagedRulesBotControlRuleSet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetLoginPath() {
	_jsii_.InvokeVoid(
		t,
		"resetLoginPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetPasswordField() {
	_jsii_.InvokeVoid(
		t,
		"resetPasswordField",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetPayloadType() {
	_jsii_.InvokeVoid(
		t,
		"resetPayloadType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetUsernameField() {
	_jsii_.InvokeVoid(
		t,
		"resetUsernameField",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

