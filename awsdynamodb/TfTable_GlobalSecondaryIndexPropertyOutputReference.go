package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdynamodb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdynamodb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTable_GlobalSecondaryIndexPropertyOutputReference interface {
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
	SetHashKey(val *string)
	// Experimental.
	HashKeyInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	KeySchema() TfTable_KeySchemaPropertyList
	// Experimental.
	KeySchemaInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NonKeyAttributes() *[]*string
	// Experimental.
	SetNonKeyAttributes(val *[]*string)
	// Experimental.
	NonKeyAttributesInput() *[]*string
	// Experimental.
	OnDemandThroughput() TfTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference
	// Experimental.
	OnDemandThroughputInput() *TfTable_GlobalSecondaryIndexOnDemandThroughputProperty
	// Experimental.
	ProjectionType() *string
	// Experimental.
	SetProjectionType(val *string)
	// Experimental.
	ProjectionTypeInput() *string
	// Experimental.
	RangeKey() *string
	// Experimental.
	SetRangeKey(val *string)
	// Experimental.
	RangeKeyInput() *string
	// Experimental.
	ReadCapacity() *float64
	// Experimental.
	SetReadCapacity(val *float64)
	// Experimental.
	ReadCapacityInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WarmThroughput() TfTable_GlobalSecondaryIndexWarmThroughputPropertyOutputReference
	// Experimental.
	WarmThroughputInput() *TfTable_GlobalSecondaryIndexWarmThroughputProperty
	// Experimental.
	WriteCapacity() *float64
	// Experimental.
	SetWriteCapacity(val *float64)
	// Experimental.
	WriteCapacityInput() *float64
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
	PutKeySchema(value interface{})
	// Experimental.
	PutOnDemandThroughput(value *TfTable_GlobalSecondaryIndexOnDemandThroughputProperty)
	// Experimental.
	PutWarmThroughput(value *TfTable_GlobalSecondaryIndexWarmThroughputProperty)
	// Experimental.
	ResetHashKey()
	// Experimental.
	ResetKeySchema()
	// Experimental.
	ResetNonKeyAttributes()
	// Experimental.
	ResetOnDemandThroughput()
	// Experimental.
	ResetRangeKey()
	// Experimental.
	ResetReadCapacity()
	// Experimental.
	ResetWarmThroughput()
	// Experimental.
	ResetWriteCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTable_GlobalSecondaryIndexPropertyOutputReference
type jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) HashKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) HashKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) KeySchema() TfTable_KeySchemaPropertyList {
	var returns TfTable_KeySchemaPropertyList
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) KeySchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) NonKeyAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonKeyAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) NonKeyAttributesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonKeyAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) OnDemandThroughput() TfTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference {
	var returns TfTable_GlobalSecondaryIndexOnDemandThroughputPropertyOutputReference
	_jsii_.Get(
		j,
		"onDemandThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) OnDemandThroughputInput() *TfTable_GlobalSecondaryIndexOnDemandThroughputProperty {
	var returns *TfTable_GlobalSecondaryIndexOnDemandThroughputProperty
	_jsii_.Get(
		j,
		"onDemandThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ProjectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ProjectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) RangeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) RangeKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ReadCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ReadCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) WarmThroughput() TfTable_GlobalSecondaryIndexWarmThroughputPropertyOutputReference {
	var returns TfTable_GlobalSecondaryIndexWarmThroughputPropertyOutputReference
	_jsii_.Get(
		j,
		"warmThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) WarmThroughputInput() *TfTable_GlobalSecondaryIndexWarmThroughputProperty {
	var returns *TfTable_GlobalSecondaryIndexWarmThroughputProperty
	_jsii_.Get(
		j,
		"warmThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) WriteCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) WriteCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeCapacityInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTable_GlobalSecondaryIndexPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfTable_GlobalSecondaryIndexPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTable_GlobalSecondaryIndexPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dynamodb.TfTable.GlobalSecondaryIndexPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTable_GlobalSecondaryIndexPropertyOutputReference_Override(t TfTable_GlobalSecondaryIndexPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dynamodb.TfTable.GlobalSecondaryIndexPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetHashKey(val *string) {
	if err := j.validateSetHashKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashKey",
		val,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetNonKeyAttributes(val *[]*string) {
	if err := j.validateSetNonKeyAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonKeyAttributes",
		val,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetProjectionType(val *string) {
	if err := j.validateSetProjectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectionType",
		val,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetRangeKey(val *string) {
	if err := j.validateSetRangeKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeKey",
		val,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetReadCapacity(val *float64) {
	if err := j.validateSetReadCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readCapacity",
		val,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference)SetWriteCapacity(val *float64) {
	if err := j.validateSetWriteCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeCapacity",
		val,
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) PutKeySchema(value interface{}) {
	if err := t.validatePutKeySchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKeySchema",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) PutOnDemandThroughput(value *TfTable_GlobalSecondaryIndexOnDemandThroughputProperty) {
	if err := t.validatePutOnDemandThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOnDemandThroughput",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) PutWarmThroughput(value *TfTable_GlobalSecondaryIndexWarmThroughputProperty) {
	if err := t.validatePutWarmThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWarmThroughput",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ResetHashKey() {
	_jsii_.InvokeVoid(
		t,
		"resetHashKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ResetKeySchema() {
	_jsii_.InvokeVoid(
		t,
		"resetKeySchema",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ResetNonKeyAttributes() {
	_jsii_.InvokeVoid(
		t,
		"resetNonKeyAttributes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ResetOnDemandThroughput() {
	_jsii_.InvokeVoid(
		t,
		"resetOnDemandThroughput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ResetRangeKey() {
	_jsii_.InvokeVoid(
		t,
		"resetRangeKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ResetReadCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetReadCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ResetWarmThroughput() {
	_jsii_.InvokeVoid(
		t,
		"resetWarmThroughput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ResetWriteCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetWriteCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTable_GlobalSecondaryIndexPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

