package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/waf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/waf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AwsManagedRulesAcfpRuleSet() AwsWebAclRuleGroupAssociation_AwsManagedRulesAcfpRuleSetPropertyList
	// Experimental.
	AwsManagedRulesAcfpRuleSetInput() interface{}
	// Experimental.
	AwsManagedRulesAntiDdosRuleSet() AwsWebAclRuleGroupAssociation_AwsManagedRulesAntiDdosRuleSetPropertyList
	// Experimental.
	AwsManagedRulesAntiDdosRuleSetInput() interface{}
	// Experimental.
	AwsManagedRulesAtpRuleSet() AwsWebAclRuleGroupAssociation_AwsManagedRulesAtpRuleSetPropertyList
	// Experimental.
	AwsManagedRulesAtpRuleSetInput() interface{}
	// Experimental.
	AwsManagedRulesBotControlRuleSet() AwsWebAclRuleGroupAssociation_AwsManagedRulesBotControlRuleSetPropertyList
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
	PutAwsManagedRulesAcfpRuleSet(value interface{})
	// Experimental.
	PutAwsManagedRulesAntiDdosRuleSet(value interface{})
	// Experimental.
	PutAwsManagedRulesAtpRuleSet(value interface{})
	// Experimental.
	PutAwsManagedRulesBotControlRuleSet(value interface{})
	// Experimental.
	ResetAwsManagedRulesAcfpRuleSet()
	// Experimental.
	ResetAwsManagedRulesAntiDdosRuleSet()
	// Experimental.
	ResetAwsManagedRulesAtpRuleSet()
	// Experimental.
	ResetAwsManagedRulesBotControlRuleSet()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference
type jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAcfpRuleSet() AwsWebAclRuleGroupAssociation_AwsManagedRulesAcfpRuleSetPropertyList {
	var returns AwsWebAclRuleGroupAssociation_AwsManagedRulesAcfpRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesAcfpRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAcfpRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesAcfpRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAntiDdosRuleSet() AwsWebAclRuleGroupAssociation_AwsManagedRulesAntiDdosRuleSetPropertyList {
	var returns AwsWebAclRuleGroupAssociation_AwsManagedRulesAntiDdosRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesAntiDdosRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAntiDdosRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesAntiDdosRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAtpRuleSet() AwsWebAclRuleGroupAssociation_AwsManagedRulesAtpRuleSetPropertyList {
	var returns AwsWebAclRuleGroupAssociation_AwsManagedRulesAtpRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesAtpRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesAtpRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesAtpRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesBotControlRuleSet() AwsWebAclRuleGroupAssociation_AwsManagedRulesBotControlRuleSetPropertyList {
	var returns AwsWebAclRuleGroupAssociation_AwsManagedRulesBotControlRuleSetPropertyList
	_jsii_.Get(
		j,
		"awsManagedRulesBotControlRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) AwsManagedRulesBotControlRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsManagedRulesBotControlRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRuleGroupAssociation.ManagedRuleGroupConfigsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference_Override(a AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRuleGroupAssociation.ManagedRuleGroupConfigsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesAcfpRuleSet(value interface{}) {
	if err := a.validatePutAwsManagedRulesAcfpRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsManagedRulesAcfpRuleSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesAntiDdosRuleSet(value interface{}) {
	if err := a.validatePutAwsManagedRulesAntiDdosRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsManagedRulesAntiDdosRuleSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesAtpRuleSet(value interface{}) {
	if err := a.validatePutAwsManagedRulesAtpRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsManagedRulesAtpRuleSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) PutAwsManagedRulesBotControlRuleSet(value interface{}) {
	if err := a.validatePutAwsManagedRulesBotControlRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsManagedRulesBotControlRuleSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesAcfpRuleSet() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsManagedRulesAcfpRuleSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesAntiDdosRuleSet() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsManagedRulesAntiDdosRuleSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesAtpRuleSet() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsManagedRulesAtpRuleSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) ResetAwsManagedRulesBotControlRuleSet() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsManagedRulesBotControlRuleSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_ManagedRuleGroupConfigsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

