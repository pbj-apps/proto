// package: heyrenee.v1
// file: heyrenee/v1/insurance_service.proto

var heyrenee_v1_insurance_service_pb = require("../../heyrenee/v1/insurance_service_pb");
var heyrenee_v1_messages_insurance_pb = require("../../heyrenee/v1/messages/insurance_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var InsuranceService = (function () {
  function InsuranceService() {}
  InsuranceService.serviceName = "heyrenee.v1.InsuranceService";
  return InsuranceService;
}());

InsuranceService.CreateInsurance = {
  methodName: "CreateInsurance",
  service: InsuranceService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_insurance_service_pb.CreateInsuranceRequest,
  responseType: heyrenee_v1_messages_insurance_pb.Insurance
};

InsuranceService.UpdateInsurance = {
  methodName: "UpdateInsurance",
  service: InsuranceService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_insurance_service_pb.UpdateInsuranceRequest,
  responseType: heyrenee_v1_messages_insurance_pb.Insurance
};

InsuranceService.ListInsurance = {
  methodName: "ListInsurance",
  service: InsuranceService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_insurance_service_pb.ListInsuranceRequest,
  responseType: heyrenee_v1_insurance_service_pb.ListInsuranceResponse
};

exports.InsuranceService = InsuranceService;

function InsuranceServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

InsuranceServiceClient.prototype.createInsurance = function createInsurance(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(InsuranceService.CreateInsurance, {
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

InsuranceServiceClient.prototype.updateInsurance = function updateInsurance(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(InsuranceService.UpdateInsurance, {
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

InsuranceServiceClient.prototype.listInsurance = function listInsurance(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(InsuranceService.ListInsurance, {
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

exports.InsuranceServiceClient = InsuranceServiceClient;

