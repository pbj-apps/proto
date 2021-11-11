// package: heyrenee.v1
// file: heyrenee/v1/cors_service.proto

var heyrenee_v1_cors_service_pb = require("../../heyrenee/v1/cors_service_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var CorsService = (function () {
  function CorsService() {}
  CorsService.serviceName = "heyrenee.v1.CorsService";
  return CorsService;
}());

CorsService.CorsPreflight = {
  methodName: "CorsPreflight",
  service: CorsService,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: google_protobuf_empty_pb.Empty
};

exports.CorsService = CorsService;

function CorsServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

CorsServiceClient.prototype.corsPreflight = function corsPreflight(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(CorsService.CorsPreflight, {
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

exports.CorsServiceClient = CorsServiceClient;

