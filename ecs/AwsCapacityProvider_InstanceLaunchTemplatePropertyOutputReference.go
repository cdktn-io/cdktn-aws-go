package ecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CapacityOptionType() *string
	// Experimental.
	SetCapacityOptionType(val *string)
	// Experimental.
	CapacityOptionTypeInput() *string
	// Experimental.
	CapacityReservations() AwsCapacityProvider_CapacityReservationsPropertyOutputReference
	// Experimental.
	CapacityReservationsInput() *AwsCapacityProvider_CapacityReservationsProperty
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
	Ec2InstanceProfileArn() *string
	// Experimental.
	SetEc2InstanceProfileArn(val *string)
	// Experimental.
	Ec2InstanceProfileArnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InstanceRequirements() AwsCapacityProvider_InstanceRequirementsPropertyOutputReference
	// Experimental.
	InstanceRequirementsInput() *AwsCapacityProvider_InstanceRequirementsProperty
	// Experimental.
	InternalValue() *AwsCapacityProvider_InstanceLaunchTemplateProperty
	// Experimental.
	SetInternalValue(val *AwsCapacityProvider_InstanceLaunchTemplateProperty)
	// Experimental.
	LocalStorageConfiguration() AwsCapacityProvider_LocalStorageConfigurationPropertyOutputReference
	// Experimental.
	LocalStorageConfigurationInput() *AwsCapacityProvider_LocalStorageConfigurationProperty
	// Experimental.
	Monitoring() *string
	// Experimental.
	SetMonitoring(val *string)
	// Experimental.
	MonitoringInput() *string
	// Experimental.
	NetworkConfiguration() AwsCapacityProvider_NetworkConfigurationPropertyOutputReference
	// Experimental.
	NetworkConfigurationInput() *AwsCapacityProvider_NetworkConfigurationProperty
	// Experimental.
	StorageConfiguration() AwsCapacityProvider_StorageConfigurationPropertyOutputReference
	// Experimental.
	StorageConfigurationInput() *AwsCapacityProvider_StorageConfigurationProperty
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
	PutCapacityReservations(value *AwsCapacityProvider_CapacityReservationsProperty)
	// Experimental.
	PutInstanceRequirements(value *AwsCapacityProvider_InstanceRequirementsProperty)
	// Experimental.
	PutLocalStorageConfiguration(value *AwsCapacityProvider_LocalStorageConfigurationProperty)
	// Experimental.
	PutNetworkConfiguration(value *AwsCapacityProvider_NetworkConfigurationProperty)
	// Experimental.
	PutStorageConfiguration(value *AwsCapacityProvider_StorageConfigurationProperty)
	// Experimental.
	ResetCapacityOptionType()
	// Experimental.
	ResetCapacityReservations()
	// Experimental.
	ResetInstanceRequirements()
	// Experimental.
	ResetLocalStorageConfiguration()
	// Experimental.
	ResetMonitoring()
	// Experimental.
	ResetStorageConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference
type jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) CapacityOptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityOptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) CapacityOptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityOptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) CapacityReservations() AwsCapacityProvider_CapacityReservationsPropertyOutputReference {
	var returns AwsCapacityProvider_CapacityReservationsPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) CapacityReservationsInput() *AwsCapacityProvider_CapacityReservationsProperty {
	var returns *AwsCapacityProvider_CapacityReservationsProperty
	_jsii_.Get(
		j,
		"capacityReservationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) Ec2InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) Ec2InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) InstanceRequirements() AwsCapacityProvider_InstanceRequirementsPropertyOutputReference {
	var returns AwsCapacityProvider_InstanceRequirementsPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) InstanceRequirementsInput() *AwsCapacityProvider_InstanceRequirementsProperty {
	var returns *AwsCapacityProvider_InstanceRequirementsProperty
	_jsii_.Get(
		j,
		"instanceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) InternalValue() *AwsCapacityProvider_InstanceLaunchTemplateProperty {
	var returns *AwsCapacityProvider_InstanceLaunchTemplateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) LocalStorageConfiguration() AwsCapacityProvider_LocalStorageConfigurationPropertyOutputReference {
	var returns AwsCapacityProvider_LocalStorageConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"localStorageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) LocalStorageConfigurationInput() *AwsCapacityProvider_LocalStorageConfigurationProperty {
	var returns *AwsCapacityProvider_LocalStorageConfigurationProperty
	_jsii_.Get(
		j,
		"localStorageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) Monitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) MonitoringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) NetworkConfiguration() AwsCapacityProvider_NetworkConfigurationPropertyOutputReference {
	var returns AwsCapacityProvider_NetworkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) NetworkConfigurationInput() *AwsCapacityProvider_NetworkConfigurationProperty {
	var returns *AwsCapacityProvider_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) StorageConfiguration() AwsCapacityProvider_StorageConfigurationPropertyOutputReference {
	var returns AwsCapacityProvider_StorageConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"storageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) StorageConfigurationInput() *AwsCapacityProvider_StorageConfigurationProperty {
	var returns *AwsCapacityProvider_StorageConfigurationProperty
	_jsii_.Get(
		j,
		"storageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsCapacityProvider.InstanceLaunchTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference_Override(a AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsCapacityProvider.InstanceLaunchTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetCapacityOptionType(val *string) {
	if err := j.validateSetCapacityOptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityOptionType",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetEc2InstanceProfileArn(val *string) {
	if err := j.validateSetEc2InstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2InstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetInternalValue(val *AwsCapacityProvider_InstanceLaunchTemplateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetMonitoring(val *string) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) PutCapacityReservations(value *AwsCapacityProvider_CapacityReservationsProperty) {
	if err := a.validatePutCapacityReservationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCapacityReservations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) PutInstanceRequirements(value *AwsCapacityProvider_InstanceRequirementsProperty) {
	if err := a.validatePutInstanceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceRequirements",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) PutLocalStorageConfiguration(value *AwsCapacityProvider_LocalStorageConfigurationProperty) {
	if err := a.validatePutLocalStorageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLocalStorageConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) PutNetworkConfiguration(value *AwsCapacityProvider_NetworkConfigurationProperty) {
	if err := a.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) PutStorageConfiguration(value *AwsCapacityProvider_StorageConfigurationProperty) {
	if err := a.validatePutStorageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStorageConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetCapacityOptionType() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityOptionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetCapacityReservations() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityReservations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetInstanceRequirements() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceRequirements",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetLocalStorageConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalStorageConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetMonitoring() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetStorageConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

