package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CapacityOptionType() *string
	// Experimental.
	SetCapacityOptionType(val *string)
	// Experimental.
	CapacityOptionTypeInput() *string
	// Experimental.
	CapacityReservations() TfCapacityProvider_CapacityReservationsPropertyOutputReference
	// Experimental.
	CapacityReservationsInput() *TfCapacityProvider_CapacityReservationsProperty
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
	InstanceRequirements() TfCapacityProvider_InstanceRequirementsPropertyOutputReference
	// Experimental.
	InstanceRequirementsInput() *TfCapacityProvider_InstanceRequirementsProperty
	// Experimental.
	InternalValue() *TfCapacityProvider_InstanceLaunchTemplateProperty
	// Experimental.
	SetInternalValue(val *TfCapacityProvider_InstanceLaunchTemplateProperty)
	// Experimental.
	LocalStorageConfiguration() TfCapacityProvider_LocalStorageConfigurationPropertyOutputReference
	// Experimental.
	LocalStorageConfigurationInput() *TfCapacityProvider_LocalStorageConfigurationProperty
	// Experimental.
	Monitoring() *string
	// Experimental.
	SetMonitoring(val *string)
	// Experimental.
	MonitoringInput() *string
	// Experimental.
	NetworkConfiguration() TfCapacityProvider_NetworkConfigurationPropertyOutputReference
	// Experimental.
	NetworkConfigurationInput() *TfCapacityProvider_NetworkConfigurationProperty
	// Experimental.
	StorageConfiguration() TfCapacityProvider_StorageConfigurationPropertyOutputReference
	// Experimental.
	StorageConfigurationInput() *TfCapacityProvider_StorageConfigurationProperty
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
	PutCapacityReservations(value *TfCapacityProvider_CapacityReservationsProperty)
	// Experimental.
	PutInstanceRequirements(value *TfCapacityProvider_InstanceRequirementsProperty)
	// Experimental.
	PutLocalStorageConfiguration(value *TfCapacityProvider_LocalStorageConfigurationProperty)
	// Experimental.
	PutNetworkConfiguration(value *TfCapacityProvider_NetworkConfigurationProperty)
	// Experimental.
	PutStorageConfiguration(value *TfCapacityProvider_StorageConfigurationProperty)
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

// The jsii proxy struct for TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference
type jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) CapacityOptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityOptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) CapacityOptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityOptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) CapacityReservations() TfCapacityProvider_CapacityReservationsPropertyOutputReference {
	var returns TfCapacityProvider_CapacityReservationsPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) CapacityReservationsInput() *TfCapacityProvider_CapacityReservationsProperty {
	var returns *TfCapacityProvider_CapacityReservationsProperty
	_jsii_.Get(
		j,
		"capacityReservationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) Ec2InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) Ec2InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) InstanceRequirements() TfCapacityProvider_InstanceRequirementsPropertyOutputReference {
	var returns TfCapacityProvider_InstanceRequirementsPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) InstanceRequirementsInput() *TfCapacityProvider_InstanceRequirementsProperty {
	var returns *TfCapacityProvider_InstanceRequirementsProperty
	_jsii_.Get(
		j,
		"instanceRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) InternalValue() *TfCapacityProvider_InstanceLaunchTemplateProperty {
	var returns *TfCapacityProvider_InstanceLaunchTemplateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) LocalStorageConfiguration() TfCapacityProvider_LocalStorageConfigurationPropertyOutputReference {
	var returns TfCapacityProvider_LocalStorageConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"localStorageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) LocalStorageConfigurationInput() *TfCapacityProvider_LocalStorageConfigurationProperty {
	var returns *TfCapacityProvider_LocalStorageConfigurationProperty
	_jsii_.Get(
		j,
		"localStorageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) Monitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) MonitoringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) NetworkConfiguration() TfCapacityProvider_NetworkConfigurationPropertyOutputReference {
	var returns TfCapacityProvider_NetworkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) NetworkConfigurationInput() *TfCapacityProvider_NetworkConfigurationProperty {
	var returns *TfCapacityProvider_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) StorageConfiguration() TfCapacityProvider_StorageConfigurationPropertyOutputReference {
	var returns TfCapacityProvider_StorageConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"storageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) StorageConfigurationInput() *TfCapacityProvider_StorageConfigurationProperty {
	var returns *TfCapacityProvider_StorageConfigurationProperty
	_jsii_.Get(
		j,
		"storageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCapacityProvider_InstanceLaunchTemplatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCapacityProvider.InstanceLaunchTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference_Override(t TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCapacityProvider.InstanceLaunchTemplatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetCapacityOptionType(val *string) {
	if err := j.validateSetCapacityOptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityOptionType",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetEc2InstanceProfileArn(val *string) {
	if err := j.validateSetEc2InstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2InstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetInternalValue(val *TfCapacityProvider_InstanceLaunchTemplateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetMonitoring(val *string) {
	if err := j.validateSetMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoring",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) PutCapacityReservations(value *TfCapacityProvider_CapacityReservationsProperty) {
	if err := t.validatePutCapacityReservationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityReservations",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) PutInstanceRequirements(value *TfCapacityProvider_InstanceRequirementsProperty) {
	if err := t.validatePutInstanceRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceRequirements",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) PutLocalStorageConfiguration(value *TfCapacityProvider_LocalStorageConfigurationProperty) {
	if err := t.validatePutLocalStorageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLocalStorageConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) PutNetworkConfiguration(value *TfCapacityProvider_NetworkConfigurationProperty) {
	if err := t.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) PutStorageConfiguration(value *TfCapacityProvider_StorageConfigurationProperty) {
	if err := t.validatePutStorageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStorageConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetCapacityOptionType() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityOptionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetCapacityReservations() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityReservations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetInstanceRequirements() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceRequirements",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetLocalStorageConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLocalStorageConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetMonitoring() {
	_jsii_.InvokeVoid(
		t,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ResetStorageConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

