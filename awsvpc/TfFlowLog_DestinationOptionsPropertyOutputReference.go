package awsvpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlowLog_DestinationOptionsPropertyOutputReference interface {
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
	FileFormat() *string
	// Experimental.
	SetFileFormat(val *string)
	// Experimental.
	FileFormatInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	HiveCompatiblePartitions() interface{}
	// Experimental.
	SetHiveCompatiblePartitions(val interface{})
	// Experimental.
	HiveCompatiblePartitionsInput() interface{}
	// Experimental.
	InternalValue() *TfFlowLog_DestinationOptionsProperty
	// Experimental.
	SetInternalValue(val *TfFlowLog_DestinationOptionsProperty)
	// Experimental.
	PerHourPartition() interface{}
	// Experimental.
	SetPerHourPartition(val interface{})
	// Experimental.
	PerHourPartitionInput() interface{}
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
	ResetFileFormat()
	// Experimental.
	ResetHiveCompatiblePartitions()
	// Experimental.
	ResetPerHourPartition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlowLog_DestinationOptionsPropertyOutputReference
type jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) FileFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) FileFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) HiveCompatiblePartitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hiveCompatiblePartitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) HiveCompatiblePartitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hiveCompatiblePartitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) InternalValue() *TfFlowLog_DestinationOptionsProperty {
	var returns *TfFlowLog_DestinationOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) PerHourPartition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"perHourPartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) PerHourPartitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"perHourPartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlowLog_DestinationOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlowLog_DestinationOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlowLog_DestinationOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.TfFlowLog.DestinationOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlowLog_DestinationOptionsPropertyOutputReference_Override(t TfFlowLog_DestinationOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.TfFlowLog.DestinationOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference)SetFileFormat(val *string) {
	if err := j.validateSetFileFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileFormat",
		val,
	)
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference)SetHiveCompatiblePartitions(val interface{}) {
	if err := j.validateSetHiveCompatiblePartitionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hiveCompatiblePartitions",
		val,
	)
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference)SetInternalValue(val *TfFlowLog_DestinationOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference)SetPerHourPartition(val interface{}) {
	if err := j.validateSetPerHourPartitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perHourPartition",
		val,
	)
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) ResetFileFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetFileFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) ResetHiveCompatiblePartitions() {
	_jsii_.InvokeVoid(
		t,
		"resetHiveCompatiblePartitions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) ResetPerHourPartition() {
	_jsii_.InvokeVoid(
		t,
		"resetPerHourPartition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlowLog_DestinationOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

