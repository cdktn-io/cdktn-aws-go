//go:build no_runtime_type_checking

package eks

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCluster_IdentityPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsCluster_IdentityPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsCluster_IdentityPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsCluster_IdentityPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsCluster_IdentityPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsCluster_IdentityPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsCluster_IdentityPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

