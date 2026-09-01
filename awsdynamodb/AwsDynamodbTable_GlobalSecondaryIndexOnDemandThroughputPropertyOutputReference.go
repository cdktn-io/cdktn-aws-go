package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdynamodb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdynamodb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference interface {
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
	InternalValue() *AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputProperty
	// Experimental.
	SetInternalValue(val *AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputProperty)
	// Experimental.
	MaxReadRequestUnits() *float64
	// Experimental.
	SetMaxReadRequestUnits(val *float64)
	// Experimental.
	MaxReadRequestUnitsInput() *float64
	// Experimental.
	MaxWriteRequestUnits() *float64
	// Experimental.
	SetMaxWriteRequestUnits(val *float64)
	// Experimental.
	MaxWriteRequestUnitsInput() *float64
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
	ResetMaxReadRequestUnits()
	// Experimental.
	ResetMaxWriteRequestUnits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference
type jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) InternalValue() *AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputProperty {
	var returns *AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) MaxReadRequestUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReadRequestUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) MaxReadRequestUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReadRequestUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) MaxWriteRequestUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWriteRequestUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) MaxWriteRequestUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWriteRequestUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dynamodb.AwsDynamodbTable.GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference_Override(a AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dynamodb.AwsDynamodbTable.GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference)SetInternalValue(val *AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference)SetMaxReadRequestUnits(val *float64) {
	if err := j.validateSetMaxReadRequestUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxReadRequestUnits",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference)SetMaxWriteRequestUnits(val *float64) {
	if err := j.validateSetMaxWriteRequestUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWriteRequestUnits",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) ResetMaxReadRequestUnits() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxReadRequestUnits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) ResetMaxWriteRequestUnits() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxWriteRequestUnits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDynamodbTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

