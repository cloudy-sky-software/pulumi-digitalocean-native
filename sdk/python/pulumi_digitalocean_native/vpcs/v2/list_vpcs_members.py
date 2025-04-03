# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'ListVpcsMembersItems',
    'AwaitableListVpcsMembersItems',
    'list_vpcs_members',
    'list_vpcs_members_output',
]

@pulumi.output_type
class ListVpcsMembersItems:
    def __init__(__self__, links=None, members=None, meta=None):
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def members(self) -> Optional[Sequence['outputs.VpcMember']]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")


class AwaitableListVpcsMembersItems(ListVpcsMembersItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListVpcsMembersItems(
            links=self.links,
            members=self.members,
            meta=self.meta)


def list_vpcs_members(vpc_id: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListVpcsMembersItems:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str vpc_id: A unique identifier for a VPC.
    """
    __args__ = dict()
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:vpcs/v2:listVpcsMembers', __args__, opts=opts, typ=ListVpcsMembersItems).value

    return AwaitableListVpcsMembersItems(
        links=pulumi.get(__ret__, 'links'),
        members=pulumi.get(__ret__, 'members'),
        meta=pulumi.get(__ret__, 'meta'))
def list_vpcs_members_output(vpc_id: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListVpcsMembersItems]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str vpc_id: A unique identifier for a VPC.
    """
    __args__ = dict()
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:vpcs/v2:listVpcsMembers', __args__, opts=opts, typ=ListVpcsMembersItems)
    return __ret__.apply(lambda __response__: ListVpcsMembersItems(
        links=pulumi.get(__response__, 'links'),
        members=pulumi.get(__response__, 'members'),
        meta=pulumi.get(__response__, 'meta')))
