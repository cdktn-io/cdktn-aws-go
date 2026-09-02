package awsfms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPolicy_NetworkAclEntrySetPropertyOutputReference interface {
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
	FirstEntry() TfPolicy_FirstEntryPropertyList
	// Experimental.
	FirstEntryInput() interface{}
	// Experimental.
	ForceRemediateForFirstEntries() interface{}
	// Experimental.
	SetForceRemediateForFirstEntries(val interface{})
	// Experimental.
	ForceRemediateForFirstEntriesInput() interface{}
	// Experimental.
	ForceRemediateForLastEntries() interface{}
	// Experimental.
	SetForceRemediateForLastEntries(val interface{})
	// Experimental.
	ForceRemediateForLastEntriesInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfPolicy_NetworkAclEntrySetProperty
	// Experimental.
	SetInternalValue(val *TfPolicy_NetworkAclEntrySetProperty)
	// Experimental.
	LastEntry() TfPolicy_LastEntryPropertyList
	// Experimental.
	LastEntryInput() interface{}
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
	PutFirstEntry(value interface{})
	// Experimental.
	PutLastEntry(value interface{})
	// Experimental.
	ResetFirstEntry()
	// Experimental.
	ResetLastEntry()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPolicy_NetworkAclEntrySetPropertyOutputReference
type jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) FirstEntry() TfPolicy_FirstEntryPropertyList {
	var returns TfPolicy_FirstEntryPropertyList
	_jsii_.Get(
		j,
		"firstEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) FirstEntryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstEntryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) ForceRemediateForFirstEntries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRemediateForFirstEntries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) ForceRemediateForFirstEntriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRemediateForFirstEntriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) ForceRemediateForLastEntries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRemediateForLastEntries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) ForceRemediateForLastEntriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRemediateForLastEntriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) InternalValue() *TfPolicy_NetworkAclEntrySetProperty {
	var returns *TfPolicy_NetworkAclEntrySetProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) LastEntry() TfPolicy_LastEntryPropertyList {
	var returns TfPolicy_LastEntryPropertyList
	_jsii_.Get(
		j,
		"lastEntry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) LastEntryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastEntryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPolicy_NetworkAclEntrySetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPolicy_NetworkAclEntrySetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPolicy_NetworkAclEntrySetPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fms.TfPolicy.NetworkAclEntrySetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPolicy_NetworkAclEntrySetPropertyOutputReference_Override(t TfPolicy_NetworkAclEntrySetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fms.TfPolicy.NetworkAclEntrySetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference)SetForceRemediateForFirstEntries(val interface{}) {
	if err := j.validateSetForceRemediateForFirstEntriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceRemediateForFirstEntries",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference)SetForceRemediateForLastEntries(val interface{}) {
	if err := j.validateSetForceRemediateForLastEntriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceRemediateForLastEntries",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference)SetInternalValue(val *TfPolicy_NetworkAclEntrySetProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) PutFirstEntry(value interface{}) {
	if err := t.validatePutFirstEntryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFirstEntry",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) PutLastEntry(value interface{}) {
	if err := t.validatePutLastEntryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLastEntry",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) ResetFirstEntry() {
	_jsii_.InvokeVoid(
		t,
		"resetFirstEntry",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) ResetLastEntry() {
	_jsii_.InvokeVoid(
		t,
		"resetLastEntry",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPolicy_NetworkAclEntrySetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

