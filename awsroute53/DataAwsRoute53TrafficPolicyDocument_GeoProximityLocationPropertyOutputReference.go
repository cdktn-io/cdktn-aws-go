package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Bias() *string
	// Experimental.
	SetBias(val *string)
	// Experimental.
	BiasInput() *string
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
	EndpointReference() *string
	// Experimental.
	SetEndpointReference(val *string)
	// Experimental.
	EndpointReferenceInput() *string
	// Experimental.
	EvaluateTargetHealth() interface{}
	// Experimental.
	SetEvaluateTargetHealth(val interface{})
	// Experimental.
	EvaluateTargetHealthInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	HealthCheck() *string
	// Experimental.
	SetHealthCheck(val *string)
	// Experimental.
	HealthCheckInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Latitude() *string
	// Experimental.
	SetLatitude(val *string)
	// Experimental.
	LatitudeInput() *string
	// Experimental.
	Longitude() *string
	// Experimental.
	SetLongitude(val *string)
	// Experimental.
	LongitudeInput() *string
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RuleReference() *string
	// Experimental.
	SetRuleReference(val *string)
	// Experimental.
	RuleReferenceInput() *string
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
	ResetBias()
	// Experimental.
	ResetEndpointReference()
	// Experimental.
	ResetEvaluateTargetHealth()
	// Experimental.
	ResetHealthCheck()
	// Experimental.
	ResetLatitude()
	// Experimental.
	ResetLongitude()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRuleReference()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference
type jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) Bias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) BiasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"biasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) EndpointReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) EndpointReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) EvaluateTargetHealth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateTargetHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) EvaluateTargetHealthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateTargetHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) HealthCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) HealthCheckInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) Latitude() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) LatitudeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) Longitude() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) LongitudeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) RuleReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) RuleReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53.DataAwsRoute53TrafficPolicyDocument.GeoProximityLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference_Override(d DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.DataAwsRoute53TrafficPolicyDocument.GeoProximityLocationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetBias(val *string) {
	if err := j.validateSetBiasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bias",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetEndpointReference(val *string) {
	if err := j.validateSetEndpointReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointReference",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetEvaluateTargetHealth(val interface{}) {
	if err := j.validateSetEvaluateTargetHealthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluateTargetHealth",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetHealthCheck(val *string) {
	if err := j.validateSetHealthCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheck",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetLatitude(val *string) {
	if err := j.validateSetLatitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latitude",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetLongitude(val *string) {
	if err := j.validateSetLongitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longitude",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetRuleReference(val *string) {
	if err := j.validateSetRuleReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleReference",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ResetBias() {
	_jsii_.InvokeVoid(
		d,
		"resetBias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ResetEndpointReference() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ResetEvaluateTargetHealth() {
	_jsii_.InvokeVoid(
		d,
		"resetEvaluateTargetHealth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ResetLatitude() {
	_jsii_.InvokeVoid(
		d,
		"resetLatitude",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ResetLongitude() {
	_jsii_.InvokeVoid(
		d,
		"resetLongitude",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ResetRuleReference() {
	_jsii_.InvokeVoid(
		d,
		"resetRuleReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_GeoProximityLocationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

