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
from ._enums import *

__all__ = [
    'ListUptimeChecksItems',
    'AwaitableListUptimeChecksItems',
    'list_uptime_checks',
    'list_uptime_checks_output',
]

@pulumi.output_type
class ListUptimeChecksItems:
    def __init__(__self__, checks=None, links=None, meta=None):
        if checks and not isinstance(checks, list):
            raise TypeError("Expected argument 'checks' to be a list")
        pulumi.set(__self__, "checks", checks)
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)

    @property
    @pulumi.getter
    def checks(self) -> Optional[Sequence['outputs.Check']]:
        return pulumi.get(self, "checks")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")


class AwaitableListUptimeChecksItems(ListUptimeChecksItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListUptimeChecksItems(
            checks=self.checks,
            links=self.links,
            meta=self.meta)


def list_uptime_checks(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListUptimeChecksItems:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:uptime/v2:listUptimeChecks', __args__, opts=opts, typ=ListUptimeChecksItems).value

    return AwaitableListUptimeChecksItems(
        checks=pulumi.get(__ret__, 'checks'),
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'))
def list_uptime_checks_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListUptimeChecksItems]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:uptime/v2:listUptimeChecks', __args__, opts=opts, typ=ListUptimeChecksItems)
    return __ret__.apply(lambda __response__: ListUptimeChecksItems(
        checks=pulumi.get(__response__, 'checks'),
        links=pulumi.get(__response__, 'links'),
        meta=pulumi.get(__response__, 'meta')))
