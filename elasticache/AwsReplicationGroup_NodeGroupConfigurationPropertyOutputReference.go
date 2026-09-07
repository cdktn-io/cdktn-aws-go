package elasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elasticache/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elasticache/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	NodeGroupId() *string
	// Experimental.
	SetNodeGroupId(val *string)
	// Experimental.
	NodeGroupIdInput() *string
	// Experimental.
	PrimaryAvailabilityZone() *string
	// Experimental.
	SetPrimaryAvailabilityZone(val *string)
	// Experimental.
	PrimaryAvailabilityZoneInput() *string
	// Experimental.
	PrimaryOutpostArn() *string
	// Experimental.
	SetPrimaryOutpostArn(val *string)
	// Experimental.
	PrimaryOutpostArnInput() *string
	// Experimental.
	ReplicaAvailabilityZones() *[]*string
	// Experimental.
	SetReplicaAvailabilityZones(val *[]*string)
	// Experimental.
	ReplicaAvailabilityZonesInput() *[]*string
	// Experimental.
	ReplicaCount() *float64
	// Experimental.
	SetReplicaCount(val *float64)
	// Experimental.
	ReplicaCountInput() *float64
	// Experimental.
	ReplicaOutpostArns() *[]*string
	// Experimental.
	SetReplicaOutpostArns(val *[]*string)
	// Experimental.
	ReplicaOutpostArnsInput() *[]*string
	// Experimental.
	Slots() *string
	// Experimental.
	SetSlots(val *string)
	// Experimental.
	SlotsInput() *string
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
	ResetNodeGroupId()
	// Experimental.
	ResetPrimaryAvailabilityZone()
	// Experimental.
	ResetPrimaryOutpostArn()
	// Experimental.
	ResetReplicaAvailabilityZones()
	// Experimental.
	ResetReplicaCount()
	// Experimental.
	ResetReplicaOutpostArns()
	// Experimental.
	ResetSlots()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference
type jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) NodeGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) NodeGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) PrimaryAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) PrimaryAvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) PrimaryOutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryOutpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) PrimaryOutpostArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryOutpostArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaAvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaAvailabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaAvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaAvailabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaOutpostArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaOutpostArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaOutpostArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaOutpostArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) Slots() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) SlotsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsReplicationGroup_NodeGroupConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elasticache.AwsReplicationGroup.NodeGroupConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference_Override(a AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elasticache.AwsReplicationGroup.NodeGroupConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetNodeGroupId(val *string) {
	if err := j.validateSetNodeGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroupId",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetPrimaryAvailabilityZone(val *string) {
	if err := j.validateSetPrimaryAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetPrimaryOutpostArn(val *string) {
	if err := j.validateSetPrimaryOutpostArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryOutpostArn",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetReplicaAvailabilityZones(val *[]*string) {
	if err := j.validateSetReplicaAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaAvailabilityZones",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetReplicaCount(val *float64) {
	if err := j.validateSetReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaCount",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetReplicaOutpostArns(val *[]*string) {
	if err := j.validateSetReplicaOutpostArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaOutpostArns",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetSlots(val *string) {
	if err := j.validateSetSlotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slots",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetNodeGroupId() {
	_jsii_.InvokeVoid(
		a,
		"resetNodeGroupId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetPrimaryAvailabilityZone() {
	_jsii_.InvokeVoid(
		a,
		"resetPrimaryAvailabilityZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetPrimaryOutpostArn() {
	_jsii_.InvokeVoid(
		a,
		"resetPrimaryOutpostArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetReplicaAvailabilityZones() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicaAvailabilityZones",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetReplicaCount() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicaCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetReplicaOutpostArns() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicaOutpostArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetSlots() {
	_jsii_.InvokeVoid(
		a,
		"resetSlots",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

