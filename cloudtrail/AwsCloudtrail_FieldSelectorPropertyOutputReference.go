package cloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudtrail/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudtrail/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudtrail_FieldSelectorPropertyOutputReference interface {
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
	EndsWith() *[]*string
	// Experimental.
	SetEndsWith(val *[]*string)
	// Experimental.
	EndsWithInput() *[]*string
	// Experimental.
	EqualTo() *[]*string
	// Experimental.
	SetEqualTo(val *[]*string)
	// Experimental.
	EqualToInput() *[]*string
	// Experimental.
	Field() *string
	// Experimental.
	SetField(val *string)
	// Experimental.
	FieldInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	NotEndsWith() *[]*string
	// Experimental.
	SetNotEndsWith(val *[]*string)
	// Experimental.
	NotEndsWithInput() *[]*string
	// Experimental.
	NotEquals() *[]*string
	// Experimental.
	SetNotEquals(val *[]*string)
	// Experimental.
	NotEqualsInput() *[]*string
	// Experimental.
	NotStartsWith() *[]*string
	// Experimental.
	SetNotStartsWith(val *[]*string)
	// Experimental.
	NotStartsWithInput() *[]*string
	// Experimental.
	StartsWith() *[]*string
	// Experimental.
	SetStartsWith(val *[]*string)
	// Experimental.
	StartsWithInput() *[]*string
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
	ResetEndsWith()
	// Experimental.
	ResetEqualTo()
	// Experimental.
	ResetNotEndsWith()
	// Experimental.
	ResetNotEquals()
	// Experimental.
	ResetNotStartsWith()
	// Experimental.
	ResetStartsWith()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCloudtrail_FieldSelectorPropertyOutputReference
type jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) EndsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) EndsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) EqualTo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) EqualToInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) NotEndsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) NotEndsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) NotEquals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) NotEqualsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) NotStartsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notStartsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) NotStartsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notStartsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) StartsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) StartsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCloudtrail_FieldSelectorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCloudtrail_FieldSelectorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCloudtrail_FieldSelectorPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudtrail.AwsCloudtrail.FieldSelectorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCloudtrail_FieldSelectorPropertyOutputReference_Override(a AwsCloudtrail_FieldSelectorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudtrail.AwsCloudtrail.FieldSelectorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetEndsWith(val *[]*string) {
	if err := j.validateSetEndsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endsWith",
		val,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetEqualTo(val *[]*string) {
	if err := j.validateSetEqualToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"equalTo",
		val,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetNotEndsWith(val *[]*string) {
	if err := j.validateSetNotEndsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notEndsWith",
		val,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetNotEquals(val *[]*string) {
	if err := j.validateSetNotEqualsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notEquals",
		val,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetNotStartsWith(val *[]*string) {
	if err := j.validateSetNotStartsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notStartsWith",
		val,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetStartsWith(val *[]*string) {
	if err := j.validateSetStartsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startsWith",
		val,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) ResetEndsWith() {
	_jsii_.InvokeVoid(
		a,
		"resetEndsWith",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) ResetEqualTo() {
	_jsii_.InvokeVoid(
		a,
		"resetEqualTo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) ResetNotEndsWith() {
	_jsii_.InvokeVoid(
		a,
		"resetNotEndsWith",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) ResetNotEquals() {
	_jsii_.InvokeVoid(
		a,
		"resetNotEquals",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) ResetNotStartsWith() {
	_jsii_.InvokeVoid(
		a,
		"resetNotStartsWith",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) ResetStartsWith() {
	_jsii_.InvokeVoid(
		a,
		"resetStartsWith",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCloudtrail_FieldSelectorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

