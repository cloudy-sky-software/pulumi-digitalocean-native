# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetRegistryDockerCredentialResult',
    'AwaitableGetRegistryDockerCredentialResult',
    'get_registry_docker_credential',
    'get_registry_docker_credential_output',
]

@pulumi.output_type
class GetRegistryDockerCredentialResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.DockerCredentials':
        return pulumi.get(self, "items")


class AwaitableGetRegistryDockerCredentialResult(GetRegistryDockerCredentialResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistryDockerCredentialResult(
            items=self.items)


def get_registry_docker_credential(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegistryDockerCredentialResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:registry/v2:getRegistryDockerCredential', __args__, opts=opts, typ=GetRegistryDockerCredentialResult).value

    return AwaitableGetRegistryDockerCredentialResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_registry_docker_credential)
def get_registry_docker_credential_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegistryDockerCredentialResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
