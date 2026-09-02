package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReplicationConfig_ComputeConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AvailabilityZone() *string
	// Experimental.
	SetAvailabilityZone(val *string)
	// Experimental.
	AvailabilityZoneInput() *string
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
	DnsNameServers() *string
	// Experimental.
	SetDnsNameServers(val *string)
	// Experimental.
	DnsNameServersInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfReplicationConfig_ComputeConfigProperty
	// Experimental.
	SetInternalValue(val *TfReplicationConfig_ComputeConfigProperty)
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	SetKmsKeyId(val *string)
	// Experimental.
	KmsKeyIdInput() *string
	// Experimental.
	MaxCapacityUnits() *float64
	// Experimental.
	SetMaxCapacityUnits(val *float64)
	// Experimental.
	MaxCapacityUnitsInput() *float64
	// Experimental.
	MinCapacityUnits() *float64
	// Experimental.
	SetMinCapacityUnits(val *float64)
	// Experimental.
	MinCapacityUnitsInput() *float64
	// Experimental.
	MultiAz() interface{}
	// Experimental.
	SetMultiAz(val interface{})
	// Experimental.
	MultiAzInput() interface{}
	// Experimental.
	PreferredMaintenanceWindow() *string
	// Experimental.
	SetPreferredMaintenanceWindow(val *string)
	// Experimental.
	PreferredMaintenanceWindowInput() *string
	// Experimental.
	ReplicationSubnetGroupId() *string
	// Experimental.
	SetReplicationSubnetGroupId(val *string)
	// Experimental.
	ReplicationSubnetGroupIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcSecurityGroupIds() *[]*string
	// Experimental.
	SetVpcSecurityGroupIds(val *[]*string)
	// Experimental.
	VpcSecurityGroupIdsInput() *[]*string
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
	ResetAvailabilityZone()
	// Experimental.
	ResetDnsNameServers()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetMaxCapacityUnits()
	// Experimental.
	ResetMinCapacityUnits()
	// Experimental.
	ResetMultiAz()
	// Experimental.
	ResetPreferredMaintenanceWindow()
	// Experimental.
	ResetVpcSecurityGroupIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfReplicationConfig_ComputeConfigPropertyOutputReference
type jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) DnsNameServers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) DnsNameServersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) InternalValue() *TfReplicationConfig_ComputeConfigProperty {
	var returns *TfReplicationConfig_ComputeConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) MaxCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) MaxCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) MinCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) MinCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) MultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ReplicationSubnetGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ReplicationSubnetGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfReplicationConfig_ComputeConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfReplicationConfig_ComputeConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfReplicationConfig_ComputeConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.TfReplicationConfig.ComputeConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfReplicationConfig_ComputeConfigPropertyOutputReference_Override(t TfReplicationConfig_ComputeConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.TfReplicationConfig.ComputeConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetDnsNameServers(val *string) {
	if err := j.validateSetDnsNameServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsNameServers",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetInternalValue(val *TfReplicationConfig_ComputeConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetMaxCapacityUnits(val *float64) {
	if err := j.validateSetMaxCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetMinCapacityUnits(val *float64) {
	if err := j.validateSetMinCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetMultiAz(val interface{}) {
	if err := j.validateSetMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetReplicationSubnetGroupId(val *string) {
	if err := j.validateSetReplicationSubnetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationSubnetGroupId",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ResetDnsNameServers() {
	_jsii_.InvokeVoid(
		t,
		"resetDnsNameServers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ResetMaxCapacityUnits() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxCapacityUnits",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ResetMinCapacityUnits() {
	_jsii_.InvokeVoid(
		t,
		"resetMinCapacityUnits",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ResetMultiAz() {
	_jsii_.InvokeVoid(
		t,
		"resetMultiAz",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		t,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfReplicationConfig_ComputeConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

