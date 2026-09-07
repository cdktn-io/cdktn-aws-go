//go:build no_runtime_type_checking

package eks

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCluster_OidcPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsCluster_OidcPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsCluster_OidcPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsCluster_OidcPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsCluster_OidcPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsCluster_OidcPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsCluster_OidcPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

