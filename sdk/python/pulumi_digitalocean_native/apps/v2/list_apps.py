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
    'AppsResponse',
    'AwaitableAppsResponse',
    'list_apps',
    'list_apps_output',
]

@pulumi.output_type
class AppsResponse:
    def __init__(__self__, apps=None, links=None, meta=None):
        if apps and not isinstance(apps, list):
            raise TypeError("Expected argument 'apps' to be a list")
        pulumi.set(__self__, "apps", apps)
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)

    @property
    @pulumi.getter
    def apps(self) -> Optional[Sequence['outputs.App']]:
        return pulumi.get(self, "apps")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")


class AwaitableAppsResponse(AppsResponse):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AppsResponse(
            apps=self.apps,
            links=self.links,
            meta=self.meta)


def list_apps(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAppsResponse:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:apps/v2:listApps', __args__, opts=opts, typ=AppsResponse).value

    return AwaitableAppsResponse(
        apps=pulumi.get(__ret__, 'apps'),
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'))
def list_apps_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[AppsResponse]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:apps/v2:listApps', __args__, opts=opts, typ=AppsResponse)
    return __ret__.apply(lambda __response__: AppsResponse(
        apps=pulumi.get(__response__, 'apps'),
        links=pulumi.get(__response__, 'links'),
        meta=pulumi.get(__response__, 'meta')))
