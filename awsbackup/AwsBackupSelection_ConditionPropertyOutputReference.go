package awsbackup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbackup/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbackup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBackupSelection_ConditionPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	StringEquals() AwsBackupSelection_StringEqualsPropertyList
	// Experimental.
	StringEqualsInput() interface{}
	// Experimental.
	StringLike() AwsBackupSelection_StringLikePropertyList
	// Experimental.
	StringLikeInput() interface{}
	// Experimental.
	StringNotEquals() AwsBackupSelection_StringNotEqualsPropertyList
	// Experimental.
	StringNotEqualsInput() interface{}
	// Experimental.
	StringNotLike() AwsBackupSelection_StringNotLikePropertyList
	// Experimental.
	StringNotLikeInput() interface{}
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
	PutStringEquals(value interface{})
	// Experimental.
	PutStringLike(value interface{})
	// Experimental.
	PutStringNotEquals(value interface{})
	// Experimental.
	PutStringNotLike(value interface{})
	// Experimental.
	ResetStringEquals()
	// Experimental.
	ResetStringLike()
	// Experimental.
	ResetStringNotEquals()
	// Experimental.
	ResetStringNotLike()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBackupSelection_ConditionPropertyOutputReference
type jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) StringEquals() AwsBackupSelection_StringEqualsPropertyList {
	var returns AwsBackupSelection_StringEqualsPropertyList
	_jsii_.Get(
		j,
		"stringEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) StringEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) StringLike() AwsBackupSelection_StringLikePropertyList {
	var returns AwsBackupSelection_StringLikePropertyList
	_jsii_.Get(
		j,
		"stringLike",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) StringLikeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringLikeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) StringNotEquals() AwsBackupSelection_StringNotEqualsPropertyList {
	var returns AwsBackupSelection_StringNotEqualsPropertyList
	_jsii_.Get(
		j,
		"stringNotEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) StringNotEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) StringNotLike() AwsBackupSelection_StringNotLikePropertyList {
	var returns AwsBackupSelection_StringNotLikePropertyList
	_jsii_.Get(
		j,
		"stringNotLike",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) StringNotLikeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotLikeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBackupSelection_ConditionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBackupSelection_ConditionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBackupSelection_ConditionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-backup.AwsBackupSelection.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBackupSelection_ConditionPropertyOutputReference_Override(a AwsBackupSelection_ConditionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-backup.AwsBackupSelection.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) PutStringEquals(value interface{}) {
	if err := a.validatePutStringEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStringEquals",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) PutStringLike(value interface{}) {
	if err := a.validatePutStringLikeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStringLike",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) PutStringNotEquals(value interface{}) {
	if err := a.validatePutStringNotEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStringNotEquals",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) PutStringNotLike(value interface{}) {
	if err := a.validatePutStringNotLikeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStringNotLike",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) ResetStringEquals() {
	_jsii_.InvokeVoid(
		a,
		"resetStringEquals",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) ResetStringLike() {
	_jsii_.InvokeVoid(
		a,
		"resetStringLike",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) ResetStringNotEquals() {
	_jsii_.InvokeVoid(
		a,
		"resetStringNotEquals",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) ResetStringNotLike() {
	_jsii_.InvokeVoid(
		a,
		"resetStringNotLike",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBackupSelection_ConditionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

