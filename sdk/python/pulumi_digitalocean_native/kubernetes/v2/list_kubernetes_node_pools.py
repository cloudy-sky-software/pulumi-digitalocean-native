# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'ListKubernetesNodePoolsProperties',
    'AwaitableListKubernetesNodePoolsProperties',
    'list_kubernetes_node_pools',
    'list_kubernetes_node_pools_output',
]

@pulumi.output_type
class ListKubernetesNodePoolsProperties:
    def __init__(__self__, node_pools=None):
        if node_pools and not isinstance(node_pools, list):
            raise TypeError("Expected argument 'node_pools' to be a list")
        pulumi.set(__self__, "node_pools", node_pools)

    @property
    @pulumi.getter(name="nodePools")
    def node_pools(self) -> Optional[Sequence['outputs.KubernetesNodePool']]:
        return pulumi.get(self, "node_pools")


class AwaitableListKubernetesNodePoolsProperties(ListKubernetesNodePoolsProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListKubernetesNodePoolsProperties(
            node_pools=self.node_pools)


def list_kubernetes_node_pools(cluster_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListKubernetesNodePoolsProperties:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:kubernetes/v2:listKubernetesNodePools', __args__, opts=opts, typ=ListKubernetesNodePoolsProperties).value

    return AwaitableListKubernetesNodePoolsProperties(
        node_pools=pulumi.get(__ret__, 'node_pools'))


@_utilities.lift_output_func(list_kubernetes_node_pools)
def list_kubernetes_node_pools_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListKubernetesNodePoolsProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    ...
