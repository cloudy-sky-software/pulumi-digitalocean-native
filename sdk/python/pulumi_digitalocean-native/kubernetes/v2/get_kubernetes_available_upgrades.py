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
    'GetKubernetesAvailableUpgradesResult',
    'AwaitableGetKubernetesAvailableUpgradesResult',
    'get_kubernetes_available_upgrades',
    'get_kubernetes_available_upgrades_output',
]

@pulumi.output_type
class GetKubernetesAvailableUpgradesResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.GetKubernetesAvailableUpgradesProperties':
        return pulumi.get(self, "items")


class AwaitableGetKubernetesAvailableUpgradesResult(GetKubernetesAvailableUpgradesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubernetesAvailableUpgradesResult(
            items=self.items)


def get_kubernetes_available_upgrades(cluster_id: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKubernetesAvailableUpgradesResult:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:kubernetes/v2:getKubernetesAvailableUpgrades', __args__, opts=opts, typ=GetKubernetesAvailableUpgradesResult).value

    return AwaitableGetKubernetesAvailableUpgradesResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_kubernetes_available_upgrades)
def get_kubernetes_available_upgrades_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKubernetesAvailableUpgradesResult]:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    ...