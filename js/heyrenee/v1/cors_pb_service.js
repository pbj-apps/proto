// package: heyrenee.v1
// file: heyrenee/v1/cors.proto

var heyrenee_v1_cors_pb = require("../../heyrenee/v1/cors_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Cors = (function () {
  function Cors() {}
  Cors.serviceName = "heyrenee.v1.Cors";
  return Cors;
}());

Cors.Preflight = {
  methodName: "Preflight",
  service: Cors,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: google_protobuf_empty_pb.Empty
};

exports.Cors = Cors;

function CorsClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

CorsClient.prototype.preflight = function preflight(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Cors.Preflight, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.CorsClient = CorsClient;

