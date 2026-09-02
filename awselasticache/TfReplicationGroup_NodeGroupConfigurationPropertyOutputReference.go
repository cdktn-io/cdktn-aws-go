package awselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselasticache/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselasticache/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference interface {
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

// The jsii proxy struct for TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference
type jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) NodeGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) NodeGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) PrimaryAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) PrimaryAvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) PrimaryOutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryOutpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) PrimaryOutpostArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryOutpostArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaAvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaAvailabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaAvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaAvailabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaOutpostArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaOutpostArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ReplicaOutpostArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaOutpostArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) Slots() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) SlotsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfReplicationGroup_NodeGroupConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfReplicationGroup_NodeGroupConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elasticache.TfReplicationGroup.NodeGroupConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfReplicationGroup_NodeGroupConfigurationPropertyOutputReference_Override(t TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elasticache.TfReplicationGroup.NodeGroupConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetNodeGroupId(val *string) {
	if err := j.validateSetNodeGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroupId",
		val,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetPrimaryAvailabilityZone(val *string) {
	if err := j.validateSetPrimaryAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetPrimaryOutpostArn(val *string) {
	if err := j.validateSetPrimaryOutpostArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryOutpostArn",
		val,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetReplicaAvailabilityZones(val *[]*string) {
	if err := j.validateSetReplicaAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaAvailabilityZones",
		val,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetReplicaCount(val *float64) {
	if err := j.validateSetReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaCount",
		val,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetReplicaOutpostArns(val *[]*string) {
	if err := j.validateSetReplicaOutpostArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaOutpostArns",
		val,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetSlots(val *string) {
	if err := j.validateSetSlotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slots",
		val,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetNodeGroupId() {
	_jsii_.InvokeVoid(
		t,
		"resetNodeGroupId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetPrimaryAvailabilityZone() {
	_jsii_.InvokeVoid(
		t,
		"resetPrimaryAvailabilityZone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetPrimaryOutpostArn() {
	_jsii_.InvokeVoid(
		t,
		"resetPrimaryOutpostArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetReplicaAvailabilityZones() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicaAvailabilityZones",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetReplicaCount() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicaCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetReplicaOutpostArns() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicaOutpostArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ResetSlots() {
	_jsii_.InvokeVoid(
		t,
		"resetSlots",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfReplicationGroup_NodeGroupConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

