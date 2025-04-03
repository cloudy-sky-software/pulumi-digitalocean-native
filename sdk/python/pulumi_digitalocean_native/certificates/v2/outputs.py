# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'Certificate',
    'MetaMeta',
    'PageLinks',
    'PageLinksPagesProperties',
]

@pulumi.output_type
class Certificate(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "_sha1Fingerprint":
            suggest = "_sha1_fingerprint"
        elif key == "createdAt":
            suggest = "created_at"
        elif key == "dnsNames":
            suggest = "dns_names"
        elif key == "notAfter":
            suggest = "not_after"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Certificate. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Certificate.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Certificate.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 _sha1_fingerprint: Optional[builtins.str] = None,
                 created_at: Optional[builtins.str] = None,
                 dns_names: Optional[Sequence[builtins.str]] = None,
                 id: Optional[builtins.str] = None,
                 name: Optional[builtins.str] = None,
                 not_after: Optional[builtins.str] = None,
                 state: Optional['CertificateState'] = None,
                 type: Optional['CertificateType'] = None):
        """
        :param builtins.str _sha1_fingerprint: A unique identifier generated from the SHA-1 fingerprint of the certificate.
        :param builtins.str created_at: A time value given in ISO8601 combined date and time format that represents when the certificate was created.
        :param Sequence[builtins.str] dns_names: An array of fully qualified domain names (FQDNs) for which the certificate was issued.
        :param builtins.str id: A unique ID that can be used to identify and reference a certificate.
        :param builtins.str name: A unique human-readable name referring to a certificate.
        :param builtins.str not_after: A time value given in ISO8601 combined date and time format that represents the certificate's expiration date.
        :param 'CertificateState' state: A string representing the current state of the certificate. It may be `pending`, `verified`, or `error`.
        :param 'CertificateType' type: A string representing the type of the certificate. The value will be `custom` for a user-uploaded certificate or `lets_encrypt` for one automatically generated with Let's Encrypt.
        """
        if _sha1_fingerprint is not None:
            pulumi.set(__self__, "_sha1_fingerprint", _sha1_fingerprint)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if dns_names is not None:
            pulumi.set(__self__, "dns_names", dns_names)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if not_after is not None:
            pulumi.set(__self__, "not_after", not_after)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="_sha1Fingerprint")
    def _sha1_fingerprint(self) -> Optional[builtins.str]:
        """
        A unique identifier generated from the SHA-1 fingerprint of the certificate.
        """
        return pulumi.get(self, "_sha1_fingerprint")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[builtins.str]:
        """
        A time value given in ISO8601 combined date and time format that represents when the certificate was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dnsNames")
    def dns_names(self) -> Optional[Sequence[builtins.str]]:
        """
        An array of fully qualified domain names (FQDNs) for which the certificate was issued.
        """
        return pulumi.get(self, "dns_names")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        A unique ID that can be used to identify and reference a certificate.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        A unique human-readable name referring to a certificate.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notAfter")
    def not_after(self) -> Optional[builtins.str]:
        """
        A time value given in ISO8601 combined date and time format that represents the certificate's expiration date.
        """
        return pulumi.get(self, "not_after")

    @property
    @pulumi.getter
    def state(self) -> Optional['CertificateState']:
        """
        A string representing the current state of the certificate. It may be `pending`, `verified`, or `error`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> Optional['CertificateType']:
        """
        A string representing the type of the certificate. The value will be `custom` for a user-uploaded certificate or `lets_encrypt` for one automatically generated with Let's Encrypt.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class MetaMeta(dict):
    def __init__(__self__, *,
                 total: Optional[builtins.int] = None):
        """
        :param builtins.int total: Number of objects returned by the request.
        """
        if total is not None:
            pulumi.set(__self__, "total", total)

    @property
    @pulumi.getter
    def total(self) -> Optional[builtins.int]:
        """
        Number of objects returned by the request.
        """
        return pulumi.get(self, "total")


@pulumi.output_type
class PageLinks(dict):
    def __init__(__self__, *,
                 pages: Optional['outputs.PageLinksPagesProperties'] = None):
        if pages is not None:
            pulumi.set(__self__, "pages", pages)

    @property
    @pulumi.getter
    def pages(self) -> Optional['outputs.PageLinksPagesProperties']:
        return pulumi.get(self, "pages")


@pulumi.output_type
class PageLinksPagesProperties(dict):
    def __init__(__self__, *,
                 first: Optional[builtins.str] = None,
                 last: Optional[builtins.str] = None,
                 next: Optional[builtins.str] = None,
                 prev: Optional[builtins.str] = None):
        if first is not None:
            pulumi.set(__self__, "first", first)
        if last is not None:
            pulumi.set(__self__, "last", last)
        if next is not None:
            pulumi.set(__self__, "next", next)
        if prev is not None:
            pulumi.set(__self__, "prev", prev)

    @property
    @pulumi.getter
    def first(self) -> Optional[builtins.str]:
        return pulumi.get(self, "first")

    @property
    @pulumi.getter
    def last(self) -> Optional[builtins.str]:
        return pulumi.get(self, "last")

    @property
    @pulumi.getter
    def next(self) -> Optional[builtins.str]:
        return pulumi.get(self, "next")

    @property
    @pulumi.getter
    def prev(self) -> Optional[builtins.str]:
        return pulumi.get(self, "prev")


