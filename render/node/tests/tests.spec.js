/* eslint-disable no-unused-expressions */
/* global describe, beforeEach, it */
const { expect } = require('chai');
const grpc = require('grpc');
const pug = require('pug');
const msgpack = require('msgpack');
const Service = require('../index');


describe('Pug Template Rendering Service', () => {
  let args;
  let srv;
  beforeEach(() => {
    srv = new Service();
    srv.templates.testMsg = pug.compile('p #{message}');
    args = msgpack.pack({ message: 'Hello World' });
  });
  describe('With template', () => {
    it('Should compile the template properly', (done) => {
      srv.render({
        request: { tmpName: 'testMsg', argumentMap: args },
      }, (err, resp) => {
        expect(err).to.be.null;
        expect(resp).to.eql({ data: '<p>Hello World</p>' });
        done();
      });
    });
  });
  describe('Without template', () => {
    it('Should raise an error.', (done) => {
      srv.render({
        request: { tmpName: 'template', argumentMap: args },
      }, (err, resp) => {
        expect(err).to.exist;
        expect(srv.NoTemplateError).to.exist;
        expect(err).to.equal(srv.NoTemplateError);
        expect(resp).to.null;
        done();
      });
    });
  });
  it('Should Get the service', () => {
    const svr = srv.getService();
    expect(svr).to.instanceOf(grpc.Server);
  });
});
