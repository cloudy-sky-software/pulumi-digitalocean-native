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
    'User',
    'AwaitableUser',
    'get_kubernetes_cluster_user',
    'get_kubernetes_cluster_user_output',
]

@pulumi.output_type
class User:
    def __init__(__self__, kubernetes_cluster_user=None):
        if kubernetes_cluster_user and not isinstance(kubernetes_cluster_user, dict):
            raise TypeError("Expected argument 'kubernetes_cluster_user' to be a dict")
        pulumi.set(__self__, "kubernetes_cluster_user", kubernetes_cluster_user)

    @property
    @pulumi.getter(name="kubernetesClusterUser")
    def kubernetes_cluster_user(self) -> Optional['outputs.UserKubernetesClusterUserProperties']:
        return pulumi.get(self, "kubernetes_cluster_user")


class AwaitableUser(User):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return User(
            kubernetes_cluster_user=self.kubernetes_cluster_user)


def get_kubernetes_cluster_user(cluster_id: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableUser:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:kubernetes/v2:getKubernetesClusterUser', __args__, opts=opts, typ=User).value

    return AwaitableUser(
        kubernetes_cluster_user=pulumi.get(__ret__, 'kubernetes_cluster_user'))
def get_kubernetes_cluster_user_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[User]:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_id: A unique ID that can be used to reference a Kubernetes cluster.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:kubernetes/v2:getKubernetesClusterUser', __args__, opts=opts, typ=User)
    return __ret__.apply(lambda __response__: User(
        kubernetes_cluster_user=pulumi.get(__response__, 'kubernetes_cluster_user')))
