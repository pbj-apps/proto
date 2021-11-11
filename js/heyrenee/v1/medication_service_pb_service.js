// package: heyrenee.v1
// file: heyrenee/v1/medication_service.proto

var heyrenee_v1_medication_service_pb = require("../../heyrenee/v1/medication_service_pb");
var heyrenee_v1_messages_prescription_pb = require("../../heyrenee/v1/messages/prescription_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var MedicationService = (function () {
  function MedicationService() {}
  MedicationService.serviceName = "heyrenee.v1.MedicationService";
  return MedicationService;
}());

MedicationService.MedicationSuggest = {
  methodName: "MedicationSuggest",
  service: MedicationService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_medication_service_pb.MedicationSuggestRequest,
  responseType: heyrenee_v1_medication_service_pb.MedicationSuggestResponse
};

MedicationService.CreatePrescription = {
  methodName: "CreatePrescription",
  service: MedicationService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_medication_service_pb.CreatePrescriptionRequest,
  responseType: heyrenee_v1_messages_prescription_pb.Prescription
};

MedicationService.UpdatePrescription = {
  methodName: "UpdatePrescription",
  service: MedicationService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_medication_service_pb.UpdatePrescriptionRequest,
  responseType: heyrenee_v1_messages_prescription_pb.Prescription
};

MedicationService.ListPrescriptions = {
  methodName: "ListPrescriptions",
  service: MedicationService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_medication_service_pb.ListPrescriptionsRequest,
  responseType: heyrenee_v1_medication_service_pb.ListPrescriptionsResponse
};

MedicationService.ListMedicationDoses = {
  methodName: "ListMedicationDoses",
  service: MedicationService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_medication_service_pb.ListMedicationDosesRequest,
  responseType: heyrenee_v1_medication_service_pb.ListMedicationDosesResponse
};

MedicationService.ListRefills = {
  methodName: "ListRefills",
  service: MedicationService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_medication_service_pb.ListRefillsRequest,
  responseType: heyrenee_v1_medication_service_pb.ListRefillsResponse
};

exports.MedicationService = MedicationService;

function MedicationServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

MedicationServiceClient.prototype.medicationSuggest = function medicationSuggest(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MedicationService.MedicationSuggest, {
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

MedicationServiceClient.prototype.createPrescription = function createPrescription(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MedicationService.CreatePrescription, {
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

MedicationServiceClient.prototype.updatePrescription = function updatePrescription(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MedicationService.UpdatePrescription, {
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

MedicationServiceClient.prototype.listPrescriptions = function listPrescriptions(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MedicationService.ListPrescriptions, {
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

MedicationServiceClient.prototype.listMedicationDoses = function listMedicationDoses(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MedicationService.ListMedicationDoses, {
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

MedicationServiceClient.prototype.listRefills = function listRefills(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(MedicationService.ListRefills, {
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

exports.MedicationServiceClient = MedicationServiceClient;

