package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference interface {
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
	ResetEndpointReference()
	// Experimental.
	ResetEvaluateTargetHealth()
	// Experimental.
	ResetHealthCheck()
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

// The jsii proxy struct for DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference
type jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) EndpointReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) EndpointReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) EvaluateTargetHealth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateTargetHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) EvaluateTargetHealthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateTargetHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) HealthCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) HealthCheckInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) RuleReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) RuleReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53.DataAwsRoute53TrafficPolicyDocument.RegionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference_Override(d DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.DataAwsRoute53TrafficPolicyDocument.RegionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference)SetEndpointReference(val *string) {
	if err := j.validateSetEndpointReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointReference",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference)SetEvaluateTargetHealth(val interface{}) {
	if err := j.validateSetEvaluateTargetHealthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluateTargetHealth",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference)SetHealthCheck(val *string) {
	if err := j.validateSetHealthCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheck",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference)SetRuleReference(val *string) {
	if err := j.validateSetRuleReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleReference",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) ResetEndpointReference() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) ResetEvaluateTargetHealth() {
	_jsii_.InvokeVoid(
		d,
		"resetEvaluateTargetHealth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) ResetRuleReference() {
	_jsii_.InvokeVoid(
		d,
		"resetRuleReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_RegionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

