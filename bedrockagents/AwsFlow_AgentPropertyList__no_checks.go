//go:build no_runtime_type_checking

package bedrockagents

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsFlow_AgentPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsFlow_AgentPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsFlow_AgentPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsFlow_AgentPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsFlow_AgentPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsFlow_AgentPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsFlow_AgentPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsFlow_AgentPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

