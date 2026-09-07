package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference interface {
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
	InternalValue() *AwsRoute_SpecGrpcRouteMatchMetadataMatchProperty
	// Experimental.
	SetInternalValue(val *AwsRoute_SpecGrpcRouteMatchMetadataMatchProperty)
	// Experimental.
	Prefix() *string
	// Experimental.
	SetPrefix(val *string)
	// Experimental.
	PrefixInput() *string
	// Experimental.
	Range() AwsRoute_SpecGrpcRouteMatchMetadataMatchRangePropertyOutputReference
	// Experimental.
	RangeInput() *AwsRoute_SpecGrpcRouteMatchMetadataMatchRangeProperty
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
	PutRange(value *AwsRoute_SpecGrpcRouteMatchMetadataMatchRangeProperty)
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

// The jsii proxy struct for AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference
type jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) Exact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) ExactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) InternalValue() *AwsRoute_SpecGrpcRouteMatchMetadataMatchProperty {
	var returns *AwsRoute_SpecGrpcRouteMatchMetadataMatchProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) Range() AwsRoute_SpecGrpcRouteMatchMetadataMatchRangePropertyOutputReference {
	var returns AwsRoute_SpecGrpcRouteMatchMetadataMatchRangePropertyOutputReference
	_jsii_.Get(
		j,
		"range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) RangeInput() *AwsRoute_SpecGrpcRouteMatchMetadataMatchRangeProperty {
	var returns *AwsRoute_SpecGrpcRouteMatchMetadataMatchRangeProperty
	_jsii_.Get(
		j,
		"rangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) RegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) Suffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) SuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsRoute.SpecGrpcRouteMatchMetadataMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference_Override(a AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsRoute.SpecGrpcRouteMatchMetadataMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference)SetExact(val *string) {
	if err := j.validateSetExactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exact",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference)SetInternalValue(val *AwsRoute_SpecGrpcRouteMatchMetadataMatchProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference)SetRegex(val *string) {
	if err := j.validateSetRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regex",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference)SetSuffix(val *string) {
	if err := j.validateSetSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suffix",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) PutRange(value *AwsRoute_SpecGrpcRouteMatchMetadataMatchRangeProperty) {
	if err := a.validatePutRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRange",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) ResetExact() {
	_jsii_.InvokeVoid(
		a,
		"resetExact",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) ResetRange() {
	_jsii_.InvokeVoid(
		a,
		"resetRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) ResetSuffix() {
	_jsii_.InvokeVoid(
		a,
		"resetSuffix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteMatchMetadataMatchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

