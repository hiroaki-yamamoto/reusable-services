"""Jinja rendering service."""

import grpc
from jinja2 import Environment
import msgpack
from ..rpc.services_pb2_grpc import TemplateServiceServicer
from ..rpc.models_pb2 import RenderingRequest, RenderingResponse


class Server(TemplateServiceServicer):
    """Jinja2 template rendering server."""

    def __init__(self, env: Environment):
        """Init."""
        super().__init__()
        self.env = env

    def render(
        self: TemplateServiceServicer,
        request: RenderingRequest,
        context: grpc.ServicerContext,
    ):
        """Render render the template with render arguments."""
        tmp = self.env.get_template(request.tmpName)
        args = msgpack.loads(request.argumentMap)
        return RenderingResponse(data=tmp.render(**args))
