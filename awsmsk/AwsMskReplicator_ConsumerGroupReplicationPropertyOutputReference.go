package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmsk/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmsk/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference interface {
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
	// Experimental.
	ConsumerGroupOffsetSyncMode() *string
	// Experimental.
	SetConsumerGroupOffsetSyncMode(val *string)
	// Experimental.
	ConsumerGroupOffsetSyncModeInput() *string
	// Experimental.
	ConsumerGroupsToExclude() *[]*string
	// Experimental.
	SetConsumerGroupsToExclude(val *[]*string)
	// Experimental.
	ConsumerGroupsToExcludeInput() *[]*string
	// Experimental.
	ConsumerGroupsToReplicate() *[]*string
	// Experimental.
	SetConsumerGroupsToReplicate(val *[]*string)
	// Experimental.
	ConsumerGroupsToReplicateInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DetectAndCopyNewConsumerGroups() interface{}
	// Experimental.
	SetDetectAndCopyNewConsumerGroups(val interface{})
	// Experimental.
	DetectAndCopyNewConsumerGroupsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SynchroniseConsumerGroupOffsets() interface{}
	// Experimental.
	SetSynchroniseConsumerGroupOffsets(val interface{})
	// Experimental.
	SynchroniseConsumerGroupOffsetsInput() interface{}
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
	ResetConsumerGroupOffsetSyncMode()
	// Experimental.
	ResetConsumerGroupsToExclude()
	// Experimental.
	ResetDetectAndCopyNewConsumerGroups()
	// Experimental.
	ResetSynchroniseConsumerGroupOffsets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference
type jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupOffsetSyncMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupOffsetSyncMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupOffsetSyncModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupOffsetSyncModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupsToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerGroupsToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupsToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerGroupsToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupsToReplicate() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerGroupsToReplicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ConsumerGroupsToReplicateInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"consumerGroupsToReplicateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) DetectAndCopyNewConsumerGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectAndCopyNewConsumerGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) DetectAndCopyNewConsumerGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectAndCopyNewConsumerGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) SynchroniseConsumerGroupOffsets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"synchroniseConsumerGroupOffsets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) SynchroniseConsumerGroupOffsetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"synchroniseConsumerGroupOffsetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMskReplicator_ConsumerGroupReplicationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-msk.AwsMskReplicator.ConsumerGroupReplicationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference_Override(a AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-msk.AwsMskReplicator.ConsumerGroupReplicationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference)SetConsumerGroupOffsetSyncMode(val *string) {
	if err := j.validateSetConsumerGroupOffsetSyncModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroupOffsetSyncMode",
		val,
	)
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference)SetConsumerGroupsToExclude(val *[]*string) {
	if err := j.validateSetConsumerGroupsToExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroupsToExclude",
		val,
	)
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference)SetConsumerGroupsToReplicate(val *[]*string) {
	if err := j.validateSetConsumerGroupsToReplicateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroupsToReplicate",
		val,
	)
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference)SetDetectAndCopyNewConsumerGroups(val interface{}) {
	if err := j.validateSetDetectAndCopyNewConsumerGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectAndCopyNewConsumerGroups",
		val,
	)
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference)SetSynchroniseConsumerGroupOffsets(val interface{}) {
	if err := j.validateSetSynchroniseConsumerGroupOffsetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"synchroniseConsumerGroupOffsets",
		val,
	)
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ResetConsumerGroupOffsetSyncMode() {
	_jsii_.InvokeVoid(
		a,
		"resetConsumerGroupOffsetSyncMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ResetConsumerGroupsToExclude() {
	_jsii_.InvokeVoid(
		a,
		"resetConsumerGroupsToExclude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ResetDetectAndCopyNewConsumerGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetDetectAndCopyNewConsumerGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ResetSynchroniseConsumerGroupOffsets() {
	_jsii_.InvokeVoid(
		a,
		"resetSynchroniseConsumerGroupOffsets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMskReplicator_ConsumerGroupReplicationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

