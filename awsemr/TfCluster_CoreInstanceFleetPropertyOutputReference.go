package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCluster_CoreInstanceFleetPropertyOutputReference interface {
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
	Id() *string
	// Experimental.
	InstanceTypeConfigs() TfCluster_CoreInstanceFleetInstanceTypeConfigsPropertyList
	// Experimental.
	InstanceTypeConfigsInput() interface{}
	// Experimental.
	InternalValue() *TfCluster_CoreInstanceFleetProperty
	// Experimental.
	SetInternalValue(val *TfCluster_CoreInstanceFleetProperty)
	// Experimental.
	LaunchSpecifications() TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference
	// Experimental.
	LaunchSpecificationsInput() *TfCluster_CoreInstanceFleetLaunchSpecificationsProperty
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	ProvisionedOnDemandCapacity() *float64
	// Experimental.
	ProvisionedSpotCapacity() *float64
	// Experimental.
	TargetOnDemandCapacity() *float64
	// Experimental.
	SetTargetOnDemandCapacity(val *float64)
	// Experimental.
	TargetOnDemandCapacityInput() *float64
	// Experimental.
	TargetSpotCapacity() *float64
	// Experimental.
	SetTargetSpotCapacity(val *float64)
	// Experimental.
	TargetSpotCapacityInput() *float64
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
	PutInstanceTypeConfigs(value interface{})
	// Experimental.
	PutLaunchSpecifications(value *TfCluster_CoreInstanceFleetLaunchSpecificationsProperty)
	// Experimental.
	ResetInstanceTypeConfigs()
	// Experimental.
	ResetLaunchSpecifications()
	// Experimental.
	ResetName()
	// Experimental.
	ResetTargetOnDemandCapacity()
	// Experimental.
	ResetTargetSpotCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCluster_CoreInstanceFleetPropertyOutputReference
type jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) InstanceTypeConfigs() TfCluster_CoreInstanceFleetInstanceTypeConfigsPropertyList {
	var returns TfCluster_CoreInstanceFleetInstanceTypeConfigsPropertyList
	_jsii_.Get(
		j,
		"instanceTypeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) InstanceTypeConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceTypeConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) InternalValue() *TfCluster_CoreInstanceFleetProperty {
	var returns *TfCluster_CoreInstanceFleetProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) LaunchSpecifications() TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference {
	var returns TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference
	_jsii_.Get(
		j,
		"launchSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) LaunchSpecificationsInput() *TfCluster_CoreInstanceFleetLaunchSpecificationsProperty {
	var returns *TfCluster_CoreInstanceFleetLaunchSpecificationsProperty
	_jsii_.Get(
		j,
		"launchSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) ProvisionedOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) ProvisionedSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) TargetOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) TargetOnDemandCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) TargetSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) TargetSpotCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCluster_CoreInstanceFleetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCluster_CoreInstanceFleetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCluster_CoreInstanceFleetPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.TfCluster.CoreInstanceFleetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCluster_CoreInstanceFleetPropertyOutputReference_Override(t TfCluster_CoreInstanceFleetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.TfCluster.CoreInstanceFleetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference)SetInternalValue(val *TfCluster_CoreInstanceFleetProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference)SetTargetOnDemandCapacity(val *float64) {
	if err := j.validateSetTargetOnDemandCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetOnDemandCapacity",
		val,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference)SetTargetSpotCapacity(val *float64) {
	if err := j.validateSetTargetSpotCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSpotCapacity",
		val,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) PutInstanceTypeConfigs(value interface{}) {
	if err := t.validatePutInstanceTypeConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceTypeConfigs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) PutLaunchSpecifications(value *TfCluster_CoreInstanceFleetLaunchSpecificationsProperty) {
	if err := t.validatePutLaunchSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLaunchSpecifications",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) ResetInstanceTypeConfigs() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceTypeConfigs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) ResetLaunchSpecifications() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchSpecifications",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) ResetTargetOnDemandCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetOnDemandCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) ResetTargetSpotCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetSpotCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

