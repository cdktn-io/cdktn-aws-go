package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsroute53/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsroute53/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference interface {
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
	InternalValue() *DataAwsRoute53TrafficPolicyDocument_SecondaryProperty
	// Experimental.
	SetInternalValue(val *DataAwsRoute53TrafficPolicyDocument_SecondaryProperty)
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

// The jsii proxy struct for DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference
type jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) EndpointReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) EndpointReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) EvaluateTargetHealth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateTargetHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) EvaluateTargetHealthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateTargetHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) HealthCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) HealthCheckInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) InternalValue() *DataAwsRoute53TrafficPolicyDocument_SecondaryProperty {
	var returns *DataAwsRoute53TrafficPolicyDocument_SecondaryProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) RuleReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) RuleReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-route-53.DataAwsRoute53TrafficPolicyDocument.SecondaryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference_Override(d DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-route-53.DataAwsRoute53TrafficPolicyDocument.SecondaryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference)SetEndpointReference(val *string) {
	if err := j.validateSetEndpointReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointReference",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference)SetEvaluateTargetHealth(val interface{}) {
	if err := j.validateSetEvaluateTargetHealthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluateTargetHealth",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference)SetHealthCheck(val *string) {
	if err := j.validateSetHealthCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheck",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference)SetInternalValue(val *DataAwsRoute53TrafficPolicyDocument_SecondaryProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference)SetRuleReference(val *string) {
	if err := j.validateSetRuleReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleReference",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) ResetEndpointReference() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) ResetEvaluateTargetHealth() {
	_jsii_.InvokeVoid(
		d,
		"resetEvaluateTargetHealth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		d,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) ResetRuleReference() {
	_jsii_.InvokeVoid(
		d,
		"resetRuleReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsRoute53TrafficPolicyDocument_SecondaryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

