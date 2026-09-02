package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpclattice/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpclattice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfListenerRule_HttpMatchPropertyOutputReference interface {
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
	HeaderMatches() TfListenerRule_HeaderMatchesPropertyList
	// Experimental.
	HeaderMatchesInput() interface{}
	// Experimental.
	InternalValue() *TfListenerRule_HttpMatchProperty
	// Experimental.
	SetInternalValue(val *TfListenerRule_HttpMatchProperty)
	// Experimental.
	Method() *string
	// Experimental.
	SetMethod(val *string)
	// Experimental.
	MethodInput() *string
	// Experimental.
	PathMatch() TfListenerRule_PathMatchPropertyOutputReference
	// Experimental.
	PathMatchInput() *TfListenerRule_PathMatchProperty
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
	PutHeaderMatches(value interface{})
	// Experimental.
	PutPathMatch(value *TfListenerRule_PathMatchProperty)
	// Experimental.
	ResetHeaderMatches()
	// Experimental.
	ResetMethod()
	// Experimental.
	ResetPathMatch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfListenerRule_HttpMatchPropertyOutputReference
type jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) HeaderMatches() TfListenerRule_HeaderMatchesPropertyList {
	var returns TfListenerRule_HeaderMatchesPropertyList
	_jsii_.Get(
		j,
		"headerMatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) HeaderMatchesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerMatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) InternalValue() *TfListenerRule_HttpMatchProperty {
	var returns *TfListenerRule_HttpMatchProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) PathMatch() TfListenerRule_PathMatchPropertyOutputReference {
	var returns TfListenerRule_PathMatchPropertyOutputReference
	_jsii_.Get(
		j,
		"pathMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) PathMatchInput() *TfListenerRule_PathMatchProperty {
	var returns *TfListenerRule_PathMatchProperty
	_jsii_.Get(
		j,
		"pathMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfListenerRule_HttpMatchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfListenerRule_HttpMatchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfListenerRule_HttpMatchPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc-lattice.TfListenerRule.HttpMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfListenerRule_HttpMatchPropertyOutputReference_Override(t TfListenerRule_HttpMatchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc-lattice.TfListenerRule.HttpMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference)SetInternalValue(val *TfListenerRule_HttpMatchProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) PutHeaderMatches(value interface{}) {
	if err := t.validatePutHeaderMatchesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeaderMatches",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) PutPathMatch(value *TfListenerRule_PathMatchProperty) {
	if err := t.validatePutPathMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPathMatch",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) ResetHeaderMatches() {
	_jsii_.InvokeVoid(
		t,
		"resetHeaderMatches",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) ResetPathMatch() {
	_jsii_.InvokeVoid(
		t,
		"resetPathMatch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfListenerRule_HttpMatchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

