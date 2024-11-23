# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs

__all__ = [
    'GetKubernetesAvailableUpgradeProperties',
    'AwaitableGetKubernetesAvailableUpgradeProperties',
    'get_kubernetes_available_upgrade',
    'get_kubernetes_available_upgrade_output',
]

@pulumi.output_type
class GetKubernetesAvailableUpgradeProperties:
    def __init__(__self__, available_upgrade_versions=None):
        if available_upgrade_versions and not isinstance(available_upgrade_versions, list):
            raise TypeError("Expected argument 'available_upgrade_versions' to be a list")
        pulumi.set(__self__, "available_upgrade_versions", available_upgrade_versions)

    @property
    @pulumi.getter(name="availableUpgradeVersions")
    def available_upgrade_versions(self) -> Optional[Sequence['outputs.KubernetesVersion']]:
        return pulumi.get(self, "available_upgrade_versions")


class AwaitableGetKubernetesAvailableUpgradeProperties(GetKubernetesAvailableUpgradeProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubernetesAvailableUpgradeProperties(
            available_upgrade_versions=self.available_upgrade_versions)


def get_kubernetes_available_upgrade(cluster_id: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKubernetesAvailableUpgradeProperties:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:kubernetes/v2:getKubernetesAvailableUpgrade', __args__, opts=opts, typ=GetKubernetesAvailableUpgradeProperties).value

    return AwaitableGetKubernetesAvailableUpgradeProperties(
        available_upgrade_versions=pulumi.get(__ret__, 'available_upgrade_versions'))
def get_kubernetes_available_upgrade_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetKubernetesAvailableUpgradeProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:kubernetes/v2:getKubernetesAvailableUpgrade', __args__, opts=opts, typ=GetKubernetesAvailableUpgradeProperties)
    return __ret__.apply(lambda __response__: GetKubernetesAvailableUpgradeProperties(
        available_upgrade_versions=pulumi.get(__response__, 'available_upgrade_versions')))
