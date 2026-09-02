package awselasticsearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselasticsearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselasticsearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_ClusterConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ColdStorageOptions() TfDomain_ColdStorageOptionsPropertyOutputReference
	// Experimental.
	ColdStorageOptionsInput() *TfDomain_ColdStorageOptionsProperty
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
	DedicatedMasterCount() *float64
	// Experimental.
	SetDedicatedMasterCount(val *float64)
	// Experimental.
	DedicatedMasterCountInput() *float64
	// Experimental.
	DedicatedMasterEnabled() interface{}
	// Experimental.
	SetDedicatedMasterEnabled(val interface{})
	// Experimental.
	DedicatedMasterEnabledInput() interface{}
	// Experimental.
	DedicatedMasterType() *string
	// Experimental.
	SetDedicatedMasterType(val *string)
	// Experimental.
	DedicatedMasterTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InstanceCount() *float64
	// Experimental.
	SetInstanceCount(val *float64)
	// Experimental.
	InstanceCountInput() *float64
	// Experimental.
	InstanceType() *string
	// Experimental.
	SetInstanceType(val *string)
	// Experimental.
	InstanceTypeInput() *string
	// Experimental.
	InternalValue() *TfDomain_ClusterConfigProperty
	// Experimental.
	SetInternalValue(val *TfDomain_ClusterConfigProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WarmCount() *float64
	// Experimental.
	SetWarmCount(val *float64)
	// Experimental.
	WarmCountInput() *float64
	// Experimental.
	WarmEnabled() interface{}
	// Experimental.
	SetWarmEnabled(val interface{})
	// Experimental.
	WarmEnabledInput() interface{}
	// Experimental.
	WarmType() *string
	// Experimental.
	SetWarmType(val *string)
	// Experimental.
	WarmTypeInput() *string
	// Experimental.
	ZoneAwarenessConfig() TfDomain_ZoneAwarenessConfigPropertyOutputReference
	// Experimental.
	ZoneAwarenessConfigInput() *TfDomain_ZoneAwarenessConfigProperty
	// Experimental.
	ZoneAwarenessEnabled() interface{}
	// Experimental.
	SetZoneAwarenessEnabled(val interface{})
	// Experimental.
	ZoneAwarenessEnabledInput() interface{}
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
	PutColdStorageOptions(value *TfDomain_ColdStorageOptionsProperty)
	// Experimental.
	PutZoneAwarenessConfig(value *TfDomain_ZoneAwarenessConfigProperty)
	// Experimental.
	ResetColdStorageOptions()
	// Experimental.
	ResetDedicatedMasterCount()
	// Experimental.
	ResetDedicatedMasterEnabled()
	// Experimental.
	ResetDedicatedMasterType()
	// Experimental.
	ResetInstanceCount()
	// Experimental.
	ResetInstanceType()
	// Experimental.
	ResetWarmCount()
	// Experimental.
	ResetWarmEnabled()
	// Experimental.
	ResetWarmType()
	// Experimental.
	ResetZoneAwarenessConfig()
	// Experimental.
	ResetZoneAwarenessEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_ClusterConfigPropertyOutputReference
type jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ColdStorageOptions() TfDomain_ColdStorageOptionsPropertyOutputReference {
	var returns TfDomain_ColdStorageOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"coldStorageOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ColdStorageOptionsInput() *TfDomain_ColdStorageOptionsProperty {
	var returns *TfDomain_ColdStorageOptionsProperty
	_jsii_.Get(
		j,
		"coldStorageOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) DedicatedMasterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) DedicatedMasterCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) DedicatedMasterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedMasterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) DedicatedMasterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedMasterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) DedicatedMasterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) DedicatedMasterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) InternalValue() *TfDomain_ClusterConfigProperty {
	var returns *TfDomain_ClusterConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) WarmCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) WarmCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) WarmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) WarmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) WarmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) WarmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ZoneAwarenessConfig() TfDomain_ZoneAwarenessConfigPropertyOutputReference {
	var returns TfDomain_ZoneAwarenessConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"zoneAwarenessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ZoneAwarenessConfigInput() *TfDomain_ZoneAwarenessConfigProperty {
	var returns *TfDomain_ZoneAwarenessConfigProperty
	_jsii_.Get(
		j,
		"zoneAwarenessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ZoneAwarenessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ZoneAwarenessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessEnabledInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_ClusterConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_ClusterConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_ClusterConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elasticsearch.TfDomain.ClusterConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_ClusterConfigPropertyOutputReference_Override(t TfDomain_ClusterConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elasticsearch.TfDomain.ClusterConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetDedicatedMasterCount(val *float64) {
	if err := j.validateSetDedicatedMasterCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterCount",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetDedicatedMasterEnabled(val interface{}) {
	if err := j.validateSetDedicatedMasterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterEnabled",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetDedicatedMasterType(val *string) {
	if err := j.validateSetDedicatedMasterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterType",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetInternalValue(val *TfDomain_ClusterConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetWarmCount(val *float64) {
	if err := j.validateSetWarmCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmCount",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetWarmEnabled(val interface{}) {
	if err := j.validateSetWarmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmEnabled",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetWarmType(val *string) {
	if err := j.validateSetWarmTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmType",
		val,
	)
}

func (j *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference)SetZoneAwarenessEnabled(val interface{}) {
	if err := j.validateSetZoneAwarenessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneAwarenessEnabled",
		val,
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) PutColdStorageOptions(value *TfDomain_ColdStorageOptionsProperty) {
	if err := t.validatePutColdStorageOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putColdStorageOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) PutZoneAwarenessConfig(value *TfDomain_ZoneAwarenessConfigProperty) {
	if err := t.validatePutZoneAwarenessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putZoneAwarenessConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ResetColdStorageOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetColdStorageOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ResetDedicatedMasterCount() {
	_jsii_.InvokeVoid(
		t,
		"resetDedicatedMasterCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ResetDedicatedMasterEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetDedicatedMasterEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ResetDedicatedMasterType() {
	_jsii_.InvokeVoid(
		t,
		"resetDedicatedMasterType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ResetWarmCount() {
	_jsii_.InvokeVoid(
		t,
		"resetWarmCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ResetWarmEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetWarmEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ResetWarmType() {
	_jsii_.InvokeVoid(
		t,
		"resetWarmType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ResetZoneAwarenessConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetZoneAwarenessConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ResetZoneAwarenessEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetZoneAwarenessEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_ClusterConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

