package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdynamodb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdynamodb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfTable_GlobalSecondaryIndexPropertyOutputReference interface {
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
	HashKey() *string
	// Experimental.
	InternalValue() *DataTfTable_GlobalSecondaryIndexProperty
	// Experimental.
	SetInternalValue(val *DataTfTable_GlobalSecondaryIndexProperty)
	// Experimental.
	KeySchema() DataTfTable_KeySchemaPropertyList
	// Experimental.
	Name() *string
	// Experimental.
	NonKeyAttributes() *[]*string
	// Experimental.
	OnDemandThroughput() DataTfTable_GlobalSecondaryIndexOnDemandThroughputPropertyList
	// Experimental.
	ProjectionType() *string
	// Experimental.
	RangeKey() *string
	// Experimental.
	ReadCapacity() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WarmThroughput() DataTfTable_GlobalSecondaryIndexWarmThroughputPropertyList
	// Experimental.
	WriteCapacity() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfTable_GlobalSecondaryIndexPropertyOutputReference
type jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) HashKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) InternalValue() *DataTfTable_GlobalSecondaryIndexProperty {
	var returns *DataTfTable_GlobalSecondaryIndexProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) KeySchema() DataTfTable_KeySchemaPropertyList {
	var returns DataTfTable_KeySchemaPropertyList
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) NonKeyAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonKeyAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) OnDemandThroughput() DataTfTable_GlobalSecondaryIndexOnDemandThroughputPropertyList {
	var returns DataTfTable_GlobalSecondaryIndexOnDemandThroughputPropertyList
	_jsii_.Get(
		j,
		"onDemandThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) ProjectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) RangeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) ReadCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) WarmThroughput() DataTfTable_GlobalSecondaryIndexWarmThroughputPropertyList {
	var returns DataTfTable_GlobalSecondaryIndexWarmThroughputPropertyList
	_jsii_.Get(
		j,
		"warmThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) WriteCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeCapacity",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfTable_GlobalSecondaryIndexPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfTable_GlobalSecondaryIndexPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfTable_GlobalSecondaryIndexPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dynamodb.DataTfTable.GlobalSecondaryIndexPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfTable_GlobalSecondaryIndexPropertyOutputReference_Override(d DataTfTable_GlobalSecondaryIndexPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dynamodb.DataTfTable.GlobalSecondaryIndexPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference)SetInternalValue(val *DataTfTable_GlobalSecondaryIndexProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTable_GlobalSecondaryIndexPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

