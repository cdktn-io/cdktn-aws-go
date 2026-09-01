package awselasticsearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselasticsearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselasticsearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsElasticsearchDomain_ClusterConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ColdStorageOptions() AwsElasticsearchDomain_ColdStorageOptionsPropertyOutputReference
	// Experimental.
	ColdStorageOptionsInput() *AwsElasticsearchDomain_ColdStorageOptionsProperty
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
	InternalValue() *AwsElasticsearchDomain_ClusterConfigProperty
	// Experimental.
	SetInternalValue(val *AwsElasticsearchDomain_ClusterConfigProperty)
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
	ZoneAwarenessConfig() AwsElasticsearchDomain_ZoneAwarenessConfigPropertyOutputReference
	// Experimental.
	ZoneAwarenessConfigInput() *AwsElasticsearchDomain_ZoneAwarenessConfigProperty
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
	PutColdStorageOptions(value *AwsElasticsearchDomain_ColdStorageOptionsProperty)
	// Experimental.
	PutZoneAwarenessConfig(value *AwsElasticsearchDomain_ZoneAwarenessConfigProperty)
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

// The jsii proxy struct for AwsElasticsearchDomain_ClusterConfigPropertyOutputReference
type jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ColdStorageOptions() AwsElasticsearchDomain_ColdStorageOptionsPropertyOutputReference {
	var returns AwsElasticsearchDomain_ColdStorageOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"coldStorageOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ColdStorageOptionsInput() *AwsElasticsearchDomain_ColdStorageOptionsProperty {
	var returns *AwsElasticsearchDomain_ColdStorageOptionsProperty
	_jsii_.Get(
		j,
		"coldStorageOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) DedicatedMasterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) DedicatedMasterCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) DedicatedMasterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedMasterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) DedicatedMasterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedMasterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) DedicatedMasterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) DedicatedMasterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) InternalValue() *AwsElasticsearchDomain_ClusterConfigProperty {
	var returns *AwsElasticsearchDomain_ClusterConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) WarmCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) WarmCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) WarmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) WarmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) WarmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) WarmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ZoneAwarenessConfig() AwsElasticsearchDomain_ZoneAwarenessConfigPropertyOutputReference {
	var returns AwsElasticsearchDomain_ZoneAwarenessConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"zoneAwarenessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ZoneAwarenessConfigInput() *AwsElasticsearchDomain_ZoneAwarenessConfigProperty {
	var returns *AwsElasticsearchDomain_ZoneAwarenessConfigProperty
	_jsii_.Get(
		j,
		"zoneAwarenessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ZoneAwarenessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ZoneAwarenessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessEnabledInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsElasticsearchDomain_ClusterConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsElasticsearchDomain_ClusterConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsElasticsearchDomain_ClusterConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elasticsearch.AwsElasticsearchDomain.ClusterConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsElasticsearchDomain_ClusterConfigPropertyOutputReference_Override(a AwsElasticsearchDomain_ClusterConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elasticsearch.AwsElasticsearchDomain.ClusterConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetDedicatedMasterCount(val *float64) {
	if err := j.validateSetDedicatedMasterCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterCount",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetDedicatedMasterEnabled(val interface{}) {
	if err := j.validateSetDedicatedMasterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetDedicatedMasterType(val *string) {
	if err := j.validateSetDedicatedMasterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedMasterType",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetInternalValue(val *AwsElasticsearchDomain_ClusterConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetWarmCount(val *float64) {
	if err := j.validateSetWarmCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmCount",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetWarmEnabled(val interface{}) {
	if err := j.validateSetWarmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetWarmType(val *string) {
	if err := j.validateSetWarmTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warmType",
		val,
	)
}

func (j *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference)SetZoneAwarenessEnabled(val interface{}) {
	if err := j.validateSetZoneAwarenessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneAwarenessEnabled",
		val,
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) PutColdStorageOptions(value *AwsElasticsearchDomain_ColdStorageOptionsProperty) {
	if err := a.validatePutColdStorageOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putColdStorageOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) PutZoneAwarenessConfig(value *AwsElasticsearchDomain_ZoneAwarenessConfigProperty) {
	if err := a.validatePutZoneAwarenessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZoneAwarenessConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ResetColdStorageOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetColdStorageOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ResetDedicatedMasterCount() {
	_jsii_.InvokeVoid(
		a,
		"resetDedicatedMasterCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ResetDedicatedMasterEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDedicatedMasterEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ResetDedicatedMasterType() {
	_jsii_.InvokeVoid(
		a,
		"resetDedicatedMasterType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ResetWarmCount() {
	_jsii_.InvokeVoid(
		a,
		"resetWarmCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ResetWarmEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetWarmEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ResetWarmType() {
	_jsii_.InvokeVoid(
		a,
		"resetWarmType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ResetZoneAwarenessConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetZoneAwarenessConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ResetZoneAwarenessEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetZoneAwarenessEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsElasticsearchDomain_ClusterConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

