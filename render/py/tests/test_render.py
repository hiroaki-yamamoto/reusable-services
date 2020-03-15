#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""Rendering features tests."""

from unittest import TestCase

# import grpc
from jinja2 import Environment, DictLoader
import msgpack

from reusable_services.render.rpc.models_pb2 import RenderingRequest
from reusable_services.render.server import Server


class RenderTest(TestCase):
    """Jinja Environment Test Base."""

    def setUp(self: TestCase):
        """Set up."""
        self.env = Environment(loader=DictLoader({
            "test": "{{ test }}",
        }))
        self.svr = Server(self.env)

    def test_render(self):
        """Should render the template."""
        resp = self.svr.render(RenderingRequest(
                tmpName="test",
                argumentMap=msgpack.dumps({"test": "Hello World"}),
        ), None)
        self.assertEqual(resp.data, "Hello World")
