package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference interface {
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
	CustomResponseBodyKey() *string
	// Experimental.
	SetCustomResponseBodyKey(val *string)
	// Experimental.
	CustomResponseBodyKeyInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ResponseCode() *float64
	// Experimental.
	SetResponseCode(val *float64)
	// Experimental.
	ResponseCodeInput() *float64
	// Experimental.
	ResponseHeader() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponseResponseHeaderPropertyList
	// Experimental.
	ResponseHeaderInput() interface{}
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
	PutResponseHeader(value interface{})
	// Experimental.
	ResetCustomResponseBodyKey()
	// Experimental.
	ResetResponseHeader()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference
type jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) CustomResponseBodyKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customResponseBodyKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) CustomResponseBodyKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customResponseBodyKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) ResponseCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) ResponseCodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) ResponseHeader() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponseResponseHeaderPropertyList {
	var returns TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponseResponseHeaderPropertyList
	_jsii_.Get(
		j,
		"responseHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) ResponseHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRuleGroupAssociation.ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference_Override(t TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRuleGroupAssociation.ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference)SetCustomResponseBodyKey(val *string) {
	if err := j.validateSetCustomResponseBodyKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customResponseBodyKey",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference)SetResponseCode(val *float64) {
	if err := j.validateSetResponseCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseCode",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) PutResponseHeader(value interface{}) {
	if err := t.validatePutResponseHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResponseHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) ResetCustomResponseBodyKey() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomResponseBodyKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) ResetResponseHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetResponseHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

