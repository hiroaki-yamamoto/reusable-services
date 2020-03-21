const grpc = require('grpc');
const loader = require('@grpc/proto-loader');
const path = require('path');
const msgpack = require('msgpack');

class RenderingService {
  constructor() {
    const protoPath = path.join(__dirname, '../proto/services.proto');
    loader.load(protoPath).then((def) => {
      this.desc = grpc.loadPackageDefinition(def);
      this.service = this.desc.TemplateService;
    });
    this.templates = {};
  }

  render(call, cb) {
    const { tmpName, argumentMap } = call.request;
    const argMap = msgpack.unpack(argumentMap);
    const tmp = this.templates[tmpName];
    const resp = {
      data: tmp ? tmp(argMap) : '',
    };

    cb(null, resp);
  }

  getService() {
    const svr = grpc.Server();
    svr.addProtoService(this.service.TemplateService.service, {
      render: this.render,
    });
    return svr;
  }
}

module.exports = RenderingService;
