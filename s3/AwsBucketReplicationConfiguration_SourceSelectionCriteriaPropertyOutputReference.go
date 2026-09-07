package s3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference interface {
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
	InternalValue() *AwsBucketReplicationConfiguration_SourceSelectionCriteriaProperty
	// Experimental.
	SetInternalValue(val *AwsBucketReplicationConfiguration_SourceSelectionCriteriaProperty)
	// Experimental.
	ReplicaModifications() AwsBucketReplicationConfiguration_ReplicaModificationsPropertyOutputReference
	// Experimental.
	ReplicaModificationsInput() *AwsBucketReplicationConfiguration_ReplicaModificationsProperty
	// Experimental.
	SseKmsEncryptedObjects() AwsBucketReplicationConfiguration_SseKmsEncryptedObjectsPropertyOutputReference
	// Experimental.
	SseKmsEncryptedObjectsInput() *AwsBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty
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
	PutReplicaModifications(value *AwsBucketReplicationConfiguration_ReplicaModificationsProperty)
	// Experimental.
	PutSseKmsEncryptedObjects(value *AwsBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty)
	// Experimental.
	ResetReplicaModifications()
	// Experimental.
	ResetSseKmsEncryptedObjects()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference
type jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) InternalValue() *AwsBucketReplicationConfiguration_SourceSelectionCriteriaProperty {
	var returns *AwsBucketReplicationConfiguration_SourceSelectionCriteriaProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ReplicaModifications() AwsBucketReplicationConfiguration_ReplicaModificationsPropertyOutputReference {
	var returns AwsBucketReplicationConfiguration_ReplicaModificationsPropertyOutputReference
	_jsii_.Get(
		j,
		"replicaModifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ReplicaModificationsInput() *AwsBucketReplicationConfiguration_ReplicaModificationsProperty {
	var returns *AwsBucketReplicationConfiguration_ReplicaModificationsProperty
	_jsii_.Get(
		j,
		"replicaModificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) SseKmsEncryptedObjects() AwsBucketReplicationConfiguration_SseKmsEncryptedObjectsPropertyOutputReference {
	var returns AwsBucketReplicationConfiguration_SseKmsEncryptedObjectsPropertyOutputReference
	_jsii_.Get(
		j,
		"sseKmsEncryptedObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) SseKmsEncryptedObjectsInput() *AwsBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty {
	var returns *AwsBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty
	_jsii_.Get(
		j,
		"sseKmsEncryptedObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucketReplicationConfiguration.SourceSelectionCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference_Override(a AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucketReplicationConfiguration.SourceSelectionCriteriaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference)SetInternalValue(val *AwsBucketReplicationConfiguration_SourceSelectionCriteriaProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) PutReplicaModifications(value *AwsBucketReplicationConfiguration_ReplicaModificationsProperty) {
	if err := a.validatePutReplicaModificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReplicaModifications",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) PutSseKmsEncryptedObjects(value *AwsBucketReplicationConfiguration_SseKmsEncryptedObjectsProperty) {
	if err := a.validatePutSseKmsEncryptedObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSseKmsEncryptedObjects",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ResetReplicaModifications() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicaModifications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ResetSseKmsEncryptedObjects() {
	_jsii_.InvokeVoid(
		a,
		"resetSseKmsEncryptedObjects",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBucketReplicationConfiguration_SourceSelectionCriteriaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

