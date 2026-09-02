package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFileCache_LustreConfigurationPropertyOutputReference interface {
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
	DeploymentType() *string
	// Experimental.
	SetDeploymentType(val *string)
	// Experimental.
	DeploymentTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LogConfiguration() TfFileCache_LogConfigurationPropertyList
	// Experimental.
	MetadataConfiguration() TfFileCache_MetadataConfigurationPropertyList
	// Experimental.
	MetadataConfigurationInput() interface{}
	// Experimental.
	MountName() *string
	// Experimental.
	PerUnitStorageThroughput() *float64
	// Experimental.
	SetPerUnitStorageThroughput(val *float64)
	// Experimental.
	PerUnitStorageThroughputInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WeeklyMaintenanceStartTime() *string
	// Experimental.
	SetWeeklyMaintenanceStartTime(val *string)
	// Experimental.
	WeeklyMaintenanceStartTimeInput() *string
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
	PutMetadataConfiguration(value interface{})
	// Experimental.
	ResetWeeklyMaintenanceStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFileCache_LustreConfigurationPropertyOutputReference
type jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) LogConfiguration() TfFileCache_LogConfigurationPropertyList {
	var returns TfFileCache_LogConfigurationPropertyList
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) MetadataConfiguration() TfFileCache_MetadataConfigurationPropertyList {
	var returns TfFileCache_MetadataConfigurationPropertyList
	_jsii_.Get(
		j,
		"metadataConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) MetadataConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) MountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) PerUnitStorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) PerUnitStorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) WeeklyMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) WeeklyMaintenanceStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTimeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFileCache_LustreConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfFileCache_LustreConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFileCache_LustreConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfFileCache.LustreConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFileCache_LustreConfigurationPropertyOutputReference_Override(t TfFileCache_LustreConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfFileCache.LustreConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference)SetDeploymentType(val *string) {
	if err := j.validateSetDeploymentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference)SetPerUnitStorageThroughput(val *float64) {
	if err := j.validateSetPerUnitStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perUnitStorageThroughput",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference)SetWeeklyMaintenanceStartTime(val *string) {
	if err := j.validateSetWeeklyMaintenanceStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyMaintenanceStartTime",
		val,
	)
}

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) PutMetadataConfiguration(value interface{}) {
	if err := t.validatePutMetadataConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetadataConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) ResetWeeklyMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		t,
		"resetWeeklyMaintenanceStartTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFileCache_LustreConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

