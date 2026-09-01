package awsoracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AvailableCoreCount() *float64
	// Experimental.
	AvailableCoreCountPerNode() *float64
	// Experimental.
	AvailableDataStorageInTbs() *float64
	// Experimental.
	AvailableDataStoragePerServerInTbs() *float64
	// Experimental.
	AvailableDbNodePerNodeInGbs() *float64
	// Experimental.
	AvailableDbNodeStorageInGbs() *float64
	// Experimental.
	AvailableMemoryInGbs() *float64
	// Experimental.
	AvailableMemoryPerNodeInGbs() *float64
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
	// Experimental.
	CoreCountIncrement() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *DataAwsOdbDbSystemShapes_DbSystemShapesProperty
	// Experimental.
	SetInternalValue(val *DataAwsOdbDbSystemShapes_DbSystemShapesProperty)
	// Experimental.
	MaximumNodeCount() *float64
	// Experimental.
	MaxStorageCount() *float64
	// Experimental.
	MinCoreCountPerNode() *float64
	// Experimental.
	MinDataStorageInTbs() *float64
	// Experimental.
	MinDbNodeStoragePerNodeInGbs() *float64
	// Experimental.
	MinimumCoreCount() *float64
	// Experimental.
	MinimumNodeCount() *float64
	// Experimental.
	MinMemoryPerNodeInGbs() *float64
	// Experimental.
	MinStorageCount() *float64
	// Experimental.
	Name() *string
	// Experimental.
	RuntimeMinimumCoreCount() *float64
	// Experimental.
	ShapeFamily() *string
	// Experimental.
	ShapeType() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference
type jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) AvailableCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) AvailableCoreCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableCoreCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) AvailableDataStorageInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDataStorageInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) AvailableDataStoragePerServerInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDataStoragePerServerInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) AvailableDbNodePerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDbNodePerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) AvailableDbNodeStorageInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableDbNodeStorageInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) AvailableMemoryInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) AvailableMemoryPerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryPerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) CoreCountIncrement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coreCountIncrement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) InternalValue() *DataAwsOdbDbSystemShapes_DbSystemShapesProperty {
	var returns *DataAwsOdbDbSystemShapes_DbSystemShapesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) MaximumNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) MaxStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) MinCoreCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCoreCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) MinDataStorageInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDataStorageInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) MinDbNodeStoragePerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDbNodeStoragePerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) MinimumCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) MinimumNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) MinMemoryPerNodeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minMemoryPerNodeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) MinStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) RuntimeMinimumCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runtimeMinimumCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) ShapeFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) ShapeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.DataAwsOdbDbSystemShapes.DbSystemShapesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference_Override(d DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.DataAwsOdbDbSystemShapes.DbSystemShapesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference)SetInternalValue(val *DataAwsOdbDbSystemShapes_DbSystemShapesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsOdbDbSystemShapes_DbSystemShapesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

