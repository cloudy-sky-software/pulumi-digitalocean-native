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
    'GetCertificateProperties',
    'AwaitableGetCertificateProperties',
    'get_certificate',
    'get_certificate_output',
]

@pulumi.output_type
class GetCertificateProperties:
    def __init__(__self__, certificate=None):
        if certificate and not isinstance(certificate, dict):
            raise TypeError("Expected argument 'certificate' to be a dict")
        pulumi.set(__self__, "certificate", certificate)

    @property
    @pulumi.getter
    def certificate(self) -> Optional['outputs.Certificate']:
        return pulumi.get(self, "certificate")


class AwaitableGetCertificateProperties(GetCertificateProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificateProperties(
            certificate=self.certificate)


def get_certificate(certificate_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCertificateProperties:
    """
    Use this data source to access information about an existing resource.

    :param str certificate_id: A unique identifier for a certificate.
    """
    __args__ = dict()
    __args__['certificateId'] = certificate_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:certificates/v2:getCertificate', __args__, opts=opts, typ=GetCertificateProperties).value

    return AwaitableGetCertificateProperties(
        certificate=pulumi.get(__ret__, 'certificate'))


@_utilities.lift_output_func(get_certificate)
def get_certificate_output(certificate_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCertificateProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str certificate_id: A unique identifier for a certificate.
    """
    ...
