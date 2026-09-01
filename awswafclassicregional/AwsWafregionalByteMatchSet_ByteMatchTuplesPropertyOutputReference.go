package awswafclassicregional

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswafclassicregional/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswafclassicregional/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference interface {
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
	FieldToMatch() AwsWafregionalByteMatchSet_FieldToMatchPropertyOutputReference
	// Experimental.
	FieldToMatchInput() *AwsWafregionalByteMatchSet_FieldToMatchProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PositionalConstraint() *string
	// Experimental.
	SetPositionalConstraint(val *string)
	// Experimental.
	PositionalConstraintInput() *string
	// Experimental.
	TargetString() *string
	// Experimental.
	SetTargetString(val *string)
	// Experimental.
	TargetStringInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TextTransformation() *string
	// Experimental.
	SetTextTransformation(val *string)
	// Experimental.
	TextTransformationInput() *string
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
	PutFieldToMatch(value *AwsWafregionalByteMatchSet_FieldToMatchProperty)
	// Experimental.
	ResetTargetString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference
type jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) FieldToMatch() AwsWafregionalByteMatchSet_FieldToMatchPropertyOutputReference {
	var returns AwsWafregionalByteMatchSet_FieldToMatchPropertyOutputReference
	_jsii_.Get(
		j,
		"fieldToMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) FieldToMatchInput() *AwsWafregionalByteMatchSet_FieldToMatchProperty {
	var returns *AwsWafregionalByteMatchSet_FieldToMatchProperty
	_jsii_.Get(
		j,
		"fieldToMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) PositionalConstraint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"positionalConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) PositionalConstraintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"positionalConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) TargetString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) TargetStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) TextTransformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textTransformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) TextTransformationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textTransformationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf-classic-regional.AwsWafregionalByteMatchSet.ByteMatchTuplesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference_Override(a AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf-classic-regional.AwsWafregionalByteMatchSet.ByteMatchTuplesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetPositionalConstraint(val *string) {
	if err := j.validateSetPositionalConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"positionalConstraint",
		val,
	)
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetTargetString(val *string) {
	if err := j.validateSetTargetStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetString",
		val,
	)
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetTextTransformation(val *string) {
	if err := j.validateSetTextTransformationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textTransformation",
		val,
	)
}

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) PutFieldToMatch(value *AwsWafregionalByteMatchSet_FieldToMatchProperty) {
	if err := a.validatePutFieldToMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFieldToMatch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) ResetTargetString() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWafregionalByteMatchSet_ByteMatchTuplesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

