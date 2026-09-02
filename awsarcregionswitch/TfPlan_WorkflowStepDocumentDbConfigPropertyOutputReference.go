package awsarcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Behavior() *string
	// Experimental.
	SetBehavior(val *string)
	// Experimental.
	BehaviorInput() *string
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
	CrossAccountRole() *string
	// Experimental.
	SetCrossAccountRole(val *string)
	// Experimental.
	CrossAccountRoleInput() *string
	// Experimental.
	DatabaseClusterArns() *[]*string
	// Experimental.
	SetDatabaseClusterArns(val *[]*string)
	// Experimental.
	DatabaseClusterArnsInput() *[]*string
	// Experimental.
	ExternalId() *string
	// Experimental.
	SetExternalId(val *string)
	// Experimental.
	ExternalIdInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	GlobalClusterIdentifier() *string
	// Experimental.
	SetGlobalClusterIdentifier(val *string)
	// Experimental.
	GlobalClusterIdentifierInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutMinutes() *float64
	// Experimental.
	SetTimeoutMinutes(val *float64)
	// Experimental.
	TimeoutMinutesInput() *float64
	// Experimental.
	Ungraceful() TfPlan_WorkflowStepDocumentDbConfigUngracefulPropertyList
	// Experimental.
	UngracefulInput() interface{}
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
	PutUngraceful(value interface{})
	// Experimental.
	ResetCrossAccountRole()
	// Experimental.
	ResetExternalId()
	// Experimental.
	ResetTimeoutMinutes()
	// Experimental.
	ResetUngraceful()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference
type jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) BehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) CrossAccountRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossAccountRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) CrossAccountRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossAccountRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) DatabaseClusterArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"databaseClusterArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) DatabaseClusterArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"databaseClusterArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) GlobalClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) GlobalClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) TimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) TimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) Ungraceful() TfPlan_WorkflowStepDocumentDbConfigUngracefulPropertyList {
	var returns TfPlan_WorkflowStepDocumentDbConfigUngracefulPropertyList
	_jsii_.Get(
		j,
		"ungraceful",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) UngracefulInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ungracefulInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPlan_WorkflowStepDocumentDbConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.TfPlan.WorkflowStepDocumentDbConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference_Override(t TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.TfPlan.WorkflowStepDocumentDbConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference)SetBehavior(val *string) {
	if err := j.validateSetBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"behavior",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference)SetCrossAccountRole(val *string) {
	if err := j.validateSetCrossAccountRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossAccountRole",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference)SetDatabaseClusterArns(val *[]*string) {
	if err := j.validateSetDatabaseClusterArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseClusterArns",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference)SetGlobalClusterIdentifier(val *string) {
	if err := j.validateSetGlobalClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference)SetTimeoutMinutes(val *float64) {
	if err := j.validateSetTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMinutes",
		val,
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) PutUngraceful(value interface{}) {
	if err := t.validatePutUngracefulParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUngraceful",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) ResetCrossAccountRole() {
	_jsii_.InvokeVoid(
		t,
		"resetCrossAccountRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) ResetExternalId() {
	_jsii_.InvokeVoid(
		t,
		"resetExternalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) ResetTimeoutMinutes() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeoutMinutes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) ResetUngraceful() {
	_jsii_.InvokeVoid(
		t,
		"resetUngraceful",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPlan_WorkflowStepDocumentDbConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

