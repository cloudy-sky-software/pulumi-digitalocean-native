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
    'ListDomainsRecordsItems',
    'AwaitableListDomainsRecordsItems',
    'list_domains_records',
    'list_domains_records_output',
]

@pulumi.output_type
class ListDomainsRecordsItems:
    def __init__(__self__, domain_records=None, links=None, meta=None):
        if domain_records and not isinstance(domain_records, list):
            raise TypeError("Expected argument 'domain_records' to be a list")
        pulumi.set(__self__, "domain_records", domain_records)
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)

    @property
    @pulumi.getter(name="domainRecords")
    def domain_records(self) -> Optional[Sequence['outputs.DomainRecord']]:
        return pulumi.get(self, "domain_records")

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")


class AwaitableListDomainsRecordsItems(ListDomainsRecordsItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDomainsRecordsItems(
            domain_records=self.domain_records,
            links=self.links,
            meta=self.meta)


def list_domains_records(domain_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDomainsRecordsItems:
    """
    Use this data source to access information about an existing resource.

    :param str domain_name: The name of the domain itself.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:domains/v2:listDomainsRecords', __args__, opts=opts, typ=ListDomainsRecordsItems).value

    return AwaitableListDomainsRecordsItems(
        domain_records=pulumi.get(__ret__, 'domain_records'),
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'))
def list_domains_records_output(domain_name: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListDomainsRecordsItems]:
    """
    Use this data source to access information about an existing resource.

    :param str domain_name: The name of the domain itself.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:domains/v2:listDomainsRecords', __args__, opts=opts, typ=ListDomainsRecordsItems)
    return __ret__.apply(lambda __response__: ListDomainsRecordsItems(
        domain_records=pulumi.get(__response__, 'domain_records'),
        links=pulumi.get(__response__, 'links'),
        meta=pulumi.get(__response__, 'meta')))
