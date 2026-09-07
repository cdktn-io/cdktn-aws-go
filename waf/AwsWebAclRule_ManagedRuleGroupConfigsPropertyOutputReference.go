package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/waf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/waf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsManagedRulesAcfpRuleSet() AwsWebAclRule_AwsManagedRulesAcfpRuleSetPropertyList
	// Experimental.
	AwsManagedRulesAcfpRuleSetInput() interface{}
	// Experimental.
	AwsManagedRulesAntiDdosRuleSet() AwsWebAclRule_AwsManagedRulesAntiDdosRuleSetPropertyList
	// Experimental.
	AwsManagedRulesAntiDdosRuleSetInput() interface{}
	// Experimental.
	AwsManagedRulesAtpRuleSet() AwsWebAclRule_AwsManagedRulesAtpRuleSetPropertyList
	// Experimental.
	AwsManagedRulesAtpRuleSetInput() interface{}
	// Experimental.
	AwsManagedRulesBotControlRuleSet() AwsWebAclRule_AwsManagedRulesBotControlRuleSetPropertyList
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
	PasswordField() AwsWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsPasswordFieldPropertyList
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
	UsernameField() AwsWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsUsernameFieldPropertyList
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

// The jsii proxy struct for AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference
type jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAcfpRuleSet() AwsWebAclRule_AwsManagedRulesAcfpRuleSetPropertyList {
	var returns AwsWebAclRule_AwsManagedRulesAcfpRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesAcfpRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAcfpRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesAcfpRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAntiDdosRuleSet() AwsWebAclRule_AwsManagedRulesAntiDdosRuleSetPropertyList {
	var returns AwsWebAclRule_AwsManagedRulesAntiDdosRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesAntiDdosRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAntiDdosRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesAntiDdosRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAtpRuleSet() AwsWebAclRule_AwsManagedRulesAtpRuleSetPropertyList {
	var returns AwsWebAclRule_AwsManagedRulesAtpRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesAtpRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAtpRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesAtpRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesBotControlRuleSet() AwsWebAclRule_AwsManagedRulesBotControlRuleSetPropertyList {
	var returns AwsWebAclRule_AwsManagedRulesBotControlRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesBotControlRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesBotControlRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesBotControlRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) LoginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) LoginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PasswordField() AwsWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsPasswordFieldPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsPasswordFieldPropertyList
	_jsii_.Get(
		j,
		"passwordField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PasswordFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"passwordFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PayloadType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PayloadTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) UsernameField() AwsWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsUsernameFieldPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsUsernameFieldPropertyList
	_jsii_.Get(
		j,
		"usernameField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) UsernameFieldInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usernameFieldInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.ManagedRuleGroupConfigsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference_Override(a AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.ManagedRuleGroupConfigsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetLoginPath(val *string) {
	if err := j.validateSetLoginPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginPath",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetPayloadType(val *string) {
	if err := j.validateSetPayloadTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadType",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesAcfpRuleSet(value interface{}) {
	if err := a.validatePutAwsManagedRulesAcfpRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsManagedRulesAcfpRuleSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesAntiDdosRuleSet(value interface{}) {
	if err := a.validatePutAwsManagedRulesAntiDdosRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsManagedRulesAntiDdosRuleSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesAtpRuleSet(value interface{}) {
	if err := a.validatePutAwsManagedRulesAtpRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsManagedRulesAtpRuleSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesBotControlRuleSet(value interface{}) {
	if err := a.validatePutAwsManagedRulesBotControlRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsManagedRulesBotControlRuleSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutPasswordField(value interface{}) {
	if err := a.validatePutPasswordFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPasswordField",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) PutUsernameField(value interface{}) {
	if err := a.validatePutUsernameFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUsernameField",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesAcfpRuleSet() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsManagedRulesAcfpRuleSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesAntiDdosRuleSet() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsManagedRulesAntiDdosRuleSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesAtpRuleSet() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsManagedRulesAtpRuleSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesBotControlRuleSet() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsManagedRulesBotControlRuleSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetLoginPath() {
	_jsii_.InvokeVoid(
		a,
		"resetLoginPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetPasswordField() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetPayloadType() {
	_jsii_.InvokeVoid(
		a,
		"resetPayloadType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ResetUsernameField() {
	_jsii_.InvokeVoid(
		a,
		"resetUsernameField",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_ManagedRuleGroupConfigsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

