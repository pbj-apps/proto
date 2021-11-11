// package: heyrenee.v1
// file: heyrenee/v1/appointment_service.proto

var heyrenee_v1_appointment_service_pb = require("../../heyrenee/v1/appointment_service_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var AppointmentService = (function () {
  function AppointmentService() {}
  AppointmentService.serviceName = "heyrenee.v1.AppointmentService";
  return AppointmentService;
}());

AppointmentService.CorsPreflight = {
  methodName: "CorsPreflight",
  service: AppointmentService,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: google_protobuf_empty_pb.Empty
};

AppointmentService.ListAppointments = {
  methodName: "ListAppointments",
  service: AppointmentService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_appointment_service_pb.ListAppointmentsRequest,
  responseType: heyrenee_v1_appointment_service_pb.ListAppointmentsResponse
};

exports.AppointmentService = AppointmentService;

function AppointmentServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

AppointmentServiceClient.prototype.corsPreflight = function corsPreflight(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AppointmentService.CorsPreflight, {
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

AppointmentServiceClient.prototype.listAppointments = function listAppointments(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AppointmentService.ListAppointments, {
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

exports.AppointmentServiceClient = AppointmentServiceClient;

