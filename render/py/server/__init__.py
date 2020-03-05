"""Server module."""

import grpc
from ..rpc.services_pb2_grpc import TemplateServiceServicer


class Server(TemplateServiceServicer):
    """Jinja2 template rendering server."""

    def render(
        self: TemplateServiceServicer,
        request,
        context: grpc.ServicerContext
    ):
    """Renders render the template with render arguments."""
    raise NotImplementedError("Not Implemented yet")
