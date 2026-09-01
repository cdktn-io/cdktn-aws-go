package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEmrCluster_CoreInstanceFleetPropertyOutputReference interface {
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
	InstanceTypeConfigs() AwsEmrCluster_CoreInstanceFleetInstanceTypeConfigsPropertyList
	// Experimental.
	InstanceTypeConfigsInput() interface{}
	// Experimental.
	InternalValue() *AwsEmrCluster_CoreInstanceFleetProperty
	// Experimental.
	SetInternalValue(val *AwsEmrCluster_CoreInstanceFleetProperty)
	// Experimental.
	LaunchSpecifications() AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference
	// Experimental.
	LaunchSpecificationsInput() *AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsProperty
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
	PutLaunchSpecifications(value *AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsProperty)
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

// The jsii proxy struct for AwsEmrCluster_CoreInstanceFleetPropertyOutputReference
type jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) InstanceTypeConfigs() AwsEmrCluster_CoreInstanceFleetInstanceTypeConfigsPropertyList {
	var returns AwsEmrCluster_CoreInstanceFleetInstanceTypeConfigsPropertyList
	_jsii_.Get(
		j,
		"instanceTypeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) InstanceTypeConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceTypeConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) InternalValue() *AwsEmrCluster_CoreInstanceFleetProperty {
	var returns *AwsEmrCluster_CoreInstanceFleetProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) LaunchSpecifications() AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference {
	var returns AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference
	_jsii_.Get(
		j,
		"launchSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) LaunchSpecificationsInput() *AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsProperty {
	var returns *AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsProperty
	_jsii_.Get(
		j,
		"launchSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) ProvisionedOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) ProvisionedSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) TargetOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) TargetOnDemandCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) TargetSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) TargetSpotCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEmrCluster_CoreInstanceFleetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEmrCluster_CoreInstanceFleetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEmrCluster_CoreInstanceFleetPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrCluster.CoreInstanceFleetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEmrCluster_CoreInstanceFleetPropertyOutputReference_Override(a AwsEmrCluster_CoreInstanceFleetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrCluster.CoreInstanceFleetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference)SetInternalValue(val *AwsEmrCluster_CoreInstanceFleetProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference)SetTargetOnDemandCapacity(val *float64) {
	if err := j.validateSetTargetOnDemandCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetOnDemandCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference)SetTargetSpotCapacity(val *float64) {
	if err := j.validateSetTargetSpotCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSpotCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) PutInstanceTypeConfigs(value interface{}) {
	if err := a.validatePutInstanceTypeConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceTypeConfigs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) PutLaunchSpecifications(value *AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsProperty) {
	if err := a.validatePutLaunchSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchSpecifications",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) ResetInstanceTypeConfigs() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceTypeConfigs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) ResetLaunchSpecifications() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchSpecifications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) ResetTargetOnDemandCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetOnDemandCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) ResetTargetSpotCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetSpotCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

