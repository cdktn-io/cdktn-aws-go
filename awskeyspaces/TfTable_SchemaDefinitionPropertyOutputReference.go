package awskeyspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskeyspaces/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskeyspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTable_SchemaDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ClusteringKey() TfTable_ClusteringKeyPropertyList
	// Experimental.
	ClusteringKeyInput() interface{}
	// Experimental.
	Column() TfTable_ColumnPropertyList
	// Experimental.
	ColumnInput() interface{}
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
	InternalValue() *TfTable_SchemaDefinitionProperty
	// Experimental.
	SetInternalValue(val *TfTable_SchemaDefinitionProperty)
	// Experimental.
	PartitionKey() TfTable_PartitionKeyPropertyList
	// Experimental.
	PartitionKeyInput() interface{}
	// Experimental.
	StaticColumn() TfTable_StaticColumnPropertyList
	// Experimental.
	StaticColumnInput() interface{}
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
	PutClusteringKey(value interface{})
	// Experimental.
	PutColumn(value interface{})
	// Experimental.
	PutPartitionKey(value interface{})
	// Experimental.
	PutStaticColumn(value interface{})
	// Experimental.
	ResetClusteringKey()
	// Experimental.
	ResetStaticColumn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTable_SchemaDefinitionPropertyOutputReference
type jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) ClusteringKey() TfTable_ClusteringKeyPropertyList {
	var returns TfTable_ClusteringKeyPropertyList
	_jsii_.Get(
		j,
		"clusteringKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) ClusteringKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusteringKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) Column() TfTable_ColumnPropertyList {
	var returns TfTable_ColumnPropertyList
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) ColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) InternalValue() *TfTable_SchemaDefinitionProperty {
	var returns *TfTable_SchemaDefinitionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) PartitionKey() TfTable_PartitionKeyPropertyList {
	var returns TfTable_PartitionKeyPropertyList
	_jsii_.Get(
		j,
		"partitionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) PartitionKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) StaticColumn() TfTable_StaticColumnPropertyList {
	var returns TfTable_StaticColumnPropertyList
	_jsii_.Get(
		j,
		"staticColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) StaticColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTable_SchemaDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTable_SchemaDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTable_SchemaDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-keyspaces.TfTable.SchemaDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTable_SchemaDefinitionPropertyOutputReference_Override(t TfTable_SchemaDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-keyspaces.TfTable.SchemaDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference)SetInternalValue(val *TfTable_SchemaDefinitionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) PutClusteringKey(value interface{}) {
	if err := t.validatePutClusteringKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putClusteringKey",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) PutColumn(value interface{}) {
	if err := t.validatePutColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putColumn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) PutPartitionKey(value interface{}) {
	if err := t.validatePutPartitionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPartitionKey",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) PutStaticColumn(value interface{}) {
	if err := t.validatePutStaticColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStaticColumn",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) ResetClusteringKey() {
	_jsii_.InvokeVoid(
		t,
		"resetClusteringKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) ResetStaticColumn() {
	_jsii_.InvokeVoid(
		t,
		"resetStaticColumn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTable_SchemaDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

