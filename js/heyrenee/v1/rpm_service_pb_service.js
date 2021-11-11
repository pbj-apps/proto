// package: heyrenee.v1
// file: heyrenee/v1/rpm_service.proto

var heyrenee_v1_rpm_service_pb = require("../../heyrenee/v1/rpm_service_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var RpmService = (function () {
  function RpmService() {}
  RpmService.serviceName = "heyrenee.v1.RpmService";
  return RpmService;
}());

RpmService.ListRpmSchedules = {
  methodName: "ListRpmSchedules",
  service: RpmService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_rpm_service_pb.ListRpmSchedulesRequest,
  responseType: heyrenee_v1_rpm_service_pb.ListRpmSchedulesResponse
};

RpmService.ListRpmMeasurements = {
  methodName: "ListRpmMeasurements",
  service: RpmService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_rpm_service_pb.ListRpmMeasurementsRequest,
  responseType: heyrenee_v1_rpm_service_pb.ListRpmMeasurementsResponse
};

exports.RpmService = RpmService;

function RpmServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

RpmServiceClient.prototype.listRpmSchedules = function listRpmSchedules(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RpmService.ListRpmSchedules, {
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

RpmServiceClient.prototype.listRpmMeasurements = function listRpmMeasurements(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RpmService.ListRpmMeasurements, {
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

exports.RpmServiceClient = RpmServiceClient;

