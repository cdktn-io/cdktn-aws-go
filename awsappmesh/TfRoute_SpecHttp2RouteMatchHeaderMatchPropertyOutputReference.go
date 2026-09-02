package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference interface {
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
	Exact() *string
	// Experimental.
	SetExact(val *string)
	// Experimental.
	ExactInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfRoute_SpecHttp2RouteMatchHeaderMatchProperty
	// Experimental.
	SetInternalValue(val *TfRoute_SpecHttp2RouteMatchHeaderMatchProperty)
	// Experimental.
	Prefix() *string
	// Experimental.
	SetPrefix(val *string)
	// Experimental.
	PrefixInput() *string
	// Experimental.
	Range() TfRoute_SpecHttp2RouteMatchHeaderMatchRangePropertyOutputReference
	// Experimental.
	RangeInput() *TfRoute_SpecHttp2RouteMatchHeaderMatchRangeProperty
	// Experimental.
	Regex() *string
	// Experimental.
	SetRegex(val *string)
	// Experimental.
	RegexInput() *string
	// Experimental.
	Suffix() *string
	// Experimental.
	SetSuffix(val *string)
	// Experimental.
	SuffixInput() *string
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
	PutRange(value *TfRoute_SpecHttp2RouteMatchHeaderMatchRangeProperty)
	// Experimental.
	ResetExact()
	// Experimental.
	ResetPrefix()
	// Experimental.
	ResetRange()
	// Experimental.
	ResetRegex()
	// Experimental.
	ResetSuffix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference
type jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) Exact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) ExactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) InternalValue() *TfRoute_SpecHttp2RouteMatchHeaderMatchProperty {
	var returns *TfRoute_SpecHttp2RouteMatchHeaderMatchProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) Range() TfRoute_SpecHttp2RouteMatchHeaderMatchRangePropertyOutputReference {
	var returns TfRoute_SpecHttp2RouteMatchHeaderMatchRangePropertyOutputReference
	_jsii_.Get(
		j,
		"range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) RangeInput() *TfRoute_SpecHttp2RouteMatchHeaderMatchRangeProperty {
	var returns *TfRoute_SpecHttp2RouteMatchHeaderMatchRangeProperty
	_jsii_.Get(
		j,
		"rangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) RegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) Suffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) SuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfRoute.SpecHttp2RouteMatchHeaderMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference_Override(t TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfRoute.SpecHttp2RouteMatchHeaderMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference)SetExact(val *string) {
	if err := j.validateSetExactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exact",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference)SetInternalValue(val *TfRoute_SpecHttp2RouteMatchHeaderMatchProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference)SetRegex(val *string) {
	if err := j.validateSetRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regex",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference)SetSuffix(val *string) {
	if err := j.validateSetSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suffix",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) PutRange(value *TfRoute_SpecHttp2RouteMatchHeaderMatchRangeProperty) {
	if err := t.validatePutRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRange",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) ResetExact() {
	_jsii_.InvokeVoid(
		t,
		"resetExact",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) ResetRange() {
	_jsii_.InvokeVoid(
		t,
		"resetRange",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		t,
		"resetRegex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) ResetSuffix() {
	_jsii_.InvokeVoid(
		t,
		"resetSuffix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchHeaderMatchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

