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
from ._enums import *

__all__ = [
    'ListKubernetesNodePoolsResult',
    'AwaitableListKubernetesNodePoolsResult',
    'list_kubernetes_node_pools',
    'list_kubernetes_node_pools_output',
]

@pulumi.output_type
class ListKubernetesNodePoolsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListKubernetesNodePoolsProperties':
        return pulumi.get(self, "items")


class AwaitableListKubernetesNodePoolsResult(ListKubernetesNodePoolsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListKubernetesNodePoolsResult(
            items=self.items)


def list_kubernetes_node_pools(cluster_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListKubernetesNodePoolsResult:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:kubernetes/v2:listKubernetesNodePools', __args__, opts=opts, typ=ListKubernetesNodePoolsResult).value

    return AwaitableListKubernetesNodePoolsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_kubernetes_node_pools)
def list_kubernetes_node_pools_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListKubernetesNodePoolsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    ...
