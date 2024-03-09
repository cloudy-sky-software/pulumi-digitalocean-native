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
    'GetKubernetesClusterUserResult',
    'AwaitableGetKubernetesClusterUserResult',
    'get_kubernetes_cluster_user',
    'get_kubernetes_cluster_user_output',
]

@pulumi.output_type
class GetKubernetesClusterUserResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.User':
        return pulumi.get(self, "items")


class AwaitableGetKubernetesClusterUserResult(GetKubernetesClusterUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubernetesClusterUserResult(
            items=self.items)


def get_kubernetes_cluster_user(cluster_id: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKubernetesClusterUserResult:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:kubernetes/v2:getKubernetesClusterUser', __args__, opts=opts, typ=GetKubernetesClusterUserResult).value

    return AwaitableGetKubernetesClusterUserResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_kubernetes_cluster_user)
def get_kubernetes_cluster_user_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKubernetesClusterUserResult]:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    ...
