const grpc = require('grpc');
const loader = require('@grpc/proto-loader');
const path = require('path');

class RenderingService {
  constructor() {
    const protoPath = path.join(__dirname, '../proto/services.proto');
    loader.load(protoPath).then((def) => {
      this.desc = grpc.loadPackageDefinition(def);
      this.render = this.desc.TemplateService;
    });
  }
}

module.exports = RenderingService;
