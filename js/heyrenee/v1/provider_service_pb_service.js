// package: heyrenee.v1
// file: heyrenee/v1/provider_service.proto

var heyrenee_v1_provider_service_pb = require("../../heyrenee/v1/provider_service_pb");
var heyrenee_v1_messages_patient_provider_pb = require("../../heyrenee/v1/messages/patient_provider_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ProviderService = (function () {
  function ProviderService() {}
  ProviderService.serviceName = "heyrenee.v1.ProviderService";
  return ProviderService;
}());

ProviderService.SearchProviders = {
  methodName: "SearchProviders",
  service: ProviderService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_provider_service_pb.SearchProvidersRequest,
  responseType: heyrenee_v1_provider_service_pb.SearchProvidersResponse
};

ProviderService.CreatePatientProvider = {
  methodName: "CreatePatientProvider",
  service: ProviderService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_provider_service_pb.CreatePatientProviderRequest,
  responseType: heyrenee_v1_messages_patient_provider_pb.PatientProvider
};

ProviderService.UpdatePatientProvider = {
  methodName: "UpdatePatientProvider",
  service: ProviderService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_provider_service_pb.UpdatePatientProviderRequest,
  responseType: heyrenee_v1_messages_patient_provider_pb.PatientProvider
};

ProviderService.ListPatientProviders = {
  methodName: "ListPatientProviders",
  service: ProviderService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_provider_service_pb.ListPatientProvidersRequest,
  responseType: heyrenee_v1_provider_service_pb.ListPatientProvidersResponse
};

exports.ProviderService = ProviderService;

function ProviderServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ProviderServiceClient.prototype.searchProviders = function searchProviders(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ProviderService.SearchProviders, {
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

ProviderServiceClient.prototype.createPatientProvider = function createPatientProvider(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ProviderService.CreatePatientProvider, {
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

ProviderServiceClient.prototype.updatePatientProvider = function updatePatientProvider(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ProviderService.UpdatePatientProvider, {
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

ProviderServiceClient.prototype.listPatientProviders = function listPatientProviders(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ProviderService.ListPatientProviders, {
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

exports.ProviderServiceClient = ProviderServiceClient;

