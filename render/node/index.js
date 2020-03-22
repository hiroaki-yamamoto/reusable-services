const grpc = require('grpc');
const loader = require('@grpc/proto-loader');
const path = require('path');
const msgpack = require('msgpack');

class RenderingService {
  constructor() {
    const protoPath = path.join(__dirname, '../proto/services.proto');
    const def = loader.loadSync(protoPath);
    this.desc = grpc.loadPackageDefinition(def);
    this.service = this.desc.render;
    this.templates = {};
    this.NoTemplateError = new Error('No such template.');
  }

  render(call, cb) {
    const { tmpName, argumentMap } = call.request;
    const argMap = msgpack.unpack(argumentMap);
    const tmp = this.templates[tmpName];
    let resp = null;
    let err = null;
    if (typeof tmp === 'function') {
      resp = { data: tmp(argMap) };
    } else {
      err = this.NoTemplateError;
    }

    cb(err, resp);
  }

  getService() {
    const svr = new grpc.Server();
    svr.addProtoService(this.service.TemplateService.service, {
      render: this.render,
    });
    return svr;
  }
}

module.exports = RenderingService;
