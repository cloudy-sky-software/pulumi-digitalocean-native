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
    'ListProjectsResourcesDefaultItems',
    'AwaitableListProjectsResourcesDefaultItems',
    'list_projects_resources_default',
    'list_projects_resources_default_output',
]

@pulumi.output_type
class ListProjectsResourcesDefaultItems:
    def __init__(__self__, links=None, meta=None, resources=None):
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)
        if resources and not isinstance(resources, list):
            raise TypeError("Expected argument 'resources' to be a list")
        pulumi.set(__self__, "resources", resources)

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter
    def resources(self) -> Optional[Sequence['outputs.Resource']]:
        return pulumi.get(self, "resources")


class AwaitableListProjectsResourcesDefaultItems(ListProjectsResourcesDefaultItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListProjectsResourcesDefaultItems(
            links=self.links,
            meta=self.meta,
            resources=self.resources)


def list_projects_resources_default(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListProjectsResourcesDefaultItems:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:projects/v2:listProjectsResourcesDefault', __args__, opts=opts, typ=ListProjectsResourcesDefaultItems).value

    return AwaitableListProjectsResourcesDefaultItems(
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'),
        resources=pulumi.get(__ret__, 'resources'))
def list_projects_resources_default_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListProjectsResourcesDefaultItems]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:projects/v2:listProjectsResourcesDefault', __args__, opts=opts, typ=ListProjectsResourcesDefaultItems)
    return __ret__.apply(lambda __response__: ListProjectsResourcesDefaultItems(
        links=pulumi.get(__response__, 'links'),
        meta=pulumi.get(__response__, 'meta'),
        resources=pulumi.get(__response__, 'resources')))
