package awskeyspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskeyspaces/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskeyspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ClusteringKey() AwsKeyspacesTable_ClusteringKeyPropertyList
	// Experimental.
	ClusteringKeyInput() interface{}
	// Experimental.
	Column() AwsKeyspacesTable_ColumnPropertyList
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
	InternalValue() *AwsKeyspacesTable_SchemaDefinitionProperty
	// Experimental.
	SetInternalValue(val *AwsKeyspacesTable_SchemaDefinitionProperty)
	// Experimental.
	PartitionKey() AwsKeyspacesTable_PartitionKeyPropertyList
	// Experimental.
	PartitionKeyInput() interface{}
	// Experimental.
	StaticColumn() AwsKeyspacesTable_StaticColumnPropertyList
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

// The jsii proxy struct for AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference
type jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) ClusteringKey() AwsKeyspacesTable_ClusteringKeyPropertyList {
	var returns AwsKeyspacesTable_ClusteringKeyPropertyList
	_jsii_.Get(
		j,
		"clusteringKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) ClusteringKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusteringKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) Column() AwsKeyspacesTable_ColumnPropertyList {
	var returns AwsKeyspacesTable_ColumnPropertyList
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) ColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) InternalValue() *AwsKeyspacesTable_SchemaDefinitionProperty {
	var returns *AwsKeyspacesTable_SchemaDefinitionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) PartitionKey() AwsKeyspacesTable_PartitionKeyPropertyList {
	var returns AwsKeyspacesTable_PartitionKeyPropertyList
	_jsii_.Get(
		j,
		"partitionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) PartitionKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) StaticColumn() AwsKeyspacesTable_StaticColumnPropertyList {
	var returns AwsKeyspacesTable_StaticColumnPropertyList
	_jsii_.Get(
		j,
		"staticColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) StaticColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKeyspacesTable_SchemaDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKeyspacesTable_SchemaDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-keyspaces.AwsKeyspacesTable.SchemaDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKeyspacesTable_SchemaDefinitionPropertyOutputReference_Override(a AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-keyspaces.AwsKeyspacesTable.SchemaDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference)SetInternalValue(val *AwsKeyspacesTable_SchemaDefinitionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) PutClusteringKey(value interface{}) {
	if err := a.validatePutClusteringKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClusteringKey",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) PutColumn(value interface{}) {
	if err := a.validatePutColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putColumn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) PutPartitionKey(value interface{}) {
	if err := a.validatePutPartitionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPartitionKey",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) PutStaticColumn(value interface{}) {
	if err := a.validatePutStaticColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStaticColumn",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) ResetClusteringKey() {
	_jsii_.InvokeVoid(
		a,
		"resetClusteringKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) ResetStaticColumn() {
	_jsii_.InvokeVoid(
		a,
		"resetStaticColumn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

