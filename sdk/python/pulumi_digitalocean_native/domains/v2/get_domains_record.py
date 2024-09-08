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

__all__ = [
    'GetDomainsRecordProperties',
    'AwaitableGetDomainsRecordProperties',
    'get_domains_record',
    'get_domains_record_output',
]

@pulumi.output_type
class GetDomainsRecordProperties:
    def __init__(__self__, domain_record=None):
        if domain_record and not isinstance(domain_record, dict):
            raise TypeError("Expected argument 'domain_record' to be a dict")
        pulumi.set(__self__, "domain_record", domain_record)

    @property
    @pulumi.getter(name="domainRecord")
    def domain_record(self) -> Optional['outputs.DomainRecord']:
        return pulumi.get(self, "domain_record")


class AwaitableGetDomainsRecordProperties(GetDomainsRecordProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainsRecordProperties(
            domain_record=self.domain_record)


def get_domains_record(domain_name: Optional[str] = None,
                       domain_record_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainsRecordProperties:
    """
    Use this data source to access information about an existing resource.

    :param str domain_name: The name of the domain itself.
    :param str domain_record_id: The unique identifier of the domain record.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    __args__['domainRecordId'] = domain_record_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:domains/v2:getDomainsRecord', __args__, opts=opts, typ=GetDomainsRecordProperties).value

    return AwaitableGetDomainsRecordProperties(
        domain_record=pulumi.get(__ret__, 'domain_record'))


@_utilities.lift_output_func(get_domains_record)
def get_domains_record_output(domain_name: Optional[pulumi.Input[str]] = None,
                              domain_record_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainsRecordProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str domain_name: The name of the domain itself.
    :param str domain_record_id: The unique identifier of the domain record.
    """
    ...
