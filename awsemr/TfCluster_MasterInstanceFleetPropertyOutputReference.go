package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCluster_MasterInstanceFleetPropertyOutputReference interface {
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
	InstanceTypeConfigs() TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyList
	// Experimental.
	InstanceTypeConfigsInput() interface{}
	// Experimental.
	InternalValue() *TfCluster_MasterInstanceFleetProperty
	// Experimental.
	SetInternalValue(val *TfCluster_MasterInstanceFleetProperty)
	// Experimental.
	LaunchSpecifications() TfCluster_MasterInstanceFleetLaunchSpecificationsPropertyOutputReference
	// Experimental.
	LaunchSpecificationsInput() *TfCluster_MasterInstanceFleetLaunchSpecificationsProperty
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
	PutLaunchSpecifications(value *TfCluster_MasterInstanceFleetLaunchSpecificationsProperty)
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

// The jsii proxy struct for TfCluster_MasterInstanceFleetPropertyOutputReference
type jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) InstanceTypeConfigs() TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyList {
	var returns TfCluster_MasterInstanceFleetInstanceTypeConfigsPropertyList
	_jsii_.Get(
		j,
		"instanceTypeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) InstanceTypeConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceTypeConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) InternalValue() *TfCluster_MasterInstanceFleetProperty {
	var returns *TfCluster_MasterInstanceFleetProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) LaunchSpecifications() TfCluster_MasterInstanceFleetLaunchSpecificationsPropertyOutputReference {
	var returns TfCluster_MasterInstanceFleetLaunchSpecificationsPropertyOutputReference
	_jsii_.Get(
		j,
		"launchSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) LaunchSpecificationsInput() *TfCluster_MasterInstanceFleetLaunchSpecificationsProperty {
	var returns *TfCluster_MasterInstanceFleetLaunchSpecificationsProperty
	_jsii_.Get(
		j,
		"launchSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) ProvisionedOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) ProvisionedSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) TargetOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) TargetOnDemandCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) TargetSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) TargetSpotCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCluster_MasterInstanceFleetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCluster_MasterInstanceFleetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCluster_MasterInstanceFleetPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.TfCluster.MasterInstanceFleetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCluster_MasterInstanceFleetPropertyOutputReference_Override(t TfCluster_MasterInstanceFleetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.TfCluster.MasterInstanceFleetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference)SetInternalValue(val *TfCluster_MasterInstanceFleetProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference)SetTargetOnDemandCapacity(val *float64) {
	if err := j.validateSetTargetOnDemandCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetOnDemandCapacity",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference)SetTargetSpotCapacity(val *float64) {
	if err := j.validateSetTargetSpotCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSpotCapacity",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) PutInstanceTypeConfigs(value interface{}) {
	if err := t.validatePutInstanceTypeConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceTypeConfigs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) PutLaunchSpecifications(value *TfCluster_MasterInstanceFleetLaunchSpecificationsProperty) {
	if err := t.validatePutLaunchSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLaunchSpecifications",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) ResetInstanceTypeConfigs() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceTypeConfigs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) ResetLaunchSpecifications() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchSpecifications",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) ResetTargetOnDemandCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetOnDemandCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) ResetTargetSpotCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetSpotCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCluster_MasterInstanceFleetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

